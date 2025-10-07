package rcon

import (
	"reflect"
	"testing"
	"time"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

func TestGameStateDiffToEvents(t *testing.T) {
	defaultData := hll.GameState{
		PlayerCount:      hll.TeamData{Axis: 49, Allies: 49},
		GameScore:        hll.TeamData{Axis: 2, Allies: 3},
		RemainingSeconds: 100,
		CurrentMap:       hll.ParseLayer("carentan_warfare"),
		NextMap:          hll.ParseLayer("carentan_warfare"),
	}

	t.Run("Empty oldData should return no events", func(t *testing.T) {
		oldData := hll.GameState{}
		newData := defaultData
		expected := []hll.Event{}

		result := gameStateDiffToEvents(oldData, newData)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Game just started, AxisScore == AlliesScore should return no events", func(t *testing.T) {
		oldData := defaultData
		newData := defaultData
		newData.GameScore.Axis = 2
		newData.GameScore.Allies = 2
		expected := []hll.Event{}

		result := gameStateDiffToEvents(oldData, newData)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Game score changes should return ObjectiveCaptureEvent", func(t *testing.T) {
		oldData := defaultData
		newData := defaultData
		newData.GameScore.Axis = 3
		newData.GameScore.Allies = 2

		events := gameStateDiffToEvents(oldData, newData)
		if len(events) != 1 {
			t.Fatalf("Expected 1 event, but got %d", len(events))
		}

		event, ok := events[0].(hll.ObjectiveCaptureEvent)
		if !ok {
			t.Fatalf("Expected ObjectiveCaptureEvent, but got %T", events[0])
		}

		if event.OldScore != oldData.GameScore || event.NewScore != newData.GameScore {
			t.Errorf("Expected old score %v and new score %v, but got old score %v and new score %v",
				oldData.GameScore, newData.GameScore, event.OldScore, event.NewScore)
		}

		if event.Type() != hll.EVENT_OBJECTIVE_CAPPED {
			t.Errorf("Expected event type %s, but got %s", hll.EVENT_OBJECTIVE_CAPPED, event.Type())
		}
	})

	t.Run("No score change should return no events", func(t *testing.T) {
		oldData := defaultData
		newData := defaultData
		expected := []hll.Event{}

		result := gameStateDiffToEvents(oldData, newData)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}

var (
	malformedLine   = "[Invalid timestamp] CONNECTED A Player Name (12345678901234567)"
	unparseableLine = "[10:00:00 hours (1639106251)] UNKNOWN EVENT TYPE"
	joinLine        = "[10:00:00 hours (1639106251)] CONNECTED A Player Name (12345678901234567)"
	disconnectLine  = "[10:00:00 hours (1639122640)] DISCONNECTED A Player Name (12345678901234567)"
	// teamSwitchLine    = "[6.14 sec (1645012374)] TEAMSWITCH T17 Scott (None > Allies)"
	killLine          = "[10:00:00 hours (1639143555)] KILL: A Player Name(Axis/12345678901234567) -> Another Player name(Allies/98765432109876543) with MP40"
	teamKillLine      = "[10:00:00 hours (1639144073)] TEAM KILL: A Player Name(Allies/12345678901234567) -> Another Player name(Allies/98765432109876543) with M1 GARAND"
	teamChatLine      = "[30:00 min (1639144118)] CHAT[Team][A Player Name(Allies/12345678901234567)]: Please build garrisons!"
	unitChatLine      = "[30:00 min (1639145775)] CHAT[Unit][A Player Name(Axis/12345678901234567)]: comms working?"
	enterCamLine      = "[15.03 sec (1639148961)] Player [A Player Name (12345678901234567)] Entered Admin Camera"
	leaveCamLine      = "[15.03 sec (1639148961)] Player [A Player Name (12345678901234567)] Left Admin Camera"
	banLine           = "[15.03 sec (1639148961)] BAN: [A Player Name] has been banned. [BANNED FOR 2 HOURS BY THE ADMINISTRATOR!]"
	kickLine          = "[15.03 sec (1639148961)] KICK: [A Player Name] has been kicked. [BANNED FOR 2 HOURS BY THE ADMINISTRATOR!]"
	messageLine       = "[15.03 sec (1639148961)] MESSAGE: player [A Player Name(12345678901234567)], content [Stop teamkilling, you donkey!\nWhat u doing?]"
	matchStartLine    = "[805 ms (1639148969)] MATCH START SAINTE-MÈRE-ÉGLISE Warfare"
	matchEndLine      = "[805 ms (1639148969)] MATCH ENDED `SAINTE-MÈRE-ÉGLISE Warfare` ALLIED (2 - 3) AXIS"
	voteStartedLine   = "[15.5 sec (1675360329)] VOTESYS: Player [NoodleArms] Started a vote of type (PVR_Kick_Abuse) against [buscÃ´O-sensei]. VoteID: [2]"
	voteSubmittedLine = "[9.85 sec (1675360334)] VOTESYS: Player [Dingbat252] voted [PV_Favour] for VoteID[2]"
	voteCompleted     = "[4.56 sec (1675360340)] VOTESYS: Vote [2] completed. Result: PVR_Passed"
	voteKick          = "[4.56 sec (1675360340)] VOTESYS: Vote Kick {buscÃ´O-sensei} successfully passed. [For: 2/1 - Against: 0]"
)

func TestLogToEvents(t *testing.T) {
	t.Run("Malformed log line", func(t *testing.T) {
		expected := []hll.Event{
			hll.GenericEvent{
				EventType: hll.EVENT_GENERIC,
				EventTime: time.Now(),
			},
		}

		result := logToEvents(malformedLine)
		if len(result) != 1 || result[0].(hll.GenericEvent).Type() != expected[0].Type() {
			t.Errorf("Expected a generic event due to malformed log, but got %v", result)
		}
	})

	t.Run("Unparseable log line", func(t *testing.T) {
		expected := []hll.Event{}

		result := logToEvents(unparseableLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Parse CONNECTED log line", func(t *testing.T) {
		expected := []hll.Event{
			hll.ConnectEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_CONNECTED,
					EventTime: time.Unix(1639106251, 0),
				},
				Player: hll.PlayerInfo{
					Name: "A Player Name",
					ID:   "12345678901234567",
				},
			},
		}

		result := logToEvents(joinLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Parse DISCONNECTED log line", func(t *testing.T) {
		expected := []hll.Event{
			hll.DisconnectEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_DISCONNECTED,
					EventTime: time.Unix(1639122640, 0),
				},
				Player: hll.PlayerInfo{
					Name: "A Player Name",
					ID:   "12345678901234567",
				},
			},
		}

		result := logToEvents(disconnectLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	// t.Run("Parse TEAMSWITCH log line", func(t *testing.T) {
	// 	expected := []hll.Event{
	// 		hll.TeamSwitchEvent{
	// 			GenericEvent: hll.GenericEvent{
	// 				EventType: event_teamswitch,
	// 				EventTime: time.Unix(1645012374, 0),
	// 			},
	// 			Player: hll.PlayerInfo{
	// 				Name: "T17 Scott",
	// 				ID:   hll.NoPlayerID,
	// 			},
	// 			From: hll.TmNone,
	// 			To:   hll.TmAllies,
	// 		},
	// 	}

	// 	result := logToEvents(teamSwitchLine)
	// 	if !reflect.DeepEqual(result, expected) {
	// 		t.Errorf("Expected %v, but got %v", expected, result)
	// 	}
	// })

	t.Run("Parse KILL log line", func(t *testing.T) {
		expected := []hll.Event{
			hll.KillEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_KILL,
					EventTime: time.Unix(1639143555, 0),
				},
				Killer: hll.PlayerInfo{
					Name: "A Player Name",
					ID:   "12345678901234567",
				},
				Victim: hll.PlayerInfo{
					Name: "Another Player name",
					ID:   "98765432109876543",
				},
				Weapon: hll.ParseWeapon("MP40"),
			},
			hll.DeathEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_DEATH,
					EventTime: time.Unix(1639143555, 0),
				},
				Killer: hll.PlayerInfo{
					Name: "A Player Name",
					ID:   "12345678901234567",
				},
				Victim: hll.PlayerInfo{
					Name: "Another Player name",
					ID:   "98765432109876543",
				},
				Weapon: hll.ParseWeapon("MP40"),
			},
		}

		result := logToEvents(killLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Parse TEAMKILL log line", func(t *testing.T) {
		expected := []hll.Event{
			hll.TeamKillEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_TEAMKILL,
					EventTime: time.Unix(1639144073, 0),
				},
				Killer: hll.PlayerInfo{
					Name: "A Player Name",
					ID:   "12345678901234567",
				},
				Victim: hll.PlayerInfo{
					Name: "Another Player name",
					ID:   "98765432109876543",
				},
				Weapon: hll.ParseWeapon("M1 GARAND"),
			},
			hll.TeamDeathEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_TEAMDEATH,
					EventTime: time.Unix(1639144073, 0),
				},
				Killer: hll.PlayerInfo{
					Name: "A Player Name",
					ID:   "12345678901234567",
				},
				Victim: hll.PlayerInfo{
					Name: "Another Player name",
					ID:   "98765432109876543",
				},
				Weapon: hll.ParseWeapon("M1 GARAND"),
			},
		}

		result := logToEvents(teamKillLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Parse CHAT log line", func(t *testing.T) {
		expected := []hll.Event{
			hll.ChatEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_CHAT,
					EventTime: time.Unix(1639144118, 0),
				},
				Player: hll.PlayerInfo{
					Name: "A Player Name",
					ID:   "12345678901234567",
				},
				Team:    hll.TmAllies,
				Scope:   hll.CsTeam,
				Message: "Please build garrisons!",
			},
		}

		result := logToEvents(teamChatLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}

		expected = []hll.Event{
			hll.ChatEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_CHAT,
					EventTime: time.Unix(1639145775, 0),
				},
				Player: hll.PlayerInfo{
					Name: "A Player Name",
					ID:   "12345678901234567",
				},
				Team:    hll.TmAxis,
				Scope:   hll.CsUnit,
				Message: "comms working?",
			},
		}

		result = logToEvents(unitChatLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Parse CAM log line", func(t *testing.T) {
		expected := []hll.Event{
			hll.AdminCamEnteredEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_ENTER_ADMINCAM,
					EventTime: time.Unix(1639148961, 0),
				},
				Player: hll.PlayerInfo{
					Name: "A Player Name",
					ID:   "12345678901234567",
				},
			},
		}

		result := logToEvents(enterCamLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}

		expected = []hll.Event{
			hll.AdminCamLeftEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_LEAVE_ADMINCAM,
					EventTime: time.Unix(1639148961, 0),
				},
				Player: hll.PlayerInfo{
					Name: "A Player Name",
					ID:   "12345678901234567",
				},
			},
		}

		result = logToEvents(leaveCamLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Parse BAN log line", func(t *testing.T) {
		expected := []hll.Event{
			hll.BanEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_BAN,
					EventTime: time.Unix(1639148961, 0),
				},
				Player: hll.PlayerInfo{
					Name: "A Player Name",
					ID:   hll.NoPlayerID,
				},
				Reason: "BANNED FOR 2 HOURS BY THE ADMINISTRATOR!",
			},
		}

		result := logToEvents(banLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Parse KICK log line", func(t *testing.T) {
		expected := []hll.Event{
			hll.KickEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_KICK,
					EventTime: time.Unix(1639148961, 0),
				},
				Player: hll.PlayerInfo{
					Name: "A Player Name",
					ID:   hll.NoPlayerID,
				},
				Reason: "BANNED FOR 2 HOURS BY THE ADMINISTRATOR!",
			},
		}

		result := logToEvents(kickLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Parse Message log line", func(t *testing.T) {
		expected := []hll.Event{
			hll.MessageEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_MESSAGE,
					EventTime: time.Unix(1639148961, 0),
				},
				Player: hll.PlayerInfo{
					Name: "A Player Name",
					ID:   "12345678901234567",
				},
				Message: "Stop teamkilling, you donkey!\nWhat u doing?",
			},
		}

		result := logToEvents(messageLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Parse Match Start log line", func(t *testing.T) {
		expected := []hll.Event{
			hll.MatchStartEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_MATCHSTART,
					EventTime: time.Unix(1639148969, 0),
				},
				Map: hll.LogMapNameToMap("SAINTE-MÈRE-ÉGLISE Warfare"),
			},
		}

		result := logToEvents(matchStartLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Parse Match End log line", func(t *testing.T) {
		expected := []hll.Event{
			hll.MatchEndEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_MATCHEND,
					EventTime: time.Unix(1639148969, 0),
				},
				Map: hll.LogMapNameToMap("SAINTE-MÈRE-ÉGLISE Warfare"),
				Score: hll.TeamData{
					Allies: 2,
					Axis:   3,
				},
			},
		}

		result := logToEvents(matchEndLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Parse VOTESYS log lines", func(t *testing.T) {
		expected := []hll.Event{
			hll.VoteStartedEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_VOTE_KICK_STARTED,
					EventTime: time.Unix(1675360329, 0),
				},
				Reason: "PVR_Kick_Abuse",
				ID:     2,
				Initiator: hll.PlayerInfo{
					Name: "NoodleArms",
					ID:   hll.NoPlayerID,
				},
				Target: hll.PlayerInfo{
					Name: "buscÃ´O-sensei",
					ID:   hll.NoPlayerID,
				},
			},
			hll.VoteSubmittedEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_VOTE_SUBMITTED,
					EventTime: time.Unix(1675360334, 0),
				},
				Submitter: hll.PlayerInfo{
					Name: "Dingbat252",
					ID:   hll.NoPlayerID,
				},
				ID:   2,
				Vote: "PV_Favour",
			},
			hll.VoteCompletedEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_VOTE_KICK_COMPLETED,
					EventTime: time.Unix(1675360340, 0),
				},
				Reason: "PVR_Kick_Abuse",
				ID:     2,
				Result: "PVR_Passed",
				Initiator: hll.PlayerInfo{
					Name: "NoodleArms",
					ID:   hll.NoPlayerID,
				},
				Target: hll.PlayerInfo{
					Name: "buscÃ´O-sensei",
					ID:   hll.NoPlayerID,
				},
			},
		}

		result := logToEvents(voteStartedLine)
		result = append(result, logToEvents(voteSubmittedLine)...)
		result = append(result, logToEvents(voteCompleted)...)
		result = append(result, logToEvents(voteKick)...)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}

func TestPlayerInfoDiffToEvents(t *testing.T) {
	t.Run("Empty oldData should return no events", func(t *testing.T) {
		oldData := hll.DetailedPlayerInfo{}
		newData := hll.DetailedPlayerInfo{
			PlayerInfo: hll.PlayerInfo{Name: "Player1"},
			Team:       "Allies",
		}
		expected := []hll.Event{}

		result := playerInfoDiffToEvents(oldData, newData)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Player switches team", func(t *testing.T) {
		oldData := hll.DetailedPlayerInfo{
			PlayerInfo: hll.PlayerInfo{Name: "Player1"},
			Team:       "Axis",
		}
		newData := hll.DetailedPlayerInfo{
			PlayerInfo: hll.PlayerInfo{Name: "Player1"},
			Team:       "Allies",
		}

		events := playerInfoDiffToEvents(oldData, newData)
		if len(events) != 1 {
			t.Fatalf("Expected 1 event, but got %d", len(events))
		}

		event, ok := events[0].(hll.PlayerSwitchTeamEvent)
		if !ok {
			t.Fatalf("Expected PlayerSwitchTeamEvent, but got %T", events[0])
		}

		if event.OldTeam != oldData.Team || event.NewTeam != newData.Team {
			t.Errorf("Expected team switch from %s to %s, but got from %s to %s", oldData.Team, newData.Team, event.OldTeam, event.NewTeam)
		}
	})

	t.Run("Player switches squad", func(t *testing.T) {
		oldData := hll.DetailedPlayerInfo{
			PlayerInfo: hll.PlayerInfo{Name: "Player1"},
			Unit: hll.Unit{
				Name: "Alpha",
				ID:   1,
			},
		}
		newData := hll.DetailedPlayerInfo{
			PlayerInfo: hll.PlayerInfo{Name: "Player1"},
			Unit: hll.Unit{
				Name: "Bravo",
				ID:   2,
			},
		}

		events := playerInfoDiffToEvents(oldData, newData)
		if len(events) != 1 {
			t.Fatalf("Expected 1 event, but got %d", len(events))
		}

		event, ok := events[0].(hll.PlayerSwitchSquadEvent)
		if !ok {
			t.Fatalf("Expected PlayerSwitchSquadEvent, but got %T", events[0])
		}

		if event.OldSquad != oldData.Unit || event.NewSquad != newData.Unit {
			t.Errorf("Expected squad switch from %v to %v, but got from %v to %v", oldData.Unit, newData.Unit, event.OldSquad, event.NewSquad)
		}
	})

	t.Run("Player changes role", func(t *testing.T) {
		oldData := hll.DetailedPlayerInfo{
			PlayerInfo: hll.PlayerInfo{Name: "Player1"},
			Role:       "Rifleman",
		}
		newData := hll.DetailedPlayerInfo{
			PlayerInfo: hll.PlayerInfo{Name: "Player1"},
			Role:       "Officer",
		}

		events := playerInfoDiffToEvents(oldData, newData)
		if len(events) != 1 {
			t.Fatalf("Expected 1 event, but got %d", len(events))
		}

		event, ok := events[0].(hll.PlayerChangeRoleEvent)
		if !ok {
			t.Fatalf("Expected PlayerChangeRoleEvent, but got %T", events[0])
		}

		if event.OldRole != oldData.Role || event.NewRole != newData.Role {
			t.Errorf("Expected role change from %s to %s, but got from %s to %s", oldData.Role, newData.Role, event.OldRole, event.NewRole)
		}
	})

	t.Run("Player changes loadout", func(t *testing.T) {
		oldData := hll.DetailedPlayerInfo{
			PlayerInfo: hll.PlayerInfo{Name: "Player1"},
			Loadout:    "Standard",
		}
		newData := hll.DetailedPlayerInfo{
			PlayerInfo: hll.PlayerInfo{Name: "Player1"},
			Loadout:    "Sniper",
		}

		events := playerInfoDiffToEvents(oldData, newData)
		if len(events) != 1 {
			t.Fatalf("Expected 1 event, but got %d", len(events))
		}

		event, ok := events[0].(hll.PlayerChangeLoadoutEvent)
		if !ok {
			t.Fatalf("Expected PlayerChangeLoadoutEvent, but got %T", events[0])
		}

		if event.OldLoadout != oldData.Loadout || event.NewLoadout != newData.Loadout {
			t.Errorf("Expected loadout change from %s to %s, but got from %s to %s", oldData.Loadout, newData.Loadout, event.OldLoadout, event.NewLoadout)
		}
	})

	t.Run("Player score update", func(t *testing.T) {
		oldData := hll.DetailedPlayerInfo{
			PlayerInfo: hll.PlayerInfo{Name: "Player1"},
			Score:      hll.Score{Combat: 10, Offense: 5, Defense: 0, Support: 3},
		}
		newData := hll.DetailedPlayerInfo{
			PlayerInfo: hll.PlayerInfo{Name: "Player1"},
			Score:      hll.Score{Combat: 15, Offense: 5, Defense: 10, Support: 3},
		}

		events := playerInfoDiffToEvents(oldData, newData)
		if len(events) != 1 {
			t.Fatalf("Expected 1 event, but got %d", len(events))
		}

		event, ok := events[0].(hll.PlayerScoreUpdateEvent)
		if !ok {
			t.Fatalf("Expected PlayerScoreUpdateEvent, but got %T", events[0])
		}

		if event.OldScore.Combat == event.NewScore.Combat || event.OldScore.Offense != event.NewScore.Offense ||
			event.OldScore.Defense == event.NewScore.Defense || event.OldScore.Support != event.NewScore.Support {
			t.Errorf("Expected score update, but got old: %v, new: %v", event.OldScore, event.NewScore)
		}
	})
}
