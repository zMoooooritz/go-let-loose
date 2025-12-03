package rcon

import (
	"fmt"
	"strings"

	"github.com/zMoooooritz/go-let-loose/internal/socket/api"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

func (r *Rcon) GetQueuedPlayers() (int, error) {
	resp, err := getSessionInfo(r)
	return int(resp.QueueCount), err
}

func (r *Rcon) GetMaxQueuedPlayers() (int, error) {
	resp, err := getSessionInfo(r)
	return int(resp.MaxQueueCount), err
}

func (r *Rcon) GetQueuedVips() (int, error) {
	resp, err := getSessionInfo(r)
	return int(resp.VipQueueCount), err
}

func (r *Rcon) GetNumVipSlots() (int, error) {
	resp, err := getSessionInfo(r)
	return int(resp.MaxVipQueueCount), err
}

func (r *Rcon) GetIdleTime() (int, error) {
	resp, err := runCommand[api.GetKickIdleDuration, api.RespKickIdleDuration](r,
		api.GetKickIdleDuration{},
	)
	if err != nil {
		return 0, err
	}
	return int(resp.IdleTimeoutMinutes), nil
}

func (r *Rcon) SetKickIdleTime(threshold int) error {
	_, err := runCommand[api.SetIdleKickDuration, any](r,
		api.SetIdleKickDuration{
			IdleTimeoutMinutes: int32(threshold),
		},
	)
	return err
}

func (r *Rcon) GetHighPing() (int, error) {
	resp, err := runCommand[api.GetHighPingThreshold, api.RespHighPingThreshold](r,
		api.GetHighPingThreshold{},
	)
	if err != nil {
		return 0, err
	}
	return int(resp.HighPingThresholdMs), nil
}

func (r *Rcon) SetHighPing(threshold int) error {
	_, err := runCommand[api.SetHighPingThreshold, any](r,
		api.SetHighPingThreshold{
			HighPingThresholdMs: int32(threshold),
		},
	)
	return err
}

func (r *Rcon) GetTeamSwitchCooldown() (int, error) {
	resp, err := runCommand[api.GetTeamSwitchCooldown, api.RespTeamSwitchCooldown](r,
		api.GetTeamSwitchCooldown{},
	)
	if err != nil {
		return 0, err
	}
	return int(resp.TeamSwitchTimer), nil
}

func (r *Rcon) SetTeamSwitchCooldown(cooldown int) error {
	_, err := runCommand[api.SetTeamSwitchCooldown, any](r,
		api.SetTeamSwitchCooldown{
			TeamSwitchTimer: int32(cooldown),
		},
	)
	return err
}

func (r *Rcon) IsAutoBalanceEnabled() (bool, error) {
	resp, err := runCommand[api.GetAutoBalanceEnabled, api.RespAutoBalanceEnabled](r,
		api.GetAutoBalanceEnabled{},
	)
	if err != nil {
		return false, err
	}
	return resp.Enable, nil
}

func (r *Rcon) SetAutoBalanceEnabled(enabled bool) error {
	_, err := runCommand[api.SetAutoBalanceEnabled, any](r,
		api.SetAutoBalanceEnabled{
			Enable: enabled,
		},
	)
	return err
}

func (r *Rcon) GetAutoBalanceThreshold() (int, error) {
	resp, err := runCommand[api.GetAutoBalanceThreshold, api.RespAutoBalanceThreshold](r,
		api.GetAutoBalanceThreshold{},
	)
	if err != nil {
		return 0, err
	}
	return int(resp.AutoBalanceThreshold), nil
}

func (r *Rcon) SetAutoBalanceThreshold(threshold int) error {
	_, err := runCommand[api.SetAutoBalanceThreshold, any](r,
		api.SetAutoBalanceThreshold{
			AutoBalanceThreshold: int32(threshold),
		},
	)
	return err
}

func (r *Rcon) IsVoteKickEnabled() (bool, error) {
	resp, err := runCommand[api.GetVoteKickEnabled, api.RespVoteKickEnabled](r,
		api.GetVoteKickEnabled{},
	)
	if err != nil {
		return false, err
	}
	return resp.Enable, nil
}

func (r *Rcon) SetVoteKickEnabled(enabled bool) error {
	_, err := runCommand[api.SetVoteKickEnabled, any](r,
		api.SetVoteKickEnabled{
			Enable: enabled,
		},
	)
	return err
}

func (r *Rcon) GetVoteKickThresholds() ([]hll.ThresholdPair, error) {
	resp, err := runCommand[api.GetVoteKickThreshold, api.RespVoteKickThreshold](r,
		api.GetVoteKickThreshold{},
	)
	if err != nil {
		return []hll.ThresholdPair{}, err
	}
	thresholdPairs := []hll.ThresholdPair{}
	for _, pair := range resp.Entries {
		thresholdPairs = append(thresholdPairs, hll.ThresholdPair{
			PlayerCount: int(pair.PlayerCount),
			Threshold:   int(pair.VoteThreshold),
		})
	}
	return thresholdPairs, nil
}

func (r *Rcon) SetVoteKickThresholds(thresholdPairs []hll.ThresholdPair) error {
	thresholdStrs := []string{}
	for _, pair := range thresholdPairs {
		thresholdStrs = append(thresholdStrs, fmt.Sprintf("%d,%d", pair.PlayerCount, pair.Threshold))
	}
	thresholdValue := strings.Join(thresholdStrs, ",")

	_, err := runCommand[api.SetVoteKickThreshold, any](r,
		api.SetVoteKickThreshold{
			ThresholdValue: thresholdValue,
		},
	)
	return err
}

func (r *Rcon) ResetVoteKickThreshold() error {
	_, err := runCommand[api.ResetVoteKickThreshold, any](r,
		api.ResetVoteKickThreshold{},
	)
	return err
}

func (r *Rcon) SetMaxQueuedPlayers(size int) error {
	_, err := runCommand[api.SetMaxQueuedPlayers, any](r,
		api.SetMaxQueuedPlayers{
			MaxQueuedPlayers: int32(size),
		},
	)
	return err
}

func (r *Rcon) SetNumVipSlots(amount int) error {
	_, err := runCommand[api.SetVipSlotCount, any](r,
		api.SetVipSlotCount{
			VipSlotCount: int32(amount),
		},
	)
	return err
}

func (r *Rcon) BanProfanities(profanities []string) error {
	_, err := runCommand[api.AddBannedWords, any](r,
		api.AddBannedWords{
			BannedWords: strings.Join(profanities, ","),
		})
	return err
}

func (r *Rcon) UnbanProfanities(profanities []string) error {
	_, err := runCommand[api.RemoveBannedWords, any](r,
		api.RemoveBannedWords{
			BannedWords: strings.Join(profanities, ","),
		})
	return err
}

func (r *Rcon) GetProfanities() ([]string, error) {
	resp, err := getBannedWords(r)
	if err != nil {
		return []string{}, err
	}
	return resp.BannedWords, nil
}
