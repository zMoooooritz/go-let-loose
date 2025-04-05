package rconv2

import (
	"fmt"

	"github.com/zMoooooritz/go-let-loose/internal/socketv2/api"
)

func (r *Rcon) GetQueuedPlayers() (int, error) {
	resp, err := getSessionInfo(r)
	return resp.QueueCount, err
}

func (r *Rcon) GetMaxQueuedPlayers() (int, error) {
	resp, err := getSessionInfo(r)
	return resp.MaxQueueCount, err
}

func (r *Rcon) GetQueuedVips() (int, error) {
	resp, err := getSessionInfo(r)
	return resp.VIPQueueCount, err
}

func (r *Rcon) GetNumVipSlots() (int, error) {
	resp, err := getSessionInfo(r)
	return resp.MaxVIPQueueCount, err
}

func (r *Rcon) GetIdleTime() (int, error) {
	return 0, fmt.Errorf("not implemented")
}

func (r *Rcon) GetHighPing() (int, error) {
	return 0, fmt.Errorf("not implemented")
}

func (r *Rcon) GetTeamSwitchCooldown() (int, error) {
	return 0, fmt.Errorf("not implemented")
}

func (r *Rcon) IsAutoBalanceEnabled() (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func (r *Rcon) GetAutoBalanceThreshold() (int, error) {
	return 0, fmt.Errorf("not implemented")
}

func (r *Rcon) IsVoteKickEnabled() (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func (r *Rcon) GetVoteKickThreshold() (int, error) {
	return 0, fmt.Errorf("not implemented")
}

func (r *Rcon) GetProfanities() ([]string, error) {
	return []string{}, fmt.Errorf("not implemented")
}

func (r *Rcon) SetMaxQueuedPlayers(size int) error {
	_, err := runCommand[api.SetMaxQueuedPlayers, any](r,
		api.SetMaxQueuedPlayers{
			MaxQueuedPlayers: size,
		},
	)
	return err
}

func (r *Rcon) SetNumVipSlots(amount int) error {
	return fmt.Errorf("not implemented")
}

func (r *Rcon) SetKickIdleTime(threshold int) error {
	_, err := runCommand[api.SetIdleKickDuration, any](r,
		api.SetIdleKickDuration{
			IdleTimeoutMinutes: threshold,
		},
	)
	return err
}

func (r *Rcon) SetHighPing(threshold int) error {
	_, err := runCommand[api.SetHighPingThreshold, any](r,
		api.SetHighPingThreshold{
			HighPingThresholdMs: threshold,
		},
	)
	return err
}

func (r *Rcon) SetTeamSwitchCooldown(cooldown int) error {
	_, err := runCommand[api.SetTeamSwitchCooldown, any](r,
		api.SetTeamSwitchCooldown{
			TeamSwitchTimer: cooldown,
		},
	)
	return err
}

func (r *Rcon) SetAutoBalanceEnabled(enabled bool) error {
	_, err := runCommand[api.AutoBalance, any](r,
		api.AutoBalance{
			EnableAutoBalance: enabled,
		},
	)
	return err
}

func (r *Rcon) SetAutoBalanceThreshold(threshold int) error {
	_, err := runCommand[api.AutoBalanceThreshold, any](r,
		api.AutoBalanceThreshold{
			AutoBalanceThreshold: threshold,
		},
	)
	return err
}

func (r *Rcon) SetVoteKickEnabled(enabled bool) error {
	_, err := runCommand[api.VoteKickEnabled, any](r,
		api.VoteKickEnabled{
			Enabled: enabled,
		},
	)
	return err
}

func (r *Rcon) SetVoteKickThreshold(thresholdPairs string) error {
	_, err := runCommand[api.VoteKickThreshold, any](r,
		api.VoteKickThreshold{
			ThresholdValue: thresholdPairs,
		},
	)
	return err
}

func (r *Rcon) ResetVoteKickThreshold() error {
	_, err := runCommand[api.ResetKickThreshold, any](r,
		api.ResetKickThreshold{},
	)
	return err
}

func (r *Rcon) BanProfanities(profanities []string) error {
	return fmt.Errorf("not implemented")
}

func (r *Rcon) UnbanProfanities(profanities []string) error {
	return fmt.Errorf("not implemented")
}
