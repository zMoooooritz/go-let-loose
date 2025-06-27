package rconv2

import (
	"fmt"

	"github.com/zMoooooritz/go-let-loose/internal/socketv2/api"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

func (r *Rcon) GetTempBans() ([]hll.ServerBan, error) {
	// TODO: via ClientReferenceData-API?
	return []hll.ServerBan{}, fmt.Errorf("not implemented")
}

func (r *Rcon) GetPermaBans() ([]hll.ServerBan, error) {
	// TODO: via ClientReferenceData-API?
	return []hll.ServerBan{}, fmt.Errorf("not implemented")
}

func (r *Rcon) MessagePlayer(playerID string, message string) error {
	_, err := runCommand[api.MessagePlayer, any](r,
		api.MessagePlayer{
			Message:  message,
			PlayerID: playerID,
		},
	)
	return err
}

func (r *Rcon) PunishPlayer(player, reason string) error {
	_, err := runCommand[api.PunishPlayer, any](r,
		api.PunishPlayer{
			Reason:   reason,
			PlayerID: player,
		},
	)
	return err
}

func (r *Rcon) SwitchPlayerOnDeath(player string) error {
	return fmt.Errorf("not implemented")
}

func (r *Rcon) SwitchPlayerNow(player string) error {
	return fmt.Errorf("not implemented")
}

func (r *Rcon) KickPlayer(player, reason string) error {
	_, err := runCommand[api.Kick, any](r,
		api.Kick{
			Reason:   reason,
			PlayerID: player,
		},
	)
	return err
}

func (r *Rcon) TempBanPlayer(player string, duration int, reason, admin string) error {
	_, err := runCommand[api.TemporaryBanPlayer, any](r,
		api.TemporaryBanPlayer{
			Reason:    reason,
			PlayerID:  player,
			Duration:  duration,
			AdminName: admin,
		},
	)
	return err
}

func (r *Rcon) PardonTempBanPlayer(ban hll.ServerBan) error {
	_, err := runCommand[api.RemoveTemporaryBan, any](r,
		api.RemoveTemporaryBan{
			PlayerID: ban.Player.ID,
		},
	)
	return err
}

func (r *Rcon) PermaBanPlayer(player, reason, admin string) error {
	_, err := runCommand[api.PermanentBanPlayer, any](r,
		api.PermanentBanPlayer{
			Reason:    reason,
			PlayerID:  player,
			AdminName: admin,
		},
	)
	return err
}

func (r *Rcon) PardonPermaBanPlayer(ban hll.ServerBan) error {
	_, err := runCommand[api.RemovePermanentBan, any](r,
		api.RemovePermanentBan{
			PlayerID: ban.Player.ID,
		},
	)
	return err
}
