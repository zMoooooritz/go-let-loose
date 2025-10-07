package api

import (
	"time"
)

type SetMaxQueuedPlayers struct {
	MaxQueuedPlayers int32 `json:"MaxQueuedPlayers"`
}

type ServerBroadcast struct {
	Message string `json:"Message"`
}

type SendServerMessage struct {
	Message string `json:"Message"`
}

type GetAdminLog struct {
	LogBackTrackTime int32  `json:"LogBackTrackTime"` // in seconds
	Filters          string `json:"Filters"`
}

type ResponseAdminLog struct {
	Entries []AdminLogEntry `json:"Entries"`
}

type AdminLogEntry struct {
	Timestamp string `json:"Timestamp"`
	Message   string `json:"Message"`
}

func (a AdminLogEntry) Time() time.Time {
	t, err := time.Parse(time.RFC3339, a.Timestamp)
	if err != nil {
		return time.Time{}
	}
	return t
}

type GetServerChangelist struct {
}

type ResponseServerChangelist struct {
	Changelist string `json:"Changelist"`
}

type GetDisplayableCommands struct {
}

type ResponseDisplayableCommands struct {
	Entries []DisplayableCommandEntry `json:"entries"`
}

type DisplayableCommandEntry struct {
	ID                string `json:"ID"`
	FriendlyName      string `json:"FriendlyName"`
	IsClientSupported bool   `json:"IsClientSupported"`
}

type GetClientReferenceData string

type ResponseClientReferenceData struct {
	Name               string              `json:"Name"`
	Text               string              `json:"Text"`
	Description        string              `json:"Description"`
	DialogueParameters []DialogueParameter `json:"DialogueParameters"`
}

type DialogueParameter struct {
	Type          string `json:"Type"`
	Name          string `json:"Name"`
	ID            string `json:"Id"`
	DisplayMember string `json:"DisplayMember"`
	ValueMember   string `json:"ValueMember"`
}
