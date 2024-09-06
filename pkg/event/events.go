package event

import (
	"regexp"
	"strings"
	"time"

	"github.com/zMoooooritz/go-let-loose/internal/util"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
)

const (
	channel_size = 100
)

var (
	logPattern    = regexp.MustCompile(`^\[.+? \((\d+)\)] (.*)`)
	connPattern   = regexp.MustCompile(`CONNECTED (.+) \((.+)\)`)
	killPattern   = regexp.MustCompile(`KILL: (.+)\(.+/(.+)\) -> (.+)\(.+/(.+)\) with (.+)`)
	switchPattern = regexp.MustCompile(`TEAMSWITCH (.+) \((.+) > (.+)\)`)
	chatPattern   = regexp.MustCompile(`CHAT\[.+\]\[(.+)\(.+/(.+)\)\]: (.+)`)
	banPattern    = regexp.MustCompile(`BAN: \[(.+)\].+\[(.+.+)\]`)
	kickPattern   = regexp.MustCompile(`KICK: \[(.+)\].+\[(.+.+)\]`)
	msgPattern    = regexp.MustCompile(`MESSAGE: player \[(.+)\((.+)\)\], content \[(.*)\]`)
	startPattern  = regexp.MustCompile(`MATCH START (.+)`)
	endPattern    = regexp.MustCompile(`MATCH ENDED \x60(.+)\x60.+\((\d) - (\d)\) AXIS`)
	camPattern    = regexp.MustCompile(`Player \[(.+) \((.+)\)\] (.+)`)
)

type EventType string

const (
	EVENT_CONNECTED        EventType = "CONNECTED"
	EVENT_DISCONNECTED     EventType = "DISCONNECTED"
	EVENT_KILL             EventType = "KILL"
	EVENT_DEATH            EventType = "DEATH"
	EVENT_TEAMKILL         EventType = "TEAM KILL"
	EVENT_TEAMDEATH        EventType = "TEAM DEATH"
	EVENT_TEAMSWITCH       EventType = "TEAMSWITCH" // WARN: generic hll event; this event does not have a player id
	EVENT_CHAT             EventType = "CHAT"
	EVENT_BAN              EventType = "BAN"
	EVENT_KICK             EventType = "KICK"
	EVENT_MESSAGE          EventType = "MESSAGE"
	EVENT_MATCHSTART       EventType = "MATCH START"
	EVENT_MATCHEND         EventType = "MATCH END"
	EVENT_ADMINCAM         EventType = "Player"
	EVENT_TEAM_SWITCHED    EventType = "TEAM SWITCHED" // WARN: custom event; this event does have the player id
	EVENT_SQUAD_SWITCHED   EventType = "SQUAD SWITCHED"
	EVENT_SCORE_UPDATE     EventType = "SCORE UPDATE"
	EVENT_ROLE_CHANGED     EventType = "ROLE CHANGED"
	EVENT_LOADOUT_CHANGED  EventType = "LOADOUT_CHANGED"
	EVENT_OBJECTIVE_CAPPED EventType = "OBJECTIVE_CAPPED"
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

type ConnectEvent struct {
	GenericEvent
	Player hll.PlayerInfo
}

func logToConnectEvent(time time.Time, eventdata string) ConnectEvent {
	match := connPattern.FindStringSubmatch(eventdata)
	if len(match) < 3 {
		return ConnectEvent{}
	}
	return ConnectEvent{
		GenericEvent{EVENT_CONNECTED, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
	}
}

type DisconnectEvent struct {
	GenericEvent
	Player hll.PlayerInfo
}

func logToDisconnectEvent(time time.Time, eventdata string) DisconnectEvent {
	match := connPattern.FindStringSubmatch(eventdata)
	if len(match) < 3 {
		return DisconnectEvent{}
	}
	return DisconnectEvent{
		GenericEvent{EVENT_DISCONNECTED, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
	}
}

type KillEvent struct {
	GenericEvent
	Killer hll.PlayerInfo
	Victim hll.PlayerInfo
	Weapon hll.Weapon
}

func logToKillEvent(time time.Time, eventdata string) KillEvent {
	match := killPattern.FindStringSubmatch(eventdata)
	if len(match) < 6 {
		return KillEvent{}
	}
	return KillEvent{
		GenericEvent{EVENT_KILL, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
		hll.PlayerInfo{
			Name: match[3],
			ID:   match[4],
		},
		hll.ParseWeapon(match[5]),
	}
}

type DeathEvent struct {
	GenericEvent
	Victim hll.PlayerInfo
	Killer hll.PlayerInfo
	Weapon hll.Weapon
}

func killToDeatchEvent(killEvent KillEvent) DeathEvent {
	return DeathEvent{
		GenericEvent{EVENT_DEATH, killEvent.time},
		killEvent.Victim,
		killEvent.Killer,
		killEvent.Weapon,
	}
}

type TeamKillEvent struct {
	GenericEvent
	Killer hll.PlayerInfo
	Victim hll.PlayerInfo
	Weapon hll.Weapon
}

func logToTeamKillEvent(time time.Time, eventdata string) TeamKillEvent {
	match := killPattern.FindStringSubmatch(eventdata)
	if len(match) < 6 {
		return TeamKillEvent{}
	}
	return TeamKillEvent{
		GenericEvent{EVENT_TEAMKILL, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
		hll.PlayerInfo{
			Name: match[3],
			ID:   match[4],
		},
		hll.ParseWeapon(match[5]),
	}
}

type TeamDeathEvent struct {
	GenericEvent
	Victim hll.PlayerInfo
	Killer hll.PlayerInfo
	Weapon hll.Weapon
}

func teamKillToTeamDeatchEvent(teamKillEvent TeamKillEvent) TeamDeathEvent {
	return TeamDeathEvent{
		GenericEvent{EVENT_TEAMDEATH, teamKillEvent.time},
		teamKillEvent.Victim,
		teamKillEvent.Killer,
		teamKillEvent.Weapon,
	}
}

type TeamSwitchEvent struct {
	GenericEvent
	Player hll.PlayerInfo
	From   hll.Team
	To     hll.Team
}

func logToTeamSwitchEvent(time time.Time, eventdata string) TeamSwitchEvent {
	match := switchPattern.FindStringSubmatch(eventdata)
	if len(match) < 4 {
		return TeamSwitchEvent{}
	}
	return TeamSwitchEvent{
		GenericEvent{EVENT_TEAMSWITCH, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   hll.NoPlayerID,
		},
		hll.Team(match[2]),
		hll.Team(match[3]),
	}
}

type ChatEvent struct {
	GenericEvent
	Player  hll.PlayerInfo
	Message string
}

func logToChatEvent(time time.Time, eventdata string) ChatEvent {
	match := chatPattern.FindStringSubmatch(eventdata)
	if len(match) < 4 {
		return ChatEvent{}
	}
	return ChatEvent{
		GenericEvent{EVENT_CHAT, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
		match[3],
	}
}

type BanEvent struct {
	GenericEvent
	Player hll.PlayerInfo
	Reason string
}

func logToBanEvent(time time.Time, eventdata string) BanEvent {
	match := banPattern.FindStringSubmatch(eventdata)
	if len(match) < 3 {
		return BanEvent{}
	}
	return BanEvent{
		GenericEvent{EVENT_BAN, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   hll.NoPlayerID,
		},
		match[2],
	}
}

type KickEvent struct {
	GenericEvent
	Player hll.PlayerInfo
	Reason string
}

func logToKickEvent(time time.Time, eventdata string) KickEvent {
	match := kickPattern.FindStringSubmatch(eventdata)
	if len(match) < 3 {
		return KickEvent{}
	}
	return KickEvent{
		GenericEvent{EVENT_KICK, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   hll.NoPlayerID,
		},
		match[2],
	}
}

type MessageEvent struct {
	GenericEvent
	Player  hll.PlayerInfo
	Message string
}

func logToMessageEvent(time time.Time, eventdata string) MessageEvent {
	match := msgPattern.FindStringSubmatch(eventdata)
	if len(match) < 4 {
		return MessageEvent{}
	}
	return MessageEvent{
		GenericEvent{EVENT_MESSAGE, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
		match[3],
	}
}

type MatchStartEvent struct {
	GenericEvent
	Map hll.GameMap
}

func logToMatchStartEvent(time time.Time, eventdata string) MatchStartEvent {
	match := startPattern.FindStringSubmatch(eventdata)
	if len(match) < 2 {
		return MatchStartEvent{}
	}
	return MatchStartEvent{
		GenericEvent{EVENT_MATCHSTART, time},
		hll.LogMapNameToMap(match[1]),
	}
}

type MatchEndEvent struct {
	GenericEvent
	Map   hll.GameMap
	Score hll.TeamData
}

func logToMatchEndEvent(time time.Time, eventdata string) MatchEndEvent {
	match := endPattern.FindStringSubmatch(eventdata)
	if len(match) < 4 {
		return MatchEndEvent{}
	}
	return MatchEndEvent{
		GenericEvent{EVENT_MATCHEND, time},
		hll.LogMapNameToMap(match[1]),
		hll.TeamData{
			Allies: util.ToInt(match[2]),
			Axis:   util.ToInt(match[3]),
		},
	}
}

type AdminCamEvent struct {
	GenericEvent
	Player  hll.PlayerInfo
	Content string
}

func logToAdminCamEvent(time time.Time, eventdata string) AdminCamEvent {
	match := camPattern.FindStringSubmatch(eventdata)
	if len(match) < 4 {
		return AdminCamEvent{}
	}
	return AdminCamEvent{
		GenericEvent{EVENT_ADMINCAM, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
		match[3],
	}
}

type ObjectiveCaptureEvent struct {
	GenericEvent
	OldScore hll.TeamData
	NewScore hll.TeamData
}

type PlayerScoreUpdateEvent struct {
	GenericEvent
	Player      hll.PlayerInfo
	ScoreChange hll.Score
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

func PlayerInfoDiffToEvents(oldData hll.DetailedPlayerInfo, newData hll.DetailedPlayerInfo) []Event {
	events := []Event{}

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
		events = append(events, PlayerScoreUpdateEvent{GenericEvent{EVENT_SCORE_UPDATE, time.Now()}, newData.PlayerInfo, ds})
	}
	return events
}

func GameStateDiffToEvents(oldData hll.GameState, newData hll.GameState) []Event {
	events := []Event{}

	emptyGameState := hll.GameState{}
	if oldData == emptyGameState {
		return events
	}

	if oldData.GameScore != newData.GameScore {
		if newData.GameScore.Axis == newData.PlayerCount.Allies { // game just started
			return events
		}
		events = append(events, ObjectiveCaptureEvent{GenericEvent{EVENT_OBJECTIVE_CAPPED, time.Now()}, oldData.GameScore, newData.GameScore})
	}
	return events
}

func logToEvents(logline string) []Event {
	match := logPattern.FindStringSubmatch(logline)
	if len(match) < 3 {
		logger.Error("Logline invalid format:", logline)
		return []Event{GenericEvent{EVENT_GENERIC, time.Now()}}
	}
	time := time.Unix(util.ToInt64(match[1]), 0)
	data := match[2]

	if strings.HasPrefix(data, string(EVENT_CONNECTED)) {
		return []Event{logToConnectEvent(time, data)}
	} else if strings.HasPrefix(data, string(EVENT_DISCONNECTED)) {
		return []Event{logToDisconnectEvent(time, data)}
	} else if strings.HasPrefix(data, string(EVENT_KILL)) {
		killEvent := logToKillEvent(time, data)
		deathEvent := killToDeatchEvent(killEvent)
		return []Event{killEvent, deathEvent}
	} else if strings.HasPrefix(data, string(EVENT_TEAMKILL)) {
		teamKillEvent := logToTeamKillEvent(time, data)
		teamDeathEvent := teamKillToTeamDeatchEvent(teamKillEvent)
		return []Event{teamKillEvent, teamDeathEvent}
	} else if strings.HasPrefix(data, string(EVENT_TEAMSWITCH)) {
		return []Event{logToTeamSwitchEvent(time, data)}
	} else if strings.HasPrefix(data, string(EVENT_CHAT)) {
		return []Event{logToChatEvent(time, data)}
	} else if strings.HasPrefix(data, string(EVENT_BAN)) {
		return []Event{logToBanEvent(time, data)}
	} else if strings.HasPrefix(data, string(EVENT_KICK)) {
		return []Event{logToKickEvent(time, data)}
	} else if strings.HasPrefix(data, string(EVENT_MESSAGE)) {
		return []Event{logToMessageEvent(time, data)}
	} else if strings.HasPrefix(data, string(EVENT_MATCHSTART)) {
		return []Event{logToMatchStartEvent(time, data)}
	} else if strings.HasPrefix(data, string(EVENT_MATCHEND)) {
		return []Event{logToMatchEndEvent(time, data)}
	} else if strings.HasPrefix(data, string(EVENT_ADMINCAM)) {
		return []Event{logToAdminCamEvent(time, data)}
	}

	logger.Error("Logline unparseable:", logline)
	return []Event{GenericEvent{EVENT_GENERIC, time}}
}

func GetAffectedPlayer(e Event) (hll.PlayerInfo, bool) {
	switch e := e.(type) {
	case KillEvent:
		return e.Killer, true
	case DeathEvent:
		return e.Victim, true
	case TeamKillEvent:
		return e.Killer, true
	case TeamDeathEvent:
		return e.Victim, true
	case ConnectEvent:
		return e.Player, true
	case DisconnectEvent:
		return e.Player, true
	case TeamSwitchEvent:
		return e.Player, true
	case ChatEvent:
		return e.Player, true
	case BanEvent:
		return e.Player, true
	case KickEvent:
		return e.Player, true
	case MessageEvent:
		return e.Player, true
	case AdminCamEvent:
		return e.Player, true
	case PlayerSwitchTeamEvent:
		return e.Player, true
	case PlayerSwitchSquadEvent:
		return e.Player, true
	case PlayerScoreUpdateEvent:
		return e.Player, true
	case PlayerChangeRoleEvent:
		return e.Player, true
	case PlayerChangeLoadoutEvent:
		return e.Player, true
	}
	return hll.PlayerInfo{}, false
}
