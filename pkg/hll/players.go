package hll

import (
	"fmt"
	"strings"
	"time"
)

type SquadType string

const (
	SQUAD_TYPE_COMMANDER SquadType = "Commander"
	SQUAD_TYPE_INFANTRY  SquadType = "Infantry"
	SQUAD_TYPE_RECON     SquadType = "Recon"
	SQUAD_TYPE_ARMOR     SquadType = "Armor"
	SQUAD_TYPE_ARTILLERY SquadType = "Artillery"
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

var (
	CommandUnit = Unit{Name: CommandUnitName, ID: CommandUnitID}
	NoUnit      = Unit{Name: NoUnitName, ID: NoUnitID}
	unitNames   = []string{"Able", "Baker", "Charlie", "Dog", "Easy", "Fox", "George", "How", "Item", "Jig", "King", "Love", "Mike", "Negat", "Option", "Prep", "Queen", "Roger", "Sugar", "Tare"}
)

func UnitIDToName(unitID int) string {
	if unitID == CommandUnitID {
		return CommandUnitName
	}
	if unitID >= 0 && unitID < len(unitNames) {
		return unitNames[unitID]
	}
	return NoUnitName
}

func UnitNameToID(name string) int {
	if len(name) == 0 {
		return NoUnitID
	}
	ch := strings.ToLower(name)[0]
	if ch >= 'a' && ch <= 'z' {
		return int(ch - 'a')
	}
	return NoUnitID
}

func UnitFromString(name string) Unit {
	unitID := UnitNameToID(name)
	return Unit{
		ID:   unitID,
		Name: UnitIDToName(unitID),
	}
}

type ScoreCategory int

const (
	SCORE_CATEGORY_COMBAT ScoreCategory = iota
	SCORE_CATEGORY_OFFENSE
	SCORE_CATEGORY_DEFENSE
	SCORE_CATEGORY_SUPPORT
)

type Score struct {
	Combat  int
	Offense int
	Defense int
	Support int
}

func (s Score) GetScoreValue(scoreCategory ScoreCategory) int {
	switch scoreCategory {
	case SCORE_CATEGORY_COMBAT:
		return s.Combat
	case SCORE_CATEGORY_OFFENSE:
		return s.Offense
	case SCORE_CATEGORY_DEFENSE:
		return s.Defense
	case SCORE_CATEGORY_SUPPORT:
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
	ADMIN_ROLE_OWNER     AdminRole = "owner"
	ADMIN_ROLE_SENIOR    AdminRole = "senior"
	ADMIN_ROLE_JUNIOR    AdminRole = "junior"
	ADMIN_ROLE_SPECTATOR AdminRole = "spectator"
)

type Admin struct {
	PlayerInfo
	Role    AdminRole
	Comment string
}

type ChatScope string

const (
	CHAT_SCOPE_TEAM ChatScope = "Team"
	CHAT_SCOPE_UNIT ChatScope = "Unit"
	CHAT_SCOPE_NONE ChatScope = "None"
)

func ChatScopeFromString(name string) ChatScope {
	typed := ChatScope(name)
	switch typed {
	case CHAT_SCOPE_TEAM, CHAT_SCOPE_UNIT:
		return typed
	default:
		return CHAT_SCOPE_NONE
	}
}

type BanType int

const (
	BAN_TYPE_TEMP BanType = iota
	BAN_TYPE_PERMA
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

type PlayerPlatform string

const (
	PLAYER_PLATFORM_STEAM PlayerPlatform = "steam"
	PLAYER_PLATFORM_EPIC  PlayerPlatform = "epic"
	PLAYER_PLATFORM_XBOS  PlayerPlatform = "xbl"
	PLAYER_PLATFORM_NONE  PlayerPlatform = "none"
)

func PlayerPlatformFromString(name string) PlayerPlatform {
	typed := PlayerPlatform(name)
	switch typed {
	case PLAYER_PLATFORM_STEAM, PLAYER_PLATFORM_EPIC, PLAYER_PLATFORM_XBOS:
		return typed
	default:
		return PLAYER_PLATFORM_NONE
	}
}

type DetailedPlayerInfo struct {
	PlayerInfo
	ClanTag           string
	Platform          PlayerPlatform
	Team              TeamIdentifier
	Faction           FactionIdentifier
	Role              RoleIdentifier
	Unit              Unit
	Loadout           string
	Kills             int
	Deaths            int
	TeamKills         int
	VehicleKills      int
	VehiclesDestroyed int
	Score             Score
	Level             int
	Position          Position
}

func EmptyDetailedPlayerInfo() DetailedPlayerInfo {
	dpi := DetailedPlayerInfo{}
	dpi.Team = TEAM_NONE
	dpi.Role = ROLE_UNKNOWN
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

func (pi DetailedPlayerInfo) IsSquadLeader() bool {
	role := RoleFromString(string(pi.Role))
	return role.IsSquadLeader
}

// the distance is measured in 1 unit = 1 cm on the 2x2km map
func (pi DetailedPlayerInfo) SpacialDistanceTo(coords Position) int {
	return pi.Position.SpacialDistanceTo(coords)
}

// the distance is measured in 1 unit = 1 cm on the 2x2km map
func (pi DetailedPlayerInfo) PlanarDistanceTo(coords Position) int {
	return pi.Position.PlanarDistanceTo(coords)
}
