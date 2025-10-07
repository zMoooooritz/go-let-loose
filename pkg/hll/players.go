package hll

import (
	"fmt"
	"math"
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
	leaderRoles = []Role{ArmyCommander, Officer, Spotter, TankCommander}
	roleMap     = map[string]Role{
		"ArmyCommander":      ArmyCommander,
		"Officer":            Officer,
		"Rifleman":           Rifleman,
		"Assault":            Assault,
		"AutomaticRifleman":  AutomaticRifleman,
		"Medic":              Medic,
		"Support":            Support,
		"HeavyMachineGunner": HeavyMachinegunner,
		"AntiTank":           AntiTank,
		"Engineer":           Engineer,
		"TankCommander":      TankCommander,
		"Crewman":            Crewman,
		"Spotter":            Spotter,
		"Sniper":             Sniper,
	}
	roleIntMap = []Role{
		Rifleman, Assault, AutomaticRifleman, Medic, Spotter, Support,
		HeavyMachinegunner, AntiTank, Engineer, Officer, Sniper,
		Crewman, TankCommander, ArmyCommander,
	}
)

func RoleFromString(name string) Role {
	if role, ok := roleMap[name]; ok {
		return role
	}
	return NoRole
}

func RoleFromInt(id int) Role {
	if id >= 0 && id < len(roleIntMap) {
		return roleIntMap[id]
	}
	return NoRole
}

func (r Role) ToInt() int {
	for i, role := range roleIntMap {
		if r == role {
			return i
		}
	}
	return NoRoleID
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
	NoRoleID         = -1
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
	if unitID > 0 && unitID <= len(unitNames) {
		return unitNames[unitID-1]
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
	Role    AdminRole
	Comment string
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

type PlayerPlatform string

const (
	PlayerPlatformSteam PlayerPlatform = "steam"
	PlayerPlatformEpic  PlayerPlatform = "epic"
	PlayerPlatformXbos  PlayerPlatform = "xbl"
	PlayerPlatformNone  PlayerPlatform = "none"
)

func PlayerPlatformFromString(name string) PlayerPlatform {
	typed := PlayerPlatform(name)
	switch typed {
	case PlayerPlatformSteam, PlayerPlatformEpic, PlayerPlatformXbos:
		return typed
	default:
		return PlayerPlatformNone
	}
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
	Platform PlayerPlatform
	Team     Team
	Faction  Faction
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
