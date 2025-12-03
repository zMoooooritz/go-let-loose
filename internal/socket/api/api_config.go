package api

type GetAutoBalanceEnabled struct {
}

type RespAutoBalanceEnabled struct {
	Enable bool `json:"Enable"`
}

type SetAutoBalanceEnabled struct {
	Enable bool `json:"EnableAutoBalance"`
}

type GetAutoBalanceThreshold struct {
}

type RespAutoBalanceThreshold struct {
	AutoBalanceThreshold int32 `json:"AutoBalanceThreshold"`
}

type SetAutoBalanceThreshold struct {
	AutoBalanceThreshold int32 `json:"AutoBalanceThreshold"`
}

type ResetVoteKickThreshold struct {
}

type GetHighPingThreshold struct {
}

type RespHighPingThreshold struct {
	HighPingThresholdMs int32 `json:"HighPingThresholdMs"`
}

type SetHighPingThreshold struct {
	HighPingThresholdMs int32 `json:"HighPingThresholdMs"`
}

type GetKickIdleDuration struct {
}

type RespKickIdleDuration struct {
	IdleTimeoutMinutes int32 `json:"IdleTimeoutMinutes"`
}

type SetIdleKickDuration struct {
	IdleTimeoutMinutes int32 `json:"IdleTimeoutMinutes"`
}

type GetTeamSwitchCooldown struct {
}

type RespTeamSwitchCooldown struct {
	TeamSwitchTimer int32 `json:"TeamSwitchTimer"`
}

type GetVoteKickEnabled struct {
}

type RespVoteKickEnabled struct {
	Enable bool `json:"Enable"`
}

type SetTeamSwitchCooldown struct {
	TeamSwitchTimer int32 `json:"TeamSwitchTimer"`
}

type SetVoteKickEnabled struct {
	Enable bool `json:"Enable"`
}

type GetVoteKickThreshold struct {
}

type RespVoteKickThresholdEntry struct {
	PlayerCount   int32 `json:"PlayerCount"`
	VoteThreshold int32 `json:"VoteThreshold"`
}

type RespVoteKickThreshold struct {
	Entries []RespVoteKickThresholdEntry `json:"Entries"`
}

type SetVoteKickThreshold struct {
	ThresholdValue string `json:"ThresholdValue"`
}

type SetMaxQueuedPlayers struct {
	MaxQueuedPlayers int32 `json:"MaxQueuedPlayers"`
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
