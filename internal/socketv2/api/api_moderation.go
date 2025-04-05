package api

type Kick struct {
	Reason   string `json:"Reason"`
	PlayerID string `json:"PlayerID"`
}

type MessagePlayer struct {
	Message  string `json:"Message"`
	PlayerID string `json:"PlayerID"`
}

type PermanentBan struct {
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

type RemoveTempBan struct {
	PlayerID string `json:"PlayerID"`
}

type SendServerMessage struct {
	Message string `json:"Message"`
}

type TempBan struct {
	Reason    string `json:"Reason"`
	PlayerID  string `json:"PlayerID"`
	Duration  int    `json:"Duration"`
	AdminName string `json:"AdminName"`
}
