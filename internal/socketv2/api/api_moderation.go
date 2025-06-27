package api

type Kick struct {
	Reason   string `json:"Reason"`
	PlayerID string `json:"PlayerID"`
}

type MessagePlayer struct {
	Message  string `json:"Message"`
	PlayerID string `json:"PlayerID"`
}

type PermanentBanPlayer struct {
	Reason    string `json:"Reason"`
	PlayerID  string `json:"PlayerID"`
	AdminName string `json:"AdminName"`
}

type PunishPlayer struct {
	Reason   string `json:"Reason"`
	PlayerID string `json:"PlayerID"`
}

type RemovePermanentBan struct {
	PlayerID string `json:"PlayerID"`
}

type RemoveTemporaryBan struct {
	PlayerID string `json:"PlayerID"`
}

type TemporaryBanPlayer struct {
	Reason    string `json:"Reason"`
	PlayerID  string `json:"PlayerID"`
	Duration  int    `json:"Duration"`
	AdminName string `json:"AdminName"`
}

type AddAdmin struct {
	PlayerId   string `json:"playerId"`
	AdminGroup string `json:"adminGroup"`
	Comment    string `json:"comment"`
}

type RemoveAdmin struct {
	PlayerId string `json:"playerId"`
}

type AddVipPlayer struct {
	PlayerId    string `json:"PlayerId"`
	Description string `json:"Description"`
}

type RemoveVipPlayer struct {
	PlayerId string `json:"PlayerId"`
}
