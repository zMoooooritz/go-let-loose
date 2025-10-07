package rcon

import (
	"github.com/zMoooooritz/go-let-loose/internal/socket/api"
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
	resp, err := runCommand[api.GetAdminUsers, api.RespAdminUsers](r,
		api.GetAdminUsers{},
	)
	if err != nil {
		return []hll.Admin{}, err
	}
	admins := []hll.Admin{}
	for _, admin := range resp.AdminUsers {
		admins = append(admins,
			hll.Admin{
				PlayerInfo: hll.PlayerInfo{
					Name: "",
					ID:   admin.UserId,
				},
				Role:    hll.AdminRole(admin.Group),
				Comment: admin.Comment,
			},
		)
	}
	return admins, nil
}

func (r *Rcon) GetAdminRoles() ([]hll.AdminRole, error) {
	resp, err := runCommand[api.GetAdminGroups, api.RespAdminGroups](r,
		api.GetAdminGroups{},
	)
	if err != nil {
		return []hll.AdminRole{}, err
	}
	adminRoles := []hll.AdminRole{}
	for _, group := range resp.GroupNames {
		adminRoles = append(adminRoles, hll.AdminRole(group))
	}
	return adminRoles, nil
}

func (r *Rcon) GetVIPs() ([]hll.PlayerInfo, error) {
	vipPlayers := []hll.PlayerInfo{}
	data, err := getVipPlayers(r)
	if err != nil {
		return []hll.PlayerInfo{}, err
	}
	for _, playerID := range data.VipPlayerIDs {
		vipPlayers = append(vipPlayers, hll.PlayerInfo{
			Name: "",
			ID:   playerID,
		})
	}
	return vipPlayers, nil
}

func (r *Rcon) GetPlayerInfo(playerID string) (hll.DetailedPlayerInfo, error) {
	data, err := getPlayer(r, playerID)
	if err != nil {
		return hll.DetailedPlayerInfo{}, err
	}
	return toDetailedPlayerInfo(data), nil
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

func (r *Rcon) GetServerView() (hll.ServerView, error) {
	detailedPlayers, err := r.GetPlayersInfo()
	if err != nil {
		return hll.ServerView{}, err
	}
	return *hll.PlayersToServerView(detailedPlayers), nil
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
	_, err := runCommand[api.AddVip, any](r,
		api.AddVip{
			PlayerId:    id,
			Description: comment,
		},
	)
	return err
}

func (r *Rcon) RemoveVip(id string) error {
	_, err := runCommand[api.RemoveVip, any](r,
		api.RemoveVip{
			PlayerId: id,
		},
	)
	return err
}
