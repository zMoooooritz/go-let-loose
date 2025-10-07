package rcon

import (
	"time"

	"github.com/zMoooooritz/go-let-loose/internal/socket/api"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

func (r *Rcon) GetTempBans() ([]hll.ServerBan, error) {
	data, err := runCommand[api.GetTemporaryBans, api.RespTemporaryBans](r,
		api.GetTemporaryBans{},
	)
	if err != nil {
		return []hll.ServerBan{}, err
	}
	bans := []hll.ServerBan{}
	for _, ban := range data.BanList {
		timestamp, _ := time.Parse(time.RFC3339, ban.TimeOfBanning)
		bans = append(bans, hll.ServerBan{
			Player: hll.PlayerInfo{
				Name: ban.UserName,
				ID:   ban.UserID,
			},
			Reason:    ban.Reason,
			AdminName: ban.AdminName,
			Timestamp: timestamp,
			Duration:  time.Duration(ban.Duration) * time.Minute,
			Type:      hll.TempBan,
		})
	}
	return bans, nil
}

func (r *Rcon) GetPermaBans() ([]hll.ServerBan, error) {
	data, err := runCommand[api.GetTemporaryBans, api.RespTemporaryBans](r,
		api.GetTemporaryBans{},
	)
	if err != nil {
		return []hll.ServerBan{}, err
	}
	bans := []hll.ServerBan{}
	for _, ban := range data.BanList {
		timestamp, _ := time.Parse(time.RFC3339, ban.TimeOfBanning)
		bans = append(bans, hll.ServerBan{
			Player: hll.PlayerInfo{
				Name: ban.UserName,
				ID:   ban.UserID,
			},
			Reason:    ban.Reason,
			AdminName: ban.AdminName,
			Timestamp: timestamp,
			Duration:  time.Duration(ban.Duration) * time.Minute,
			Type:      hll.PermaBan,
		})
	}
	return bans, nil
}

func (r *Rcon) MessagePlayer(playerID string, message string) error {
	_, err := runCommand[api.MessagePlayer, any](r,
		api.MessagePlayer{
			PlayerID: playerID,
			Message:  message,
		},
	)
	return err
}

func (r *Rcon) PunishPlayer(player, reason string) error {
	_, err := runCommand[api.PunishPlayer, any](r,
		api.PunishPlayer{
			PlayerID: player,
			Reason:   reason,
		},
	)
	return err
}

func (r *Rcon) RemovePlayerFromPlatoon(player, reason string) error {
	_, err := runCommand[api.RemovePlayerFromPlatoon, any](r,
		api.RemovePlayerFromPlatoon{
			PlayerID: player,
			Reason:   reason,
		},
	)
	return err
}

func (r *Rcon) DisbandPlatoon(team hll.Team, unit hll.Unit, reason string) error {
	_, err := runCommand[api.DisbandPlatoon, any](r,
		api.DisbandPlatoon{
			TeamIndex:  int8(team.ToInt()),
			SquadIndex: int32(unit.ID),
			Reason:     reason,
		},
	)
	return err
}

func (r *Rcon) SwitchPlayerOnDeath(player string) error {
	_, err := runCommand[api.ForceTeamSwitch, any](r,
		api.ForceTeamSwitch{
			PlayerID:  player,
			ForceMode: 0, // 0 = Switch on Death, 1 = Switch Immediately
		},
	)
	return err
}

func (r *Rcon) SwitchPlayerNow(player string) error {
	_, err := runCommand[api.ForceTeamSwitch, any](r,
		api.ForceTeamSwitch{
			PlayerID:  player,
			ForceMode: 1, // 0 = Switch on Death, 1 = Switch Immediately
		},
	)
	return err
}

func (r *Rcon) KickPlayer(player, reason string) error {
	_, err := runCommand[api.KickPlayer, any](r,
		api.KickPlayer{
			PlayerID: player,
			Reason:   reason,
		},
	)
	return err
}

func (r *Rcon) TempBanPlayer(player string, duration int, reason, admin string) error {
	_, err := runCommand[api.TemporaryBanPlayer, any](r,
		api.TemporaryBanPlayer{
			PlayerID:  player,
			Reason:    reason,
			Duration:  int32(duration),
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
			PlayerID:  player,
			Reason:    reason,
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
