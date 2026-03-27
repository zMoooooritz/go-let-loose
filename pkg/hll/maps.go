package hll

import (
	"strings"

	"github.com/zMoooooritz/go-let-loose/pkg/logger"
)

type MapIdentifier string

const (
	MAP_STMEREEGLISE    MapIdentifier = "stmereeglise"
	MAP_STMARIEDUMONT   MapIdentifier = "stmariedumont"
	MAP_UTAHBEACH       MapIdentifier = "utahbeach"
	MAP_OMAHABEACH      MapIdentifier = "omahabeach"
	MAP_PURPLEHEARTLANE MapIdentifier = "purpleheartlane"
	MAP_CARENTAN        MapIdentifier = "carentan"
	MAP_HURTGENFOREST   MapIdentifier = "hurtgenforest"
	MAP_HILL400         MapIdentifier = "hill400"
	MAP_FOY             MapIdentifier = "foy"
	MAP_KURSK           MapIdentifier = "kursk"
	MAP_SMOLENSK        MapIdentifier = "smolensk"
	MAP_STALINGRAD      MapIdentifier = "stalingrad"
	MAP_REMAGEN         MapIdentifier = "remagen"
	MAP_KHARKOV         MapIdentifier = "kharkov"
	MAP_DRIEL           MapIdentifier = "driel"
	MAP_ELALAMEIN       MapIdentifier = "elalamein"
	MAP_MORTAIN         MapIdentifier = "mortain"
	MAP_ELSENBORNRIDGE  MapIdentifier = "elsenbornridge"
	MAP_TOBRUK          MapIdentifier = "tobruk"
	MAP_INVALID         MapIdentifier = "invalid"
)

type Orientation string

const (
	ORIENTATION_HORIZONTAL Orientation = "Horizontal"
	ORIENTATION_VERTICAL   Orientation = "Vertical"
)

type Map struct {
	ID               MapIdentifier
	Name             string
	Tag              string
	PrettyName       string
	ShortName        string
	Allies           FactionIdentifier
	Axis             FactionIdentifier
	Orientation      Orientation // Whether the sectors are arranged horizontally (left-to-right) or vertically (top-to-bottom)
	MirroredFactions bool        // If the start side of the factions is mirrored or not. By default, Allies start left/top and Axis start right/bottom
}

var mapMap = map[MapIdentifier]Map{
	MAP_CARENTAN:        {ID: MAP_CARENTAN, Name: "CARENTAN", Tag: "CAR", PrettyName: "Carentan", ShortName: "Carentan", Allies: FACTION_US, Axis: FACTION_GER, Orientation: ORIENTATION_HORIZONTAL, MirroredFactions: false},
	MAP_DRIEL:           {ID: MAP_DRIEL, Name: "DRIEL", Tag: "DRL", PrettyName: "Driel", ShortName: "Driel", Allies: FACTION_CW, Axis: FACTION_GER, Orientation: ORIENTATION_VERTICAL, MirroredFactions: true},
	MAP_ELALAMEIN:       {ID: MAP_ELALAMEIN, Name: "EL ALAMEIN", Tag: "ELA", PrettyName: "El Alamein", ShortName: "Alamein", Allies: FACTION_CW, Axis: FACTION_GER, Orientation: ORIENTATION_HORIZONTAL, MirroredFactions: true},
	MAP_ELSENBORNRIDGE:  {ID: MAP_ELSENBORNRIDGE, Name: "ELSENBORN RIDGE", Tag: "EBR", PrettyName: "Elsenborn Ridge", ShortName: "Elsenborn", Allies: FACTION_US, Axis: FACTION_GER, Orientation: ORIENTATION_VERTICAL, MirroredFactions: false},
	MAP_FOY:             {ID: MAP_FOY, Name: "FOY", Tag: "FOY", PrettyName: "Foy", ShortName: "Foy", Allies: FACTION_US, Axis: FACTION_GER, Orientation: ORIENTATION_VERTICAL, MirroredFactions: true},
	MAP_HILL400:         {ID: MAP_HILL400, Name: "HILL 400", Tag: "HIL", PrettyName: "Hill 400", ShortName: "Hill 400", Allies: FACTION_US, Axis: FACTION_GER, Orientation: ORIENTATION_HORIZONTAL, MirroredFactions: false},
	MAP_HURTGENFOREST:   {ID: MAP_HURTGENFOREST, Name: "HÜRTGEN FOREST", Tag: "HUR", PrettyName: "Hurtgen Forest", ShortName: "Hurtgen", Allies: FACTION_US, Axis: FACTION_GER, Orientation: ORIENTATION_HORIZONTAL, MirroredFactions: false},
	MAP_KHARKOV:         {ID: MAP_KHARKOV, Name: "Kharkov", Tag: "KHA", PrettyName: "Kharkov", ShortName: "Kharkov", Allies: FACTION_SOV, Axis: FACTION_GER, Orientation: ORIENTATION_VERTICAL, MirroredFactions: false},
	MAP_KURSK:           {ID: MAP_KURSK, Name: "KURSK", Tag: "KUR", PrettyName: "Kursk", ShortName: "Kursk", Allies: FACTION_SOV, Axis: FACTION_GER, Orientation: ORIENTATION_VERTICAL, MirroredFactions: false},
	MAP_MORTAIN:         {ID: MAP_MORTAIN, Name: "MORTAIN", Tag: "MOR", PrettyName: "Mortain", ShortName: "Mortain", Allies: FACTION_US, Axis: FACTION_GER, Orientation: ORIENTATION_HORIZONTAL, MirroredFactions: false},
	MAP_OMAHABEACH:      {ID: MAP_OMAHABEACH, Name: "OMAHA BEACH", Tag: "OMA", PrettyName: "Omaha Beach", ShortName: "Omaha", Allies: FACTION_US, Axis: FACTION_GER, Orientation: ORIENTATION_HORIZONTAL, MirroredFactions: true},
	MAP_PURPLEHEARTLANE: {ID: MAP_PURPLEHEARTLANE, Name: "PURPLE HEART LANE", Tag: "PHL", PrettyName: "Purple Heart Lane", ShortName: "PHL", Allies: FACTION_US, Axis: FACTION_GER, Orientation: ORIENTATION_VERTICAL, MirroredFactions: false},
	MAP_REMAGEN:         {ID: MAP_REMAGEN, Name: "REMAGEN", Tag: "REM", PrettyName: "Remagen", ShortName: "Remagen", Allies: FACTION_US, Axis: FACTION_GER, Orientation: ORIENTATION_VERTICAL, MirroredFactions: true},
	MAP_SMOLENSK:        {ID: MAP_SMOLENSK, Name: "SMOLENSK", Tag: "SMO", PrettyName: "Smolensk", ShortName: "Smolensk", Allies: FACTION_SOV, Axis: FACTION_GER, Orientation: ORIENTATION_HORIZONTAL, MirroredFactions: true},
	MAP_STMARIEDUMONT:   {ID: MAP_STMARIEDUMONT, Name: "ST MARIE DU MONT", Tag: "BRC", PrettyName: "St. Marie Du Mont", ShortName: "SMDM", Allies: FACTION_US, Axis: FACTION_GER, Orientation: ORIENTATION_VERTICAL, MirroredFactions: false},
	MAP_STMEREEGLISE:    {ID: MAP_STMEREEGLISE, Name: "SAINTE-MÈRE-ÉGLISE", Tag: "SME", PrettyName: "St. Mere Eglise", ShortName: "SME", Allies: FACTION_US, Axis: FACTION_GER, Orientation: ORIENTATION_HORIZONTAL, MirroredFactions: true},
	MAP_STALINGRAD:      {ID: MAP_STALINGRAD, Name: "STALINGRAD", Tag: "STA", PrettyName: "Stalingrad", ShortName: "Stalingrad", Allies: FACTION_SOV, Axis: FACTION_GER, Orientation: ORIENTATION_HORIZONTAL, MirroredFactions: true},
	MAP_TOBRUK:          {ID: MAP_TOBRUK, Name: "TOBRUK", Tag: "TBK", PrettyName: "Tobruk", ShortName: "Tobruk", Allies: FACTION_CW, Axis: FACTION_GER, Orientation: ORIENTATION_HORIZONTAL, MirroredFactions: true},
	MAP_UTAHBEACH:       {ID: MAP_UTAHBEACH, Name: "UTAH BEACH", Tag: "UTA", PrettyName: "Utah Beach", ShortName: "Utah", Allies: FACTION_US, Axis: FACTION_GER, Orientation: ORIENTATION_HORIZONTAL, MirroredFactions: true},
}

var fallback_map = Map{ID: MAP_INVALID, Name: "INVALID", Tag: "INV", PrettyName: "Invalid", ShortName: "Invalid", Allies: FACTION_US, Axis: FACTION_GER}

func (m MapIdentifier) Map() Map {
	return ParseMap(string(m))
}

func ParseMap(mapName string) Map {
	if gameMap, ok := mapMap[MapIdentifier(mapName)]; ok {
		return gameMap
	}
	logger.Warn("Map not found:", mapName)
	return fallback_map
}

func LogMapNameToMap(logMapName string) Map {
	for _, v := range mapMap {
		if strings.HasPrefix(logMapName, v.Name) {
			return v
		}
	}
	logger.Warn("LogMapName not found:", logMapName)
	return fallback_map
}

func AllMaps() []Map {
	maps := []Map{}
	for _, m := range mapMap {
		maps = append(maps, m)
	}
	return maps
}
