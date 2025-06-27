package api

type AddMapToRotation struct {
	MapName string `json:"MapName"`
	Index   int    `json:"Index"`
}

type AddMapToSequence struct {
	MapName string `json:"MapName"`
	Index   int    `json:"Index"`
}

type SetSectorLayout struct {
	SectorOne   string `json:"Sector_1"`
	SectorTwo   string `json:"Sector_2"`
	SectorThree string `json:"Sector_3"`
	SectorFour  string `json:"Sector_4"`
	SectorFive  string `json:"Sector_5"`
}

type ChangeMap struct {
	MapName string `json:"MapName"`
}

type MoveMapInSequence struct {
	CurrentIndex int `json:"CurrentIndex"`
	NewIndex     int `json:"NewIndex"`
}

type RemoveMapFromRotation struct {
	Index int `json:"Index"`
}

type RemoveMapFromSequence struct {
	Index int `json:"index"`
}

type SetShuffleMapSequence struct {
	Enable bool `json:"Enable"`
}

type SetMapWeatherToggle struct {
	MapId  string `json:"MapId"`
	Enable bool   `json:"Enable"`
}

type SetMatchTimer struct {
	GameMode    string `json:"GameMode"`
	MatchLength int    `json:"MatchLength"` // in minutes
}
type RemoveMatchTimer struct {
	GameMode string `json:"GameMode"`
}

type SetWarmupTimer struct {
	GameMode     string `json:"GameMode"`
	WarmupLength int    `json:"WarmupLength"` // in minutes
}

type RemoveWarmupTimer struct {
	GameMode string `json:"GameMode"`
}
