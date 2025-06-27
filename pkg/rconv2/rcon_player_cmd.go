package rconv2

import (
	"fmt"

	"github.com/zMoooooritz/go-let-loose/internal/socketv2/api"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

func (r *Rcon) GetPlayers() ([]hll.PlayerInfo, error) {
	players := []hll.PlayerInfo{}
	data, err := getPlayers(r)
	if err != nil {
		return players, err
	}
	for _, player := range data.Players {
		players = append(players, hll.PlayerInfo{
			Name: player.Name,
			ID:   player.ID,
		})
	}
	return players, nil
}

func (r *Rcon) GetPlayerNames() ([]string, error) {
	names := []string{}
	data, err := getPlayers(r)
	if err != nil {
		return names, err
	}
	for _, player := range data.Players {
		names = append(names, player.Name)
	}
	return names, nil
}

func (r *Rcon) GetPlayerIDs() ([]string, error) {
	playerIDs := []string{}
	data, err := getPlayers(r)
	if err != nil {
		return playerIDs, err
	}
	for _, player := range data.Players {
		playerIDs = append(playerIDs, player.ID)
	}
	return playerIDs, nil
}

func (r *Rcon) GetAdmins() ([]hll.Admin, error) {
	return []hll.Admin{}, fmt.Errorf("not implemented")
}

func (r *Rcon) GetAdminRoles() ([]hll.AdminRole, error) {
	return []hll.AdminRole{}, fmt.Errorf("not implemented")
}

func (r *Rcon) GetVIPs() ([]hll.PlayerInfo, error) {
	return []hll.PlayerInfo{}, fmt.Errorf("not implemented")
}

func (r *Rcon) GetPlayerInfo(playerID string) (hll.DetailedPlayerInfo, error) {
	return hll.DetailedPlayerInfo{}, fmt.Errorf("not implemented, BUG in official API")
	// data, err := getPlayer(r, playerID)
	// if err != nil {
	// 	return hll.DetailedPlayerInfo{}, err
	// }
	// return toDetailedPlayerInfo(data), nil
}

func (r *Rcon) GetPlayersInfo() ([]hll.DetailedPlayerInfo, error) {
	detailedPlayers := []hll.DetailedPlayerInfo{}
	data, err := getPlayers(r)
	if err != nil {
		return detailedPlayers, err
	}
	for _, player := range data.Players {
		detailedPlayers = append(detailedPlayers, toDetailedPlayerInfo(&player))
	}
	return detailedPlayers, nil
}

func (r *Rcon) AddAdmin(id, comment string, role hll.AdminRole) error {
	_, err := runCommand[api.AddAdmin, any](r,
		api.AddAdmin{
			PlayerId:   id,
			Comment:    comment,
			AdminGroup: string(role),
		},
	)
	return err
}

func (r *Rcon) RemoveAdmin(id string) error {
	_, err := runCommand[api.RemoveAdmin, any](r,
		api.RemoveAdmin{
			PlayerId: id,
		},
	)
	return err
}

func (r *Rcon) AddVip(id, comment string) error {
	_, err := runCommand[api.AddVipPlayer, any](r,
		api.AddVipPlayer{
			PlayerId:    id,
			Description: comment,
		},
	)
	return err
}

func (r *Rcon) RemoveVip(id string) error {
	_, err := runCommand[api.RemoveVipPlayer, any](r,
		api.RemoveVipPlayer{
			PlayerId: id,
		},
	)
	return err
}
