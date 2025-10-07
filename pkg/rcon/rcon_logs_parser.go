package rcon

import (
	"regexp"
	"strings"
	"time"

	"github.com/zMoooooritz/go-let-loose/internal/util"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
)

var (
	logPattern  = regexp.MustCompile(`(?s)^\[.+? \((\d+)\)] (.*)`)
	connPattern = regexp.MustCompile(`CONNECTED (.+) \((.+)\)`)
	killPattern = regexp.MustCompile(`KILL: (.+)\(.+/(.+)\) -> (.+)\(.+/(.+)\) with (.+)`)
	// switchPattern        = regexp.MustCompile(`TEAMSWITCH (.+) \((.+) > (.+)\)`)
	chatPattern          = regexp.MustCompile(`(?s)CHAT\[(.+)\]\[(.+)\((.+)/(.+)\)\]: (.+)`)
	banPattern           = regexp.MustCompile(`BAN: \[(.+)\].+\[(.+.+)\]`)
	kickPattern          = regexp.MustCompile(`KICK: \[(.+)\].+\[(.+.+)\]`)
	msgPattern           = regexp.MustCompile(`(?s)MESSAGE: player \[(.+)\((.+)\)\], content \[(.*)\]`)
	startPattern         = regexp.MustCompile(`MATCH START (.+)`)
	endPattern           = regexp.MustCompile(`MATCH ENDED \x60(.+)\x60.+\((\d) - (\d)\) AXIS`)
	camPattern           = regexp.MustCompile(`Player \[(.+) \((.+)\)\] (.+)`)
	voteStartedPattern   = regexp.MustCompile(`VOTESYS: Player \[(.*)\] Started a vote of type \((.*)\) against \[(.*)\]. VoteID: \[(\d+)\]`)
	voteSubmittedPattern = regexp.MustCompile(`VOTESYS: Player \[(.*)\] voted \[(.*)\] for VoteID\[(\d+)\]`)
	voteCompletePattern  = regexp.MustCompile(`VOTESYS: Vote \[(\d+)\] completed. Result: (.*)`)
)

var openVoteKicksMap = make(map[int]hll.VoteStartedEvent)

const (
	event_admincam   hll.EventType = "Player"
	event_vote       hll.EventType = "VOTESYS"
	event_teamswitch hll.EventType = "TEAMSWITCH" // WARN: generic hll event; this event does not have a player id
)

var logEventParsers = map[hll.EventType]func(time.Time, string) []hll.Event{
	hll.EVENT_CONNECTED:    logToConnectEvent,
	hll.EVENT_DISCONNECTED: logToDisconnectEvent,
	hll.EVENT_KILL:         logToKillEvents,
	hll.EVENT_TEAMKILL:     logToTeamKillEvents,
	hll.EVENT_CHAT:         logToChatEvent,
	hll.EVENT_BAN:          logToBanEvent,
	hll.EVENT_KICK:         logToKickEvent,
	hll.EVENT_MESSAGE:      logToMessageEvent,
	hll.EVENT_MATCHSTART:   logToMatchStartEvent,
	hll.EVENT_MATCHEND:     logToMatchEndEvent,
	event_admincam:         logToAdminCamEvent,
	event_vote:             logToVoteEvents,
	event_teamswitch:       logToTeamSwitchEvent,
}

func logToConnectEvent(time time.Time, eventdata string) []hll.Event {
	match := connPattern.FindStringSubmatch(eventdata)
	if len(match) < 3 {
		logger.Error("Event data unparseable:", eventdata)
		return []hll.Event{}
	}
	return []hll.Event{hll.ConnectEvent{
		GenericEvent: hll.GenericEvent{
			EventType: hll.EVENT_CONNECTED,
			EventTime: time,
		},
		Player: hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
	}}
}

func logToDisconnectEvent(time time.Time, eventdata string) []hll.Event {
	match := connPattern.FindStringSubmatch(eventdata)
	if len(match) < 3 {
		logger.Error("Event data unparseable:", eventdata)
		return []hll.Event{}
	}
	return []hll.Event{hll.DisconnectEvent{
		GenericEvent: hll.GenericEvent{
			EventType: hll.EVENT_DISCONNECTED,
			EventTime: time,
		},
		Player: hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
	}}
}

func logToKillEvents(time time.Time, eventdata string) []hll.Event {
	match := killPattern.FindStringSubmatch(eventdata)
	if len(match) < 6 {
		logger.Error("Event data unparseable:", eventdata)
		return []hll.Event{}
	}
	killEvent := hll.KillEvent{
		GenericEvent: hll.GenericEvent{
			EventType: hll.EVENT_KILL,
			EventTime: time,
		},
		Killer: hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
		Victim: hll.PlayerInfo{
			Name: match[3],
			ID:   match[4],
		},
		Weapon: hll.ParseWeapon(match[5]),
	}
	return []hll.Event{killEvent, killToDeatchEvent(killEvent)}
}

func killToDeatchEvent(killEvent hll.KillEvent) hll.DeathEvent {
	return hll.DeathEvent{
		GenericEvent: hll.GenericEvent{
			EventType: hll.EVENT_DEATH,
			EventTime: killEvent.EventTime,
		},
		Victim: killEvent.Victim,
		Killer: killEvent.Killer,
		Weapon: killEvent.Weapon,
	}
}

func logToTeamKillEvents(time time.Time, eventdata string) []hll.Event {
	match := killPattern.FindStringSubmatch(eventdata)
	if len(match) < 6 {
		logger.Error("Event data unparseable:", eventdata)
		return []hll.Event{}
	}
	teamKillEvent := hll.TeamKillEvent{
		GenericEvent: hll.GenericEvent{
			EventType: hll.EVENT_TEAMKILL,
			EventTime: time,
		},
		Killer: hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
		Victim: hll.PlayerInfo{
			Name: match[3],
			ID:   match[4],
		},
		Weapon: hll.ParseWeapon(match[5]),
	}
	return []hll.Event{teamKillEvent, teamKillToTeamDeathEvent(teamKillEvent)}
}

func teamKillToTeamDeathEvent(teamKillEvent hll.TeamKillEvent) hll.TeamDeathEvent {
	return hll.TeamDeathEvent{
		GenericEvent: hll.GenericEvent{
			EventType: hll.EVENT_TEAMDEATH,
			EventTime: teamKillEvent.EventTime,
		},
		Victim: teamKillEvent.Victim,
		Killer: teamKillEvent.Killer,
		Weapon: teamKillEvent.Weapon,
	}
}

func logToTeamSwitchEvent(time time.Time, eventdata string) []hll.Event {
	// Dont use since a custom event with the player id is implemented
	return []hll.Event{}

	// match := switchPattern.FindStringSubmatch(eventdata)
	// if len(match) < 4 {
	// 	logger.Error("Event data unparseable:", eventdata)
	// 	return []hll.Event{}
	// }
	// return []hll.Event{hll.TeamSwitchEvent{
	// 	GenericEvent: hll.GenericEvent{
	// 		EventType: event_teamswitch,
	// 		EventTime: time,
	// 	},
	// 	Player: hll.PlayerInfo{
	// 		Name: match[1],
	// 		ID:   hll.NoPlayerID,
	// 	},
	// 	From: hll.Team(match[2]),
	// 	To:   hll.Team(match[3]),
	// }}
}

func logToChatEvent(time time.Time, eventdata string) []hll.Event {
	match := chatPattern.FindStringSubmatch(eventdata)
	if len(match) < 6 {
		logger.Error("Event data unparseable:", eventdata)
		return []hll.Event{}
	}
	return []hll.Event{hll.ChatEvent{
		GenericEvent: hll.GenericEvent{
			EventType: hll.EVENT_CHAT,
			EventTime: time,
		},
		Player: hll.PlayerInfo{
			Name: match[2],
			ID:   match[4],
		},
		Team:    hll.TeamFromString(match[3]),
		Scope:   hll.ChatScopeFromString(match[1]),
		Message: match[5],
	}}
}

func logToBanEvent(time time.Time, eventdata string) []hll.Event {
	match := banPattern.FindStringSubmatch(eventdata)
	if len(match) < 3 {
		logger.Error("Event data unparseable:", eventdata)
		return []hll.Event{}
	}
	return []hll.Event{hll.BanEvent{
		GenericEvent: hll.GenericEvent{
			EventType: hll.EVENT_BAN,
			EventTime: time,
		},
		Player: hll.PlayerInfo{
			Name: match[1],
			ID:   hll.NoPlayerID,
		},
		Reason: match[2],
	}}
}

func logToKickEvent(time time.Time, eventdata string) []hll.Event {
	match := kickPattern.FindStringSubmatch(eventdata)
	if len(match) < 3 {
		logger.Error("Event data unparseable:", eventdata)
		return []hll.Event{}
	}
	return []hll.Event{hll.KickEvent{
		GenericEvent: hll.GenericEvent{
			EventType: hll.EVENT_KICK,
			EventTime: time,
		},
		Player: hll.PlayerInfo{
			Name: match[1],
			ID:   hll.NoPlayerID,
		},
		Reason: match[2],
	}}
}

func logToMessageEvent(time time.Time, eventdata string) []hll.Event {
	match := msgPattern.FindStringSubmatch(eventdata)
	if len(match) < 4 {
		logger.Error("Event data unparseable:", eventdata)
		return []hll.Event{}
	}
	return []hll.Event{hll.MessageEvent{
		GenericEvent: hll.GenericEvent{
			EventType: hll.EVENT_MESSAGE,
			EventTime: time,
		},
		Player: hll.PlayerInfo{
			Name: match[1],
			ID:   match[2],
		},
		Message: match[3],
	}}
}

func logToMatchStartEvent(time time.Time, eventdata string) []hll.Event {
	match := startPattern.FindStringSubmatch(eventdata)
	if len(match) < 2 {
		logger.Error("Event data unparseable:", eventdata)
		return []hll.Event{}
	}
	return []hll.Event{hll.MatchStartEvent{
		GenericEvent: hll.GenericEvent{
			EventType: hll.EVENT_MATCHSTART,
			EventTime: time,
		},
		Map: hll.LogMapNameToMap(match[1]),
	}}
}

func logToMatchEndEvent(time time.Time, eventdata string) []hll.Event {
	match := endPattern.FindStringSubmatch(eventdata)
	if len(match) < 4 {
		logger.Error("Event data unparseable:", eventdata)
		return []hll.Event{}
	}
	return []hll.Event{hll.MatchEndEvent{
		GenericEvent: hll.GenericEvent{
			EventType: hll.EVENT_MATCHEND,
			EventTime: time,
		},
		Map: hll.LogMapNameToMap(match[1]),
		Score: hll.TeamData{
			Allies: util.ToInt(match[2]),
			Axis:   util.ToInt(match[3]),
		},
	}}
}

func logToAdminCamEvent(time time.Time, eventdata string) []hll.Event {
	events := []hll.Event{}
	match := camPattern.FindStringSubmatch(eventdata)
	if len(match) < 4 {
		logger.Error("Event data unparseable:", eventdata)
		return events
	}
	if strings.HasPrefix(match[3], "Entered") {
		events = append(events,
			hll.AdminCamEnteredEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_ENTER_ADMINCAM,
					EventTime: time,
				},
				Player: hll.PlayerInfo{
					Name: match[1],
					ID:   match[2],
				},
			})
	} else {
		events = append(events,
			hll.AdminCamLeftEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_LEAVE_ADMINCAM,
					EventTime: time,
				},
				Player: hll.PlayerInfo{
					Name: match[1],
					ID:   match[2],
				},
			})
	}
	return events
}

func logToVoteEvents(time time.Time, eventdata string) []hll.Event {
	events := []hll.Event{}
	if match := voteStartedPattern.FindStringSubmatch(eventdata); len(match) > 4 {
		voteStartEvent := hll.VoteStartedEvent{
			GenericEvent: hll.GenericEvent{
				EventType: hll.EVENT_VOTE_KICK_STARTED,
				EventTime: time,
			},
			Reason: match[2],
			ID:     util.ToInt(match[4]),
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
	} else if match := voteSubmittedPattern.FindStringSubmatch(eventdata); len(match) > 3 {
		events = append(events,
			hll.VoteSubmittedEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_VOTE_SUBMITTED,
					EventTime: time,
				},
				Submitter: hll.PlayerInfo{
					Name: match[1],
					ID:   hll.NoPlayerID,
				},
				ID:   util.ToInt(match[3]),
				Vote: match[2],
			},
		)
	} else if match := voteCompletePattern.FindStringSubmatch(eventdata); len(match) > 2 {
		voteID := util.ToInt(match[1])

		if voteStartEvent, ok := openVoteKicksMap[voteID]; ok {
			events = append(events,
				hll.VoteCompletedEvent{
					GenericEvent: hll.GenericEvent{
						EventType: hll.EVENT_VOTE_KICK_COMPLETED,
						EventTime: time,
					},
					Reason:    voteStartEvent.Reason,
					Result:    match[2],
					ID:        voteStartEvent.ID,
					Initiator: voteStartEvent.Initiator,
					Target:    voteStartEvent.Target,
				},
			)
		}
		delete(openVoteKicksMap, voteID)
	}
	return events
}

func logToEvents(logline string) []hll.Event {
	match := logPattern.FindStringSubmatch(logline)
	if len(match) < 3 {
		logger.Error("Logline invalid format:", logline)
		return []hll.Event{hll.GenericEvent{
			EventType: hll.EVENT_GENERIC,
			EventTime: time.Now(),
		}}
	}
	timestamp := time.Unix(util.ToInt64(match[1]), 0)
	data := match[2]

	for eventPrefix, parser := range logEventParsers {
		if strings.HasPrefix(data, string(eventPrefix)) {
			return parser(timestamp, data)
		}
	}

	logger.Error("Logline unparseable:", logline)
	return []hll.Event{}
}
