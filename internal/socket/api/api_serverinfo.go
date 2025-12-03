package api

import (
	"encoding/json"
	"time"
)

const (
	ServerInfoPlayers      = "players"
	ServerInfoPlayer       = "player"
	ServerInfoMapRotation  = "maprotation"
	ServerInfoMapSequence  = "mapsequence"
	ServerInfoSession      = "session"
	ServerInfoServerConfig = "serverconfig"
	ServerInfoBannedWords  = "bannedwords"
	ServerInfoVipPlayers   = "vipplayers"
)

type ServerInformationName string

type GetServerInformation struct {
	Name  ServerInformationName `json:"Name"`
	Value string                `json:"Value"`
}

func (s *GetServerInformation) Pack() []byte {
	body, err := json.Marshal(s)
	if err != nil {
		return []byte{}
	}
	return body
}

type RespScoreData struct {
	Combat  int32 `json:"Combat"`
	Offense int32 `json:"Offense"`
	Defense int32 `json:"Defense"`
	Support int32 `json:"Support"`
}

type RespWorldPosition struct {
	X float64 `json:"X"`
	Y float64 `json:"Y"`
	Z float64 `json:"Z"`
}

type RespPlayerStats struct {
	Deaths            int32 `json:"Deaths"`
	InfantryKills     int32 `json:"InfantryKills"`
	VehicleKills      int32 `json:"VehicleKills"`
	TeamKills         int32 `json:"TeamKills"`
	VehiclesDestroyed int32 `json:"VehiclesDestroyed"`
}

type RespPlayerInformation struct {
	Name     string            `json:"Name"`
	ClanTag  string            `json:"ClanTag"`
	ID       string            `json:"ID"`
	Platform string            `json:"Platform"`
	EpicID   string            `json:"EOSID"`
	Level    int32             `json:"Level"`
	Team     int32             `json:"Team"`
	Role     int32             `json:"Role"`
	Platoon  string            `json:"Platoon"`
	Stats    RespPlayerStats   `json:"Stats"`
	Score    RespScoreData     `json:"ScoreData"`
	Loadout  string            `json:"Loadout"`
	Position RespWorldPosition `json:"WorldPosition"`
}

func (r RespPlayerInformation) CacheTTL() time.Duration {
	return 500 * time.Millisecond
}

type RespPlayersInformation struct {
	Players []RespPlayerInformation `json:"Players"`
}

func (r RespPlayersInformation) CacheTTL() time.Duration {
	return 500 * time.Millisecond
}

type RespMapInformation struct {
	Name      string `json:"Name"`
	GameMode  string `json:"GameMode"`
	TimeOfDay string `json:"TimeOfDay"`
	ID        string `json:"Id"`
	Position  int32  `json:"Position"`
}

type RespMapRotation struct {
	CurrentIndex int32                `json:"CurrentIndex"`
	Maps         []RespMapInformation `json:"Maps"`
}

func (r RespMapRotation) CacheTTL() time.Duration {
	return 1 * time.Second
}

type RespMapSequence struct {
	Maps []RespMapInformation `json:"Maps"`
}

func (r RespMapSequence) CacheTTL() time.Duration {
	return 1 * time.Second
}

type RespSessionInformation struct {
	ServerName         string `json:"ServerName"`
	MapName            string `json:"MapName"`
	MapID              string `json:"MapId"`
	GameMode           string `json:"GameMode"`
	RemainingMatchTime int32  `json:"RemainingMatchTime"`
	MatchTime          int32  `json:"MatchTime"`
	AlliedFaction      int32  `json:"AlliedFaction"`
	AxisFaction        int32  `json:"AxisFaction"`
	MaxPlayerCount     int32  `json:"MaxPlayerCount"`
	AlliedScore        int32  `json:"AlliedScore"`
	AxisScore          int32  `json:"AxisScore"`
	PlayerCount        int32  `json:"PlayerCount"`
	AlliedPlayerCount  int32  `json:"AlliedPlayerCount"`
	AxisPlayerCount    int32  `json:"AxisPlayerCount"`
	MaxQueueCount      int32  `json:"MaxQueueCount"`
	QueueCount         int32  `json:"QueueCount"`
	MaxVipQueueCount   int32  `json:"MaxVipQueueCount"`
	VipQueueCount      int32  `json:"VipQueueCount"`
}

func (r RespSessionInformation) CacheTTL() time.Duration {
	return 500 * time.Millisecond
}

type RespServerConfiguration struct {
	ServerName         string   `json:"ServerName"`
	BuildNumber        string   `json:"BuildNumber"`
	BuildRevision      string   `json:"BuildRevision"`
	SupportedPlatforms []string `json:"SupportedPlatforms"`
	PasswordProtected  bool     `json:"PasswordProtected"`
}

func (r RespServerConfiguration) CacheTTL() time.Duration {
	return 1 * time.Second
}

type RespBannedWords struct {
	BannedWords []string `json:"BannedWords"`
}

func (r RespBannedWords) CacheTTL() time.Duration {
	return 1 * time.Second
}

type RespVipPlayerEntry struct {
	ID      string `json:"ID"`
	Comment string `json:"Comment"`
}

type RespVipPlayers struct {
	VipPlayers []RespVipPlayerEntry `json:"VipPlayers"`
}

func (r RespVipPlayers) CacheTTL() time.Duration {
	return 1 * time.Second
}
