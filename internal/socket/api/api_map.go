package api

type AddMapToRotation struct {
	MapName string `json:"MapName"`
	Index   int32  `json:"Index"`
}

type AddMapToSequence struct {
	MapName string `json:"MapName"`
	Index   int32  `json:"Index"`
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
	CurrentIndex int32 `json:"CurrentIndex"`
	NewIndex     int32 `json:"NewIndex"`
}

type RemoveMapFromRotation struct {
	Index int32 `json:"Index"`
}

type RemoveMapFromSequence struct {
	Index int32 `json:"Index"`
}

type SetMapShuffleEnabled struct {
	Enable bool `json:"Enable"`
}

type SetDynamicWeatherEnabled struct {
	MapId  string `json:"MapId"`
	Enable bool   `json:"Enable"`
}

type SetMatchTimer struct {
	GameMode    string `json:"GameMode"`
	MatchLength int32  `json:"MatchLength"` // in minutes
}
type RemoveMatchTimer struct {
	GameMode string `json:"GameMode"`
}

type SetWarmupTimer struct {
	GameMode     string `json:"GameMode"`
	WarmupLength int32  `json:"WarmupLength"` // in minutes
}

type RemoveWarmupTimer struct {
	GameMode string `json:"GameMode"`
}
