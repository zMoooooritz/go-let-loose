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
	MP_INVALID         Map = "invalid"
)

type Gamemode string

const (
	GmWarfare   Gamemode = "Warfare"
	GmOffensive Gamemode = "Offensive"
	GmSkirmish  Gamemode = "Skirmish"
)

type Team string

const (
	TmAllies Team = "Allies"
	TmAxis   Team = "Axis"
	TmNone   Team = "None"
)

type Environment string

const (
	EnvDay      Environment = "Day"
	EnvDusk     Environment = "Dusk"
	EnvDawn     Environment = "Dawn"
	EnvNight    Environment = "Night"
	EnvOvercast Environment = "Overcast"
	EnvRain     Environment = "Rain"
)

type Faction string

const (
	FctUS  Faction = "US"
	FctGER Faction = "GER"
	FctRUS Faction = "RUS"
	FctGB  Faction = "GB"
)

type GameMap struct {
	ID         Map
	Name       string
	Tag        string
	PrettyName string
	ShortName  string
	Allies     Faction
	Axis       Faction
}

type Layer struct {
	ID          string
	GameMap     GameMap
	GameMode    Gamemode
	Attackers   Team
	Environment Environment
}

var maps = map[Map]GameMap{
	MP_STMEREEGLISE:    {ID: MP_STMEREEGLISE, Name: "SAINTE-MÈRE-ÉGLISE", Tag: "SME", PrettyName: "St. Mere Eglise", ShortName: "SME", Allies: FctUS, Axis: FctGER},
	MP_STMARIEDUMONT:   {ID: MP_STMARIEDUMONT, Name: "ST MARIE DU MONT", Tag: "BRC", PrettyName: "St. Marie Du Mont", ShortName: "SMDM", Allies: FctUS, Axis: FctGER},
	MP_UTAHBEACH:       {ID: MP_UTAHBEACH, Name: "UTAH BEACH", Tag: "UTA", PrettyName: "Utah Beach", ShortName: "Utah", Allies: FctUS, Axis: FctGER},
	MP_OMAHABEACH:      {ID: MP_OMAHABEACH, Name: "OMAHA BEACH", Tag: "OMA", PrettyName: "Omaha Beach", ShortName: "Omaha", Allies: FctUS, Axis: FctGER},
	MP_PURPLEHEARTLANE: {ID: MP_PURPLEHEARTLANE, Name: "PURPLE HEART LANE", Tag: "PHL", PrettyName: "Purple Heart Lane", ShortName: "PHL", Allies: FctUS, Axis: FctGER},
	MP_CARENTAN:        {ID: MP_CARENTAN, Name: "CARENTAN", Tag: "CAR", PrettyName: "Carentan", ShortName: "Carentan", Allies: FctUS, Axis: FctGER},
	MP_HURTGENFOREST:   {ID: MP_HURTGENFOREST, Name: "HÜRTGEN FOREST", Tag: "HUR", PrettyName: "Hurtgen Forest", ShortName: "Hurtgen", Allies: FctUS, Axis: FctGER},
	MP_HILL400:         {ID: MP_HILL400, Name: "HILL 400", Tag: "HIL", PrettyName: "Hill 400", ShortName: "Hill 400", Allies: FctUS, Axis: FctGER},
	MP_FOY:             {ID: MP_FOY, Name: "FOY", Tag: "FOY", PrettyName: "Foy", ShortName: "Foy", Allies: FctUS, Axis: FctGER},
	MP_KURSK:           {ID: MP_KURSK, Name: "KURSK", Tag: "KUR", PrettyName: "Kursk", ShortName: "Kursk", Allies: FctRUS, Axis: FctGER},
	MP_STALINGRAD:      {ID: MP_STALINGRAD, Name: "STALINGRAD", Tag: "STA", PrettyName: "Stalingrad", ShortName: "Stalingrad", Allies: FctRUS, Axis: FctGER},
	MP_REMAGEN:         {ID: MP_REMAGEN, Name: "REMAGEN", Tag: "REM", PrettyName: "Remagen", ShortName: "Remagen", Allies: FctUS, Axis: FctGER},
	MP_KHARKOV:         {ID: MP_KHARKOV, Name: "Kharkov", Tag: "KHA", PrettyName: "Kharkov", ShortName: "Kharkov", Allies: FctRUS, Axis: FctGER},
	MP_DRIEL:           {ID: MP_DRIEL, Name: "DRIEL", Tag: "DRL", PrettyName: "Driel", ShortName: "Driel", Allies: FctGB, Axis: FctGER},
	MP_ELALAMEIN:       {ID: MP_ELALAMEIN, Name: "EL ALAMEIN", Tag: "ELA", PrettyName: "El Alamein", ShortName: "Alamein", Allies: FctGB, Axis: FctGER},
	MP_MORTAIN:         {ID: MP_MORTAIN, Name: "MORTAIN", Tag: "MTN", PrettyName: "Mortain", ShortName: "Mortain", Allies: FctUS, Axis: FctGER},
}

var fallback_gamemap = GameMap{ID: MP_INVALID, Name: "INVALID", Tag: "INV", PrettyName: "Invalid", ShortName: "Invalid", Allies: FctUS, Axis: FctGER}

func MapToGameMap(mapName Map) GameMap {
	if gameMap, ok := maps[mapName]; ok {
		return gameMap
	}
	logger.Warn("Map not found:", mapName)
	return fallback_gamemap
}

func LogMapNameToMap(logMapName string) GameMap {
	for _, v := range maps {
		if strings.HasPrefix(logMapName, v.Name) {
			return v
		}
	}
	logger.Warn("LogMapName not found:", logMapName)
	return fallback_gamemap
}

var layers = map[string]Layer{
	"stmereeglise_warfare":           {ID: "stmereeglise_warfare", GameMap: maps[MP_STMEREEGLISE], GameMode: GmWarfare, Environment: EnvDay},
	"stmereeglise_warfare_night":     {ID: "stmereeglise_warfare_night", GameMap: maps[MP_STMEREEGLISE], Environment: EnvNight},
	"stmereeglise_offensive_us":      {ID: "stmereeglise_offensive_us", GameMap: maps[MP_STMEREEGLISE], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"stmereeglise_offensive_ger":     {ID: "stmereeglise_offensive_ger", GameMap: maps[MP_STMEREEGLISE], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"SME_S_1944_Day_P_Skirmish":      {ID: "SME_S_1944_Day_P_Skirmish", GameMap: maps[MP_STMEREEGLISE], GameMode: GmSkirmish, Environment: EnvDay},
	"SME_S_1944_Morning_P_Skirmish":  {ID: "SME_S_1944_Morning_P_Skirmish", GameMap: maps[MP_STMEREEGLISE], GameMode: GmSkirmish, Environment: EnvDawn},
	"SME_S_1944_Night_P_Skirmish":    {ID: "SME_S_1944_Night_P_Skirmish", GameMap: maps[MP_STMEREEGLISE], GameMode: GmSkirmish, Environment: EnvNight},
	"stmariedumont_warfare":          {ID: "stmariedumont_warfare", GameMap: maps[MP_STMARIEDUMONT], GameMode: GmWarfare, Environment: EnvDay},
	"stmariedumont_warfare_night":    {ID: "stmariedumont_warfare_night", GameMap: maps[MP_STMARIEDUMONT], GameMode: GmWarfare, Environment: EnvNight},
	"stmariedumont_off_us":           {ID: "stmariedumont_off_us", GameMap: maps[MP_STMARIEDUMONT], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"stmariedumont_off_ger":          {ID: "stmariedumont_off_ger", GameMap: maps[MP_STMARIEDUMONT], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"SMDM_S_1944_Day_P_Skirmish":     {ID: "SMDM_S_1944_Day_P_Skirmish", GameMap: maps[MP_STMARIEDUMONT], GameMode: GmSkirmish, Environment: EnvDay},
	"SMDM_S_1944_Night_P_Skirmish":   {ID: "SMDM_S_1944_Night_P_Skirmish", GameMap: maps[MP_STMARIEDUMONT], GameMode: GmSkirmish, Environment: EnvNight},
	"SMDM_S_1944_Rain_P_Skirmish":    {ID: "SMDM_S_1944_Rain_P_Skirmish", GameMap: maps[MP_STMARIEDUMONT], GameMode: GmSkirmish, Environment: EnvRain},
	"utahbeach_warfare":              {ID: "utahbeach_warfare", GameMap: maps[MP_UTAHBEACH], GameMode: GmWarfare, Environment: EnvDay},
	"utahbeach_warfare_night":        {ID: "utahbeach_warfare_night", GameMap: maps[MP_UTAHBEACH], GameMode: GmWarfare, Environment: EnvNight},
	"utahbeach_offensive_us":         {ID: "utahbeach_offensive_us", GameMap: maps[MP_UTAHBEACH], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"utahbeach_offensive_ger":        {ID: "utahbeach_offensive_ger", GameMap: maps[MP_UTAHBEACH], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"omahabeach_warfare":             {ID: "omahabeach_warfare", GameMap: maps[MP_OMAHABEACH], GameMode: GmWarfare, Environment: EnvDay},
	"omahabeach_warfare_night":       {ID: "omahabeach_warfare_night", GameMap: maps[MP_OMAHABEACH], GameMode: GmWarfare, Environment: EnvNight},
	"omahabeach_offensive_us":        {ID: "omahabeach_offensive_us", GameMap: maps[MP_OMAHABEACH], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"omahabeach_offensive_ger":       {ID: "omahabeach_offensive_ger", GameMap: maps[MP_OMAHABEACH], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"purpleheartlane_warfare":        {ID: "purpleheartlane_warfare", GameMap: maps[MP_PURPLEHEARTLANE], GameMode: GmWarfare, Environment: EnvDay},
	"purpleheartlane_warfare_night":  {ID: "purpleheartlane_warfare_night", GameMap: maps[MP_PURPLEHEARTLANE], GameMode: GmWarfare, Environment: EnvNight},
	"purpleheartlane_offensive_us":   {ID: "purpleheartlane_offensive_us", GameMap: maps[MP_PURPLEHEARTLANE], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"purpleheartlane_offensive_ger":  {ID: "purpleheartlane_offensive_ger", GameMap: maps[MP_PURPLEHEARTLANE], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"carentan_warfare":               {ID: "carentan_warfare", GameMap: maps[MP_CARENTAN], GameMode: GmWarfare, Environment: EnvDay},
	"carentan_warfare_night":         {ID: "carentan_warfare_night", GameMap: maps[MP_CARENTAN], GameMode: GmWarfare, Environment: EnvNight},
	"carentan_offensive_us":          {ID: "carentan_offensive_us", GameMap: maps[MP_CARENTAN], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"carentan_offensive_ger":         {ID: "carentan_offensive_ger", GameMap: maps[MP_CARENTAN], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"CAR_S_1944_Day_P_Skirmish":      {ID: "CAR_S_1944_Day_P_Skirmish", GameMap: maps[MP_CARENTAN], GameMode: GmSkirmish, Environment: EnvDay},
	"CAR_S_1944_Rain_P_Skirmish":     {ID: "CAR_S_1944_Rain_P_Skirmish", GameMap: maps[MP_CARENTAN], GameMode: GmSkirmish, Environment: EnvRain},
	"CAR_S_1944_Dusk_P_Skirmish":     {ID: "CAR_S_1944_Dusk_P_Skirmish", GameMap: maps[MP_CARENTAN], GameMode: GmSkirmish, Environment: EnvDusk},
	"hurtgenforest_warfare_V2":       {ID: "hurtgenforest_warfare_V2", GameMap: maps[MP_HURTGENFOREST], GameMode: GmWarfare, Environment: EnvDay},
	"hurtgenforest_warfare_V2_night": {ID: "hurtgenforest_warfare_V2_night", GameMap: maps[MP_HURTGENFOREST], GameMode: GmWarfare, Environment: EnvNight},
	"hurtgenforest_offensive_US":     {ID: "hurtgenforest_offensive_US", GameMap: maps[MP_HURTGENFOREST], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"hurtgenforest_offensive_ger":    {ID: "hurtgenforest_offensive_ger", GameMap: maps[MP_HURTGENFOREST], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"hill400_warfare":                {ID: "hill400_warfare", GameMap: maps[MP_HILL400], GameMode: GmWarfare, Environment: EnvDay},
	"hill400_warfare_night":          {ID: "hill400_warfare_night", GameMap: maps[MP_HILL400], GameMode: GmWarfare, Environment: EnvNight},
	"hill400_offensive_US":           {ID: "hill400_offensive_US", GameMap: maps[MP_HILL400], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"hill400_offensive_ger":          {ID: "hill400_offensive_ger", GameMap: maps[MP_HILL400], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"foy_warfare":                    {ID: "foy_warfare", GameMap: maps[MP_FOY], GameMode: GmWarfare, Environment: EnvDay},
	"foy_warfare_night":              {ID: "foy_warfare_night", GameMap: maps[MP_FOY], GameMode: GmWarfare, Environment: EnvNight},
	"foy_offensive_us":               {ID: "foy_offensive_us", GameMap: maps[MP_FOY], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"foy_offensive_ger":              {ID: "foy_offensive_ger", GameMap: maps[MP_FOY], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"kursk_warfare":                  {ID: "kursk_warfare", GameMap: maps[MP_KURSK], GameMode: GmWarfare, Environment: EnvDay},
	"kursk_warfare_night":            {ID: "kursk_warfare_night", GameMap: maps[MP_KURSK], GameMode: GmWarfare, Environment: EnvNight},
	"kursk_offensive_rus":            {ID: "kursk_offensive_rus", GameMap: maps[MP_KURSK], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"kursk_offensive_ger":            {ID: "kursk_offensive_ger", GameMap: maps[MP_KURSK], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"stalingrad_warfare":             {ID: "stalingrad_warfare", GameMap: maps[MP_STALINGRAD], GameMode: GmWarfare, Environment: EnvDay},
	"stalingrad_warfare_night":       {ID: "stalingrad_warfare_night", GameMap: maps[MP_STALINGRAD], GameMode: GmWarfare, Environment: EnvNight},
	"stalingrad_offensive_rus":       {ID: "stalingrad_offensive_rus", GameMap: maps[MP_STALINGRAD], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"stalingrad_offensive_ger":       {ID: "stalingrad_offensive_ger", GameMap: maps[MP_STALINGRAD], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"remagen_warfare":                {ID: "remagen_warfare", GameMap: maps[MP_REMAGEN], GameMode: GmWarfare, Environment: EnvDay},
	"remagen_warfare_night":          {ID: "remagen_warfare_night", GameMap: maps[MP_REMAGEN], GameMode: GmWarfare, Environment: EnvNight},
	"remagen_offensive_us":           {ID: "remagen_offensive_us", GameMap: maps[MP_REMAGEN], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"remagen_offensive_ger":          {ID: "remagen_offensive_ger", GameMap: maps[MP_REMAGEN], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"kharkov_warfare":                {ID: "kharkov_warfare", GameMap: maps[MP_KHARKOV], GameMode: GmWarfare, Environment: EnvDay},
	"kharkov_warfare_night":          {ID: "kharkov_warfare_night", GameMap: maps[MP_KHARKOV], GameMode: GmWarfare, Environment: EnvNight},
	"kharkov_offensive_rus":          {ID: "kharkov_offensive_rus", GameMap: maps[MP_KHARKOV], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"kharkov_offensive_ger":          {ID: "kharkov_offensive_ger", GameMap: maps[MP_KHARKOV], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"driel_warfare":                  {ID: "driel_warfare", GameMap: maps[MP_DRIEL], GameMode: GmWarfare, Environment: EnvDay},
	"driel_warfare_night":            {ID: "driel_warfare_night", GameMap: maps[MP_DRIEL], GameMode: GmWarfare, Environment: EnvNight},
	"driel_offensive_us":             {ID: "driel_offensive_us", GameMap: maps[MP_DRIEL], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"driel_offensive_ger":            {ID: "driel_offensive_ger", GameMap: maps[MP_DRIEL], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"DRL_S_1944_P_Skirmish":          {ID: "DRL_S_1944_P_Skirmish", GameMap: maps[MP_DRIEL], GameMode: GmSkirmish, Environment: EnvDawn},
	"DRL_S_1944_Night_P_Skirmish":    {ID: "DRL_S_1944_Night_P_Skirmish", GameMap: maps[MP_DRIEL], GameMode: GmSkirmish, Environment: EnvNight},
	"DRL_S_1944_Day_P_Skirmish":      {ID: "DRL_S_1944_Day_P_Skirmish", GameMap: maps[MP_DRIEL], GameMode: GmSkirmish, Environment: EnvDay},
	"elalamein_warfare":              {ID: "elalamein_warfare", GameMap: maps[MP_ELALAMEIN], GameMode: GmWarfare, Environment: EnvDay},
	"elalamein_warfare_night":        {ID: "elalamein_warfare_night", GameMap: maps[MP_ELALAMEIN], GameMode: GmWarfare, Environment: EnvDusk},
	"elalamein_offensive_CW":         {ID: "elalamein_offensive_CW", GameMap: maps[MP_ELALAMEIN], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"elalamein_offensive_ger":        {ID: "elalamein_offensive_ger", GameMap: maps[MP_ELALAMEIN], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"ELA_S_1942_P_Skirmish":          {ID: "ELA_S_1942_P_Skirmish", GameMap: maps[MP_ELALAMEIN], GameMode: GmSkirmish, Environment: EnvDay},
	"ELA_S_1942_Night_P_Skirmish":    {ID: "ELA_S_1942_Night_P_Skirmish", GameMap: maps[MP_ELALAMEIN], GameMode: GmSkirmish, Environment: EnvDusk},
	"mortain_warfare_day":            {ID: "mortain_warfare_day", GameMap: maps[MP_MORTAIN], GameMode: GmWarfare, Environment: EnvDay},
	"mortain_warfare_overcast":       {ID: "mortain_warfare_overcast", GameMap: maps[MP_MORTAIN], GameMode: GmOffensive, Environment: EnvOvercast},
	"mortain_warfare_evening":        {ID: "mortain_warfare_evening", GameMap: maps[MP_MORTAIN], GameMode: GmWarfare, Environment: EnvDawn},
	"mortain_offensiveUS_day":        {ID: "mortain_offensiveUS_day", GameMap: maps[MP_MORTAIN], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDay},
	"mortain_offensiveger_day":       {ID: "mortain_offensiveger_day", GameMap: maps[MP_MORTAIN], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDay},
	"mortain_offensiveUS_overcast":   {ID: "mortain_offensiveUS_overcast", GameMap: maps[MP_MORTAIN], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvOvercast},
	"mortain_offensiveger_overcast":  {ID: "mortain_offensiveger_overcast", GameMap: maps[MP_MORTAIN], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvOvercast},
	"mortain_offensiveUS_evening":    {ID: "mortain_offensiveUS_evening", GameMap: maps[MP_MORTAIN], GameMode: GmOffensive, Attackers: TmAllies, Environment: EnvDawn},
	"mortain_offensiveger_evening":   {ID: "mortain_offensiveger_evening", GameMap: maps[MP_MORTAIN], GameMode: GmOffensive, Attackers: TmAxis, Environment: EnvDawn},
	"mortain_skirmish_day":           {ID: "mortain_skirmish_day", GameMap: maps[MP_MORTAIN], GameMode: GmSkirmish, Environment: EnvDay},
	"mortain_skirmish_overcast":      {ID: "mortain_skirmish_overcast", GameMap: maps[MP_MORTAIN], GameMode: GmSkirmish, Environment: EnvOvercast},
	"mortain_skirmish_evening":       {ID: "mortain_skirmish_evening", GameMap: maps[MP_MORTAIN], GameMode: GmSkirmish, Environment: EnvDawn},
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

	if lay, ok := layers[layerName]; ok {
		previousLayer = lay
		return lay
	}
	logger.Warn("Layer not found:", layerName)
	return fallback_layer
}
