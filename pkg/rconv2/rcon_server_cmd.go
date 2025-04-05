package rconv2

import (
	"fmt"

	"github.com/zMoooooritz/go-let-loose/internal/socketv2/api"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

func (r *Rcon) GetServerName() (string, error) {
	resp, err := getSessionInfo(r)
	return resp.ServerName, err
}

func (r *Rcon) GetSlots() (int, int, error) {
	resp, err := getSessionInfo(r)
	return resp.PlayerCount, resp.MaxPlayerCount, err
}

func (r *Rcon) GetGameState() (hll.GameState, error) {
	return hll.GameState{}, fmt.Errorf("not implemented")
}

func (r *Rcon) SetWelcomeMessage(message string) error {
	return fmt.Errorf("not implemented")
}

func (r *Rcon) SetBroadcastMessage(message string) error {
	_, err := runCommand[api.SetBroadcastMessage, any](r,
		api.SetBroadcastMessage{
			Message: message,
		},
	)
	return err
}

func (r *Rcon) ClearBroadcastMessage(message string) error {
	return r.SetBroadcastMessage("")
}

func (r *Rcon) GetLogs(spanMins int) ([]string, error) {
	return []string{}, fmt.Errorf("not implemented")
}

func (r *Rcon) GetServerInfo() (hll.ServerInfo, error) {
	resp, err := getServerConfig(r)
	if err != nil {
		return hll.ServerInfo{}, err
	}
	platforms := []hll.Platform{}
	for _, entry := range resp.SupportedPlatforms {
		platforms = append(platforms, hll.PlatformFromString(entry))
	}
	return hll.ServerInfo{
		Name:               resp.ServerName,
		BuildNumber:        resp.BuildNumber,
		BuildRevision:      resp.BuildRevision,
		SupportedPlatforms: platforms,
	}, nil
}

func (r *Rcon) GetCommands() ([]hll.Command, error) {
	commands := []hll.Command{}
	resp, err := runCommand[api.DisplayableCommands, api.ResponseDisplayableCommands](r,
		api.DisplayableCommands{},
	)
	if err != nil {
		return commands, err
	}
	for _, entry := range resp.Entries {
		commands = append(commands, hll.Command{
			ID:              entry.ID,
			Name:            entry.FriendlyName,
			ClientSupported: entry.IsClientSupported,
		})
	}
	return commands, nil
}
