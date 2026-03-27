package hll

type FactionIdentifier string

const (
	FACTION_US         FactionIdentifier = "US"
	FACTION_GER        FactionIdentifier = "GER"
	FACTION_SOV        FactionIdentifier = "SOV"
	FACTION_CW         FactionIdentifier = "GB"
	FACTION_DAK        FactionIdentifier = "DAK"
	FACTION_B8A        FactionIdentifier = "B8A"
	FACTION_UNASSIGNED FactionIdentifier = "NON"
)

func (f FactionIdentifier) Faction() Faction {
	if faction, ok := factionMap[f]; ok {
		return faction
	}
	return Faction{ID: -1, Name: "Unassigned", ShortName: FACTION_UNASSIGNED, Team: TEAM_NONE}
}

func (f FactionIdentifier) Team() TeamIdentifier {
	if faction, ok := factionMap[f]; ok {
		return faction.Team
	}
	return TEAM_NONE
}

type Faction struct {
	ID        int
	Name      string
	ShortName FactionIdentifier
	Team      TeamIdentifier
}

var factionMap = map[FactionIdentifier]Faction{
	FACTION_GER: {
		ID:        0,
		Name:      "Germany",
		ShortName: FACTION_GER,
		Team:      TEAM_AXIS,
	},
	FACTION_US: {
		ID:        1,
		Name:      "United States",
		ShortName: FACTION_US,
		Team:      TEAM_ALLIES,
	},
	FACTION_SOV: {
		ID:        2,
		Name:      "Soviet Union",
		ShortName: FACTION_SOV,
		Team:      TEAM_ALLIES,
	},
	FACTION_CW: {
		ID:        3,
		Name:      "Allies",
		ShortName: FACTION_CW,
		Team:      TEAM_ALLIES,
	},
	FACTION_DAK: {
		ID:        4,
		Name:      "German Africa Corps",
		ShortName: FACTION_DAK,
		Team:      TEAM_AXIS,
	},
	FACTION_B8A: {
		ID:        5,
		Name:      "British Eighth Army",
		ShortName: FACTION_B8A,
		Team:      TEAM_ALLIES,
	},
}

var AllFactions = []FactionIdentifier{FACTION_US, FACTION_GER, FACTION_SOV, FACTION_CW, FACTION_DAK, FACTION_B8A}

func FactionFromInt(id int) FactionIdentifier {
	for _, faction := range factionMap {
		if faction.ID == id {
			return faction.ShortName
		}
	}
	return FACTION_UNASSIGNED
}
