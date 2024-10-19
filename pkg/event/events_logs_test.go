package event

import (
	"reflect"
	"testing"
	"time"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

var (
	malformedLine     = "[Invalid timestamp] CONNECTED A Player Name (12345678901234567)"
	unparseableLine   = "[10:00:00 hours (1639106251)] UNKNOWN EVENT TYPE"
	joinLine          = "[10:00:00 hours (1639106251)] CONNECTED A Player Name (12345678901234567)"
	disconnectLine    = "[10:00:00 hours (1639122640)] DISCONNECTED A Player Name (12345678901234567)"
	teamSwitchLine    = "[6.14 sec (1645012374)] TEAMSWITCH T17 Scott (None > Allies)"
	killLine          = "[10:00:00 hours (1639143555)] KILL: A Player Name(Axis/12345678901234567) -> Another Player name(Allies/98765432109876543) with MP40"
	teamKillLine      = "[10:00:00 hours (1639144073)] TEAM KILL: A Player Name(Allies/12345678901234567) -> Another Player name(Allies/98765432109876543) with M1 GARAND"
	teamChatLine      = "[30:00 min (1639144118)] CHAT[Team][A Player Name(Allies/12345678901234567)]: Please build garrisons!"
	unitChatLine      = "[30:00 min (1639145775)] CHAT[Unit][A Player Name(Axis/12345678901234567)]: comms working?"
	enterCamLine      = "[15.03 sec (1639148961)] Player [A Player Name (12345678901234567)] Entered Admin Camera"
	leaveCamLine      = "[15.03 sec (1639148961)] Player [A Player Name (12345678901234567)] Left Admin Camera"
	banLine           = "[15.03 sec (1639148961)] BAN: [A Player Name] has been banned. [BANNED FOR 2 HOURS BY THE ADMINISTRATOR!]"
	kickLine          = "[15.03 sec (1639148961)] KICK: [A Player Name] has been kicked. [BANNED FOR 2 HOURS BY THE ADMINISTRATOR!]"
	messageLine       = "[15.03 sec (1639148961)] MESSAGE: player [A Player Name(12345678901234567)], content [Stop teamkilling, you donkey!]"
	matchStartLine    = "[805 ms (1639148969)] MATCH START SAINTE-MÈRE-ÉGLISE Warfare"
	matchEndLine      = "[805 ms (1639148969)] MATCH ENDED `SAINTE-MÈRE-ÉGLISE Warfare` ALLIED (2 - 3) AXIS"
	voteStartedLine   = "[15.5 sec (1675360329)] VOTESYS: Player [NoodleArms] Started a vote of type (PVR_Kick_Abuse) against [buscÃ´O-sensei]. VoteID: [2]"
	voteSubmittedLine = "[9.85 sec (1675360334)] VOTESYS: Player [Dingbat252] voted [PV_Favour] for VoteID[2]"
	voteCompleted     = "[4.56 sec (1675360340)] VOTESYS: Vote [2] completed. Result: PVR_Passed"
	voteKick          = "[4.56 sec (1675360340)] VOTESYS: Vote Kick {buscÃ´O-sensei} successfully passed. [For: 2/1 - Against: 0]"
)

// TODO: add tests for non successfull votes
// VOTESYS: Vote [3] prematurely expired.
// VOTESYS: Vote [10] expired before completion

func TestLogToEvents(t *testing.T) {
	t.Run("Malformed log line", func(t *testing.T) {
		expected := []Event{
			GenericEvent{
				eventType: EVENT_GENERIC,
				time:      time.Now(),
			},
		}

		result := logToEvents(malformedLine)
		if len(result) != 1 || result[0].(GenericEvent).Type() != expected[0].Type() {
			t.Errorf("Expected a generic event due to malformed log, but got %v", result)
		}
	})

	t.Run("Unparseable log line", func(t *testing.T) {
		expected := []Event{}

		result := logToEvents(unparseableLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Parse CONNECTED log line", func(t *testing.T) {
		expected := []Event{
			ConnectEvent{
				GenericEvent: GenericEvent{EVENT_CONNECTED, time.Unix(1639106251, 0)},
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
		expected := []Event{
			DisconnectEvent{
				GenericEvent: GenericEvent{EVENT_DISCONNECTED, time.Unix(1639122640, 0)},
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

	t.Run("Parse TEAMSWITCH log line", func(t *testing.T) {
		expected := []Event{
			TeamSwitchEvent{
				GenericEvent: GenericEvent{EVENT_TEAMSWITCH, time.Unix(1645012374, 0)},
				Player: hll.PlayerInfo{
					Name: "T17 Scott",
					ID:   hll.NoPlayerID,
				},
				From: hll.TmNone,
				To:   hll.TmAllies,
			},
		}

		result := logToEvents(teamSwitchLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Parse KILL log line", func(t *testing.T) {
		expected := []Event{
			KillEvent{
				GenericEvent: GenericEvent{EVENT_KILL, time.Unix(1639143555, 0)},
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
			DeathEvent{
				GenericEvent: GenericEvent{EVENT_DEATH, time.Unix(1639143555, 0)},
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
		expected := []Event{
			TeamKillEvent{
				GenericEvent: GenericEvent{EVENT_TEAMKILL, time.Unix(1639144073, 0)},
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
			TeamDeathEvent{
				GenericEvent: GenericEvent{EVENT_TEAMDEATH, time.Unix(1639144073, 0)},
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
		expected := []Event{
			ChatEvent{
				GenericEvent: GenericEvent{EVENT_CHAT, time.Unix(1639144118, 0)},
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

		expected = []Event{
			ChatEvent{
				GenericEvent: GenericEvent{EVENT_CHAT, time.Unix(1639145775, 0)},
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
		expected := []Event{
			AdminCamEnteredEvent{
				GenericEvent: GenericEvent{EVENT_ENTER_ADMINCAM, time.Unix(1639148961, 0)},
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

		expected = []Event{
			AdminCamLeftEvent{
				GenericEvent: GenericEvent{EVENT_LEAVE_ADMINCAM, time.Unix(1639148961, 0)},
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
		expected := []Event{
			BanEvent{
				GenericEvent: GenericEvent{EVENT_BAN, time.Unix(1639148961, 0)},
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
		expected := []Event{
			KickEvent{
				GenericEvent: GenericEvent{EVENT_KICK, time.Unix(1639148961, 0)},
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
		expected := []Event{
			MessageEvent{
				GenericEvent: GenericEvent{EVENT_MESSAGE, time.Unix(1639148961, 0)},
				Player: hll.PlayerInfo{
					Name: "A Player Name",
					ID:   "12345678901234567",
				},
				Message: "Stop teamkilling, you donkey!",
			},
		}

		result := logToEvents(messageLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Parse Match Start log line", func(t *testing.T) {
		expected := []Event{
			MatchStartEvent{
				GenericEvent: GenericEvent{EVENT_MATCHSTART, time.Unix(1639148969, 0)},
				Map:          hll.LogMapNameToMap("SAINTE-MÈRE-ÉGLISE Warfare"),
			},
		}

		result := logToEvents(matchStartLine)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Parse Match End log line", func(t *testing.T) {
		expected := []Event{
			MatchEndEvent{
				GenericEvent: GenericEvent{EVENT_MATCHEND, time.Unix(1639148969, 0)},
				Map:          hll.LogMapNameToMap("SAINTE-MÈRE-ÉGLISE Warfare"),
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
		expected := []Event{
			VoteStartedEvent{
				GenericEvent: GenericEvent{EVENT_VOTE_KICK_STARTED, time.Unix(1675360329, 0)},
				Reason:       "PVR_Kick_Abuse",
				ID:           2,
				Initiator: hll.PlayerInfo{
					Name: "NoodleArms",
					ID:   hll.NoPlayerID,
				},
				Target: hll.PlayerInfo{
					Name: "buscÃ´O-sensei",
					ID:   hll.NoPlayerID,
				},
			},
			VoteSubmittedEvent{
				GenericEvent: GenericEvent{EVENT_VOTE_SUBMITTED, time.Unix(1675360334, 0)},
				Submitter: hll.PlayerInfo{
					Name: "Dingbat252",
					ID:   hll.NoPlayerID,
				},
				ID:   2,
				Vote: "PV_Favour",
			},
			VoteCompletedEvent{
				GenericEvent: GenericEvent{EVENT_VOTE_KICK_COMPLETED, time.Unix(1675360340, 0)},
				Reason:       "PVR_Kick_Abuse",
				ID:           2,
				Result:       "PVR_Passed",
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
