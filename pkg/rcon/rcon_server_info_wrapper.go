package rcon

import (
	"strings"

	"github.com/zMoooooritz/go-let-loose/internal/socket/api"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

func getPlayer(r *Rcon, playerID string) (*api.RespPlayerInformation, error) {
	resp, err := runCommand[api.GetServerInformation, api.RespPlayerInformation](
		r,
		api.GetServerInformation{
			Name:  api.ServerInfoPlayer,
			Value: playerID,
		},
	)
	if err != nil {
		return &api.RespPlayerInformation{}, err
	}
	return resp, nil
}

func getPlayers(r *Rcon) (*api.RespPlayersInformation, error) {
	resp, err := runCommand[api.GetServerInformation, api.RespPlayersInformation](
		r,
		api.GetServerInformation{
			Name: api.ServerInfoPlayers,
		},
	)
	if err != nil {
		return &api.RespPlayersInformation{}, err
	}
	return resp, nil
}

func getSessionInfo(r *Rcon) (*api.RespSessionInformation, error) {
	resp, err := runCommand[api.GetServerInformation, api.RespSessionInformation](
		r,
		api.GetServerInformation{
			Name: api.ServerInfoSession,
		},
	)
	if err != nil {
		return &api.RespSessionInformation{}, err
	}
	return resp, nil
}

func getMapRotation(r *Rcon) (*api.RespMapRotation, error) {
	resp, err := runCommand[api.GetServerInformation, api.RespMapRotation](
		r,
		api.GetServerInformation{
			Name: api.ServerInfoMapRotation,
		},
	)
	if err != nil {
		return &api.RespMapRotation{}, err
	}
	return resp, nil
}

func getMapSequence(r *Rcon) (*api.RespMapSequence, error) {
	resp, err := runCommand[api.GetServerInformation, api.RespMapSequence](
		r,
		api.GetServerInformation{
			Name: api.ServerInfoMapSequence,
		},
	)
	if err != nil {
		return &api.RespMapSequence{}, err
	}
	for i := range resp.Maps {
		resp.Maps[i].ID = strings.TrimPrefix(resp.Maps[i].ID, "/Game/Maps/")
	}
	return resp, nil
}

func getServerConfig(r *Rcon) (*api.RespServerConfiguration, error) {
	resp, err := runCommand[api.GetServerInformation, api.RespServerConfiguration](
		r,
		api.GetServerInformation{
			Name: api.ServerInfoServerConfig,
		},
	)
	if err != nil {
		return &api.RespServerConfiguration{}, err
	}
	return resp, nil
}

func getBannedWords(r *Rcon) (*api.RespBannedWords, error) {
	resp, err := runCommand[api.GetServerInformation, api.RespBannedWords](
		r,
		api.GetServerInformation{
			Name: api.ServerInfoBannedWords,
		},
	)
	if err != nil {
		return &api.RespBannedWords{}, err
	}
	return resp, nil
}

func getVipPlayers(r *Rcon) (*api.RespVipPlayers, error) {
	resp, err := runCommand[api.GetServerInformation, api.RespVipPlayers](
		r,
		api.GetServerInformation{
			Name: api.ServerInfoVipPlayers,
		},
	)
	if err != nil {
		return &api.RespVipPlayers{}, err
	}
	return resp, nil
}

func toDetailedPlayerInfo(player *api.RespPlayerInformation) hll.DetailedPlayerInfo {
	return hll.DetailedPlayerInfo{
		PlayerInfo: hll.PlayerInfo{
			Name: player.Name,
			ID:   player.ID,
		},
		ClanTag:  player.ClanTag,
		Platform: hll.PlayerPlatformFromString(player.Platform),
		Team:     hll.FactionFromInt(int(player.Team)).Team(),
		Faction:  hll.FactionFromInt(int(player.Team)),
		Role:     hll.RoleFromInt(int(player.Role)),
		Unit:     constructUnit(player.Platoon, int(player.Role)),
		Loadout:  player.Loadout,
		Kills:    int(player.Kills),
		Deaths:   int(player.Deaths),
		Score: hll.Score{
			Combat:  int(player.Score.Combat),
			Offense: int(player.Score.Offense),
			Defense: int(player.Score.Defense),
			Support: int(player.Score.Support),
		},
		Level: int(player.Level),
		Position: hll.Position{
			X: player.Position.X,
			Y: player.Position.Y,
			Z: player.Position.Z,
		},
	}
}

func constructUnit(playerPlatoon string, playerRole int) hll.Unit {
	role := hll.RoleFromInt(playerRole)

	unit := hll.Unit{}
	if playerPlatoon == "" {
		if role == hll.ArmyCommander {
			unit = hll.CommandUnit
		} else {
			unit = hll.NoUnit
		}
	} else {
		unit = hll.UnitFromString(playerPlatoon)
	}

	return unit
}
