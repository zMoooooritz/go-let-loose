package hll

import (
	"fmt"
	"strings"
	"time"
)

type Role string

const (
	ArmyCommander      Role = "armycommander"
	Officer            Role = "officer"
	Rifleman           Role = "rifleman"
	Assault            Role = "assault"
	AutomaticRifleman  Role = "automaticrifleman"
	Medic              Role = "medic"
	Support            Role = "support"
	HeavyMachinegunner Role = "heavymachinegunner"
	AntiTank           Role = "antitank"
	Engineer           Role = "engineer"
	TankCommander      Role = "tankcommander"
	Crewman            Role = "crewman"
	Spotter            Role = "spotter"
	Sniper             Role = "sniper"
	NoRole             Role = "none"
)

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
	NoUnitName       = "NONE"
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

type TeamData struct {
	Allies int
	Axis   int
}

type GameState struct {
	PlayerCount      TeamData
	GameScore        TeamData
	RemainingSeconds int
	CurrentMap       Layer
	NextMap          Layer
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

type DetailedPlayerInfo struct {
	PlayerInfo
	Team    Team
	Role    Role
	Unit    Unit
	Loadout string
	Kills   int
	Deaths  int
	Score   Score
	Level   int
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

func (s SquadView) String() string {
	str := fmt.Sprintf("  %s: [", s.Name)
	for _, player := range s.Players {
		str += fmt.Sprintf("%v, ", player)
	}
	str += "]\n"
	return str
}

func PlayerstoServerView(players []DetailedPlayerInfo) ServerView {
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

			if detailedPlayer.Unit.ID == CommandUnitID {
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
