package api

import "encoding/json"

const (
	ServerInfoPlayers      = "players"
	ServerInfoPlayer       = "player"
	ServerInfoMapRotation  = "maprotation"
	ServerInfoMapSequence  = "mapsequence"
	ServerInfoSession      = "session"
	ServerInfoServerConfig = "serverconfig"
)

type ServerInformationName string

type ServerInformation struct {
	Name  ServerInformationName `json:"Name"`
	Value string                `json:"Value"`
}

func (s *ServerInformation) Pack() []byte {
	body, err := json.Marshal(s)
	if err != nil {
		return []byte{}
	}
	return body
}

type ScoreData struct {
	Combat  int `json:"Combat"`
	Offense int `json:"Offense"`
	Defense int `json:"Defense"`
	Support int `json:"Support"`
}

type WorldPosition struct {
	X float64 `json:"X"`
	Y float64 `json:"Y"`
	Z float64 `json:"Z"`
}

type RespPlayerInformation struct {
	Name     string        `json:"Name"`
	ClanTag  string        `json:"ClanTag"`
	ID       string        `json:"ID"`
	Platform string        `json:"Platform"`
	EpicID   string        `json:"EOSID"`
	Level    int           `json:"Level"`
	Team     int           `json:"Team"`
	Role     int           `json:"Role"`
	Platoon  string        `json:"Platoon"`
	Kills    int           `json:"Kills"`
	Deaths   int           `json:"Deaths"`
	Score    ScoreData     `json:"ScoreData"`
	Loadout  string        `json:"Loadout"`
	Position WorldPosition `json:"WorldPosition"`
}

type RespPlayersInformation struct {
	Players []RespPlayerInformation `json:"Players"`
}

type MapInformation struct {
	Name      string `json:"Name"`
	GameMode  string `json:"GameMode"`
	TimeOfDay string `json:"TimeOfDay"`
	ID        string `json:"ID"`
	Position  int    `json:"Position"`
}

type RespMapRotation struct {
	Maps []MapInformation `json:"Maps"`
}

type RespMapSequence struct {
	Maps []MapInformation `json:"Maps"`
}

type RespSessionInformation struct {
	ServerName       string `json:"ServerName"`
	MapName          string `json:"MapName"`
	GameMode         string `json:"GameMode"`
	MaxPlayerCount   int    `json:"MaxPlayerCount"`
	PlayerCount      int    `json:"PlayerCount"`
	MaxQueueCount    int    `json:"MaxQueueCount"`
	QueueCount       int    `json:"QueueCount"`
	MaxVIPQueueCount int    `json:"MaxVIPQueueCount"`
	VIPQueueCount    int    `json:"VIPQueueCount"`
}

type RespServerConfiguration struct {
	ServerName         string   `json:"ServerName"`
	BuildNumber        string   `json:"BuildNumber"`
	BuildRevision      string   `json:"BuildRevision"`
	SupportedPlatforms []string `json:"SupportedPlatforms"`
}
