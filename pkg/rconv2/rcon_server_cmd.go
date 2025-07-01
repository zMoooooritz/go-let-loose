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
	_, err := runCommand[api.SendServerMessage, any](r,
		api.SendServerMessage{
			Message: message,
		},
	)
	return err
}

func (r *Rcon) SetBroadcastMessage(message string) error {
	_, err := runCommand[api.ServerBroadcast, any](r,
		api.ServerBroadcast{
			Message: message,
		},
	)
	return err
}

func (r *Rcon) ClearBroadcastMessage(message string) error {
	return r.SetBroadcastMessage("")
}

func (r *Rcon) GetLogs(spanMins int) ([]string, error) {
	response, err := runCommand[api.GetAdminLog, api.ResponseAdminLog](r,
		api.GetAdminLog{
			LogBackTrackTime: int32(spanMins * 60),
			Filters:          "",
		},
	)
	logLines := []string{}
	if err == nil {
		for _, entry := range response.Entries {
			logLines = append(logLines, fmt.Sprintf("%s: %s", entry.Timestamp, entry.Message))
		}
	}
	return logLines, err
}

func (r *Rcon) GetLogEntries(seconds int, filters string) ([]hll.LogEntry, error) {
	response, err := runCommand[api.GetAdminLog, api.ResponseAdminLog](r,
		api.GetAdminLog{
			LogBackTrackTime: int32(seconds),
			Filters:          filters,
		},
	)
	logEntries := []hll.LogEntry{}
	if err == nil {
		for _, entry := range response.Entries {
			logEntries = append(logEntries, hll.LogEntry{
				Timestamp: entry.Time(),
				Message:   entry.Message,
			})
		}
	}
	return logEntries, err
}

func (r *Rcon) GetServerConfig() (hll.ServerConfig, error) {
	resp, err := getServerConfig(r)
	if err != nil {
		return hll.ServerConfig{}, err
	}
	platforms := []hll.SupportedPlatform{}
	for _, entry := range resp.SupportedPlatforms {
		platforms = append(platforms, hll.SupportedPlatformFromString(entry))
	}
	return hll.ServerConfig{
		Name:               resp.ServerName,
		BuildNumber:        resp.BuildNumber,
		BuildRevision:      resp.BuildRevision,
		SupportedPlatforms: platforms,
	}, nil
}

func (r *Rcon) GetSessionInfo() (hll.SessionInfo, error) {
	resp, err := getSessionInfo(r)
	if err != nil {
		return hll.SessionInfo{}, err
	}
	return hll.SessionInfo{
		ServerName:       resp.ServerName,
		MapName:          resp.MapName,
		GameMode:         hll.GameMode(resp.GameMode),
		MaxPlayerCount:   resp.MaxPlayerCount,
		PlayerCount:      resp.PlayerCount,
		MaxQueueCount:    resp.MaxQueueCount,
		QueueCount:       resp.QueueCount,
		MaxVIPQueueCount: resp.MaxVIPQueueCount,
		VIPQueueCount:    resp.VIPQueueCount,
	}, nil
}

func (r *Rcon) GetCommands() ([]hll.Command, error) {
	commands := []hll.Command{}
	resp, err := runCommand[api.GetDisplayableCommands, api.ResponseDisplayableCommands](r,
		api.GetDisplayableCommands{},
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

func (r *Rcon) GetCommandDetails(commandID string) (hll.CommandDetails, error) {
	resp, err := runCommand[api.GetClientReferenceData, api.ResponseClientReferenceData](r,
		api.GetClientReferenceData(commandID),
	)
	if err != nil {
		return hll.CommandDetails{}, err
	}
	parameters := []hll.DialogueParameter{}
	for _, param := range resp.DialogueParameters {
		parameters = append(parameters, hll.DialogueParameter{
			Type:          param.Type,
			Name:          param.Name,
			ID:            param.ID,
			DisplayMember: param.DisplayMember,
			ValueMember:   param.ValueMember,
		})
	}
	return hll.CommandDetails{
		Name:               resp.Name,
		Text:               resp.Text,
		Description:        resp.Description,
		DialogueParameters: parameters,
	}, nil
}
