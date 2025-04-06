package hll

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"time"
)

type Role string

const (
	ArmyCommander      Role = "ArmyCommander"
	Officer            Role = "Officer"
	Rifleman           Role = "Rifleman"
	Assault            Role = "Assault"
	AutomaticRifleman  Role = "AutomaticRifleman"
	Medic              Role = "Medic"
	Support            Role = "Support"
	HeavyMachinegunner Role = "HeavyMachineGunner"
	AntiTank           Role = "AntiTank"
	Engineer           Role = "Engineer"
	TankCommander      Role = "TankCommander"
	Crewman            Role = "Crewman"
	Spotter            Role = "Spotter"
	Sniper             Role = "Sniper"
	NoRole             Role = "None"
)

var (
	leaderRoles = []Role{
		ArmyCommander,
		Officer,
		Spotter,
		TankCommander,
	}
)

func RoleFromString(name string) Role {
	typed := Role(name)
	switch typed {
	case ArmyCommander, Officer, Rifleman, Assault, AutomaticRifleman, Medic, Support, HeavyMachinegunner, AntiTank, Engineer, TankCommander, Crewman, Spotter, Sniper:
		return typed
	default:
		return NoRole
	}
}

func RoleFromInt(id int) Role {
	switch id {
	case 0:
		return Rifleman
	case 1:
		return Assault
	case 2:
		return AutomaticRifleman
	case 3:
		return Medic
	case 4:
		return Spotter
	case 5:
		return Support
	case 6:
		return HeavyMachinegunner
	case 7:
		return AntiTank
	case 8:
		return Engineer
	case 9:
		return Officer
	case 10:
		return Sniper
	case 11:
		return Crewman
	case 12:
		return TankCommander
	case 13:
		return ArmyCommander
	default:
		return NoRole
	}
}

type SquadType string

const (
	StInfanty SquadType = "Infantry"
	StRecon   SquadType = "Recon"
	StArmor   SquadType = "Armor"
)

const (
	CommandUnitID    = 100
	CommandUnitName  = "Command"
	NoUnitID         = -1
	NoUnitName       = "None"
	NoPlayerID       = "NONE"
	NoLoadout        = "NONE"
	NeutralSquadName = "Neutral"
)

type Unit struct {
	Name string
	ID   int
}

var CommandUnit = Unit{
	Name: CommandUnitName,
	ID:   CommandUnitID,
}

var NoUnit = Unit{
	Name: NoUnitName,
	ID:   NoUnitID,
}

type ScoreCategory int

const (
	ScCombat ScoreCategory = iota
	ScOffense
	ScDefense
	ScSupport
)

type Score struct {
	Combat  int
	Offense int
	Defense int
	Support int
}

func (s Score) GetScoreValue(scoreCategory ScoreCategory) int {
	switch scoreCategory {
	case ScCombat:
		return s.Combat
	case ScOffense:
		return s.Offense
	case ScDefense:
		return s.Defense
	case ScSupport:
		return s.Support
	}
	return 0
}

type PlayerInfo struct {
	Name string
	ID   string
}

type AdminRole string

const (
	ArOwner     AdminRole = "owner"
	ArSenior    AdminRole = "senior"
	ArJunior    AdminRole = "junior"
	ArSpectator AdminRole = "spectator"
)

type Admin struct {
	PlayerInfo
	Role AdminRole
}

type ChatScope string

const (
	CsTeam ChatScope = "Team"
	CsUnit ChatScope = "Unit"
	CsNone ChatScope = "None"
)

func ChatScopeFromString(name string) ChatScope {
	typed := ChatScope(name)
	switch typed {
	case CsTeam, CsUnit:
		return typed
	default:
		return CsNone
	}
}

type BanType int

const (
	TempBan BanType = iota
	PermaBan
)

type ServerBan struct {
	Type      BanType
	Player    PlayerInfo
	Timestamp time.Time
	Duration  time.Duration
	Reason    string
	AdminName string
	RawLog    string
}

func IsNameProblematic(name string) bool {
	if strings.HasSuffix(name, " ") || strings.HasSuffix(name, "?") && len([]rune(name)) > 20 {
		return true
	}
	return false
}

type Position struct {
	X float64
	Y float64
	Z float64
}

func (p Position) IsActive() bool {
	return p.X != 0 || p.Y != 0 || p.Z != 0
}

type DetailedPlayerInfo struct {
	PlayerInfo
	ClanTag  string
	Platform Platform
	Team     Team
	Role     Role
	Unit     Unit
	Loadout  string
	Kills    int
	Deaths   int
	Score    Score
	Level    int
	Position Position
}

func EmptyDetailedPlayerInfo() DetailedPlayerInfo {
	dpi := DetailedPlayerInfo{}
	dpi.Team = TmNone
	dpi.Role = NoRole
	dpi.Unit = NoUnit
	dpi.Loadout = NoLoadout
	return dpi
}

func (pi DetailedPlayerInfo) String() string {
	return fmt.Sprintf("%s [%d] (%s)", pi.Name, pi.Level, pi.Role)
}

func (pi DetailedPlayerInfo) IsSpawned() bool {
	return pi.Position.IsActive()
}

// the distance is measured in 1 unit = 1 cm on the 2x2km map
func (pi DetailedPlayerInfo) SpacialDistanceTo(coords Position) int {
	diffX := pi.Position.X - coords.X
	diffY := pi.Position.Y - coords.Y
	diffZ := pi.Position.Z - coords.Z
	return int(math.Sqrt(diffX*diffX + diffY*diffY + diffZ*diffZ))
}

// the distance is measured in 1 unit = 1 cm on the 2x2km map
func (pi DetailedPlayerInfo) PlanarDistanceTo(coords Position) int {
	diffX := pi.Position.X - coords.X
	diffY := pi.Position.Y - coords.Y
	return int(math.Sqrt(diffX*diffX + diffY*diffY))
}

type ServerView struct {
	Allies  TeamView
	Axis    TeamView
	Neutral SquadView
}

func (sv ServerView) String() string {
	return fmt.Sprintf("Allies:\n %v\nAxis:\n %v\nNeutral: %v", sv.Allies, sv.Axis, sv.Neutral)
}

type TeamView struct {
	Commander DetailedPlayerInfo
	Squads    map[string]SquadView
}

func (tv TeamView) PlayerCount() int {
	sum := 0
	for _, squad := range tv.Squads {
		sum += squad.PlayerCount()
	}
	return sum
}

func (tv TeamView) KillCount() int {
	sum := 0
	for _, squad := range tv.Squads {
		sum += squad.KillCount()
	}
	return sum
}

func (tv TeamView) DeathCount() int {
	sum := 0
	for _, squad := range tv.Squads {
		sum += squad.DeathCount()
	}
	return sum
}

func (tv TeamView) AverageLevel() int {
	sum, count := 0, 0
	for _, squad := range tv.Squads {
		count += len(squad.Players)
		for _, player := range squad.Players {
			sum += player.Level
		}
	}
	return int(sum / count)
}

func (tv TeamView) String() string {
	str := fmt.Sprintf("Commander: %v\n", tv.Commander)
	for _, squad := range tv.Squads {
		str += fmt.Sprint(squad)
	}
	return str
}

type SquadView struct {
	Team      Team
	SquadType SquadType
	Name      string
	Players   []DetailedPlayerInfo
}

func (sv SquadView) PlayerCount() int {
	return len(sv.Players)
}

func (sv SquadView) KillCount() int {
	sum := 0
	for _, player := range sv.Players {
		sum += player.Kills
	}
	return sum
}

func (sv SquadView) DeathCount() int {
	sum := 0
	for _, player := range sv.Players {
		sum += player.Deaths
	}
	return sum
}

func (sv SquadView) AverageLevel() int {
	sum := 0
	for _, player := range sv.Players {
		sum += player.Level
	}
	return int(sum / len(sv.Players))
}

func (sv SquadView) HasSquadLead() bool {
	for _, player := range sv.Players {
		if slices.Contains(leaderRoles, player.Role) {
			return true
		}
	}
	return false
}

// the distance is measured in 1 unit = 1 cm on the 2x2km map
func (sv SquadView) CalculateSpread() int {
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
func (sv SquadView) CalculateCohesion() int {
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
func (sv SquadView) CalculateLeaderDistance() int {
	if len(sv.Players) < 2 {
		return 0
	}

	var leader *DetailedPlayerInfo
	for _, p := range sv.Players {
		if slices.Contains(leaderRoles, p.Role) {
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
		if !slices.Contains(leaderRoles, p.Role) {
			totalDist += p.PlanarDistanceTo(leader.Position)
			count++
		}
	}

	if count == 0 {
		return 0
	}
	return totalDist / count
}

func (s SquadView) String() string {
	str := fmt.Sprintf("  %s: [", s.Name)
	for _, player := range s.Players {
		str += fmt.Sprintf("%v, ", player)
	}
	str += "]\n"
	return str
}

func PlayersToServerView(players []DetailedPlayerInfo) ServerView {
	allies := TeamView{DetailedPlayerInfo{}, make(map[string]SquadView)}
	axis := TeamView{DetailedPlayerInfo{}, make(map[string]SquadView)}
	sv := ServerView{allies, axis, SquadView{TmNone, StInfanty, NeutralSquadName, []DetailedPlayerInfo{}}}

	for _, detailedPlayer := range players {
		if detailedPlayer.Team == TmNone {
			sv.Neutral.Players = append(sv.Neutral.Players, detailedPlayer)
		} else {
			tv := sv.Allies
			if detailedPlayer.Team == TmAxis {
				tv = sv.Axis
			}

			if detailedPlayer.Unit.ID == CommandUnitID || detailedPlayer.Role == ArmyCommander {
				tv.Commander = detailedPlayer
			} else {
				_, ok := tv.Squads[detailedPlayer.Unit.Name]
				if !ok {
					tv.Squads[detailedPlayer.Unit.Name] = SquadView{detailedPlayer.Team, StInfanty, detailedPlayer.Unit.Name, []DetailedPlayerInfo{}}
				}
				squad := tv.Squads[detailedPlayer.Unit.Name]
				squad.Players = append(squad.Players, detailedPlayer)
				tv.Squads[detailedPlayer.Unit.Name] = squad
			}

			if detailedPlayer.Team == TmAllies {
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

	return sv
}

func guessSquadType(players []DetailedPlayerInfo) SquadType {
	for _, player := range players {
		if player.Role == TankCommander || player.Role == Crewman {
			return StArmor
		}
		if player.Role == Spotter || player.Role == Sniper {
			return StRecon
		}
	}
	return StInfanty
}

func GetOppositeSide(team Team) Team {
	if team == TmAllies {
		return TmAxis
	}
	return TmAllies
}

func FactionToTeam(faction Faction) Team {
	if faction == FctGER {
		return TmAxis
	}
	return TmAllies
}
