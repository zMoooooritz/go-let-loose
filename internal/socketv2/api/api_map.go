package api

type AddMapToRotation struct {
	MapName string `json:"MapName"`
	Index   int    `json:"Index"`
}

type AddMapToSequence struct {
	MapName string `json:"MapName"`
	Index   int    `json:"Index"`
}

type ChangeSectorLayout struct {
	SectorOne   string `json:"Sector_1"`
	SectorTwo   string `json:"Sector_2"`
	SectorThree string `json:"Sector_3"`
	SectorFour  string `json:"Sector_4"`
	SectorFive  string `json:"Sector_5"`
}

type MapChange struct {
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

type ShuffleMapSequence struct {
	Enable bool `json:"Enable"`
}
