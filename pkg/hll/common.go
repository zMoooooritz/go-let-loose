package hll

import "time"

type TeamData struct {
	Allies int
	Axis   int
}

type GameState struct {
	PlayerCount      TeamData
	GameScore        TeamData
	RemainingSeconds int
	CurrentMap       Layer
	NextMap          Layer
}

type Command struct {
	ID              string
	Name            string
	ClientSupported bool
}

type CommandDetails struct {
	Name               string
	Text               string
	Description        string
	DialogueParameters []DialogueParameter
}

type DialogueParameter struct {
	Type          string
	Name          string
	ID            string
	DisplayMember string
	ValueMember   string
}

type SupportedPlatform string

const (
	SupportedPlatformSteam   SupportedPlatform = "Steam"
	SupportedPlatformWindows SupportedPlatform = "WinGDK"
	SupportedPlatformEpic    SupportedPlatform = "eos"
	SupportedPlatformNone    SupportedPlatform = "None"
)

func SupportedPlatformFromString(name string) SupportedPlatform {
	typed := SupportedPlatform(name)
	switch typed {
	case SupportedPlatformSteam, SupportedPlatformWindows, SupportedPlatformEpic:
		return typed
	default:
		return SupportedPlatformNone
	}
}

type ServerConfig struct {
	Name               string
	BuildNumber        string
	BuildRevision      string
	SupportedPlatforms []SupportedPlatform
}

type SessionInfo struct {
	ServerName       string
	MapName          string
	GameMode         GameMode
	MaxPlayerCount   int
	PlayerCount      int
	MaxQueueCount    int
	QueueCount       int
	MaxVIPQueueCount int
	VIPQueueCount    int
}

type LogEntry struct {
	Timestamp time.Time
	Message   string
}
