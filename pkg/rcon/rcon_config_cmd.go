package rcon

import (
	"fmt"
	"strings"
)

func (r *Rcon) GetIdleTime() (int, error) {
	return getNumVal(r, "get idletime")
}

func (r *Rcon) GetHighPing() (int, error) {
	return getNumVal(r, "get highping")
}

func (r *Rcon) GetTeamSwitchCooldown() (int, error) {
	return getNumVal(r, "get teamswitchcooldown")
}

func (r *Rcon) IsAutoBalanceEnabled() (bool, error) {
	return getBoolVal(r, "get autobalanceenabled")
}

func (r *Rcon) GetAutoBalanceThreshold() (int, error) {
	return getNumVal(r, "get autobalancethreshold")
}

func (r *Rcon) IsVoteKickEnabled() (bool, error) {
	return getBoolVal(r, "get votekickenabled")
}

func (r *Rcon) GetVoteKickThreshold() (int, error) {
	return getNumVal(r, "get votekickthreshold")
}

func (r *Rcon) GetProfanities() ([]string, error) {
	return r.runListCommand("get profanity")
}

func (r *Rcon) SetKickIdleTime(threshold int) error {
	return runSetCommand(r, fmt.Sprintf("setkickidletime %d", max(0, threshold)))
}

func (r *Rcon) SetHighPing(threshold int) error {
	return runSetCommand(r, fmt.Sprintf("sethighping %d", max(0, threshold)))
}

func (r *Rcon) SetTeamSwitchCooldown(cooldown int) error {
	return runSetCommand(r, fmt.Sprintf("setteamswitchcooldown %d", max(0, cooldown)))
}

func (r *Rcon) SetAutoBalanceEnabled(enabled bool) error {
	return runSetCommand(r, fmt.Sprintf("setautobalanceenabled %s", boolToToggleStr(enabled)))
}

func (r *Rcon) SetAutoBalanceThreshold(threshold int) error {
	return runSetCommand(r, fmt.Sprintf("setautobalancethreshold %d", threshold))
}

func (r *Rcon) SetVoteKickEnabled(enabled bool) error {
	return runSetCommand(r, fmt.Sprintf("setvotekickenabled %s", boolToToggleStr(enabled)))
}

func (r *Rcon) Setvotekickthreshold(thresholdPairs string) error {
	return runSetCommand(r, fmt.Sprintf("setvotekickthreshold %s", thresholdPairs))
}

func (r *Rcon) Resetvotekickthreshold(thresholdPairs string) error {
	return runSetCommand(r, "resetvotekickthreshold")
}

func (r *Rcon) BanProfanities(profanities []string) error {
	return runSetCommand(r, fmt.Sprintf("banprofanity %s", strings.Join(profanities, ",")))
}

func (r *Rcon) UnbanProfanities(profanities []string) error {
	return runSetCommand(r, fmt.Sprintf("unbanprofanity %s", strings.Join(profanities, ",")))
}
