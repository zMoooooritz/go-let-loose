package api

type AutoBalance struct {
	EnableAutoBalance bool `json:"EnableAutoBalance"`
}

type AutoBalanceThreshold struct {
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

type VoteKickEnabled struct {
	Enabled bool `json:"Enabled"`
}

type VoteKickThreshold struct {
	ThresholdValue string `json:"ThresholdValue"`
}
