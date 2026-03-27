package hll

import (
	"strings"

	"github.com/zMoooooritz/go-let-loose/pkg/logger"
)

type GameModeIdentifier string

const (
	GAMEMODE_WARFARE   GameModeIdentifier = "Warfare"
	GAMEMODE_OFFENSIVE GameModeIdentifier = "Offensive"
	GAMEMODE_SKIRMISH  GameModeIdentifier = "Skirmish"
	GAMEMODE_CONQUEST  GameModeIdentifier = "Conquest"
)

var ObjectiveCount = map[GameModeIdentifier]int{
	GAMEMODE_WARFARE:   5,
	GAMEMODE_OFFENSIVE: 5,
	GAMEMODE_SKIRMISH:  1,
}

var OptionsPerObjective = map[GameModeIdentifier]int{
	GAMEMODE_WARFARE:   3,
	GAMEMODE_OFFENSIVE: 3,
	GAMEMODE_SKIRMISH:  1,
}

type TimeOfDay string

const (
	TOD_NIGHT TimeOfDay = "Night"
	TOD_DAY   TimeOfDay = "Day"
	TOD_DUSK  TimeOfDay = "Dusk"
	TOD_DAWN  TimeOfDay = "Dawn"
)

type Weather string

const (
	WEATHER_CLEAR    Weather = "Clear"
	WEATHER_RAIN     Weather = "Rain"
	WEATHER_OVERCAST Weather = "Overcast"
	WEATHER_SNOW     Weather = "Snow"
)

type LayerIdentifier string

const (
	LAYER_CARENTAN_WARFARE                    LayerIdentifier = "carentan_warfare"
	LAYER_CARENTAN_WARFARE_NIGHT              LayerIdentifier = "carentan_warfare_night"
	LAYER_CARENTAN_OFFENSIVE_US               LayerIdentifier = "carentan_offensive_us"
	LAYER_CARENTAN_OFFENSIVE_GER              LayerIdentifier = "carentan_offensive_ger"
	LAYER_CAR_S_1944_DAY_P_SKIRMISH           LayerIdentifier = "CAR_S_1944_Day_P_Skirmish"
	LAYER_CAR_S_1944_RAIN_P_SKIRMISH          LayerIdentifier = "CAR_S_1944_Rain_P_Skirmish"
	LAYER_CAR_S_1944_DUSK_P_SKIRMISH          LayerIdentifier = "CAR_S_1944_Dusk_P_Skirmish"
	LAYER_DRIEL_WARFARE                       LayerIdentifier = "driel_warfare"
	LAYER_DRIEL_WARFARE_NIGHT                 LayerIdentifier = "driel_warfare_night"
	LAYER_DRIEL_OFFENSIVE_US                  LayerIdentifier = "driel_offensive_us"
	LAYER_DRIEL_OFFENSIVE_GER                 LayerIdentifier = "driel_offensive_ger"
	LAYER_DRL_S_1944_P_SKIRMISH               LayerIdentifier = "DRL_S_1944_P_Skirmish"
	LAYER_DRL_S_1944_NIGHT_P_SKIRMISH         LayerIdentifier = "DRL_S_1944_Night_P_Skirmish"
	LAYER_DRL_S_1944_DAY_P_SKIRMISH           LayerIdentifier = "DRL_S_1944_Day_P_Skirmish"
	LAYER_ELALAMEIN_WARFARE                   LayerIdentifier = "elalamein_warfare"
	LAYER_ELALAMEIN_WARFARE_NIGHT             LayerIdentifier = "elalamein_warfare_night"
	LAYER_ELALAMEIN_OFFENSIVE_CW              LayerIdentifier = "elalamein_offensive_CW"
	LAYER_ELALAMEIN_OFFENSIVE_GER             LayerIdentifier = "elalamein_offensive_ger"
	LAYER_ELA_S_1942_P_SKIRMISH               LayerIdentifier = "ELA_S_1942_P_Skirmish"
	LAYER_ELA_S_1942_NIGHT_P_SKIRMISH         LayerIdentifier = "ELA_S_1942_Night_P_Skirmish"
	LAYER_ELSENBORNRIDGE_WARFARE_DAY          LayerIdentifier = "elsenbornridge_warfare_day"
	LAYER_ELSENBORNRIDGE_WARFARE_MORNING      LayerIdentifier = "elsenbornridge_warfare_morning"
	LAYER_ELSENBORNRIDGE_WARFARE_NIGHT        LayerIdentifier = "elsenbornridge_warfare_night"
	LAYER_ELSENBORNRIDGE_OFFENSIVEUS_DAY      LayerIdentifier = "elsenbornridge_offensiveUS_day"
	LAYER_ELSENBORNRIDGE_OFFENSIVEUS_MORNING  LayerIdentifier = "elsenbornridge_offensiveUS_morning"
	LAYER_ELSENBORNRIDGE_OFFENSIVEUS_NIGHT    LayerIdentifier = "elsenbornridge_offensiveUS_night"
	LAYER_ELSENBORNRIDGE_OFFENSIVEGER_DAY     LayerIdentifier = "elsenbornridge_offensiveger_day"
	LAYER_ELSENBORNRIDGE_OFFENSIVEGER_MORNING LayerIdentifier = "elsenbornridge_offensiveger_morning"
	LAYER_ELSENBORNRIDGE_OFFENSIVEGER_NIGHT   LayerIdentifier = "elsenbornridge_offensiveger_night"
	LAYER_ELSENBORNRIDGE_SKIRMISH_DAY         LayerIdentifier = "elsenbornridge_skirmish_day"
	LAYER_ELSENBORNRIDGE_SKIRMISH_MORNING     LayerIdentifier = "elsenbornridge_skirmish_morning"
	LAYER_ELSENBORNRIDGE_SKIRMISH_NIGHT       LayerIdentifier = "elsenbornridge_skirmish_night"
	LAYER_FOY_WARFARE                         LayerIdentifier = "foy_warfare"
	LAYER_FOY_WARFARE_NIGHT                   LayerIdentifier = "foy_warfare_night"
	LAYER_FOY_OFFENSIVE_US                    LayerIdentifier = "foy_offensive_us"
	LAYER_FOY_OFFENSIVE_GER                   LayerIdentifier = "foy_offensive_ger"
	LAYER_HURTGENFOREST_WARFARE_V2            LayerIdentifier = "hurtgenforest_warfare_V2"
	LAYER_HURTGENFOREST_WARFARE_V2_NIGHT      LayerIdentifier = "hurtgenforest_warfare_V2_night"
	LAYER_HURTGENFOREST_OFFENSIVE_US          LayerIdentifier = "hurtgenforest_offensive_US"
	LAYER_HURTGENFOREST_OFFENSIVE_GER         LayerIdentifier = "hurtgenforest_offensive_ger"
	LAYER_HILL400_WARFARE                     LayerIdentifier = "hill400_warfare"
	LAYER_HILL400_OFFENSIVE_US                LayerIdentifier = "hill400_offensive_US"
	LAYER_HILL400_OFFENSIVE_GER               LayerIdentifier = "hill400_offensive_ger"
	LAYER_HIL_S_1944_DAY_P_SKIRMISH           LayerIdentifier = "HIL_S_1944_Day_P_Skirmish"
	LAYER_HIL_S_1944_DUSK_P_SKIRMISH          LayerIdentifier = "HIL_S_1944_Dusk_P_Skirmish"
	LAYER_KHARKOV_WARFARE                     LayerIdentifier = "kharkov_warfare"
	LAYER_KHARKOV_WARFARE_NIGHT               LayerIdentifier = "kharkov_warfare_night"
	LAYER_KHARKOV_OFFENSIVE_RUS               LayerIdentifier = "kharkov_offensive_rus"
	LAYER_KHARKOV_OFFENSIVE_GER               LayerIdentifier = "kharkov_offensive_ger"
	LAYER_KURSK_WARFARE                       LayerIdentifier = "kursk_warfare"
	LAYER_KURSK_WARFARE_NIGHT                 LayerIdentifier = "kursk_warfare_night"
	LAYER_KURSK_OFFENSIVE_RUS                 LayerIdentifier = "kursk_offensive_rus"
	LAYER_KURSK_OFFENSIVE_GER                 LayerIdentifier = "kursk_offensive_ger"
	LAYER_MORTAIN_WARFARE_DAY                 LayerIdentifier = "mortain_warfare_day"
	LAYER_MORTAIN_WARFARE_DUSK                LayerIdentifier = "mortain_warfare_dusk"
	LAYER_MORTAIN_WARFARE_OVERCAST            LayerIdentifier = "mortain_warfare_overcast"
	LAYER_MORTAIN_OFFENSIVEUS_DAY             LayerIdentifier = "mortain_offensiveUS_day"
	LAYER_MORTAIN_OFFENSIVEUS_OVERCAST        LayerIdentifier = "mortain_offensiveUS_overcast"
	LAYER_MORTAIN_OFFENSIVEUS_DUSK            LayerIdentifier = "mortain_offensiveUS_dusk"
	LAYER_MORTAIN_OFFENSIVEGER_DAY            LayerIdentifier = "mortain_offensiveger_day"
	LAYER_MORTAIN_OFFENSIVEGER_OVERCAST       LayerIdentifier = "mortain_offensiveger_overcast"
	LAYER_MORTAIN_OFFENSIVEGER_DUSK           LayerIdentifier = "mortain_offensiveger_dusk"
	LAYER_MORTAIN_SKIRMISH_DAY                LayerIdentifier = "mortain_skirmish_day"
	LAYER_MORTAIN_SKIRMISH_OVERCAST           LayerIdentifier = "mortain_skirmish_overcast"
	LAYER_MORTAIN_SKIRMISH_DUSK               LayerIdentifier = "mortain_skirmish_dusk"
	LAYER_OMAHABEACH_WARFARE                  LayerIdentifier = "omahabeach_warfare"
	LAYER_OMAHABEACH_WARFARE_NIGHT            LayerIdentifier = "omahabeach_warfare_night"
	LAYER_OMAHABEACH_OFFENSIVE_US             LayerIdentifier = "omahabeach_offensive_us"
	LAYER_OMAHABEACH_OFFENSIVE_GER            LayerIdentifier = "omahabeach_offensive_ger"
	LAYER_PHL_L_1944_WARFARE                  LayerIdentifier = "PHL_L_1944_Warfare"
	LAYER_PHL_L_1944_WARFARE_NIGHT            LayerIdentifier = "PHL_L_1944_Warfare_Night"
	LAYER_PHL_L_1944_OFFENSIVEUS              LayerIdentifier = "PHL_L_1944_OffensiveUS"
	LAYER_PHL_L_1944_OFFENSIVEGER             LayerIdentifier = "PHL_L_1944_OffensiveGER"
	LAYER_PHL_S_1944_RAIN_P_SKIRMISH          LayerIdentifier = "PHL_S_1944_Rain_P_Skirmish"
	LAYER_PHL_S_1944_MORNING_P_SKIRMISH       LayerIdentifier = "PHL_S_1944_Morning_P_Skirmish"
	LAYER_PHL_S_1944_NIGHT_P_SKIRMISH         LayerIdentifier = "PHL_S_1944_Night_P_Skirmish"
	LAYER_REM_L_1945_WARFARE                  LayerIdentifier = "REM_L_1945_Warfare"
	LAYER_REM_L_1945_WARFARENIGHT             LayerIdentifier = "REM_L_1945_WarfareNight"
	LAYER_REM_L_1945_OFFENSIVEUS              LayerIdentifier = "REM_L_1945_OffensiveUS"
	LAYER_REM_L_1945_OFFENSIVEGER             LayerIdentifier = "REM_L_1945_OffensiveGER"
	LAYER_REM_S_1945_P_SKIRMISH_DAY           LayerIdentifier = "REM_S_1945_P_Skirmish_Day"
	LAYER_REM_S_1945_P_SKIRMISH_NIGHT         LayerIdentifier = "REM_S_1945_P_Skirmish_Night"
	LAYER_SMOLENSK_WARFARE_DAY                LayerIdentifier = "smolensk_warfare_day"
	LAYER_SMOLENSK_WARFARE_DUSK               LayerIdentifier = "smolensk_warfare_dusk"
	LAYER_SMOLENSK_WARFARE_NIGHT              LayerIdentifier = "smolensk_warfare_night"
	LAYER_SMOLENSK_OFFENSIVERUS_DAY           LayerIdentifier = "smolensk_offensiveRus_day"
	LAYER_SMOLENSK_OFFENSIVERUS_DUSK          LayerIdentifier = "smolensk_offensiveRus_dusk"
	LAYER_SMOLENSK_OFFENSIVERUS_NIGHT         LayerIdentifier = "smolensk_offensiveRus_night"
	LAYER_SMOLENSK_OFFENSIVEGER_DAY           LayerIdentifier = "smolensk_offensiveGer_Day"
	LAYER_SMOLENSK_OFFENSIVEGER_DUSK          LayerIdentifier = "smolensk_offensiveGer_dusk"
	LAYER_SMOLENSK_OFFENSIVEGER_NIGHT         LayerIdentifier = "smolensk_offensiveGer_night"
	LAYER_SMOLENSK_SKIRMISH_DAY               LayerIdentifier = "smolensk_skirmish_day"
	LAYER_SMOLENSK_SKIRMISH_DUSK              LayerIdentifier = "smolensk_skirmish_dusk"
	LAYER_SMOLENSK_SKIRMISH_NIGHT             LayerIdentifier = "smolensk_skirmish_night"
	LAYER_STA_L_1942_WARFARE                  LayerIdentifier = "STA_L_1942_Warfare"
	LAYER_STA_L_1942_WARFARE_NIGHT            LayerIdentifier = "STA_L_1942_Warfare_Night"
	LAYER_STA_L_1942_OFFENSIVERUS             LayerIdentifier = "STA_L_1942_OffensiveRUS"
	LAYER_STA_L_1942_OFFENSIVEGER             LayerIdentifier = "STA_L_1942_OffensiveGER"
	LAYER_STA_S_1942_P_SKIRMISH_DUSK          LayerIdentifier = "STA_S_1942_P_Skirmish_Dusk"
	LAYER_STA_S_1942_P_SKIRMISH_OVERCAST      LayerIdentifier = "STA_S_1942_P_Skirmish_Overcast"
	LAYER_STMARIEDUMONT_WARFARE               LayerIdentifier = "stmariedumont_warfare"
	LAYER_STMARIEDUMONT_WARFARE_NIGHT         LayerIdentifier = "stmariedumont_warfare_night"
	LAYER_STMARIEDUMONT_OFF_US                LayerIdentifier = "stmariedumont_off_us"
	LAYER_STMARIEDUMONT_OFF_GER               LayerIdentifier = "stmariedumont_off_ger"
	LAYER_SMDM_S_1944_DAY_P_SKIRMISH          LayerIdentifier = "SMDM_S_1944_Day_P_Skirmish"
	LAYER_SMDM_S_1944_NIGHT_P_SKIRMISH        LayerIdentifier = "SMDM_S_1944_Night_P_Skirmish"
	LAYER_SMDM_S_1944_RAIN_P_SKIRMISH         LayerIdentifier = "SMDM_S_1944_Rain_P_Skirmish"
	LAYER_STMEREEGLISE_WARFARE                LayerIdentifier = "stmereeglise_warfare"
	LAYER_STMEREEGLISE_WARFARE_NIGHT          LayerIdentifier = "stmereeglise_warfare_night"
	LAYER_STMEREEGLISE_OFFENSIVE_US           LayerIdentifier = "stmereeglise_offensive_us"
	LAYER_STMEREEGLISE_OFFENSIVE_GER          LayerIdentifier = "stmereeglise_offensive_ger"
	LAYER_SME_S_1944_DAY_P_SKIRMISH           LayerIdentifier = "SME_S_1944_Day_P_Skirmish"
	LAYER_SME_S_1944_MORNING_P_SKIRMISH       LayerIdentifier = "SME_S_1944_Morning_P_Skirmish"
	LAYER_SME_S_1944_NIGHT_P_SKIRMISH         LayerIdentifier = "SME_S_1944_Night_P_Skirmish"
	LAYER_TOBRUK_WARFARE_DAY                  LayerIdentifier = "tobruk_warfare_day"
	LAYER_TOBRUK_WARFARE_DUSK                 LayerIdentifier = "tobruk_warfare_dusk"
	LAYER_TOBRUK_WARFARE_MORNING              LayerIdentifier = "tobruk_warfare_morning"
	LAYER_TOBRUK_OFFENSIVEBRITISH_DAY         LayerIdentifier = "tobruk_offensivebritish_day"
	LAYER_TOBRUK_OFFENSIVEGER_DAY             LayerIdentifier = "tobruk_offensiveger_day"
	LAYER_TOBRUK_OFFENSIVEBRITISH_DUSK        LayerIdentifier = "tobruk_offensivebritish_dusk"
	LAYER_TOBRUK_OFFENSIVEGER_DUSK            LayerIdentifier = "tobruk_offensiveger_dusk"
	LAYER_TOBRUK_OFFENSIVEBRITISH_MORNING     LayerIdentifier = "tobruk_offensivebritish_morning"
	LAYER_TOBRUK_OFFENSIVEGER_MORNING         LayerIdentifier = "tobruk_offensiveger_morning"
	LAYER_TOBRUK_SKIRMISH_DAY                 LayerIdentifier = "tobruk_skirmish_day"
	LAYER_TOBRUK_SKIRMISH_DUSK                LayerIdentifier = "tobruk_skirmish_dusk"
	LAYER_TOBRUK_SKIRMISH_MORNING             LayerIdentifier = "tobruk_skirmish_morning"
	LAYER_UTAHBEACH_WARFARE                   LayerIdentifier = "utahbeach_warfare"
	LAYER_UTAHBEACH_WARFARE_NIGHT             LayerIdentifier = "utahbeach_warfare_night"
	LAYER_UTAHBEACH_OFFENSIVE_US              LayerIdentifier = "utahbeach_offensive_us"
	LAYER_UTAHBEACH_OFFENSIVE_GER             LayerIdentifier = "utahbeach_offensive_ger"
)

type Layer struct {
	ID                 LayerIdentifier
	MapIdentifier      MapIdentifier
	GameModeIdentifier GameModeIdentifier
	TimeOfDay          TimeOfDay
	Weather            Weather
	PrettyName         string
	AttackingTeam      TeamIdentifier
	DefendingTeam      TeamIdentifier
	AttackingFaction   FactionIdentifier
	DefendingFaction   FactionIdentifier
}

var layerMap = map[LayerIdentifier]Layer{
	LAYER_CARENTAN_WARFARE: {
		ID:                 LAYER_CARENTAN_WARFARE,
		MapIdentifier:      MAP_CARENTAN,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Carentan Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_CARENTAN_WARFARE_NIGHT: {
		ID:                 LAYER_CARENTAN_WARFARE_NIGHT,
		MapIdentifier:      MAP_CARENTAN,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Carentan Warfare (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_CARENTAN_OFFENSIVE_US: {
		ID:                 LAYER_CARENTAN_OFFENSIVE_US,
		MapIdentifier:      MAP_CARENTAN,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Carentan Off. US",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_CARENTAN_OFFENSIVE_GER: {
		ID:                 LAYER_CARENTAN_OFFENSIVE_GER,
		MapIdentifier:      MAP_CARENTAN,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Carentan Off. GER",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
	LAYER_CAR_S_1944_DAY_P_SKIRMISH: {
		ID:                 LAYER_CAR_S_1944_DAY_P_SKIRMISH,
		MapIdentifier:      MAP_CARENTAN,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Carentan Skirmish",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_CAR_S_1944_RAIN_P_SKIRMISH: {
		ID:                 LAYER_CAR_S_1944_RAIN_P_SKIRMISH,
		MapIdentifier:      MAP_CARENTAN,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_RAIN,
		PrettyName:         "Carentan Skirmish (Rain)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_CAR_S_1944_DUSK_P_SKIRMISH: {
		ID:                 LAYER_CAR_S_1944_DUSK_P_SKIRMISH,
		MapIdentifier:      MAP_CARENTAN,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Carentan Skirmish (Dusk)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_DRIEL_WARFARE: {
		ID:                 LAYER_DRIEL_WARFARE,
		MapIdentifier:      MAP_DRIEL,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Driel Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_DRIEL_WARFARE_NIGHT: {
		ID:                 LAYER_DRIEL_WARFARE_NIGHT,
		MapIdentifier:      MAP_DRIEL,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Driel Warfare (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_DRIEL_OFFENSIVE_US: {
		ID:                 LAYER_DRIEL_OFFENSIVE_US,
		MapIdentifier:      MAP_DRIEL,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Driel Off. CW",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_CW,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_DRIEL_OFFENSIVE_GER: {
		ID:                 LAYER_DRIEL_OFFENSIVE_GER,
		MapIdentifier:      MAP_DRIEL,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Driel Off. GER",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_CW,
	},
	LAYER_DRL_S_1944_P_SKIRMISH: {
		ID:                 LAYER_DRL_S_1944_P_SKIRMISH,
		MapIdentifier:      MAP_DRIEL,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAWN,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Driel Skirmish (Dawn)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_DRL_S_1944_NIGHT_P_SKIRMISH: {
		ID:                 LAYER_DRL_S_1944_NIGHT_P_SKIRMISH,
		MapIdentifier:      MAP_DRIEL,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Driel Skirmish (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_DRL_S_1944_DAY_P_SKIRMISH: {
		ID:                 LAYER_DRL_S_1944_DAY_P_SKIRMISH,
		MapIdentifier:      MAP_DRIEL,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Driel Skirmish",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_ELALAMEIN_WARFARE: {
		ID:                 LAYER_ELALAMEIN_WARFARE,
		MapIdentifier:      MAP_ELALAMEIN,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "El Alamein Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_ELALAMEIN_WARFARE_NIGHT: {
		ID:                 LAYER_ELALAMEIN_WARFARE_NIGHT,
		MapIdentifier:      MAP_ELALAMEIN,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "El Alamein Warfare (Dusk)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_ELALAMEIN_OFFENSIVE_CW: {
		ID:                 LAYER_ELALAMEIN_OFFENSIVE_CW,
		MapIdentifier:      MAP_ELALAMEIN,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "El Alamein Off. B8A",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_B8A,
		DefendingFaction:   FACTION_DAK,
	},
	LAYER_ELALAMEIN_OFFENSIVE_GER: {
		ID:                 LAYER_ELALAMEIN_OFFENSIVE_GER,
		MapIdentifier:      MAP_ELALAMEIN,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "El Alamein Off. DAK",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_DAK,
		DefendingFaction:   FACTION_B8A,
	},
	LAYER_ELA_S_1942_P_SKIRMISH: {
		ID:                 LAYER_ELA_S_1942_P_SKIRMISH,
		MapIdentifier:      MAP_ELALAMEIN,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "El Alamein Skirmish",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_ELA_S_1942_NIGHT_P_SKIRMISH: {
		ID:                 LAYER_ELA_S_1942_NIGHT_P_SKIRMISH,
		MapIdentifier:      MAP_ELALAMEIN,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "El Alamein Skirmish (Dusk)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_ELSENBORNRIDGE_WARFARE_DAY: {
		ID:                 LAYER_ELSENBORNRIDGE_WARFARE_DAY,
		MapIdentifier:      MAP_ELSENBORNRIDGE,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_SNOW,
		PrettyName:         "Elsenborn Ridge Warfare (Snow)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_ELSENBORNRIDGE_WARFARE_MORNING: {
		ID:                 LAYER_ELSENBORNRIDGE_WARFARE_MORNING,
		MapIdentifier:      MAP_ELSENBORNRIDGE,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAWN,
		Weather:            WEATHER_SNOW,
		PrettyName:         "Elsenborn Ridge Warfare (Dawn, Snow)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_ELSENBORNRIDGE_WARFARE_NIGHT: {
		ID:                 LAYER_ELSENBORNRIDGE_WARFARE_NIGHT,
		MapIdentifier:      MAP_ELSENBORNRIDGE,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_SNOW,
		PrettyName:         "Elsenborn Ridge Warfare (Night, Snow)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_ELSENBORNRIDGE_OFFENSIVEUS_DAY: {
		ID:                 LAYER_ELSENBORNRIDGE_OFFENSIVEUS_DAY,
		MapIdentifier:      MAP_ELSENBORNRIDGE,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_SNOW,
		PrettyName:         "Elsenborn Ridge Off. US (Snow)",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_ELSENBORNRIDGE_OFFENSIVEUS_MORNING: {
		ID:                 LAYER_ELSENBORNRIDGE_OFFENSIVEUS_MORNING,
		MapIdentifier:      MAP_ELSENBORNRIDGE,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAWN,
		Weather:            WEATHER_SNOW,
		PrettyName:         "Elsenborn Ridge Off. US (Dawn, Snow)",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_ELSENBORNRIDGE_OFFENSIVEUS_NIGHT: {
		ID:                 LAYER_ELSENBORNRIDGE_OFFENSIVEUS_NIGHT,
		MapIdentifier:      MAP_ELSENBORNRIDGE,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_SNOW,
		PrettyName:         "Elsenborn Ridge Off. US (Night, Snow)",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_ELSENBORNRIDGE_OFFENSIVEGER_DAY: {
		ID:                 LAYER_ELSENBORNRIDGE_OFFENSIVEGER_DAY,
		MapIdentifier:      MAP_ELSENBORNRIDGE,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_SNOW,
		PrettyName:         "Elsenborn Ridge Off. GER (Snow)",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
	LAYER_ELSENBORNRIDGE_OFFENSIVEGER_MORNING: {
		ID:                 LAYER_ELSENBORNRIDGE_OFFENSIVEGER_MORNING,
		MapIdentifier:      MAP_ELSENBORNRIDGE,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAWN,
		Weather:            WEATHER_SNOW,
		PrettyName:         "Elsenborn Ridge Off. GER (Dawn, Snow)",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
	LAYER_ELSENBORNRIDGE_OFFENSIVEGER_NIGHT: {
		ID:                 LAYER_ELSENBORNRIDGE_OFFENSIVEGER_NIGHT,
		MapIdentifier:      MAP_ELSENBORNRIDGE,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_SNOW,
		PrettyName:         "Elsenborn Ridge Off. GER (Night, Snow)",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
	LAYER_ELSENBORNRIDGE_SKIRMISH_DAY: {
		ID:                 LAYER_ELSENBORNRIDGE_SKIRMISH_DAY,
		MapIdentifier:      MAP_ELSENBORNRIDGE,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_SNOW,
		PrettyName:         "Elsenborn Ridge Skirmish (Snow)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_ELSENBORNRIDGE_SKIRMISH_MORNING: {
		ID:                 LAYER_ELSENBORNRIDGE_SKIRMISH_MORNING,
		MapIdentifier:      MAP_ELSENBORNRIDGE,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAWN,
		Weather:            WEATHER_SNOW,
		PrettyName:         "Elsenborn Ridge Skirmish (Dawn, Snow)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_ELSENBORNRIDGE_SKIRMISH_NIGHT: {
		ID:                 LAYER_ELSENBORNRIDGE_SKIRMISH_NIGHT,
		MapIdentifier:      MAP_ELSENBORNRIDGE,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_SNOW,
		PrettyName:         "Elsenborn Ridge Skirmish (Night, Snow)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_FOY_WARFARE: {
		ID:                 LAYER_FOY_WARFARE,
		MapIdentifier:      MAP_FOY,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Foy Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_FOY_WARFARE_NIGHT: {
		ID:                 LAYER_FOY_WARFARE_NIGHT,
		MapIdentifier:      MAP_FOY,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Foy Warfare (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_FOY_OFFENSIVE_US: {
		ID:                 LAYER_FOY_OFFENSIVE_US,
		MapIdentifier:      MAP_FOY,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Foy Off. US",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_FOY_OFFENSIVE_GER: {
		ID:                 LAYER_FOY_OFFENSIVE_GER,
		MapIdentifier:      MAP_FOY,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Foy Off. GER",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
	LAYER_HURTGENFOREST_WARFARE_V2: {
		ID:                 LAYER_HURTGENFOREST_WARFARE_V2,
		MapIdentifier:      MAP_HURTGENFOREST,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Hurtgen Forest Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_HURTGENFOREST_WARFARE_V2_NIGHT: {
		ID:                 LAYER_HURTGENFOREST_WARFARE_V2_NIGHT,
		MapIdentifier:      MAP_HURTGENFOREST,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Hurtgen Forest Warfare (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_HURTGENFOREST_OFFENSIVE_US: {
		ID:                 LAYER_HURTGENFOREST_OFFENSIVE_US,
		MapIdentifier:      MAP_HURTGENFOREST,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Hurtgen Forest Off. US",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_HURTGENFOREST_OFFENSIVE_GER: {
		ID:                 LAYER_HURTGENFOREST_OFFENSIVE_GER,
		MapIdentifier:      MAP_HURTGENFOREST,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Hurtgen Forest Off. GER",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
	LAYER_HILL400_WARFARE: {
		ID:                 LAYER_HILL400_WARFARE,
		MapIdentifier:      MAP_HILL400,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Hill 400 Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_HILL400_OFFENSIVE_US: {
		ID:                 LAYER_HILL400_OFFENSIVE_US,
		MapIdentifier:      MAP_HILL400,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Hill 400 Off. US",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_HILL400_OFFENSIVE_GER: {
		ID:                 LAYER_HILL400_OFFENSIVE_GER,
		MapIdentifier:      MAP_HILL400,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Hill 400 Off. GER",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
	LAYER_HIL_S_1944_DAY_P_SKIRMISH: {
		ID:                 LAYER_HIL_S_1944_DAY_P_SKIRMISH,
		MapIdentifier:      MAP_HILL400,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Hill 400 Skirmish",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_HIL_S_1944_DUSK_P_SKIRMISH: {
		ID:                 LAYER_HIL_S_1944_DUSK_P_SKIRMISH,
		MapIdentifier:      MAP_HILL400,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Hill 400 Skirmish (Dusk)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_KHARKOV_WARFARE: {
		ID:                 LAYER_KHARKOV_WARFARE,
		MapIdentifier:      MAP_KHARKOV,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Kharkov Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_KHARKOV_WARFARE_NIGHT: {
		ID:                 LAYER_KHARKOV_WARFARE_NIGHT,
		MapIdentifier:      MAP_KHARKOV,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Kharkov Warfare (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_KHARKOV_OFFENSIVE_RUS: {
		ID:                 LAYER_KHARKOV_OFFENSIVE_RUS,
		MapIdentifier:      MAP_KHARKOV,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Kharkov Off. SOV",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_SOV,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_KHARKOV_OFFENSIVE_GER: {
		ID:                 LAYER_KHARKOV_OFFENSIVE_GER,
		MapIdentifier:      MAP_KHARKOV,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Kharkov Off. GER",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_SOV,
	},
	LAYER_KURSK_WARFARE: {
		ID:                 LAYER_KURSK_WARFARE,
		MapIdentifier:      MAP_KURSK,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Kursk Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_KURSK_WARFARE_NIGHT: {
		ID:                 LAYER_KURSK_WARFARE_NIGHT,
		MapIdentifier:      MAP_KURSK,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Kursk Warfare (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_KURSK_OFFENSIVE_RUS: {
		ID:                 LAYER_KURSK_OFFENSIVE_RUS,
		MapIdentifier:      MAP_KURSK,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Kursk Off. SOV",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_SOV,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_KURSK_OFFENSIVE_GER: {
		ID:                 LAYER_KURSK_OFFENSIVE_GER,
		MapIdentifier:      MAP_KURSK,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Kursk Off. GER",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_SOV,
	},
	LAYER_MORTAIN_WARFARE_DAY: {
		ID:                 LAYER_MORTAIN_WARFARE_DAY,
		MapIdentifier:      MAP_MORTAIN,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Mortain Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_MORTAIN_WARFARE_DUSK: {
		ID:                 LAYER_MORTAIN_WARFARE_DUSK,
		MapIdentifier:      MAP_MORTAIN,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Mortain Warfare (Dusk)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_MORTAIN_WARFARE_OVERCAST: {
		ID:                 LAYER_MORTAIN_WARFARE_OVERCAST,
		MapIdentifier:      MAP_MORTAIN,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_OVERCAST,
		PrettyName:         "Mortain Warfare (Overcast)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_MORTAIN_OFFENSIVEUS_DAY: {
		ID:                 LAYER_MORTAIN_OFFENSIVEUS_DAY,
		MapIdentifier:      MAP_MORTAIN,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Mortain Off. US",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_MORTAIN_OFFENSIVEUS_OVERCAST: {
		ID:                 LAYER_MORTAIN_OFFENSIVEUS_OVERCAST,
		MapIdentifier:      MAP_MORTAIN,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_OVERCAST,
		PrettyName:         "Mortain Off. US (Overcast)",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_MORTAIN_OFFENSIVEUS_DUSK: {
		ID:                 LAYER_MORTAIN_OFFENSIVEUS_DUSK,
		MapIdentifier:      MAP_MORTAIN,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Mortain Off. US (Dusk)",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_MORTAIN_OFFENSIVEGER_DAY: {
		ID:                 LAYER_MORTAIN_OFFENSIVEGER_DAY,
		MapIdentifier:      MAP_MORTAIN,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Mortain Off. GER",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
	LAYER_MORTAIN_OFFENSIVEGER_OVERCAST: {
		ID:                 LAYER_MORTAIN_OFFENSIVEGER_OVERCAST,
		MapIdentifier:      MAP_MORTAIN,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_OVERCAST,
		PrettyName:         "Mortain Off. GER (Overcast)",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
	LAYER_MORTAIN_OFFENSIVEGER_DUSK: {
		ID:                 LAYER_MORTAIN_OFFENSIVEGER_DUSK,
		MapIdentifier:      MAP_MORTAIN,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Mortain Off. GER (Dusk)",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
	LAYER_MORTAIN_SKIRMISH_DAY: {
		ID:                 LAYER_MORTAIN_SKIRMISH_DAY,
		MapIdentifier:      MAP_MORTAIN,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Mortain Skirmish",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_MORTAIN_SKIRMISH_OVERCAST: {
		ID:                 LAYER_MORTAIN_SKIRMISH_OVERCAST,
		MapIdentifier:      MAP_MORTAIN,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_OVERCAST,
		PrettyName:         "Mortain Skirmish (Overcast)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_MORTAIN_SKIRMISH_DUSK: {
		ID:                 LAYER_MORTAIN_SKIRMISH_DUSK,
		MapIdentifier:      MAP_MORTAIN,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Mortain Skirmish (Dusk)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_OMAHABEACH_WARFARE: {
		ID:                 LAYER_OMAHABEACH_WARFARE,
		MapIdentifier:      MAP_OMAHABEACH,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Omaha Beach Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_OMAHABEACH_WARFARE_NIGHT: {
		ID:                 LAYER_OMAHABEACH_WARFARE_NIGHT,
		MapIdentifier:      MAP_OMAHABEACH,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Omaha Beach Warfare (Dusk)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_OMAHABEACH_OFFENSIVE_US: {
		ID:                 LAYER_OMAHABEACH_OFFENSIVE_US,
		MapIdentifier:      MAP_OMAHABEACH,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Omaha Beach Off. US",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_OMAHABEACH_OFFENSIVE_GER: {
		ID:                 LAYER_OMAHABEACH_OFFENSIVE_GER,
		MapIdentifier:      MAP_OMAHABEACH,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Omaha Beach Off. GER",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
	LAYER_PHL_L_1944_WARFARE: {
		ID:                 LAYER_PHL_L_1944_WARFARE,
		MapIdentifier:      MAP_PURPLEHEARTLANE,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Purple Heart Lane Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_PHL_L_1944_WARFARE_NIGHT: {
		ID:                 LAYER_PHL_L_1944_WARFARE_NIGHT,
		MapIdentifier:      MAP_PURPLEHEARTLANE,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Purple Heart Lane Warfare (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_PHL_L_1944_OFFENSIVEUS: {
		ID:                 LAYER_PHL_L_1944_OFFENSIVEUS,
		MapIdentifier:      MAP_PURPLEHEARTLANE,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Purple Heart Lane Off. US",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_PHL_L_1944_OFFENSIVEGER: {
		ID:                 LAYER_PHL_L_1944_OFFENSIVEGER,
		MapIdentifier:      MAP_PURPLEHEARTLANE,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Purple Heart Lane Off. GER",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
	LAYER_PHL_S_1944_RAIN_P_SKIRMISH: {
		ID:                 LAYER_PHL_S_1944_RAIN_P_SKIRMISH,
		MapIdentifier:      MAP_PURPLEHEARTLANE,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_RAIN,
		PrettyName:         "Purple Heart Lane Skirmish (Rain)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_PHL_S_1944_MORNING_P_SKIRMISH: {
		ID:                 LAYER_PHL_S_1944_MORNING_P_SKIRMISH,
		MapIdentifier:      MAP_PURPLEHEARTLANE,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAWN,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Purple Heart Lane Skirmish (Dawn)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_PHL_S_1944_NIGHT_P_SKIRMISH: {
		ID:                 LAYER_PHL_S_1944_NIGHT_P_SKIRMISH,
		MapIdentifier:      MAP_PURPLEHEARTLANE,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Purple Heart Lane Skirmish (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_REM_L_1945_WARFARE: {
		ID:                 LAYER_REM_L_1945_WARFARE,
		MapIdentifier:      MAP_REMAGEN,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Remagen Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_REM_L_1945_WARFARENIGHT: {
		ID:                 LAYER_REM_L_1945_WARFARENIGHT,
		MapIdentifier:      MAP_REMAGEN,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Remagen Warfare (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_REM_L_1945_OFFENSIVEUS: {
		ID:                 LAYER_REM_L_1945_OFFENSIVEUS,
		MapIdentifier:      MAP_REMAGEN,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Remagen Off. US",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_REM_L_1945_OFFENSIVEGER: {
		ID:                 LAYER_REM_L_1945_OFFENSIVEGER,
		MapIdentifier:      MAP_REMAGEN,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Remagen Off. GER",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
	LAYER_REM_S_1945_P_SKIRMISH_DAY: {
		ID:                 LAYER_REM_S_1945_P_SKIRMISH_DAY,
		MapIdentifier:      MAP_REMAGEN,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Remagen Skirmish",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_REM_S_1945_P_SKIRMISH_NIGHT: {
		ID:                 LAYER_REM_S_1945_P_SKIRMISH_NIGHT,
		MapIdentifier:      MAP_REMAGEN,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Remagen Skirmish (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_SMOLENSK_WARFARE_DAY: {
		ID:                 LAYER_SMOLENSK_WARFARE_DAY,
		MapIdentifier:      MAP_SMOLENSK,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Smolensk Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_SMOLENSK_WARFARE_DUSK: {
		ID:                 LAYER_SMOLENSK_WARFARE_DUSK,
		MapIdentifier:      MAP_SMOLENSK,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Smolensk Warfare (Dusk)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_SMOLENSK_WARFARE_NIGHT: {
		ID:                 LAYER_SMOLENSK_WARFARE_NIGHT,
		MapIdentifier:      MAP_SMOLENSK,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Smolensk Warfare (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_SMOLENSK_OFFENSIVERUS_DAY: {
		ID:                 LAYER_SMOLENSK_OFFENSIVERUS_DAY,
		MapIdentifier:      MAP_SMOLENSK,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Smolensk Off. SOV",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_SOV,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_SMOLENSK_OFFENSIVERUS_DUSK: {
		ID:                 LAYER_SMOLENSK_OFFENSIVERUS_DUSK,
		MapIdentifier:      MAP_SMOLENSK,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Smolensk Off. SOV (Dusk)",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_SOV,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_SMOLENSK_OFFENSIVERUS_NIGHT: {
		ID:                 LAYER_SMOLENSK_OFFENSIVERUS_NIGHT,
		MapIdentifier:      MAP_SMOLENSK,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Smolensk Off. SOV (Night)",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_SOV,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_SMOLENSK_OFFENSIVEGER_DAY: {
		ID:                 LAYER_SMOLENSK_OFFENSIVEGER_DAY,
		MapIdentifier:      MAP_SMOLENSK,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Smolensk Off. GER",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_SOV,
	},
	LAYER_SMOLENSK_OFFENSIVEGER_DUSK: {
		ID:                 LAYER_SMOLENSK_OFFENSIVEGER_DUSK,
		MapIdentifier:      MAP_SMOLENSK,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Smolensk Off. GER (Dusk)",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_SOV,
	},
	LAYER_SMOLENSK_OFFENSIVEGER_NIGHT: {
		ID:                 LAYER_SMOLENSK_OFFENSIVEGER_NIGHT,
		MapIdentifier:      MAP_SMOLENSK,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Smolensk Off. GER (Night)",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_SOV,
	},
	LAYER_SMOLENSK_SKIRMISH_DAY: {
		ID:                 LAYER_SMOLENSK_SKIRMISH_DAY,
		MapIdentifier:      MAP_SMOLENSK,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Smolensk Skirmish",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_SMOLENSK_SKIRMISH_DUSK: {
		ID:                 LAYER_SMOLENSK_SKIRMISH_DUSK,
		MapIdentifier:      MAP_SMOLENSK,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Smolensk Skirmish (Dusk)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_SMOLENSK_SKIRMISH_NIGHT: {
		ID:                 LAYER_SMOLENSK_SKIRMISH_NIGHT,
		MapIdentifier:      MAP_SMOLENSK,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Smolensk Skirmish (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_STA_L_1942_WARFARE: {
		ID:                 LAYER_STA_L_1942_WARFARE,
		MapIdentifier:      MAP_STALINGRAD,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Stalingrad Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_STA_L_1942_WARFARE_NIGHT: {
		ID:                 LAYER_STA_L_1942_WARFARE_NIGHT,
		MapIdentifier:      MAP_STALINGRAD,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Stalingrad Warfare (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_STA_L_1942_OFFENSIVERUS: {
		ID:                 LAYER_STA_L_1942_OFFENSIVERUS,
		MapIdentifier:      MAP_STALINGRAD,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_OVERCAST,
		PrettyName:         "Stalingrad Off. SOV (Overcast)",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_SOV,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_STA_L_1942_OFFENSIVEGER: {
		ID:                 LAYER_STA_L_1942_OFFENSIVEGER,
		MapIdentifier:      MAP_STALINGRAD,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAWN,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Stalingrad Off. GER (Dawn)",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_SOV,
	},
	LAYER_STA_S_1942_P_SKIRMISH_DUSK: {
		ID:                 LAYER_STA_S_1942_P_SKIRMISH_DUSK,
		MapIdentifier:      MAP_STALINGRAD,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Stalingrad Skirmish (Dusk)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_STA_S_1942_P_SKIRMISH_OVERCAST: {
		ID:                 LAYER_STA_S_1942_P_SKIRMISH_OVERCAST,
		MapIdentifier:      MAP_STALINGRAD,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_OVERCAST,
		PrettyName:         "Stalingrad Skirmish (Overcast)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_STMARIEDUMONT_WARFARE: {
		ID:                 LAYER_STMARIEDUMONT_WARFARE,
		MapIdentifier:      MAP_STMARIEDUMONT,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "St. Marie Du Mont Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_STMARIEDUMONT_WARFARE_NIGHT: {
		ID:                 LAYER_STMARIEDUMONT_WARFARE_NIGHT,
		MapIdentifier:      MAP_STMARIEDUMONT,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "St. Marie Du Mont Warfare (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_STMARIEDUMONT_OFF_US: {
		ID:                 LAYER_STMARIEDUMONT_OFF_US,
		MapIdentifier:      MAP_STMARIEDUMONT,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "St. Marie Du Mont Off. US",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_STMARIEDUMONT_OFF_GER: {
		ID:                 LAYER_STMARIEDUMONT_OFF_GER,
		MapIdentifier:      MAP_STMARIEDUMONT,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "St. Marie Du Mont Off. GER",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
	LAYER_SMDM_S_1944_DAY_P_SKIRMISH: {
		ID:                 LAYER_SMDM_S_1944_DAY_P_SKIRMISH,
		MapIdentifier:      MAP_STMARIEDUMONT,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "St. Marie Du Mont Skirmish",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_SMDM_S_1944_NIGHT_P_SKIRMISH: {
		ID:                 LAYER_SMDM_S_1944_NIGHT_P_SKIRMISH,
		MapIdentifier:      MAP_STMARIEDUMONT,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "St. Marie Du Mont Skirmish (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_SMDM_S_1944_RAIN_P_SKIRMISH: {
		ID:                 LAYER_SMDM_S_1944_RAIN_P_SKIRMISH,
		MapIdentifier:      MAP_STMARIEDUMONT,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_RAIN,
		PrettyName:         "St. Marie Du Mont Skirmish (Rain)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_STMEREEGLISE_WARFARE: {
		ID:                 LAYER_STMEREEGLISE_WARFARE,
		MapIdentifier:      MAP_STMEREEGLISE,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "St. Mere Eglise Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_STMEREEGLISE_WARFARE_NIGHT: {
		ID:                 LAYER_STMEREEGLISE_WARFARE_NIGHT,
		MapIdentifier:      MAP_STMEREEGLISE,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "St. Mere Eglise Warfare (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_STMEREEGLISE_OFFENSIVE_US: {
		ID:                 LAYER_STMEREEGLISE_OFFENSIVE_US,
		MapIdentifier:      MAP_STMEREEGLISE,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "St. Mere Eglise Off. US",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_STMEREEGLISE_OFFENSIVE_GER: {
		ID:                 LAYER_STMEREEGLISE_OFFENSIVE_GER,
		MapIdentifier:      MAP_STMEREEGLISE,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "St. Mere Eglise Off. GER",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
	LAYER_SME_S_1944_DAY_P_SKIRMISH: {
		ID:                 LAYER_SME_S_1944_DAY_P_SKIRMISH,
		MapIdentifier:      MAP_STMEREEGLISE,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "St. Mere Eglise Skirmish",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_SME_S_1944_MORNING_P_SKIRMISH: {
		ID:                 LAYER_SME_S_1944_MORNING_P_SKIRMISH,
		MapIdentifier:      MAP_STMEREEGLISE,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAWN,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "St. Mere Eglise Skirmish (Dawn)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_SME_S_1944_NIGHT_P_SKIRMISH: {
		ID:                 LAYER_SME_S_1944_NIGHT_P_SKIRMISH,
		MapIdentifier:      MAP_STMEREEGLISE,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "St. Mere Eglise Skirmish (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_TOBRUK_WARFARE_DAY: {
		ID:                 LAYER_TOBRUK_WARFARE_DAY,
		MapIdentifier:      MAP_TOBRUK,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Tobruk Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_TOBRUK_WARFARE_DUSK: {
		ID:                 LAYER_TOBRUK_WARFARE_DUSK,
		MapIdentifier:      MAP_TOBRUK,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Tobruk Warfare (Dusk)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_TOBRUK_WARFARE_MORNING: {
		ID:                 LAYER_TOBRUK_WARFARE_MORNING,
		MapIdentifier:      MAP_TOBRUK,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAWN,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Tobruk Warfare (Dawn)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_TOBRUK_OFFENSIVEBRITISH_DAY: {
		ID:                 LAYER_TOBRUK_OFFENSIVEBRITISH_DAY,
		MapIdentifier:      MAP_TOBRUK,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Tobruk Off. B8A",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_B8A,
		DefendingFaction:   FACTION_DAK,
	},
	LAYER_TOBRUK_OFFENSIVEGER_DAY: {
		ID:                 LAYER_TOBRUK_OFFENSIVEGER_DAY,
		MapIdentifier:      MAP_TOBRUK,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Tobruk Off. DAK",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_DAK,
		DefendingFaction:   FACTION_B8A,
	},
	LAYER_TOBRUK_OFFENSIVEBRITISH_DUSK: {
		ID:                 LAYER_TOBRUK_OFFENSIVEBRITISH_DUSK,
		MapIdentifier:      MAP_TOBRUK,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Tobruk Off. B8A (Dusk)",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_B8A,
		DefendingFaction:   FACTION_DAK,
	},
	LAYER_TOBRUK_OFFENSIVEGER_DUSK: {
		ID:                 LAYER_TOBRUK_OFFENSIVEGER_DUSK,
		MapIdentifier:      MAP_TOBRUK,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Tobruk Off. DAK (Dusk)",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_DAK,
		DefendingFaction:   FACTION_B8A,
	},
	LAYER_TOBRUK_OFFENSIVEBRITISH_MORNING: {
		ID:                 LAYER_TOBRUK_OFFENSIVEBRITISH_MORNING,
		MapIdentifier:      MAP_TOBRUK,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAWN,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Tobruk Off. B8A (Dawn)",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_B8A,
		DefendingFaction:   FACTION_DAK,
	},
	LAYER_TOBRUK_OFFENSIVEGER_MORNING: {
		ID:                 LAYER_TOBRUK_OFFENSIVEGER_MORNING,
		MapIdentifier:      MAP_TOBRUK,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAWN,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Tobruk Off. DAK (Dawn)",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_DAK,
		DefendingFaction:   FACTION_B8A,
	},
	LAYER_TOBRUK_SKIRMISH_DAY: {
		ID:                 LAYER_TOBRUK_SKIRMISH_DAY,
		MapIdentifier:      MAP_TOBRUK,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Tobruk Skirmish",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_TOBRUK_SKIRMISH_DUSK: {
		ID:                 LAYER_TOBRUK_SKIRMISH_DUSK,
		MapIdentifier:      MAP_TOBRUK,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DUSK,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Tobruk Skirmish (Dusk)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_TOBRUK_SKIRMISH_MORNING: {
		ID:                 LAYER_TOBRUK_SKIRMISH_MORNING,
		MapIdentifier:      MAP_TOBRUK,
		GameModeIdentifier: GAMEMODE_SKIRMISH,
		TimeOfDay:          TOD_DAWN,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Tobruk Skirmish (Dawn)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_UTAHBEACH_WARFARE: {
		ID:                 LAYER_UTAHBEACH_WARFARE,
		MapIdentifier:      MAP_UTAHBEACH,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Utah Beach Warfare",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_UTAHBEACH_WARFARE_NIGHT: {
		ID:                 LAYER_UTAHBEACH_WARFARE_NIGHT,
		MapIdentifier:      MAP_UTAHBEACH,
		GameModeIdentifier: GAMEMODE_WARFARE,
		TimeOfDay:          TOD_NIGHT,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Utah Beach Warfare (Night)",
		AttackingTeam:      TEAM_NONE,
		DefendingTeam:      TEAM_NONE,
		AttackingFaction:   FACTION_UNASSIGNED,
		DefendingFaction:   FACTION_UNASSIGNED,
	},
	LAYER_UTAHBEACH_OFFENSIVE_US: {
		ID:                 LAYER_UTAHBEACH_OFFENSIVE_US,
		MapIdentifier:      MAP_UTAHBEACH,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Utah Beach Off. US",
		AttackingTeam:      TEAM_ALLIES,
		DefendingTeam:      TEAM_AXIS,
		AttackingFaction:   FACTION_US,
		DefendingFaction:   FACTION_GER,
	},
	LAYER_UTAHBEACH_OFFENSIVE_GER: {
		ID:                 LAYER_UTAHBEACH_OFFENSIVE_GER,
		MapIdentifier:      MAP_UTAHBEACH,
		GameModeIdentifier: GAMEMODE_OFFENSIVE,
		TimeOfDay:          TOD_DAY,
		Weather:            WEATHER_CLEAR,
		PrettyName:         "Utah Beach Off. GER",
		AttackingTeam:      TEAM_AXIS,
		DefendingTeam:      TEAM_ALLIES,
		AttackingFaction:   FACTION_GER,
		DefendingFaction:   FACTION_US,
	},
}

var fallback_layer = Layer{ID: "invalid", MapIdentifier: "invalid", GameModeIdentifier: GAMEMODE_WARFARE, TimeOfDay: TOD_DAY, Weather: WEATHER_CLEAR}
var previousLayer = fallback_layer

const (
	restartSuffix = "_RESTART"
	untitledMap   = "Untitled"
	loadingMap    = "Loading"
)

func (l LayerIdentifier) Layer() Layer {
	return ParseLayer(string(l))
}

func ParseLayer(layerName string) Layer {
	layerName, _ = strings.CutSuffix(layerName, restartSuffix)

	if strings.HasPrefix(layerName, loadingMap) || strings.HasPrefix(layerName, untitledMap) {
		return previousLayer
	}

	if lay, ok := layerMap[LayerIdentifier(layerName)]; ok {
		previousLayer = lay
		return lay
	}
	logger.Warn("Layer not found:", layerName)
	return fallback_layer
}

func AllLayers() []Layer {
	layers := []Layer{}
	for _, l := range layerMap {
		layers = append(layers, l)
	}
	return layers
}

func LayersByMap(mapIdentifier MapIdentifier) []Layer {
	layers := []Layer{}
	for _, l := range layerMap {
		if l.MapIdentifier == mapIdentifier {
			layers = append(layers, l)
		}
	}
	return layers
}

func LayersByMode(gameModeIdentifier GameModeIdentifier) []Layer {
	layers := []Layer{}
	for _, l := range layerMap {
		if l.GameModeIdentifier == gameModeIdentifier {
			layers = append(layers, l)
		}
	}
	return layers
}
