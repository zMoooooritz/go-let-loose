package hll

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

type Platform string

const (
	PlatformSteam   Platform = "Steam"
	PlatformWindows Platform = "WinGDK"
	PlatformEpic    Platform = "eos"
	PlatformNone    Platform = "None"
)

func PlatformFromString(name string) Platform {
	typed := Platform(name)
	switch typed {
	case PlatformSteam, PlatformWindows, PlatformEpic:
		return typed
	default:
		return PlatformNone
	}
}

type ServerInfo struct {
	Name               string
	BuildNumber        string
	BuildRevision      string
	SupportedPlatforms []Platform
}
