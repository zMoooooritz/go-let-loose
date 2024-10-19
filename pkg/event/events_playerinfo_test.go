package event

import (
	"reflect"
	"testing"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

func TestPlayerInfoDiffToEvents(t *testing.T) {
	t.Run("Empty oldData should return no events", func(t *testing.T) {
		oldData := hll.DetailedPlayerInfo{}
		newData := hll.DetailedPlayerInfo{
			PlayerInfo: hll.PlayerInfo{Name: "Player1"},
			Team:       "Allies",
		}
		expected := []Event{}

		result := PlayerInfoDiffToEvents(oldData, newData)
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

		events := PlayerInfoDiffToEvents(oldData, newData)
		if len(events) != 1 {
			t.Fatalf("Expected 1 event, but got %d", len(events))
		}

		event, ok := events[0].(PlayerSwitchTeamEvent)
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

		events := PlayerInfoDiffToEvents(oldData, newData)
		if len(events) != 1 {
			t.Fatalf("Expected 1 event, but got %d", len(events))
		}

		event, ok := events[0].(PlayerSwitchSquadEvent)
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

		events := PlayerInfoDiffToEvents(oldData, newData)
		if len(events) != 1 {
			t.Fatalf("Expected 1 event, but got %d", len(events))
		}

		event, ok := events[0].(PlayerChangeRoleEvent)
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

		events := PlayerInfoDiffToEvents(oldData, newData)
		if len(events) != 1 {
			t.Fatalf("Expected 1 event, but got %d", len(events))
		}

		event, ok := events[0].(PlayerChangeLoadoutEvent)
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

		events := PlayerInfoDiffToEvents(oldData, newData)
		if len(events) != 1 {
			t.Fatalf("Expected 1 event, but got %d", len(events))
		}

		event, ok := events[0].(PlayerScoreUpdateEvent)
		if !ok {
			t.Fatalf("Expected PlayerScoreUpdateEvent, but got %T", events[0])
		}

		expectedDiff := hll.Score{Combat: 5, Offense: 0, Defense: 10, Support: 0}
		if event.ScoreChange != expectedDiff {
			t.Errorf("Expected score diff %v, but got %v", expectedDiff, event.ScoreChange)
		}
	})
}
