package hll

import (
	"strings"

	"github.com/zMoooooritz/go-let-loose/pkg/logger"
)

type Map string

const (
	MP_STMEREEGLISE    Map = "stmereeglise"
	MP_STMARIEDUMONT   Map = "stmariedumont"
	MP_UTAHBEACH       Map = "utahbeach"
	MP_OMAHABEACH      Map = "omahabeach"
	MP_PURPLEHEARTLANE Map = "purpleheartlane"
	MP_CARENTAN        Map = "carentan"
	MP_HURTGENFOREST   Map = "hurtgenforest"
	MP_HILL400         Map = "hill400"
	MP_FOY             Map = "foy"
	MP_KURSK           Map = "kursk"
	MP_STALINGRAD      Map = "stalingrad"
	MP_REMAGEN         Map = "remagen"
	MP_KHARKOV         Map = "kharkov"
	MP_DRIEL           Map = "driel"
	MP_ELALAMEIN       Map = "elalamein"
	MP_MORTAIN         Map = "mortain"
	MP_ELSENBORNRIDGE  Map = "elsenbornridge"
	MP_TOBRUK          Map = "tobruk"
	MP_INVALID         Map = "invalid"
)

type GameMode string

const (
	GmWarfare   GameMode = "Warfare"
	GmOffensive GameMode = "Offensive"
	GmSkirmish  GameMode = "Skirmish"
)

var ObjectiveCount = map[GameMode]int{
	GmWarfare:   5,
	GmOffensive: 5,
	GmSkirmish:  1,
}

var OptionsPerObjective = map[GameMode]int{
	GmWarfare:   3,
	GmOffensive: 3,
	GmSkirmish:  1,
}

type Orientation string

const (
	OriHorizontal Orientation = "Horizontal"
	OriVertical   Orientation = "Vertical"
)

type Team string

const (
	TmAllies Team = "Allies"
	TmAxis   Team = "Axis"
	TmNone   Team = "None"

	NoTeamID = 0
)

func (t Team) ToInt() int {
	switch t {
	case TmAllies:
		return 1
	case TmAxis:
		return 2
	default:
		return NoTeamID
	}
}

func TeamFromString(name string) Team {
	typed := Team(name)
	switch typed {
	case TmAllies, TmAxis:
		return typed
	default:
		return TmNone
	}
}

func TeamFromInt(team int) Team {
	switch team {
	case 1:
		return TmAllies
	case 2:
		return TmAxis
	default:
		return TmNone
	}
}

type Environment string

const (
	EnvDay      Environment = "Day"
	EnvDusk     Environment = "Dusk"
	EnvDawn     Environment = "Dawn"
	EnvNight    Environment = "Night"
	EnvOvercast Environment = "Overcast"
	EnvRain     Environment = "Rain"
	EnvFoggy    Environment = "Foggy"
	EnvMorning  Environment = "Morning"
)

type Faction string

const (
	FctUS         Faction = "US"
	FctGER        Faction = "GER"
	FctRUS        Faction = "RUS"
	FctSOV        Faction = "SOV"
	FctGB         Faction = "GB"
	FctCW         Faction = "CW"
	FctDAK        Faction = "DAK"
	FctB8A        Faction = "B8A"
	FctUnassigned Faction = "NON"
)

var AllFactions = []Faction{FctUS, FctGER, FctRUS, FctSOV, FctGB, FctCW, FctDAK, FctB8A}

func (f Faction) Team() Team {
	if f == FctUnassigned {
		return TmNone
	}
	if f == FctGER || f == FctDAK {
		return TmAxis
	}
	return TmAllies
}

func FactionFromInt(id int) Faction {
	switch id {
	case 0:
		return FctGER
	case 1:
		return FctUS
	case 2:
		return FctRUS
	case 3:
		return FctGB
	case 4:
		return FctDAK
	case 5:
		return FctB8A
	default:
		return FctUnassigned
	}
}

type GameMap struct {
	ID               Map
	Name             string
	Tag              string
	PrettyName       string
	ShortName        string
	Allies           Faction
	Axis             Faction
	Orientation      Orientation // Whether the sectors are arranged horizontally (left-to-right) or vertically (top-to-bottom)
	MirroredFactions bool        // If the start side of the factions is mirrored or not. By default, Allies start left/top and Axis start right/bottom
}

type Layer struct {
	ID          string
	GameMap     GameMap
	GameMode    GameMode
	Attackers   Team
	Environment Environment
}

var mapMap = map[Map]GameMap{
	MP_CARENTAN:        {ID: MP_CARENTAN, Name: "CARENTAN", Tag: "CAR", PrettyName: "Carentan", ShortName: "Carentan", Allies: FctUS, Axis: FctGER, Orientation: OriHorizontal, MirroredFactions: false},
	MP_DRIEL:           {ID: MP_DRIEL, Name: "DRIEL", Tag: "DRL", PrettyName: "Driel", ShortName: "Driel", Allies: FctGB, Axis: FctGER, Orientation: OriVertical, MirroredFactions: true},
	MP_ELALAMEIN:       {ID: MP_ELALAMEIN, Name: "EL ALAMEIN", Tag: "ELA", PrettyName: "El Alamein", ShortName: "Alamein", Allies: FctGB, Axis: FctGER, Orientation: OriHorizontal, MirroredFactions: true},
	MP_ELSENBORNRIDGE:  {ID: MP_ELSENBORNRIDGE, Name: "ELSENBORN RIDGE", Tag: "EBR", PrettyName: "Elsenborn Ridge", ShortName: "Elsenborn", Allies: FctUS, Axis: FctGER, Orientation: OriVertical, MirroredFactions: false},
	MP_FOY:             {ID: MP_FOY, Name: "FOY", Tag: "FOY", PrettyName: "Foy", ShortName: "Foy", Allies: FctUS, Axis: FctGER, Orientation: OriVertical, MirroredFactions: true},
	MP_HILL400:         {ID: MP_HILL400, Name: "HILL 400", Tag: "HIL", PrettyName: "Hill 400", ShortName: "Hill 400", Allies: FctUS, Axis: FctGER, Orientation: OriHorizontal, MirroredFactions: false},
	MP_HURTGENFOREST:   {ID: MP_HURTGENFOREST, Name: "HÜRTGEN FOREST", Tag: "HUR", PrettyName: "Hurtgen Forest", ShortName: "Hurtgen", Allies: FctUS, Axis: FctGER, Orientation: OriHorizontal, MirroredFactions: false},
	MP_KHARKOV:         {ID: MP_KHARKOV, Name: "Kharkov", Tag: "KHA", PrettyName: "Kharkov", ShortName: "Kharkov", Allies: FctRUS, Axis: FctGER, Orientation: OriVertical, MirroredFactions: false},
	MP_KURSK:           {ID: MP_KURSK, Name: "KURSK", Tag: "KUR", PrettyName: "Kursk", ShortName: "Kursk", Allies: FctRUS, Axis: FctGER, Orientation: OriVertical, MirroredFactions: false},
	MP_MORTAIN:         {ID: MP_MORTAIN, Name: "MORTAIN", Tag: "MOR", PrettyName: "Mortain", ShortName: "Mortain", Allies: FctUS, Axis: FctGER, Orientation: OriHorizontal, MirroredFactions: false},
	MP_OMAHABEACH:      {ID: MP_OMAHABEACH, Name: "OMAHA BEACH", Tag: "OMA", PrettyName: "Omaha Beach", ShortName: "Omaha", Allies: FctUS, Axis: FctGER, Orientation: OriHorizontal, MirroredFactions: true},
	MP_PURPLEHEARTLANE: {ID: MP_PURPLEHEARTLANE, Name: "PURPLE HEART LANE", Tag: "PHL", PrettyName: "Purple Heart Lane", ShortName: "PHL", Allies: FctUS, Axis: FctGER, Orientation: OriVertical, MirroredFactions: false},
	MP_REMAGEN:         {ID: MP_REMAGEN, Name: "REMAGEN", Tag: "REM", PrettyName: "Remagen", ShortName: "Remagen", Allies: FctUS, Axis: FctGER, Orientation: OriVertical, MirroredFactions: true},
	MP_STMARIEDUMONT:   {ID: MP_STMARIEDUMONT, Name: "ST MARIE DU MONT", Tag: "BRC", PrettyName: "St. Marie Du Mont", ShortName: "SMDM", Allies: FctUS, Axis: FctGER, Orientation: OriVertical, MirroredFactions: false},
	MP_STMEREEGLISE:    {ID: MP_STMEREEGLISE, Name: "SAINTE-MÈRE-ÉGLISE", Tag: "SME", PrettyName: "St. Mere Eglise", ShortName: "SME", Allies: FctUS, Axis: FctGER, Orientation: OriHorizontal, MirroredFactions: true},
	MP_STALINGRAD:      {ID: MP_STALINGRAD, Name: "STALINGRAD", Tag: "STA", PrettyName: "Stalingrad", ShortName: "Stalingrad", Allies: FctRUS, Axis: FctGER, Orientation: OriHorizontal, MirroredFactions: true},
	MP_TOBRUK:          {ID: MP_TOBRUK, Name: "TOBRUK", Tag: "TBK", PrettyName: "Tobruk", ShortName: "Tobruk", Allies: FctGB, Axis: FctGER, Orientation: OriHorizontal, MirroredFactions: true},
	MP_UTAHBEACH:       {ID: MP_UTAHBEACH, Name: "UTAH BEACH", Tag: "UTA", PrettyName: "Utah Beach", ShortName: "Utah", Allies: FctUS, Axis: FctGER, Orientation: OriHorizontal, MirroredFactions: true},
}

var fallback_gamemap = GameMap{ID: MP_INVALID, Name: "INVALID", Tag: "INV", PrettyName: "Invalid", ShortName: "Invalid", Allies: FctUS, Axis: FctGER}

func MapToGameMap(mapName Map) GameMap {
	if gameMap, ok := mapMap[mapName]; ok {
		return gameMap
	}
	logger.Warn("Map not found:", mapName)
	return fallback_gamemap
}

func LogMapNameToMap(logMapName string) GameMap {
	for _, v := range mapMap {
		if strings.HasPrefix(logMapName, v.Name) {
			return v
		}
	}
	logger.Warn("LogMapName not found:", logMapName)
	return fallback_gamemap
}

func AllMaps() []GameMap {
	maps := []GameMap{}
	for _, m := range mapMap {
		maps = append(maps, m)
	}
	return maps
}

var layerMap = map[string]Layer{
	"CAR_S_1944_Day_P_Skirmish":           {ID: "CAR_S_1944_Day_P_Skirmish", GameMap: mapMap[MP_CARENTAN], GameMode: GmSkirmish, Environment: EnvDay},
	"CAR_S_1944_Dusk_P_Skirmish":          {ID: "CAR_S_1944_Dusk_P_Skirmish", GameMap: mapMap[MP_CARENTAN], GameMode: GmSkirmish, Environment: EnvDusk},
	"CAR_S_1944_Rain_P_Skirmish":          {ID: "CAR_S_1944_Rain_P_Skirmish", GameMap: mapMap[MP_CARENTAN], GameMode: GmSkirmish, Environment: EnvRain},
	"carentan_offensive_ger":              {ID: "carentan_offensive_ger", GameMap: mapMap[MP_CARENTAN], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"carentan_offensive_us":               {ID: "carentan_offensive_us", GameMap: mapMap[MP_CARENTAN], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"carentan_warfare":                    {ID: "carentan_warfare", GameMap: mapMap[MP_CARENTAN], GameMode: GmWarfare, Environment: EnvDay},
	"carentan_warfare_night":              {ID: "carentan_warfare_night", GameMap: mapMap[MP_CARENTAN], GameMode: GmWarfare, Environment: EnvNight},
	"driel_offensive_ger":                 {ID: "driel_offensive_ger", GameMap: mapMap[MP_DRIEL], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"driel_offensive_us":                  {ID: "driel_offensive_us", GameMap: mapMap[MP_DRIEL], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"driel_warfare":                       {ID: "driel_warfare", GameMap: mapMap[MP_DRIEL], GameMode: GmWarfare, Environment: EnvDawn},
	"driel_warfare_night":                 {ID: "driel_warfare_night", GameMap: mapMap[MP_DRIEL], GameMode: GmWarfare, Environment: EnvNight},
	"DRL_S_1944_P_Skirmish":               {ID: "DRL_S_1944_P_Skirmish", GameMap: mapMap[MP_DRIEL], GameMode: GmSkirmish, Environment: EnvDawn},
	"DRL_S_1944_Day_P_Skirmish":           {ID: "DRL_S_1944_Day_P_Skirmish", GameMap: mapMap[MP_DRIEL], GameMode: GmSkirmish, Environment: EnvDay},
	"DRL_S_1944_Night_P_Skirmish":         {ID: "DRL_S_1944_Night_P_Skirmish", GameMap: mapMap[MP_DRIEL], GameMode: GmSkirmish, Environment: EnvNight},
	"ELA_S_1942_P_Skirmish":               {ID: "ELA_S_1942_P_Skirmish", GameMap: mapMap[MP_ELALAMEIN], GameMode: GmSkirmish, Environment: EnvDay},
	"ELA_S_1942_Night_P_Skirmish":         {ID: "ELA_S_1942_Night_P_Skirmish", GameMap: mapMap[MP_ELALAMEIN], GameMode: GmSkirmish, Environment: EnvDusk},
	"elalamein_offensive_CW":              {ID: "elalamein_offensive_CW", GameMap: mapMap[MP_ELALAMEIN], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"elalamein_offensive_ger":             {ID: "elalamein_offensive_ger", GameMap: mapMap[MP_ELALAMEIN], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"elalamein_warfare":                   {ID: "elalamein_warfare", GameMap: mapMap[MP_ELALAMEIN], GameMode: GmWarfare, Environment: EnvDay},
	"elalamein_warfare_night":             {ID: "elalamein_warfare_night", GameMap: mapMap[MP_ELALAMEIN], GameMode: GmWarfare, Environment: EnvDusk},
	"elsenbornridge_offensiveger_day":     {ID: "elsenbornridge_offensiveger_day", GameMap: mapMap[MP_ELSENBORNRIDGE], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"elsenbornridge_offensiveger_morning": {ID: "elsenbornridge_offensiveger_morning", GameMap: mapMap[MP_ELSENBORNRIDGE], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDawn},
	"elsenbornridge_offensiveger_night":   {ID: "elsenbornridge_offensiveger_night", GameMap: mapMap[MP_ELSENBORNRIDGE], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvNight},
	"elsenbornridge_offensiveUS_day":      {ID: "elsenbornridge_offensiveUS_day", GameMap: mapMap[MP_ELSENBORNRIDGE], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"elsenbornridge_offensiveUS_morning":  {ID: "elsenbornridge_offensiveUS_morning", GameMap: mapMap[MP_ELSENBORNRIDGE], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDawn},
	"elsenbornridge_offensiveUS_night":    {ID: "elsenbornridge_offensiveUS_night", GameMap: mapMap[MP_ELSENBORNRIDGE], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvNight},
	"elsenbornridge_skirmish_day":         {ID: "elsenbornridge_skirmish_day", GameMap: mapMap[MP_ELSENBORNRIDGE], GameMode: GmSkirmish, Environment: EnvDay},
	"elsenbornridge_skirmish_morning":     {ID: "elsenbornridge_skirmish_morning", GameMap: mapMap[MP_ELSENBORNRIDGE], GameMode: GmSkirmish, Environment: EnvDawn},
	"elsenbornridge_skirmish_night":       {ID: "elsenbornridge_skirmish_night", GameMap: mapMap[MP_ELSENBORNRIDGE], GameMode: GmSkirmish, Environment: EnvNight},
	"elsenbornridge_warfare_day":          {ID: "elsenbornridge_warfare_day", GameMap: mapMap[MP_ELSENBORNRIDGE], GameMode: GmWarfare, Environment: EnvDay},
	"elsenbornridge_warfare_morning":      {ID: "elsenbornridge_warfare_morning", GameMap: mapMap[MP_ELSENBORNRIDGE], GameMode: GmWarfare, Environment: EnvDawn},
	"elsenbornridge_warfare_night":        {ID: "elsenbornridge_warfare_night", GameMap: mapMap[MP_ELSENBORNRIDGE], GameMode: GmWarfare, Environment: EnvNight},
	"foy_offensive_ger":                   {ID: "foy_offensive_ger", GameMap: mapMap[MP_FOY], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"foy_offensive_us":                    {ID: "foy_offensive_us", GameMap: mapMap[MP_FOY], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"foy_warfare":                         {ID: "foy_warfare", GameMap: mapMap[MP_FOY], GameMode: GmWarfare, Environment: EnvDay},
	"foy_warfare_night":                   {ID: "foy_warfare_night", GameMap: mapMap[MP_FOY], GameMode: GmWarfare, Environment: EnvNight},
	"HIL_S_1944_Day_P_Skirmish":           {ID: "HIL_S_1944_Day_P_Skirmish", GameMap: mapMap[MP_HILL400], GameMode: GmSkirmish, Environment: EnvDay},
	"HIL_S_1944_Dusk_P_Skirmish":          {ID: "HIL_S_1944_Dusk_P_Skirmish", GameMap: mapMap[MP_HILL400], GameMode: GmSkirmish, Environment: EnvNight},
	"hill400_offensive_ger":               {ID: "hill400_offensive_ger", GameMap: mapMap[MP_HILL400], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvFoggy},
	"hill400_offensive_US":                {ID: "hill400_offensive_US", GameMap: mapMap[MP_HILL400], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"hill400_warfare":                     {ID: "hill400_warfare", GameMap: mapMap[MP_HILL400], GameMode: GmWarfare, Environment: EnvDay},
	"hurtgenforest_offensive_ger":         {ID: "hurtgenforest_offensive_ger", GameMap: mapMap[MP_HURTGENFOREST], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvFoggy},
	"hurtgenforest_offensive_US":          {ID: "hurtgenforest_offensive_US", GameMap: mapMap[MP_HURTGENFOREST], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"hurtgenforest_warfare_V2":            {ID: "hurtgenforest_warfare_V2", GameMap: mapMap[MP_HURTGENFOREST], GameMode: GmWarfare, Environment: EnvDay},
	"hurtgenforest_warfare_V2_night":      {ID: "hurtgenforest_warfare_V2_night", GameMap: mapMap[MP_HURTGENFOREST], GameMode: GmWarfare, Environment: EnvNight},
	"kharkov_offensive_ger":               {ID: "kharkov_offensive_ger", GameMap: mapMap[MP_KHARKOV], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"kharkov_offensive_rus":               {ID: "kharkov_offensive_rus", GameMap: mapMap[MP_KHARKOV], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"kharkov_warfare":                     {ID: "kharkov_warfare", GameMap: mapMap[MP_KHARKOV], GameMode: GmWarfare, Environment: EnvDay},
	"kharkov_warfare_night":               {ID: "kharkov_warfare_night", GameMap: mapMap[MP_KHARKOV], GameMode: GmWarfare, Environment: EnvNight},
	"kursk_offensive_ger":                 {ID: "kursk_offensive_ger", GameMap: mapMap[MP_KURSK], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"kursk_offensive_rus":                 {ID: "kursk_offensive_rus", GameMap: mapMap[MP_KURSK], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"kursk_warfare":                       {ID: "kursk_warfare", GameMap: mapMap[MP_KURSK], GameMode: GmWarfare, Environment: EnvDay},
	"kursk_warfare_night":                 {ID: "kursk_warfare_night", GameMap: mapMap[MP_KURSK], GameMode: GmWarfare, Environment: EnvNight},
	"mortain_offensiveger_day":            {ID: "mortain_offensiveger_day", GameMap: mapMap[MP_MORTAIN], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"mortain_offensiveger_dusk":           {ID: "mortain_offensiveger_dusk", GameMap: mapMap[MP_MORTAIN], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDusk},
	"mortain_offensiveger_overcast":       {ID: "mortain_offensiveger_overcast", GameMap: mapMap[MP_MORTAIN], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvOvercast},
	"mortain_offensiveUS_day":             {ID: "mortain_offensiveUS_day", GameMap: mapMap[MP_MORTAIN], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"mortain_offensiveUS_dusk":            {ID: "mortain_offensiveUS_dusk", GameMap: mapMap[MP_MORTAIN], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDusk},
	"mortain_offensiveUS_overcast":        {ID: "mortain_offensiveUS_overcast", GameMap: mapMap[MP_MORTAIN], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvOvercast},
	"mortain_warfare_day":                 {ID: "mortain_warfare_day", GameMap: mapMap[MP_MORTAIN], GameMode: GmWarfare, Environment: EnvDay},
	"mortain_warfare_dusk":                {ID: "mortain_warfare_dusk", GameMap: mapMap[MP_MORTAIN], GameMode: GmWarfare, Environment: EnvDusk},
	"mortain_warfare_overcast":            {ID: "mortain_warfare_overcast", GameMap: mapMap[MP_MORTAIN], GameMode: GmOffensive, Environment: EnvOvercast},
	"mortain_skirmish_day":                {ID: "mortain_skirmish_day", GameMap: mapMap[MP_MORTAIN], GameMode: GmSkirmish, Environment: EnvDay},
	"mortain_skirmish_dusk":               {ID: "mortain_skirmish_dusk", GameMap: mapMap[MP_MORTAIN], GameMode: GmSkirmish, Environment: EnvDusk},
	"mortain_skirmish_overcast":           {ID: "mortain_skirmish_overcast", GameMap: mapMap[MP_MORTAIN], GameMode: GmSkirmish, Environment: EnvOvercast},
	"omahabeach_offensive_ger":            {ID: "omahabeach_offensive_ger", GameMap: mapMap[MP_OMAHABEACH], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"omahabeach_offensive_us":             {ID: "omahabeach_offensive_us", GameMap: mapMap[MP_OMAHABEACH], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"omahabeach_warfare":                  {ID: "omahabeach_warfare", GameMap: mapMap[MP_OMAHABEACH], GameMode: GmWarfare, Environment: EnvDay},
	"omahabeach_warfare_night":            {ID: "omahabeach_warfare_night", GameMap: mapMap[MP_OMAHABEACH], GameMode: GmWarfare, Environment: EnvDusk},
	"PHL_L_1944_OffensiveGER":             {ID: "PHL_L_1944_OffensiveGER", GameMap: mapMap[MP_PURPLEHEARTLANE], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"PHL_L_1944_OffensiveUS":              {ID: "PHL_L_1944_OffensiveUS", GameMap: mapMap[MP_PURPLEHEARTLANE], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"PHL_L_1944_Warfare":                  {ID: "PHL_L_1944_Warfare", GameMap: mapMap[MP_PURPLEHEARTLANE], GameMode: GmWarfare, Environment: EnvRain},
	"PHL_L_1944_Warfare_Night":            {ID: "PHL_L_1944_Warfare_Night", GameMap: mapMap[MP_PURPLEHEARTLANE], GameMode: GmWarfare, Environment: EnvNight},
	"PHL_S_1944_Morning_P_Skirmish":       {ID: "PHL_S_1944_Morning_P_Skirmish", GameMap: mapMap[MP_PURPLEHEARTLANE], GameMode: GmSkirmish, Environment: EnvMorning},
	"PHL_S_1944_Night_P_Skirmish":         {ID: "PHL_S_1944_Night_P_Skirmish", GameMap: mapMap[MP_PURPLEHEARTLANE], GameMode: GmSkirmish, Environment: EnvNight},
	"PHL_S_1944_Rain_P_Skirmish":          {ID: "PHL_S_1944_Rain_P_Skirmish", GameMap: mapMap[MP_PURPLEHEARTLANE], GameMode: GmSkirmish, Environment: EnvRain},
	"remagen_offensive_ger":               {ID: "remagen_offensive_ger", GameMap: mapMap[MP_REMAGEN], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvFoggy},
	"remagen_offensive_us":                {ID: "remagen_offensive_us", GameMap: mapMap[MP_REMAGEN], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"remagen_warfare":                     {ID: "remagen_warfare", GameMap: mapMap[MP_REMAGEN], GameMode: GmWarfare, Environment: EnvDay},
	"remagen_warfare_night":               {ID: "remagen_warfare_night", GameMap: mapMap[MP_REMAGEN], GameMode: GmWarfare, Environment: EnvNight},
	"SMDM_S_1944_Day_P_Skirmish":          {ID: "SMDM_S_1944_Day_P_Skirmish", GameMap: mapMap[MP_STMARIEDUMONT], GameMode: GmSkirmish, Environment: EnvDay},
	"SMDM_S_1944_Rain_P_Skirmish":         {ID: "SMDM_S_1944_Rain_P_Skirmish", GameMap: mapMap[MP_STMARIEDUMONT], GameMode: GmSkirmish, Environment: EnvRain},
	"SMDM_S_1944_Night_P_Skirmish":        {ID: "SMDM_S_1944_Night_P_Skirmish", GameMap: mapMap[MP_STMARIEDUMONT], GameMode: GmSkirmish, Environment: EnvNight},
	"SME_S_1944_Day_P_Skirmish":           {ID: "SME_S_1944_Day_P_Skirmish", GameMap: mapMap[MP_STMEREEGLISE], GameMode: GmSkirmish, Environment: EnvDay},
	"SME_S_1944_Morning_P_Skirmish":       {ID: "SME_S_1944_Morning_P_Skirmish", GameMap: mapMap[MP_STMEREEGLISE], GameMode: GmSkirmish, Environment: EnvMorning},
	"SME_S_1944_Night_P_Skirmish":         {ID: "SME_S_1944_Night_P_Skirmish", GameMap: mapMap[MP_STMEREEGLISE], GameMode: GmSkirmish, Environment: EnvNight},
	"STA_L_1942_Warfare":                  {ID: "STA_L_1942_Warfare", GameMap: mapMap[MP_STALINGRAD], GameMode: GmWarfare, Environment: EnvDay},
	"STA_L_1942_Warfare_Night":            {ID: "STA_L_1942_Warfare_Night", GameMap: mapMap[MP_STALINGRAD], GameMode: GmWarfare, Environment: EnvNight},
	"STA_L_1942_OffensiveGER":             {ID: "STA_L_1942_OffensiveGER", GameMap: mapMap[MP_STALINGRAD], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"STA_L_1942_OffensiveRUS":             {ID: "STA_L_1942_OffensiveRUS", GameMap: mapMap[MP_STALINGRAD], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"STA_S_1942_P_Skirmish_Dusk":          {ID: "STA_S_1942_P_Skirmish_Dusk", GameMap: mapMap[MP_STALINGRAD], GameMode: GmSkirmish, Environment: EnvDusk},
	"STA_S_1942_P_Skirmish_Overcast":      {ID: "STA_S_1942_P_Skirmish_Overcast", GameMap: mapMap[MP_STALINGRAD], GameMode: GmSkirmish, Environment: EnvOvercast},
	"stmariedumont_off_ger":               {ID: "stmariedumont_off_ger", GameMap: mapMap[MP_STMARIEDUMONT], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"stmariedumont_off_us":                {ID: "stmariedumont_off_us", GameMap: mapMap[MP_STMARIEDUMONT], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"stmariedumont_warfare":               {ID: "stmariedumont_warfare", GameMap: mapMap[MP_STMARIEDUMONT], GameMode: GmWarfare, Environment: EnvDay},
	"stmariedumont_warfare_night":         {ID: "stmariedumont_warfare_night", GameMap: mapMap[MP_STMARIEDUMONT], GameMode: GmWarfare, Environment: EnvNight},
	"stmereeglise_offensive_ger":          {ID: "stmereeglise_offensive_ger", GameMap: mapMap[MP_STMEREEGLISE], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDawn},
	"stmereeglise_offensive_us":           {ID: "stmereeglise_offensive_us", GameMap: mapMap[MP_STMEREEGLISE], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"stmereeglise_warfare":                {ID: "stmereeglise_warfare", GameMap: mapMap[MP_STMEREEGLISE], GameMode: GmWarfare, Environment: EnvDay},
	"stmereeglise_warfare_night":          {ID: "stmereeglise_warfare_night", GameMap: mapMap[MP_STMEREEGLISE], Environment: EnvNight},
	"tobruk_offensivebritish_day":         {ID: "tobruk_offensivebritish_day", GameMap: mapMap[MP_TOBRUK], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"tobruk_offensivebritish_dusk":        {ID: "tobruk_offensivebritish_dusk", GameMap: mapMap[MP_TOBRUK], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDusk},
	"tobruk_offensivebritish_morning":     {ID: "tobruk_offensivebritish_morning", GameMap: mapMap[MP_TOBRUK], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDawn},
	"tobruk_offensiveger_day":             {ID: "tobruk_offensiveger_day", GameMap: mapMap[MP_TOBRUK], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"tobruk_offensiveger_dusk":            {ID: "tobruk_offensiveger_dusk", GameMap: mapMap[MP_TOBRUK], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDusk},
	"tobruk_offensiveger_morning":         {ID: "tobruk_offensiveger_morning", GameMap: mapMap[MP_TOBRUK], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDawn},
	"tobruk_skirmish_day":                 {ID: "tobruk_skirmish_day", GameMap: mapMap[MP_TOBRUK], GameMode: GmSkirmish, Environment: EnvDay},
	"tobruk_skirmish_dusk":                {ID: "tobruk_skirmish_dusk", GameMap: mapMap[MP_TOBRUK], GameMode: GmSkirmish, Environment: EnvDusk},
	"tobruk_skirmish_morning":             {ID: "tobruk_skirmish_morning", GameMap: mapMap[MP_TOBRUK], GameMode: GmSkirmish, Environment: EnvDawn},
	"tobruk_warfare_day":                  {ID: "tobruk_warfare_day", GameMap: mapMap[MP_TOBRUK], GameMode: GmWarfare, Environment: EnvDay},
	"tobruk_warfare_dusk":                 {ID: "tobruk_warfare_dusk", GameMap: mapMap[MP_TOBRUK], GameMode: GmWarfare, Environment: EnvDusk},
	"tobruk_warfare_morning":              {ID: "tobruk_warfare_morning", GameMap: mapMap[MP_TOBRUK], GameMode: GmWarfare, Environment: EnvDawn},
	"utahbeach_offensive_ger":             {ID: "utahbeach_offensive_ger", GameMap: mapMap[MP_UTAHBEACH], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"utahbeach_offensive_us":              {ID: "utahbeach_offensive_us", GameMap: mapMap[MP_UTAHBEACH], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"utahbeach_warfare":                   {ID: "utahbeach_warfare", GameMap: mapMap[MP_UTAHBEACH], GameMode: GmWarfare, Environment: EnvDay},
	"utahbeach_warfare_night":             {ID: "utahbeach_warfare_night", GameMap: mapMap[MP_UTAHBEACH], GameMode: GmWarfare, Environment: EnvNight},
}

var fallback_layer = Layer{ID: "invalid", GameMap: fallback_gamemap, GameMode: GmWarfare, Environment: EnvDay}
var previousLayer = fallback_layer

const (
	restartSuffix = "_RESTART"
	untitledMap   = "Untitled"
	loadingMap    = "Loading"
)

func ParseLayer(layerName string) Layer {
	layerName, _ = strings.CutSuffix(layerName, restartSuffix)

	if strings.HasPrefix(layerName, loadingMap) || strings.HasPrefix(layerName, untitledMap) {
		return previousLayer
	}

	if lay, ok := layerMap[layerName]; ok {
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

func LayersByMap(gmap GameMap) []Layer {
	layers := []Layer{}
	for _, l := range layerMap {
		if l.GameMap == gmap {
			layers = append(layers, l)
		}
	}
	return layers
}

func LayersByMode(gmode GameMode) []Layer {
	layers := []Layer{}
	for _, l := range layerMap {
		if l.GameMode == gmode {
			layers = append(layers, l)
		}
	}
	return layers
}
