package api

type SetAutoBalanceEnabled struct {
	Enable bool `json:"EnableAutoBalance"`
}

type SetAutoBalanceThreshold struct {
	AutoBalanceThreshold int32 `json:"AutoBalanceThreshold"`
}

type ResetVoteKickThreshold struct {
}

type SetHighPingThreshold struct {
	HighPingThresholdMs int32 `json:"HighPingThresholdMs"`
}

type SetIdleKickDuration struct {
	IdleTimeoutMinutes int32 `json:"IdleTimeoutMinutes"`
}

type SetTeamSwitchCooldown struct {
	TeamSwitchTimer int32 `json:"TeamSwitchTimer"`
}

type SetVoteKickEnabled struct {
	Enable bool `json:"Enable"`
}

type SetVoteKickThreshold struct {
	ThresholdValue string `json:"ThresholdValue"`
}

type SetVipSlotCount struct {
	VipSlotCount int32 `json:"VipSlotCount"`
}

type AddBannedWords struct {
	BannedWords string `json:"BannedWords"`
}

type RemoveBannedWords struct {
	BannedWords string `json:"BannedWords"`
}
