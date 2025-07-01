package api

type KickPlayer struct {
	Reason   string `json:"Reason"`
	PlayerID string `json:"PlayerId"`
}

type MessagePlayer struct {
	Message  string `json:"Message"`
	PlayerID string `json:"PlayerId"`
}

type PermanentBanPlayer struct {
	Reason    string `json:"Reason"`
	PlayerID  string `json:"PlayerId"`
	AdminName string `json:"AdminName"`
}

type PunishPlayer struct {
	Reason   string `json:"Reason"`
	PlayerID string `json:"PlayerId"`
}

type RemovePermanentBan struct {
	PlayerID string `json:"PlayerId"`
}

type RemoveTemporaryBan struct {
	PlayerID string `json:"PlayerId"`
}

type TemporaryBanPlayer struct {
	Reason    string `json:"Reason"`
	PlayerID  string `json:"PlayerId"`
	Duration  int    `json:"Duration"`
	AdminName string `json:"AdminName"`
}

type AddAdmin struct {
	PlayerId   string `json:"PlayerId"`
	AdminGroup string `json:"AdminGroup"`
	Comment    string `json:"Comment"`
}

type RemoveAdmin struct {
	PlayerId string `json:"PlayerId"`
}

type AddVipPlayer struct {
	PlayerId    string `json:"PlayerId"`
	Description string `json:"Description"`
}

type RemoveVipPlayer struct {
	PlayerId string `json:"PlayerId"`
}
