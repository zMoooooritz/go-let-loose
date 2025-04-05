package eventv2

import (
	"time"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

const (
	channel_size = 100
)

type EventType string

const (
	EVENT_TEAM_SWITCHED    EventType = "TEAM SWITCHED" // WARN: custom event; this event does have the player id
	EVENT_SQUAD_SWITCHED   EventType = "SQUAD SWITCHED"
	EVENT_SCORE_UPDATE     EventType = "SCORE UPDATE"
	EVENT_ROLE_CHANGED     EventType = "ROLE CHANGED"
	EVENT_LOADOUT_CHANGED  EventType = "LOADOUT CHANGED"
	EVENT_OBJECTIVE_CAPPED EventType = "OBJECTIVE CAPPED"
	EVENT_POSITION_CHANGED EventType = "POSITION CHANGED"
	EVENT_CLAN_TAG_CHANGED EventType = "CLAN TAG CHANGED"
	EVENT_GENERIC          EventType = "GENERIC"
)

type Event interface {
	Type() EventType
	Time() time.Time
}

type EventObserver interface {
	Notify(Event)
	IsSubscribed(Event) bool
}

type EventNotifier struct {
	observers map[EventObserver]struct{}
}

func NewEventNotifier() *EventNotifier {
	return &EventNotifier{
		observers: make(map[EventObserver]struct{}),
	}
}

func (n *EventNotifier) Register(o EventObserver) {
	n.observers[o] = struct{}{}
}

func (n *EventNotifier) Unregister(o EventObserver) {
	delete(n.observers, o)
}

func (n *EventNotifier) notify(e Event) {
	for observer := range n.observers {
		if observer.IsSubscribed(e) {
			observer.Notify(e)
		}
	}
}

type GenericEvent struct {
	eventType EventType
	time      time.Time
}

func (ge GenericEvent) Type() EventType {
	return ge.eventType
}

func (ge GenericEvent) Time() time.Time {
	return ge.time
}

type PlayerScoreUpdateEvent struct {
	GenericEvent
	Player   hll.PlayerInfo
	OldScore hll.Score
	NewScore hll.Score
}

type PlayerSwitchTeamEvent struct {
	GenericEvent
	Player  hll.PlayerInfo
	OldTeam hll.Team
	NewTeam hll.Team
}

type PlayerSwitchSquadEvent struct {
	GenericEvent
	Player   hll.PlayerInfo
	OldSquad hll.Unit
	NewSquad hll.Unit
}

type PlayerChangeRoleEvent struct {
	GenericEvent
	Player  hll.PlayerInfo
	OldRole hll.Role
	NewRole hll.Role
}

type PlayerChangeLoadoutEvent struct {
	GenericEvent
	Player     hll.PlayerInfo
	OldLoadout string
	NewLoadout string
}

type PlayerPositionChangedEvent struct {
	GenericEvent
	Player hll.PlayerInfo
	OldPos hll.Position
	NewPos hll.Position
}

type PlayerClanTagChangedEvent struct {
	GenericEvent
	Player     hll.PlayerInfo
	OldClanTag string
	NewClanTag string
}

func PlayerInfoDiffToEvents(oldData hll.DetailedPlayerInfo, newData hll.DetailedPlayerInfo) []Event {
	events := []Event{}

	emptyPlayerInfo := hll.DetailedPlayerInfo{}
	if oldData == emptyPlayerInfo {
		return events
	}

	if oldData.Team != newData.Team {
		events = append(events, PlayerSwitchTeamEvent{GenericEvent{EVENT_TEAM_SWITCHED, time.Now()}, newData.PlayerInfo, oldData.Team, newData.Team})
	}
	if oldData.Unit != newData.Unit {
		events = append(events, PlayerSwitchSquadEvent{GenericEvent{EVENT_SQUAD_SWITCHED, time.Now()}, newData.PlayerInfo, oldData.Unit, newData.Unit})
	}
	if oldData.Role != newData.Role {
		events = append(events, PlayerChangeRoleEvent{GenericEvent{EVENT_ROLE_CHANGED, time.Now()}, newData.PlayerInfo, oldData.Role, newData.Role})
	}
	if oldData.Loadout != newData.Loadout {
		events = append(events, PlayerChangeLoadoutEvent{GenericEvent{EVENT_LOADOUT_CHANGED, time.Now()}, newData.PlayerInfo, oldData.Loadout, newData.Loadout})
	}
	ds := hll.Score{}
	ds.Combat = max(newData.Score.Combat-oldData.Score.Combat, 0)
	ds.Defense = max(newData.Score.Defense-oldData.Score.Defense, 0)
	ds.Offense = max(newData.Score.Offense-oldData.Score.Offense, 0)
	ds.Support = max(newData.Score.Support-oldData.Score.Support, 0)
	if ds.Combat != 0 || ds.Support != 0 || ds.Offense != 0 || ds.Defense != 0 {
		events = append(events, PlayerScoreUpdateEvent{GenericEvent{EVENT_SCORE_UPDATE, time.Now()}, newData.PlayerInfo, oldData.Score, newData.Score})
	}
	if oldData.Position.IsActive() && newData.Position.IsActive() {
		if oldData.Position != newData.Position {
			events = append(events, PlayerPositionChangedEvent{GenericEvent{EVENT_POSITION_CHANGED, time.Now()}, newData.PlayerInfo, oldData.Position, newData.Position})
		}
	}
	if oldData.ClanTag != newData.ClanTag {
		events = append(events, PlayerClanTagChangedEvent{GenericEvent{EVENT_CLAN_TAG_CHANGED, time.Now()}, newData.PlayerInfo, oldData.ClanTag, newData.ClanTag})
	}
	return events
}

func GetAffectedPlayers(e Event) []hll.PlayerInfo {
	switch e := e.(type) {
	case PlayerSwitchTeamEvent:
		return []hll.PlayerInfo{e.Player}
	case PlayerSwitchSquadEvent:
		return []hll.PlayerInfo{e.Player}
	case PlayerScoreUpdateEvent:
		return []hll.PlayerInfo{e.Player}
	case PlayerChangeRoleEvent:
		return []hll.PlayerInfo{e.Player}
	case PlayerChangeLoadoutEvent:
		return []hll.PlayerInfo{e.Player}
	case PlayerPositionChangedEvent:
		return []hll.PlayerInfo{e.Player}
	case PlayerClanTagChangedEvent:
		return []hll.PlayerInfo{e.Player}
	}
	return []hll.PlayerInfo{}
}
