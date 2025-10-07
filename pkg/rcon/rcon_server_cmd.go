package rcon

import (
	"fmt"

	"github.com/zMoooooritz/go-let-loose/internal/socket/api"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

func (r *Rcon) GetServerName() (string, error) {
	resp, err := getSessionInfo(r)
	return resp.ServerName, err
}

func (r *Rcon) GetSlots() (int, int, error) {
	resp, err := getSessionInfo(r)
	return int(resp.PlayerCount), int(resp.MaxPlayerCount), err
}

func (r *Rcon) GetGameState() (hll.GameState, error) {
	resp, err := getSessionInfo(r)
	if err != nil {
		return hll.GameState{}, err
	}
	return hll.GameState{
		PlayerCount: hll.TeamData{
			Allies: int(resp.AlliedPlayerCount),
			Axis:   int(resp.AxisPlayerCount),
		},
		GameScore: hll.TeamData{
			Allies: int(resp.AlliedScore),
			Axis:   int(resp.AxisScore),
		},
		RemainingSeconds: int(resp.RemainingMatchTime),
		CurrentMap:       hll.Layer{},
		NextMap:          hll.Layer{},
	}, nil
}

func (r *Rcon) GetPlayerCounts() (hll.TeamData, error) {
	resp, err := getSessionInfo(r)
	if err != nil {
		return hll.TeamData{}, err
	}
	return hll.TeamData{
		Allies: int(resp.AlliedPlayerCount),
		Axis:   int(resp.AxisPlayerCount),
	}, nil
}

func (r *Rcon) GetScore() (hll.TeamData, error) {
	resp, err := getSessionInfo(r)
	if err != nil {
		return hll.TeamData{}, err
	}
	return hll.TeamData{
		Allies: int(resp.AlliedScore),
		Axis:   int(resp.AxisScore),
	}, nil
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
		PasswordProtected:  resp.PasswordProtected,
	}, nil
}

func (r *Rcon) GetSessionInfo() (hll.SessionInfo, error) {
	resp, err := getSessionInfo(r)
	if err != nil {
		return hll.SessionInfo{}, err
	}
	return hll.SessionInfo{
		ServerName:         resp.ServerName,
		MapName:            resp.MapName,
		GameMode:           hll.GameMode(resp.GameMode),
		RemainingMatchTime: int(resp.RemainingMatchTime),
		MatchTime:          int(resp.MatchTime),
		AlliedFaction:      hll.FactionFromInt(int(resp.AlliedFaction)),
		AxisFaction:        hll.FactionFromInt(int(resp.AxisFaction)),
		MaxPlayerCount:     int(resp.MaxPlayerCount),
		AlliedScore:        int(resp.AlliedScore),
		AxisScore:          int(resp.AxisScore),
		PlayerCount:        int(resp.PlayerCount),
		AlliedPlayerCount:  int(resp.AlliedPlayerCount),
		AxisPlayerCount:    int(resp.AxisPlayerCount),
		MaxQueueCount:      int(resp.MaxQueueCount),
		QueueCount:         int(resp.QueueCount),
		MaxVIPQueueCount:   int(resp.MaxVipQueueCount),
		VIPQueueCount:      int(resp.VipQueueCount),
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

func (r *Rcon) GetServerChangelist() (string, error) {
	resp, err := runCommand[api.GetServerChangelist, api.ResponseServerChangelist](r,
		api.GetServerChangelist{},
	)
	if err != nil {
		return "", err
	}
	return resp.Changelist, nil
}
