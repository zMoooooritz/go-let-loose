package hll

import (
	"fmt"
	"slices"
	"strings"
)

type TeamIdentifier string

const (
	TEAM_ALLIES TeamIdentifier = "Allies"
	TEAM_AXIS   TeamIdentifier = "Axis"
	TEAM_NONE   TeamIdentifier = "None"

	NoTeamID = 0
)

type Team struct {
	ID   int
	Name string
}

var teamMap = map[TeamIdentifier]Team{
	TEAM_ALLIES: {1, string(TEAM_ALLIES)},
	TEAM_AXIS:   {2, string(TEAM_AXIS)},
}

func (t TeamIdentifier) ToInt() int {
	if team, ok := teamMap[t]; ok {
		return team.ID
	}
	return NoTeamID
}

func TeamFromString(name string) TeamIdentifier {
	for id, team := range teamMap {
		if team.Name == name {
			return id
		}
	}
	return TEAM_NONE
}

func TeamFromInt(team int) TeamIdentifier {
	for id, t := range teamMap {
		if t.ID == team {
			return id
		}
	}
	return TEAM_NONE
}

type ServerView struct {
	Allies  *TeamView
	Axis    *TeamView
	Neutral *SquadView
}

func (sv ServerView) String() string {
	return fmt.Sprintf("Allies:\n %v\nAxis:\n %v\nNeutral: %v", sv.Allies, sv.Axis, sv.Neutral)
}

type TeamView struct {
	Commander DetailedPlayerInfo
	Squads    map[string]*SquadView
}

func (tv *TeamView) PlayerCount() int {
	sum := 0
	for _, squad := range tv.Squads {
		sum += squad.PlayerCount()
	}
	return sum
}

func (tv *TeamView) KillCount() int {
	sum := 0
	for _, squad := range tv.Squads {
		sum += squad.KillCount()
	}
	return sum
}

func (tv *TeamView) DeathCount() int {
	sum := 0
	for _, squad := range tv.Squads {
		sum += squad.DeathCount()
	}
	return sum
}

func (tv *TeamView) AverageLevel() int {
	sum, count := 0, 0
	for _, squad := range tv.Squads {
		count += len(squad.Players)
		for _, player := range squad.Players {
			sum += player.Level
		}
	}
	return int(sum / count)
}

func (tv *TeamView) HasPlayer(playerID string) bool {
	if tv.Commander.ID == playerID {
		return true
	}
	for _, squad := range tv.Squads {
		if squad.HasPlayer(playerID) {
			return true
		}
	}
	return false
}

func (tv *TeamView) String() string {
	str := fmt.Sprintf("Commander: %v\n", tv.Commander)
	// sort the squads by name
	squadNames := make([]string, 0, len(tv.Squads))
	for name := range tv.Squads {
		squadNames = append(squadNames, name)
	}
	slices.Sort(squadNames)

	for _, name := range squadNames {
		squad := tv.Squads[name]
		str += fmt.Sprint(squad)
	}
	return str
}

type SquadView struct {
	Team      TeamIdentifier
	SquadType SquadType
	Name      string
	Players   []DetailedPlayerInfo
}

func (sv *SquadView) PlayerCount() int {
	return len(sv.Players)
}

func (sv *SquadView) KillCount() int {
	sum := 0
	for _, player := range sv.Players {
		sum += player.Kills
	}
	return sum
}

func (sv *SquadView) DeathCount() int {
	sum := 0
	for _, player := range sv.Players {
		sum += player.Deaths
	}
	return sum
}

func (sv *SquadView) AverageLevel() int {
	sum := 0
	for _, player := range sv.Players {
		sum += player.Level
	}
	return int(sum / len(sv.Players))
}

func (sv *SquadView) HasSquadLead() bool {
	for _, player := range sv.Players {
		role := RoleFromString(string(player.Role))
		if role.IsSquadLeader {
			return true
		}
	}
	return false
}

func (sv *SquadView) HasPlayer(playerID string) bool {
	for _, player := range sv.Players {
		if player.ID == playerID {
			return true
		}
	}
	return false
}

// the distance is measured in 1 unit = 1 cm on the 2x2km map
func (sv *SquadView) CalculateSpread() int {
	if len(sv.Players) < 2 {
		return 0
	}

	var centroid Position
	for _, p := range sv.Players {
		centroid.X += p.Position.X
		centroid.Y += p.Position.Y
		centroid.Z += p.Position.Z
	}
	centroid.X /= float64(len(sv.Players))
	centroid.Y /= float64(len(sv.Players))
	centroid.Z /= float64(len(sv.Players))

	totalDist := 0
	for _, p := range sv.Players {
		totalDist += p.PlanarDistanceTo(centroid)
	}

	return totalDist / len(sv.Players)
}

// the distance is measured in 1 unit = 1 cm on the 2x2km map
func (sv *SquadView) CalculateCohesion() int {
	if len(sv.Players) < 2 {
		return 0
	}

	maxDist := 0
	for i := 0; i < len(sv.Players); i++ {
		for j := i + 1; j < len(sv.Players); j++ {
			dist := sv.Players[i].PlanarDistanceTo(sv.Players[j].Position)
			if dist > maxDist {
				maxDist = dist
			}
		}
	}
	return maxDist
}

// the distance is measured in 1 unit = 1 cm on the 2x2km map
func (sv *SquadView) CalculateLeaderDistance() int {
	if len(sv.Players) < 2 {
		return 0
	}

	var leader *DetailedPlayerInfo
	for _, p := range sv.Players {
		role := RoleFromString(string(p.Role))
		if role.IsSquadLeader {
			leader = &p
			break
		}
	}

	if leader == nil {
		return 0
	}

	totalDist := 0
	count := 0
	for _, p := range sv.Players {
		role := RoleFromString(string(p.Role))
		if !role.IsSquadLeader {
			totalDist += p.PlanarDistanceTo(leader.Position)
			count++
		}
	}

	if count == 0 {
		return 0
	}
	return totalDist / count
}

func (s *SquadView) String() string {
	str := fmt.Sprintf("  %s: [", s.Name)
	for _, player := range s.Players {
		str += fmt.Sprintf("%v, ", player)
	}
	str = strings.TrimSuffix(str, ", ")
	str += "]\n"
	return str
}

func PlayersToServerView(players []DetailedPlayerInfo) *ServerView {
	allies := &TeamView{EmptyDetailedPlayerInfo(), make(map[string]*SquadView)}
	axis := &TeamView{EmptyDetailedPlayerInfo(), make(map[string]*SquadView)}
	neutralSquad := &SquadView{TEAM_NONE, SQUAD_TYPE_INFANTRY, NeutralSquadName, []DetailedPlayerInfo{}}
	sv := ServerView{allies, axis, neutralSquad}

	for _, detailedPlayer := range players {
		if detailedPlayer.Team == TEAM_NONE {
			sv.Neutral.Players = append(sv.Neutral.Players, detailedPlayer)
		} else {
			tv := sv.Allies
			if detailedPlayer.Team == TEAM_AXIS {
				tv = sv.Axis
			}

			if detailedPlayer.Unit.ID == CommandUnitID || detailedPlayer.Role == ROLE_ARMYCOMMANDER {
				tv.Commander = detailedPlayer
			} else {
				_, ok := tv.Squads[detailedPlayer.Unit.Name]
				if !ok {
					tv.Squads[detailedPlayer.Unit.Name] = &SquadView{detailedPlayer.Team, SQUAD_TYPE_INFANTRY, detailedPlayer.Unit.Name, []DetailedPlayerInfo{}}
				}
				squad := tv.Squads[detailedPlayer.Unit.Name]
				squad.Players = append(squad.Players, detailedPlayer)
				tv.Squads[detailedPlayer.Unit.Name] = squad
			}

			if detailedPlayer.Team == TEAM_ALLIES {
				sv.Allies = tv
			} else {
				sv.Axis = tv
			}
		}
	}

	for i, s := range sv.Allies.Squads {
		s.SquadType = guessSquadType(s.Players)
		sv.Allies.Squads[i] = s
	}
	for i, s := range sv.Axis.Squads {
		s.SquadType = guessSquadType(s.Players)
		sv.Axis.Squads[i] = s
	}

	return &sv
}

func guessSquadType(players []DetailedPlayerInfo) SquadType {
	for _, player := range players {
		if player.Role == ROLE_TANKCOMMANDER || player.Role == ROLE_CREWMAN {
			return SQUAD_TYPE_ARMOR
		}
		if player.Role == ROLE_SPOTTER || player.Role == ROLE_SNIPER {
			return SQUAD_TYPE_RECON
		}
		if player.Role == ROLE_ARTILLERYOBSERVER || player.Role == ROLE_OPERATOR || player.Role == ROLE_GUNNER {
			return SQUAD_TYPE_INFANTRY
		}
	}
	return SQUAD_TYPE_INFANTRY
}

func GetOppositeSide(team TeamIdentifier) TeamIdentifier {
	if team == TEAM_ALLIES {
		return TEAM_AXIS
	}
	return TEAM_ALLIES
}

func FactionToTeam(faction FactionIdentifier) TeamIdentifier {
	if faction == FACTION_GER {
		return TEAM_AXIS
	}
	return TEAM_ALLIES
}
