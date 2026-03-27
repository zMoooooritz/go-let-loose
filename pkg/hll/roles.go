package hll

type RoleIdentifier string

const (
	ROLE_ARMYCOMMANDER      RoleIdentifier = "ArmyCommander"
	ROLE_MEDIC              RoleIdentifier = "Medic"
	ROLE_ASSAULT            RoleIdentifier = "Assault"
	ROLE_AUTOMATICRIFLEMAN  RoleIdentifier = "AutomaticRifleman"
	ROLE_HEAVYMACHINEGUNNER RoleIdentifier = "HeavyMachineGunner"
	ROLE_TANKCOMMANDER      RoleIdentifier = "TankCommander"
	ROLE_SPOTTER            RoleIdentifier = "Spotter"
	ROLE_CREWMAN            RoleIdentifier = "Crewman"
	ROLE_OPERATOR           RoleIdentifier = "Operator"
	ROLE_ANTITANK           RoleIdentifier = "AntiTank"
	ROLE_OFFICER            RoleIdentifier = "Officer"
	ROLE_RIFLEMAN           RoleIdentifier = "Rifleman"
	ROLE_SNIPER             RoleIdentifier = "Sniper"
	ROLE_SUPPORT            RoleIdentifier = "Support"
	ROLE_ARTILLERYOBSERVER  RoleIdentifier = "ArtilleryObserver"
	ROLE_GUNNER             RoleIdentifier = "Gunner"
	ROLE_ENGINEER           RoleIdentifier = "Engineer"
	ROLE_UNKNOWN            RoleIdentifier = "Unknown"
)

func (r RoleIdentifier) Role() Role {
	if role, ok := roleMap[r]; ok {
		return role
	}
	return Role{}
}

func RoleIdentifierFromInt(id int) RoleIdentifier {
	for _, role := range roleMap {
		if role.ID == id {
			return role.Name
		}
	}
	return ROLE_UNKNOWN
}

type Role struct {
	ID                int
	Name              RoleIdentifier
	PrettyName        string
	RoleType          SquadType
	IsSquadLeader     bool
	KillCombatScore   int
	AssistCombatScore int
}

var roleMap = map[RoleIdentifier]Role{
	ROLE_RIFLEMAN: {
		ID:                0,
		Name:              ROLE_RIFLEMAN,
		PrettyName:        "Rifleman",
		RoleType:          SQUAD_TYPE_INFANTRY,
		IsSquadLeader:     false,
		KillCombatScore:   3,
		AssistCombatScore: 2,
	},
	ROLE_ASSAULT: {
		ID:                1,
		Name:              ROLE_ASSAULT,
		PrettyName:        "Assault",
		RoleType:          SQUAD_TYPE_INFANTRY,
		IsSquadLeader:     false,
		KillCombatScore:   6,
		AssistCombatScore: 4,
	},
	ROLE_AUTOMATICRIFLEMAN: {
		ID:                2,
		Name:              ROLE_AUTOMATICRIFLEMAN,
		PrettyName:        "Automatic Rifleman",
		RoleType:          SQUAD_TYPE_INFANTRY,
		IsSquadLeader:     false,
		KillCombatScore:   6,
		AssistCombatScore: 4,
	},
	ROLE_MEDIC: {
		ID:                3,
		Name:              ROLE_MEDIC,
		PrettyName:        "Medic",
		RoleType:          SQUAD_TYPE_INFANTRY,
		IsSquadLeader:     false,
		KillCombatScore:   6,
		AssistCombatScore: 4,
	},
	ROLE_SPOTTER: {
		ID:                4,
		Name:              ROLE_SPOTTER,
		PrettyName:        "Spotter",
		RoleType:          SQUAD_TYPE_RECON,
		IsSquadLeader:     true,
		KillCombatScore:   6,
		AssistCombatScore: 4,
	},
	ROLE_SUPPORT: {
		ID:                5,
		Name:              ROLE_SUPPORT,
		PrettyName:        "Support",
		RoleType:          SQUAD_TYPE_INFANTRY,
		IsSquadLeader:     false,
		KillCombatScore:   6,
		AssistCombatScore: 4,
	},
	ROLE_HEAVYMACHINEGUNNER: {
		ID:                6,
		Name:              ROLE_HEAVYMACHINEGUNNER,
		PrettyName:        "Machine Gunner",
		RoleType:          SQUAD_TYPE_INFANTRY,
		IsSquadLeader:     false,
		KillCombatScore:   9,
		AssistCombatScore: 6,
	},
	ROLE_ANTITANK: {
		ID:                7,
		Name:              ROLE_ANTITANK,
		PrettyName:        "Anti-Tank",
		RoleType:          SQUAD_TYPE_INFANTRY,
		IsSquadLeader:     false,
		KillCombatScore:   9,
		AssistCombatScore: 6,
	},
	ROLE_ENGINEER: {
		ID:                8,
		Name:              ROLE_ENGINEER,
		PrettyName:        "Engineer",
		RoleType:          SQUAD_TYPE_INFANTRY,
		IsSquadLeader:     false,
		KillCombatScore:   9,
		AssistCombatScore: 6,
	},
	ROLE_OFFICER: {
		ID:                9,
		Name:              ROLE_OFFICER,
		PrettyName:        "Officer",
		RoleType:          SQUAD_TYPE_INFANTRY,
		IsSquadLeader:     true,
		KillCombatScore:   9,
		AssistCombatScore: 6,
	},
	ROLE_SNIPER: {
		ID:                10,
		Name:              ROLE_SNIPER,
		PrettyName:        "Sniper",
		RoleType:          SQUAD_TYPE_RECON,
		IsSquadLeader:     false,
		KillCombatScore:   6,
		AssistCombatScore: 4,
	},
	ROLE_CREWMAN: {
		ID:                11,
		Name:              ROLE_CREWMAN,
		PrettyName:        "Crewman",
		RoleType:          SQUAD_TYPE_ARMOR,
		IsSquadLeader:     false,
		KillCombatScore:   3,
		AssistCombatScore: 2,
	},
	ROLE_TANKCOMMANDER: {
		ID:                12,
		Name:              ROLE_TANKCOMMANDER,
		PrettyName:        "Tank Commander",
		RoleType:          SQUAD_TYPE_ARMOR,
		IsSquadLeader:     true,
		KillCombatScore:   9,
		AssistCombatScore: 6,
	},
	ROLE_ARMYCOMMANDER: {
		ID:                13,
		Name:              ROLE_ARMYCOMMANDER,
		PrettyName:        "Commander",
		RoleType:          SQUAD_TYPE_COMMANDER,
		IsSquadLeader:     true,
		KillCombatScore:   12,
		AssistCombatScore: 8,
	},
	ROLE_ARTILLERYOBSERVER: {
		ID:                14,
		Name:              ROLE_ARTILLERYOBSERVER,
		PrettyName:        "Artillery Observer",
		RoleType:          SQUAD_TYPE_ARTILLERY,
		IsSquadLeader:     true,
		KillCombatScore:   9,
		AssistCombatScore: 6,
	},
	ROLE_OPERATOR: {
		ID:                15,
		Name:              ROLE_OPERATOR,
		PrettyName:        "Operator",
		RoleType:          SQUAD_TYPE_ARTILLERY,
		IsSquadLeader:     false,
		KillCombatScore:   9,
		AssistCombatScore: 6,
	},
	ROLE_GUNNER: {
		ID:                16,
		Name:              ROLE_GUNNER,
		PrettyName:        "Gunner",
		RoleType:          SQUAD_TYPE_ARTILLERY,
		IsSquadLeader:     false,
		KillCombatScore:   6,
		AssistCombatScore: 4,
	},
}

func RoleFromString(name string) Role {
	role := RoleIdentifier(name)
	if role, ok := roleMap[role]; ok {
		return role
	}
	return Role{}
}

func RoleFromInt(id int) Role {
	for _, role := range roleMap {
		if role.ID == id {
			return role
		}
	}
	return Role{}
}
