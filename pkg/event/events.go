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
	logPattern           = regexp.MustCompile(`^\[.+? \((\d+)\)] (.*)`)
	connPattern          = regexp.MustCompile(`CONNECTED (.+) \((.+)\)`)
	killPattern          = regexp.MustCompile(`KILL: (.+)\(.+/(.+)\) -> (.+)\(.+/(.+)\) with (.+)`)
	switchPattern        = regexp.MustCompile(`TEAMSWITCH (.+) \((.+) > (.+)\)`)
	chatPattern          = regexp.MustCompile(`CHAT\[(.+)\]\[(.+)\((.+)/(.+)\)\]: (.+)`)
	banPattern           = regexp.MustCompile(`BAN: \[(.+)\].+\[(.+.+)\]`)
	kickPattern          = regexp.MustCompile(`KICK: \[(.+)\].+\[(.+.+)\]`)
	msgPattern           = regexp.MustCompile(`MESSAGE: player \[(.+)\((.+)\)\], content \[(.*)\]`)
	startPattern         = regexp.MustCompile(`MATCH START (.+)`)
	endPattern           = regexp.MustCompile(`MATCH ENDED \x60(.+)\x60.+\((\d) - (\d)\) AXIS`)
	camPattern           = regexp.MustCompile(`Player \[(.+) \((.+)\)\] (.+)`)
	voteStartedPattern   = regexp.MustCompile(`VOTESYS: Player \[(.*)\] Started a vote of type \((.*)\) against \[(.*)\]. VoteID: \[(\d+)\]`)
	voteSubmittedPattern = regexp.MustCompile(`VOTESYS: Player \[(.*)\] voted \[(.*)\] for VoteID\[(\d+)\]`)
	voteCompletePattern  = regexp.MustCompile(`VOTESYS: Vote \[(\d+)\] completed. Result: (.*)`)
)

var openVoteKicksMap = make(map[int]VoteStartedEvent)

type EventType string

const (
	EVENT_CONNECTED    EventType = "CONNECTED"
	EVENT_DISCONNECTED EventType = "DISCONNECTED"
	EVENT_KILL         EventType = "KILL"
	EVENT_DEATH        EventType = "DEATH"
	EVENT_TEAMKILL     EventType = "TEAM KILL"
	EVENT_TEAMDEATH    EventType = "TEAM DEATH"
	EVENT_TEAMSWITCH   EventType = "TEAMSWITCH" // WARN: generic hll event; this event does not have a player id
	EVENT_CHAT         EventType = "CHAT"
	EVENT_BAN          EventType = "BAN"
	EVENT_KICK         EventType = "KICK"
	EVENT_MESSAGE      EventType = "MESSAGE"
	EVENT_MATCHSTART   EventType = "MATCH START"
	EVENT_MATCHEND     EventType = "MATCH END"
	EVENT_ADMINCAM     EventType = "Player"
	event_vote         EventType = "VOTESYS"

	// native log events above; custom events below

	EVENT_VOTE_KICK_STARTED   EventType = "VOTE KICK STARTED"
	EVENT_VOTE_SUBMITTED      EventType = "VOTE SUBMITTED"
	EVENT_VOTE_KICK_COMPLETED EventType = "VOTE KICK COMPLETED"
	EVENT_TEAM_SWITCHED       EventType = "TEAM SWITCHED" // WARN: custom event; this event does have the player id
	EVENT_SQUAD_SWITCHED      EventType = "SQUAD SWITCHED"
	EVENT_SCORE_UPDATE        EventType = "SCORE UPDATE"
	EVENT_ROLE_CHANGED        EventType = "ROLE CHANGED"
	EVENT_LOADOUT_CHANGED     EventType = "LOADOUT_CHANGED"
	EVENT_OBJECTIVE_CAPPED    EventType = "OBJECTIVE_CAPPED"
	EVENT_GENERIC             EventType = "GENERIC"
)

var logEventParsers = map[EventType]func(time.Time, string) []Event{
	EVENT_CONNECTED:    logToConnectEvent,
	EVENT_DISCONNECTED: logToDisconnectEvent,
	EVENT_KILL:         logToKillEvents,
	EVENT_TEAMKILL:     logToTeamKillEvents,
	EVENT_TEAMSWITCH:   logToTeamSwitchEvent,
	EVENT_CHAT:         logToChatEvent,
	EVENT_BAN:          logToBanEvent,
	EVENT_KICK:         logToKickEvent,
	EVENT_MESSAGE:      logToMessageEvent,
	EVENT_MATCHSTART:   logToMatchStartEvent,
	EVENT_MATCHEND:     logToMatchEndEvent,
	EVENT_ADMINCAM:     logToAdminCamEvent,
	event_vote:         logToVoteEvents,
}

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

func logToConnectEvent(time time.Time, eventdata string) []Event {
	match := connPattern.FindStringSubmatch(eventdata)
	if len(match) < 3 {
		logger.Error("Event data unparseable:", eventdata)
		return []Event{}
	}
	return []Event{ConnectEvent{
		GenericEvent{EVENT_CONNECTED, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
	}}
}

type DisconnectEvent struct {
	GenericEvent
	Player hll.PlayerInfo
}

func logToDisconnectEvent(time time.Time, eventdata string) []Event {
	match := connPattern.FindStringSubmatch(eventdata)
	if len(match) < 3 {
		logger.Error("Event data unparseable:", eventdata)
		return []Event{}
	}
	return []Event{DisconnectEvent{
		GenericEvent{EVENT_DISCONNECTED, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
	}}
}

type KillEvent struct {
	GenericEvent
	Killer hll.PlayerInfo
	Victim hll.PlayerInfo
	Weapon hll.Weapon
}

func logToKillEvents(time time.Time, eventdata string) []Event {
	match := killPattern.FindStringSubmatch(eventdata)
	if len(match) < 6 {
		logger.Error("Event data unparseable:", eventdata)
		return []Event{}
	}
	killEvent := KillEvent{
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
	return []Event{killEvent, killToDeatchEvent(killEvent)}
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

func logToTeamKillEvents(time time.Time, eventdata string) []Event {
	match := killPattern.FindStringSubmatch(eventdata)
	if len(match) < 6 {
		logger.Error("Event data unparseable:", eventdata)
		return []Event{}
	}
	teamKillEvent := TeamKillEvent{
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
	return []Event{teamKillEvent, teamKillToTeamDeatchEvent(teamKillEvent)}
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

func logToTeamSwitchEvent(time time.Time, eventdata string) []Event {
	match := switchPattern.FindStringSubmatch(eventdata)
	if len(match) < 4 {
		logger.Error("Event data unparseable:", eventdata)
		return []Event{}
	}
	return []Event{TeamSwitchEvent{
		GenericEvent{EVENT_TEAMSWITCH, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   hll.NoPlayerID,
		},
		hll.Team(match[2]),
		hll.Team(match[3]),
	}}
}

type ChatEvent struct {
	GenericEvent
	Player  hll.PlayerInfo
	Team    hll.Team
	Scope   hll.ChatScope
	Message string
}

func logToChatEvent(time time.Time, eventdata string) []Event {
	match := chatPattern.FindStringSubmatch(eventdata)
	if len(match) < 6 {
		logger.Error("Event data unparseable:", eventdata)
		return []Event{}
	}
	return []Event{ChatEvent{
		GenericEvent{EVENT_CHAT, time},
		hll.PlayerInfo{
			Name: match[2],
			ID:   match[4],
		},
		hll.TeamFromString(match[3]),
		hll.ChatScopeFromString(match[1]),
		match[5],
	}}
}

type BanEvent struct {
	GenericEvent
	Player hll.PlayerInfo
	Reason string
}

func logToBanEvent(time time.Time, eventdata string) []Event {
	match := banPattern.FindStringSubmatch(eventdata)
	if len(match) < 3 {
		logger.Error("Event data unparseable:", eventdata)
		return []Event{}
	}
	return []Event{BanEvent{
		GenericEvent{EVENT_BAN, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   hll.NoPlayerID,
		},
		match[2],
	}}
}

type KickEvent struct {
	GenericEvent
	Player hll.PlayerInfo
	Reason string
}

func logToKickEvent(time time.Time, eventdata string) []Event {
	match := kickPattern.FindStringSubmatch(eventdata)
	if len(match) < 3 {
		logger.Error("Event data unparseable:", eventdata)
		return []Event{}
	}
	return []Event{KickEvent{
		GenericEvent{EVENT_KICK, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   hll.NoPlayerID,
		},
		match[2],
	}}
}

type MessageEvent struct {
	GenericEvent
	Player  hll.PlayerInfo
	Message string
}

func logToMessageEvent(time time.Time, eventdata string) []Event {
	match := msgPattern.FindStringSubmatch(eventdata)
	if len(match) < 4 {
		logger.Error("Event data unparseable:", eventdata)
		return []Event{}
	}
	return []Event{MessageEvent{
		GenericEvent{EVENT_MESSAGE, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
		match[3],
	}}
}

type MatchStartEvent struct {
	GenericEvent
	Map hll.GameMap
}

func logToMatchStartEvent(time time.Time, eventdata string) []Event {
	match := startPattern.FindStringSubmatch(eventdata)
	if len(match) < 2 {
		logger.Error("Event data unparseable:", eventdata)
		return []Event{}
	}
	return []Event{MatchStartEvent{
		GenericEvent{EVENT_MATCHSTART, time},
		hll.LogMapNameToMap(match[1]),
	}}
}

type MatchEndEvent struct {
	GenericEvent
	Map   hll.GameMap
	Score hll.TeamData
}

func logToMatchEndEvent(time time.Time, eventdata string) []Event {
	match := endPattern.FindStringSubmatch(eventdata)
	if len(match) < 4 {
		logger.Error("Event data unparseable:", eventdata)
		return []Event{}
	}
	return []Event{MatchEndEvent{
		GenericEvent{EVENT_MATCHEND, time},
		hll.LogMapNameToMap(match[1]),
		hll.TeamData{
			Allies: util.ToInt(match[2]),
			Axis:   util.ToInt(match[3]),
		},
	}}
}

type AdminCamEvent struct {
	GenericEvent
	Player  hll.PlayerInfo
	Content string
}

func logToAdminCamEvent(time time.Time, eventdata string) []Event {
	match := camPattern.FindStringSubmatch(eventdata)
	if len(match) < 4 {
		logger.Error("Event data unparseable:", eventdata)
		return []Event{}
	}
	return []Event{AdminCamEvent{
		GenericEvent{EVENT_ADMINCAM, time},
		hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
		match[3],
	}}
}

type VoteStartedEvent struct {
	GenericEvent
	Reason    string
	ID        int
	Initiator hll.PlayerInfo
	Target    hll.PlayerInfo
}

type VoteSubmittedEvent struct {
	GenericEvent
	Submitter hll.PlayerInfo
	ID        int
	Vote      string
}

type VoteCompletedEvent struct {
	GenericEvent
	Reason    string
	Result    string
	ID        int
	Initiator hll.PlayerInfo
	Target    hll.PlayerInfo
}

func logToVoteEvents(time time.Time, eventdata string) []Event {
	events := []Event{}
	if match := voteSubmittedPattern.FindStringSubmatch(eventdata); len(match) > 3 {
		events = append(events,
			VoteSubmittedEvent{
				GenericEvent: GenericEvent{EVENT_VOTE_SUBMITTED, time},
				Submitter: hll.PlayerInfo{
					Name: match[1],
					ID:   hll.NoPlayerID,
				},
				ID:   util.ToInt(match[3]),
				Vote: match[2],
			},
		)
	} else if match := voteStartedPattern.FindStringSubmatch(eventdata); len(match) > 4 {
		voteStartEvent := VoteStartedEvent{
			GenericEvent: GenericEvent{EVENT_VOTE_KICK_STARTED, time},
			Reason:       match[2],
			ID:           util.ToInt(match[4]),
			Initiator: hll.PlayerInfo{
				Name: match[1],
				ID:   hll.NoPlayerID,
			},
			Target: hll.PlayerInfo{
				Name: match[3],
				ID:   hll.NoPlayerID,
			},
		}

		openVoteKicksMap[voteStartEvent.ID] = voteStartEvent
		events = append(events, voteStartEvent)
	} else if match := voteCompletePattern.FindStringSubmatch(eventdata); len(match) > 2 {
		voteID := util.ToInt(match[1])

		if voteStartEvent, ok := openVoteKicksMap[voteID]; ok {
			events = append(events,
				VoteCompletedEvent{
					GenericEvent: GenericEvent{EVENT_VOTE_KICK_COMPLETED, time},
					Reason:       voteStartEvent.Reason,
					Result:       match[2],
					ID:           voteStartEvent.ID,
					Initiator:    voteStartEvent.Initiator,
					Target:       voteStartEvent.Target,
				},
			)
		}
		delete(openVoteKicksMap, voteID)
	}
	return events
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
		if newData.GameScore.Axis == newData.GameScore.Allies { // game just started
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
	timestamp := time.Unix(util.ToInt64(match[1]), 0)
	data := match[2]

	for eventPrefix, parser := range logEventParsers {
		if strings.HasPrefix(data, string(eventPrefix)) {
			return parser(timestamp, data)
		}
	}

	logger.Error("Logline unparseable:", logline)
	return []Event{}
}
