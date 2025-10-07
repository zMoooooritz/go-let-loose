package api

type MessagePlayer struct {
	Message  string `json:"Message"`
	PlayerID string `json:"PlayerId"`
}

type ForceTeamSwitch struct {
	PlayerID  string `json:"PlayerId"`
	ForceMode uint8  `json:"ForceMode"`
}

type RemovePlayerFromPlatoon struct {
	PlayerID string `json:"PlayerId"`
	Reason   string `json:"Reason"`
}

type DisbandPlatoon struct {
	TeamIndex  int8   `json:"TeamIndex"`
	SquadIndex int32  `json:"SquadIndex"`
	Reason     string `json:"Reason"`
}

type PunishPlayer struct {
	Reason   string `json:"Reason"`
	PlayerID string `json:"PlayerId"`
}

type KickPlayer struct {
	Reason   string `json:"Reason"`
	PlayerID string `json:"PlayerId"`
}

type BanEntry struct {
	UserID        string `json:"UserId"`
	UserName      string `json:"UserName"`
	TimeOfBanning string `json:"TimeOfBanning"`
	Duration      int32  `json:"DurationHours"`
	Reason        string `json:"BanReason"`
	AdminName     string `json:"AdminName"`
}

type GetTemporaryBans struct{}

type RespTemporaryBans struct {
	BanList []BanEntry `json:"BanList"`
}

type TemporaryBanPlayer struct {
	Reason    string `json:"Reason"`
	PlayerID  string `json:"PlayerId"`
	Duration  int32  `json:"Duration"`
	AdminName string `json:"AdminName"`
}

type RemoveTemporaryBan struct {
	PlayerID string `json:"PlayerId"`
}

type GetPermanentBans struct{}

type RespPermanentBans struct {
	BanList []BanEntry `json:"BanList"`
}

type PermanentBanPlayer struct {
	Reason    string `json:"Reason"`
	PlayerID  string `json:"PlayerId"`
	AdminName string `json:"AdminName"`
}

type RemovePermanentBan struct {
	PlayerID string `json:"PlayerId"`
}

type AddAdmin struct {
	PlayerId   string `json:"PlayerId"`
	AdminGroup string `json:"AdminGroup"`
	Comment    string `json:"Comment"`
}

type RemoveAdmin struct {
	PlayerId string `json:"PlayerId"`
}

type AddVip struct {
	PlayerId    string `json:"PlayerId"`
	Description string `json:"Description"`
}

type RemoveVip struct {
	PlayerId string `json:"PlayerId"`
}

type GetAdminGroups struct{}

type RespAdminGroups struct {
	GroupNames []string `json:"GroupNames"`
}

type GetAdminUsers struct{}

type RespAdminUsers struct {
	AdminUsers []AdminUser `json:"AdminUsers"`
}

type AdminUser struct {
	UserId  string `json:"UserId"`
	Group   string `json:"Group"`
	Comment string `json:"Comment"`
}
