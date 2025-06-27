package api

type SetAutoBalance struct {
	EnableAutoBalance bool `json:"EnableAutoBalance"`
}

type SetAutoBalanceThreshold struct {
	AutoBalanceThreshold int `json:"AutoBalanceThreshold"`
}

type ResetKickThreshold struct {
}

type SetHighPingThreshold struct {
	HighPingThresholdMs int `json:"HighPingThresholdMs"`
}

type SetIdleKickDuration struct {
	IdleTimeoutMinutes int `json:"IdleTimeoutMinutes"`
}

type SetTeamSwitchCooldown struct {
	TeamSwitchTimer int `json:"TeamSwitchTimer"`
}

type SetVoteKick struct {
	Enabled bool `json:"Enabled"`
}

type SetVoteKickThreshold struct {
	ThresholdValue string `json:"ThresholdValue"`
}

type AddBannedWords struct {
	BannedWords string `json:"BannedWords"`
}

type RemoveBannedWords struct {
	BannedWords string `json:"BannedWords"`
}
