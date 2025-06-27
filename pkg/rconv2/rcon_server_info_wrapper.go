package rconv2

import (
	"strings"

	"github.com/zMoooooritz/go-let-loose/internal/socketv2/api"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	caser = cases.Title(language.AmericanEnglish)
)

//nolint:all
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

func toDetailedPlayerInfo(player *api.RespPlayerInformation) hll.DetailedPlayerInfo {
	return hll.DetailedPlayerInfo{
		PlayerInfo: hll.PlayerInfo{
			Name: player.Name,
			ID:   player.ID,
		},
		ClanTag:  player.ClanTag,
		Platform: hll.PlayerPlatformFromString(player.Platform),
		Team:     hll.FactionFromInt(player.Team).Team(),
		Faction:  hll.FactionFromInt(player.Team),
		Role:     hll.RoleFromInt(player.Role),
		Unit:     constructUnit(player.Platoon, player.Role),
		Loadout:  player.Loadout,
		Kills:    player.Kills,
		Deaths:   player.Deaths,
		Score: hll.Score{
			Combat:  player.Score.Combat,
			Offense: player.Score.Offense,
			Defense: player.Score.Defense,
			Support: player.Score.Support,
		},
		Level: player.Level,
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
		unit.Name = caser.String(playerPlatoon)
		unit.ID = hll.UnitNameToID(playerPlatoon)
	}

	return unit
}
