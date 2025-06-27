package api

import (
	"strconv"
	"strings"
	"time"
)

type SetMaxQueuedPlayers struct {
	MaxQueuedPlayers int `json:"MaxQueuedPlayers"`
}

type ServerBroadcast struct {
	Message string `json:"Message"`
}

type SendServerMessage struct {
	Message string `json:"Message"`
}

type GetAdminLog struct {
}

type ResponseAdminLog struct {
	Entries []AdminLogEntry `json:"entries"`
}

type AdminLogEntry struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}

func (a AdminLogEntry) Time() time.Time {
	li := strings.LastIndex(a.Timestamp, ":")
	p, _ := strconv.Atoi(a.Timestamp[li+1:])
	r, _ := time.Parse("2006.01.02-15:04:05", a.Timestamp[:li])
	return time.Date(r.Year(), r.Month(), r.Day(), r.Hour(), r.Minute(), r.Second(), p*1000000, r.Location())
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
	ID            string `json:"ID"`
	DisplayMember string `json:"DisplayMember"`
	ValueMember   string `json:"ValueMember"`
}
