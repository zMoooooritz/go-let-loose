package hll

import (
	"slices"
	"strings"

	"github.com/zMoooooritz/go-let-loose/pkg/logger"
)

type VehicleIdentifier string

const (
	VEHICLE_M1_57MM                  VehicleIdentifier = "M1 57mm"
	VEHICLE_M114                     VehicleIdentifier = "M114"
	VEHICLE_M8_GREYHOUND             VehicleIdentifier = "M8 Greyhound"
	VEHICLE_STUART_M5A1              VehicleIdentifier = "Stuart M5A1"
	VEHICLE_SHERMAN_M4A3_75_W        VehicleIdentifier = "Sherman M4A3(75)W"
	VEHICLE_SHERMAN_M4A3E2           VehicleIdentifier = "Sherman M4A3E2"
	VEHICLE_SHERMAN_M4A3E2_76        VehicleIdentifier = "Sherman M4A3E2(76)"
	VEHICLE_GMC_CCKW_353_SUPPLY      VehicleIdentifier = "GMC CCKW 353 (Supply)"
	VEHICLE_GMC_CCKW_353_TRANSPORT   VehicleIdentifier = "GMC CCKW 353 (Transport)"
	VEHICLE_M3_HALF_TRACK            VehicleIdentifier = "M3 Half-track"
	VEHICLE_JEEP_WILLYS              VehicleIdentifier = "Jeep Willys"
	VEHICLE_M4A3_105MM               VehicleIdentifier = "M4A3 (105mm)"
	VEHICLE_PAK_40                   VehicleIdentifier = "PAK 40"
	VEHICLE_SFH_18                   VehicleIdentifier = "sFH 18"
	VEHICLE_SD_KFZ_234_PUMA          VehicleIdentifier = "Sd.Kfz.234 Puma"
	VEHICLE_SD_KFZ_121_LUCHS         VehicleIdentifier = "Sd.Kfz.121 Luchs"
	VEHICLE_SD_KFZ_161_PANZER_IV     VehicleIdentifier = "Sd.Kfz.161 Panzer IV"
	VEHICLE_SD_KFZ_171_PANTHER       VehicleIdentifier = "Sd.Kfz.171 Panther"
	VEHICLE_SD_KFZ_181_TIGER_1       VehicleIdentifier = "Sd.Kfz.181 Tiger 1"
	VEHICLE_OPEL_BLITZ_SUPPLY        VehicleIdentifier = "Opel Blitz (Supply)"
	VEHICLE_OPEL_BLITZ_TRANSPORT     VehicleIdentifier = "Opel Blitz (Transport)"
	VEHICLE_SD_KFZ_251_HALF_TRACK    VehicleIdentifier = "Sd.Kfz 251 Half-track"
	VEHICLE_KUBELWAGEN               VehicleIdentifier = "Kubelwagen"
	VEHICLE_STURMPANZER_IV           VehicleIdentifier = "Sturmpanzer IV"
	VEHICLE_PANZER_III_AUSF_N        VehicleIdentifier = "Panzer III Ausf.N"
	VEHICLE_ZIS_2                    VehicleIdentifier = "ZiS-2"
	VEHICLE_M1938_M_30               VehicleIdentifier = "M1938 (M-30)"
	VEHICLE_BA_10                    VehicleIdentifier = "BA-10"
	VEHICLE_T70                      VehicleIdentifier = "T70"
	VEHICLE_T34_76                   VehicleIdentifier = "T34/76"
	VEHICLE_IS_1                     VehicleIdentifier = "IS-1"
	VEHICLE_ZIS_5_SUPPLY             VehicleIdentifier = "ZIS-5 (Supply)"
	VEHICLE_ZIS_5_TRANSPORT          VehicleIdentifier = "ZIS-5 (Transport)"
	VEHICLE_GAZ_67                   VehicleIdentifier = "GAZ-67"
	VEHICLE_KV_2                     VehicleIdentifier = "KV-2"
	VEHICLE_QF_6_POUNDER             VehicleIdentifier = "QF 6-Pounder"
	VEHICLE_QF_25_POUNDER            VehicleIdentifier = "QF 25-Pounder"
	VEHICLE_DAIMLER                  VehicleIdentifier = "Daimler"
	VEHICLE_TETRARCH                 VehicleIdentifier = "Tetrarch"
	VEHICLE_M3_STUART_HONEY          VehicleIdentifier = "M3 Stuart Honey"
	VEHICLE_CROMWELL                 VehicleIdentifier = "Cromwell"
	VEHICLE_CRUSADER_MK_III          VehicleIdentifier = "Crusader Mk.III"
	VEHICLE_FIREFLY                  VehicleIdentifier = "Firefly"
	VEHICLE_CHURCHILL_MK_III         VehicleIdentifier = "Churchill Mk.III"
	VEHICLE_CHURCHILL_MK_VII         VehicleIdentifier = "Churchill Mk.VII"
	VEHICLE_BEDFORD_OYD_SUPPLY       VehicleIdentifier = "Bedford OYD (Supply)"
	VEHICLE_BEDFORD_OYD_TRANSPORT    VehicleIdentifier = "Bedford OYD (Transport)"
	VEHICLE_CHURCHILL_MK_III_A_V_R_E VehicleIdentifier = "Churchill Mk III A.V.R.E."
	VEHICLE_BISHOP_SP_25PDR          VehicleIdentifier = "Bishop SP 25pdr"
	VEHICLE_UNKNOWN                  VehicleIdentifier = "Unknown"
)

type VehicleType string

const (
	VEHICLE_TYPE_JEEP                     VehicleType = "Jeep"
	VEHICLE_TYPE_LIGHT_TANK               VehicleType = "Light Tank"
	VEHICLE_TYPE_ARTILLERY                VehicleType = "Artillery"
	VEHICLE_TYPE_MEDIUM_TANK              VehicleType = "Medium Tank"
	VEHICLE_TYPE_SUPPLY_TRUCK             VehicleType = "Supply Truck"
	VEHICLE_TYPE_SELF_PROPELLED_ARTILLERY VehicleType = "Self-Propelled Artillery"
	VEHICLE_TYPE_HEAVY_TANK               VehicleType = "Heavy Tank"
	VEHICLE_TYPE_HALF_TRACK               VehicleType = "Half-Track"
	VEHICLE_TYPE_RECON_VEHICLE            VehicleType = "Recon Vehicle"
	VEHICLE_TYPE_TRANSPORT_TRUCK          VehicleType = "Transport Truck"
	VEHICLE_TYPE_ANTI_TANK_GUN            VehicleType = "Anti-Tank Gun"
	VEHICLE_TYPE_UNKNOWN                  VehicleType = "Unknown"
)

type VehicleSeatType string

const (
	VEHICLE_SEAT_TYPE_DRIVER    VehicleSeatType = "Driver"
	VEHICLE_SEAT_TYPE_LOADER    VehicleSeatType = "Loader"
	VEHICLE_SEAT_TYPE_GUNNER    VehicleSeatType = "Gunner"
	VEHICLE_SEAT_TYPE_SPOTTER   VehicleSeatType = "Spotter"
	VEHICLE_SEAT_TYPE_PASSENGER VehicleSeatType = "Passenger"
)

type VehicleSeat struct {
	Index         int
	Type          VehicleSeatType
	Weapons       []WeaponIdentifier
	RequiresRoles []RoleIdentifier
	Exposed       bool
}

type Vehicle struct {
	ID       VehicleIdentifier
	Name     string
	Factions []FactionIdentifier
	Type     VehicleType
	Seats    []VehicleSeat
}

var vehicleMap = map[VehicleIdentifier]Vehicle{
	VEHICLE_M1_57MM: {
		ID:       VEHICLE_M1_57MM,
		Name:     "M1 57mm",
		Factions: []FactionIdentifier{FACTION_US},
		Type:     VEHICLE_TYPE_ANTI_TANK_GUN,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_57MM_CANNON_M1_57MM},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_LOADER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_M114: {
		ID:       VEHICLE_M114,
		Name:     "M114 Howitzer",
		Factions: []FactionIdentifier{FACTION_US},
		Type:     VEHICLE_TYPE_ARTILLERY,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_57MM_CANNON_M1_57MM},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_LOADER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       true,
			},
		},
	},
	VEHICLE_M8_GREYHOUND: {
		ID:       VEHICLE_M8_GREYHOUND,
		Name:     "M8 Greyhound",
		Factions: []FactionIdentifier{FACTION_US},
		Type:     VEHICLE_TYPE_RECON_VEHICLE,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_M6_37MM_M8_GREYHOUND, WEAPON_COAXIAL_M1919_M8_GREYHOUND},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         3,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       false,
			},
			{
				Index:         4,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       false,
			},
		},
	},
	VEHICLE_STUART_M5A1: {
		ID:       VEHICLE_STUART_M5A1,
		Name:     "M5A1 Stuart",
		Factions: []FactionIdentifier{FACTION_US},
		Type:     VEHICLE_TYPE_LIGHT_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_M1919_STUART_M5A1},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_37MM_CANNON_STUART_M5A1, WEAPON_COAXIAL_M1919_STUART_M5A1},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_SHERMAN_M4A3_75_W: {
		ID:       VEHICLE_SHERMAN_M4A3_75_W,
		Name:     "M4A3(75)W Sherman",
		Factions: []FactionIdentifier{FACTION_US},
		Type:     VEHICLE_TYPE_MEDIUM_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_M1919_SHERMAN_M4A3_75_W},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_M1919_SHERMAN_M4A3_75_W},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_SHERMAN_M4A3E2: {
		ID:       VEHICLE_SHERMAN_M4A3E2,
		Name:     "M4A3E2 Sherman",
		Factions: []FactionIdentifier{FACTION_US},
		Type:     VEHICLE_TYPE_HEAVY_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_M1919_SHERMAN_M4A3E2},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_75MM_M3_GUN_SHERMAN_M4A3E2, WEAPON_COAXIAL_M1919_SHERMAN_M4A3E2},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_SHERMAN_M4A3E2_76: {
		ID:       VEHICLE_SHERMAN_M4A3E2_76,
		Name:     "M4A3E2(76) Sherman",
		Factions: []FactionIdentifier{FACTION_US},
		Type:     VEHICLE_TYPE_HEAVY_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_M1919_SHERMAN_M4A3E2_76},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_76MM_M1_GUN_SHERMAN_M4A3E2_76, WEAPON_COAXIAL_M1919_SHERMAN_M4A3E2_76},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_GMC_CCKW_353_SUPPLY: {
		ID:       VEHICLE_GMC_CCKW_353_SUPPLY,
		Name:     "GMC CCKW 353",
		Factions: []FactionIdentifier{FACTION_US},
		Type:     VEHICLE_TYPE_SUPPLY_TRUCK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_GMC_CCKW_353_TRANSPORT: {
		ID:       VEHICLE_GMC_CCKW_353_TRANSPORT,
		Name:     "GMC CCKW 353",
		Factions: []FactionIdentifier{FACTION_US},
		Type:     VEHICLE_TYPE_TRANSPORT_TRUCK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         3,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         4,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         5,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         6,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         7,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         8,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         9,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         10,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         11,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_M3_HALF_TRACK: {
		ID:       VEHICLE_M3_HALF_TRACK,
		Name:     "M3 Half-track",
		Factions: []FactionIdentifier{FACTION_US},
		Type:     VEHICLE_TYPE_HALF_TRACK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_M2_BROWNING_M3_HALF_TRACK},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         3,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         4,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         5,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         6,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         7,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_JEEP_WILLYS: {
		ID:       VEHICLE_JEEP_WILLYS,
		Name:     "Willy's Jeep",
		Factions: []FactionIdentifier{FACTION_US},
		Type:     VEHICLE_TYPE_JEEP,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         3,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_M4A3_105MM: {
		ID:       VEHICLE_M4A3_105MM,
		Name:     "Sherman M4(105)",
		Factions: []FactionIdentifier{FACTION_US},
		Type:     VEHICLE_TYPE_SELF_PROPELLED_ARTILLERY,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_M1919_M4A3_105MM},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_105MM_HOWITZER_M4A3_105MM, WEAPON_COAXIAL_M1919_M4A3_105MM},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_PAK_40: {
		ID:       VEHICLE_PAK_40,
		Name:     "Pak 40",
		Factions: []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Type:     VEHICLE_TYPE_ANTI_TANK_GUN,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_75MM_CANNON_PAK_40},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_LOADER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_SFH_18: {
		ID:       VEHICLE_SFH_18,
		Name:     "sFH 18",
		Factions: []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Type:     VEHICLE_TYPE_ARTILLERY,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_150MM_HOWITZER_SFH_18},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_LOADER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       true,
			},
		},
	},
	VEHICLE_SD_KFZ_234_PUMA: {
		ID:       VEHICLE_SD_KFZ_234_PUMA,
		Name:     "Sd.Kfz.234 Puma",
		Factions: []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Type:     VEHICLE_TYPE_RECON_VEHICLE,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_50MM_KWK_39_1_SD_KFZ_234_PUMA, WEAPON_COAXIAL_MG34_SD_KFZ_234_PUMA},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         3,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       false,
			},
			{
				Index:         4,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       false,
			},
		},
	},
	VEHICLE_SD_KFZ_121_LUCHS: {
		ID:       VEHICLE_SD_KFZ_121_LUCHS,
		Name:     "Sd.Kfz.121 Luchs",
		Factions: []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Type:     VEHICLE_TYPE_LIGHT_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_20MM_KWK_30_SD_KFZ_121_LUCHS, WEAPON_COAXIAL_MG34_SD_KFZ_121_LUCHS},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_SD_KFZ_161_PANZER_IV: {
		ID:       VEHICLE_SD_KFZ_161_PANZER_IV,
		Name:     "Sd.Kfz.161 Panzer IV",
		Factions: []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Type:     VEHICLE_TYPE_MEDIUM_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_MG34_SD_KFZ_161_PANZER_IV},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER, ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_75MM_CANNON_SD_KFZ_161_PANZER_IV, WEAPON_COAXIAL_MG34_SD_KFZ_161_PANZER_IV},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER, ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER, ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_SD_KFZ_171_PANTHER: {
		ID:       VEHICLE_SD_KFZ_171_PANTHER,
		Name:     "Sd.Kfz.171 Panther",
		Factions: []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Type:     VEHICLE_TYPE_HEAVY_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_MG34_SD_KFZ_171_PANTHER},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_75MM_CANNON_SD_KFZ_171_PANTHER, WEAPON_COAXIAL_MG34_SD_KFZ_171_PANTHER},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_SD_KFZ_181_TIGER_1: {
		ID:       VEHICLE_SD_KFZ_181_TIGER_1,
		Name:     "Sd.Kfz.181 Tiger 1",
		Factions: []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Type:     VEHICLE_TYPE_HEAVY_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_MG34_SD_KFZ_181_TIGER_1},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_88_KWK_36_L_56_SD_KFZ_181_TIGER_1, WEAPON_COAXIAL_MG34_SD_KFZ_181_TIGER_1},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_OPEL_BLITZ_SUPPLY: {
		ID:       VEHICLE_OPEL_BLITZ_SUPPLY,
		Name:     "Opel Blitz",
		Factions: []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Type:     VEHICLE_TYPE_SUPPLY_TRUCK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_OPEL_BLITZ_TRANSPORT: {
		ID:       VEHICLE_OPEL_BLITZ_TRANSPORT,
		Name:     "Opel Blitz",
		Factions: []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Type:     VEHICLE_TYPE_TRANSPORT_TRUCK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         3,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         4,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         5,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         6,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         7,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         8,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         9,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         10,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         11,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_SD_KFZ_251_HALF_TRACK: {
		ID:       VEHICLE_SD_KFZ_251_HALF_TRACK,
		Name:     "Sd.Kfz 251 Half-track",
		Factions: []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Type:     VEHICLE_TYPE_HALF_TRACK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_M2_BROWNING_M3_HALF_TRACK},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         3,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         4,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         5,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         6,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         7,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_KUBELWAGEN: {
		ID:       VEHICLE_KUBELWAGEN,
		Name:     "Kubelwagen",
		Factions: []FactionIdentifier{FACTION_GER, FACTION_DAK},
		Type:     VEHICLE_TYPE_JEEP,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         3,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_STURMPANZER_IV: {
		ID:       VEHICLE_STURMPANZER_IV,
		Name:     "Sturmpanzer IV",
		Factions: []FactionIdentifier{FACTION_GER},
		Type:     VEHICLE_TYPE_SELF_PROPELLED_ARTILLERY,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_STUH_43_L_12_STURMPANZER_IV},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_PANZER_III_AUSF_N: {
		ID:       VEHICLE_PANZER_III_AUSF_N,
		Name:     "Panzer III",
		Factions: []FactionIdentifier{FACTION_DAK},
		Type:     VEHICLE_TYPE_SELF_PROPELLED_ARTILLERY,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_MG34_PANZER_III_AUSF_N},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_7_5CM_KWK_37_PANZER_III_AUSF_N, WEAPON_COAXIAL_MG34_PANZER_III_AUSF_N},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_ZIS_2: {
		ID:       VEHICLE_ZIS_2,
		Name:     "ZiS-2",
		Factions: []FactionIdentifier{FACTION_SOV},
		Type:     VEHICLE_TYPE_ANTI_TANK_GUN,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_57MM_CANNON_ZIS_2},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_LOADER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_M1938_M_30: {
		ID:       VEHICLE_M1938_M_30,
		Name:     "M-30",
		Factions: []FactionIdentifier{FACTION_SOV},
		Type:     VEHICLE_TYPE_ARTILLERY,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_122MM_HOWITZER_M1938_M_30},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_LOADER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       true,
			},
		},
	},
	VEHICLE_BA_10: {
		ID:       VEHICLE_BA_10,
		Name:     "BA-10",
		Factions: []FactionIdentifier{FACTION_SOV},
		Type:     VEHICLE_TYPE_RECON_VEHICLE,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_19_K_45MM_BA_10, WEAPON_COAXIAL_DT_BA_10},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         3,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       false,
			},
			{
				Index:         4,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       false,
			},
		},
	},
	VEHICLE_T70: {
		ID:       VEHICLE_T70,
		Name:     "T70",
		Factions: []FactionIdentifier{FACTION_SOV},
		Type:     VEHICLE_TYPE_LIGHT_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_45MM_M1937_T70, WEAPON_COAXIAL_DT_T70},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_T34_76: {
		ID:       VEHICLE_T34_76,
		Name:     "T34/76",
		Factions: []FactionIdentifier{FACTION_SOV},
		Type:     VEHICLE_TYPE_MEDIUM_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_DT_T34_76},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_76MM_ZIS_5_T34_76, WEAPON_COAXIAL_DT_T34_76},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_IS_1: {
		ID:       VEHICLE_IS_1,
		Name:     "IS-1",
		Factions: []FactionIdentifier{FACTION_SOV},
		Type:     VEHICLE_TYPE_HEAVY_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_DT_IS_1},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_D_5T_85MM_IS_1, WEAPON_COAXIAL_DT_IS_1},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_ZIS_5_SUPPLY: {
		ID:       VEHICLE_ZIS_5_SUPPLY,
		Name:     "ZIS-5",
		Factions: []FactionIdentifier{FACTION_SOV},
		Type:     VEHICLE_TYPE_SUPPLY_TRUCK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_ZIS_5_TRANSPORT: {
		ID:       VEHICLE_ZIS_5_TRANSPORT,
		Name:     "ZIS-5",
		Factions: []FactionIdentifier{FACTION_SOV},
		Type:     VEHICLE_TYPE_TRANSPORT_TRUCK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         3,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         4,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         5,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         6,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         7,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         8,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         9,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         10,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         11,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_GAZ_67: {
		ID:       VEHICLE_GAZ_67,
		Name:     "GAZ-67",
		Factions: []FactionIdentifier{FACTION_SOV},
		Type:     VEHICLE_TYPE_JEEP,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         3,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_KV_2: {
		ID:       VEHICLE_KV_2,
		Name:     "KV-2",
		Factions: []FactionIdentifier{FACTION_SOV},
		Type:     VEHICLE_TYPE_SELF_PROPELLED_ARTILLERY,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_DT_KV_2},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_152MM_M_10T_KV_2},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_QF_6_POUNDER: {
		ID:       VEHICLE_QF_6_POUNDER,
		Name:     "QF 6-Pounder",
		Factions: []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Type:     VEHICLE_TYPE_ANTI_TANK_GUN,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_QF_6_POUNDER_QF_6_POUNDER},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_LOADER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_QF_25_POUNDER: {
		ID:       VEHICLE_QF_25_POUNDER,
		Name:     "QF 25-Pounder",
		Factions: []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Type:     VEHICLE_TYPE_ARTILLERY,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_QF_25_POUNDER_QF_25_POUNDER},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_LOADER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_DAIMLER: {
		ID:       VEHICLE_DAIMLER,
		Name:     "Daimler",
		Factions: []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Type:     VEHICLE_TYPE_RECON_VEHICLE,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_QF_2_POUNDER_DAIMLER, WEAPON_COAXIAL_BESA_DAIMLER},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         3,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       false,
			},
			{
				Index:         4,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       false,
			},
		},
	},
	VEHICLE_TETRARCH: {
		ID:       VEHICLE_TETRARCH,
		Name:     "Tetrarch",
		Factions: []FactionIdentifier{FACTION_CW},
		Type:     VEHICLE_TYPE_LIGHT_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_QF_2_POUNDER_TETRARCH, WEAPON_COAXIAL_BESA_TETRARCH},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_M3_STUART_HONEY: {
		ID:       VEHICLE_M3_STUART_HONEY,
		Name:     "M3 Stuart Honey",
		Factions: []FactionIdentifier{FACTION_B8A},
		Type:     VEHICLE_TYPE_LIGHT_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_M1919_M3_STUART_HONEY},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_37MM_CANNON_M3_STUART_HONEY, WEAPON_COAXIAL_M1919_M3_STUART_HONEY},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_CROMWELL: {
		ID:       VEHICLE_CROMWELL,
		Name:     "Cromwell",
		Factions: []FactionIdentifier{FACTION_CW},
		Type:     VEHICLE_TYPE_MEDIUM_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_BESA_CROMWELL},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_OQF_75MM_CROMWELL, WEAPON_COAXIAL_BESA_CROMWELL},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_CRUSADER_MK_III: {
		ID:       VEHICLE_CRUSADER_MK_III,
		Name:     "Crusader Mk III",
		Factions: []FactionIdentifier{FACTION_B8A},
		Type:     VEHICLE_TYPE_MEDIUM_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_OQF_57MM_CRUSADER_MK_III, WEAPON_COAXIAL_BESA_CRUSADER_MK_III},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_FIREFLY: {
		ID:       VEHICLE_FIREFLY,
		Name:     "Sherman Firefly",
		Factions: []FactionIdentifier{FACTION_CW},
		Type:     VEHICLE_TYPE_HEAVY_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_QF_17_POUNDER_FIREFLY, WEAPON_COAXIAL_M1919_FIREFLY},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_CHURCHILL_MK_III: {
		ID:       VEHICLE_CHURCHILL_MK_III,
		Name:     "Churchill Mk III",
		Factions: []FactionIdentifier{FACTION_B8A},
		Type:     VEHICLE_TYPE_HEAVY_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_BESA_7_92MM_CHURCHILL_MK_III},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_OQF_57MM_CHURCHILL_MK_III, WEAPON_COAXIAL_BESA_7_92MM_CHURCHILL_MK_III},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_CHURCHILL_MK_VII: {
		ID:       VEHICLE_CHURCHILL_MK_VII,
		Name:     "Churchill Mk VII",
		Factions: []FactionIdentifier{FACTION_CW},
		Type:     VEHICLE_TYPE_HEAVY_TANK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_BESA_7_92MM_CHURCHILL_MK_VII},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_OQF_75MM_CHURCHILL_MK_VII, WEAPON_COAXIAL_BESA_7_92MM_CHURCHILL_MK_VII},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_CREWMAN, ROLE_TANKCOMMANDER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_BEDFORD_OYD_SUPPLY: {
		ID:       VEHICLE_BEDFORD_OYD_SUPPLY,
		Name:     "Bedford OYD",
		Factions: []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Type:     VEHICLE_TYPE_SUPPLY_TRUCK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_BEDFORD_OYD_TRANSPORT: {
		ID:       VEHICLE_BEDFORD_OYD_TRANSPORT,
		Name:     "Bedford OYD",
		Factions: []FactionIdentifier{FACTION_CW, FACTION_B8A},
		Type:     VEHICLE_TYPE_TRANSPORT_TRUCK,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         3,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         4,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         5,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         6,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         7,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         8,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         9,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         10,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
			{
				Index:         11,
				Type:          VEHICLE_SEAT_TYPE_PASSENGER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{},
				Exposed:       true,
			},
		},
	},
	VEHICLE_CHURCHILL_MK_III_A_V_R_E: {
		ID:       VEHICLE_CHURCHILL_MK_III_A_V_R_E,
		Name:     "Churchill AVRE",
		Factions: []FactionIdentifier{FACTION_CW},
		Type:     VEHICLE_TYPE_SELF_PROPELLED_ARTILLERY,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{WEAPON_HULL_BESA_7_92MM_CHURCHILL_MK_III_A_V_R_E},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_230MM_PETARD_CHURCHILL_MK_III_A_V_R_E, WEAPON_COAXIAL_BESA_7_92MM_CHURCHILL_MK_III_A_V_R_E},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
		},
	},
	VEHICLE_BISHOP_SP_25PDR: {
		ID:       VEHICLE_BISHOP_SP_25PDR,
		Name:     "Bishop",
		Factions: []FactionIdentifier{FACTION_B8A},
		Type:     VEHICLE_TYPE_SELF_PROPELLED_ARTILLERY,
		Seats: []VehicleSeat{
			{
				Index:         0,
				Type:          VEHICLE_SEAT_TYPE_DRIVER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
			{
				Index:         1,
				Type:          VEHICLE_SEAT_TYPE_GUNNER,
				Weapons:       []WeaponIdentifier{WEAPON_QF_25_POUNDER_BISHOP_SP_25PDR},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
			{
				Index:         2,
				Type:          VEHICLE_SEAT_TYPE_SPOTTER,
				Weapons:       []WeaponIdentifier{},
				RequiresRoles: []RoleIdentifier{ROLE_ARTILLERYOBSERVER, ROLE_OPERATOR, ROLE_GUNNER},
				Exposed:       false,
			},
		},
	},
}

var fallback_vehicle = Vehicle{
	ID:       VEHICLE_UNKNOWN,
	Name:     string(VEHICLE_UNKNOWN),
	Factions: []FactionIdentifier{},
	Type:     VEHICLE_TYPE_UNKNOWN,
	Seats:    []VehicleSeat{},
}

func (v VehicleIdentifier) Vehicle() Vehicle {
	return ParseVehicle(string(v))
}

func ParseVehicle(vehicleIdentifier string) Vehicle {
	vi := VehicleIdentifier(vehicleIdentifier)
	if vehicle, ok := vehicleMap[vi]; ok {
		return vehicle
	}
	for _, v := range vehicleMap {
		if strings.HasPrefix(string(v.ID), vehicleIdentifier) {
			logger.Info("Using", v.ID, "as fallback for", vehicleIdentifier)
			return v
		}
	}
	logger.Error("Vehicle unparseable:", vehicleIdentifier)
	return fallback_vehicle
}

func AllVehicles() []Vehicle {
	vehicles := []Vehicle{}
	for _, v := range vehicleMap {
		vehicles = append(vehicles, v)
	}
	return vehicles
}

func VehiclesByFaction(faction FactionIdentifier) []Vehicle {
	vehicles := []Vehicle{}
	for _, v := range vehicleMap {
		if slices.Contains(v.Factions, faction) {
			vehicles = append(vehicles, v)
		}
	}
	return vehicles
}

func VehiclesByType(vehicleType VehicleType) []Vehicle {
	vehicles := []Vehicle{}
	for _, v := range vehicleMap {
		if v.Type == vehicleType {
			vehicles = append(vehicles, v)
		}
	}
	return vehicles
}
