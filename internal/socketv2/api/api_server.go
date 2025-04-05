package api

type SetMaxQueuedPlayers struct {
	MaxQueuedPlayers int `json:"MaxQueuedPlayers"`
}

type SetBroadcastMessage struct {
	Message string `json:"Message"`
}

type DisplayableCommands struct {
}

type ResponseDisplayableCommands struct {
	Entries []DisplayableCommandEntry `json:"entries"`
}

type DisplayableCommandEntry struct {
	ID                string `json:"ID"`
	FriendlyName      string `json:"FriendlyName"`
	IsClientSupported bool   `json:"IsClientSupported"`
}
