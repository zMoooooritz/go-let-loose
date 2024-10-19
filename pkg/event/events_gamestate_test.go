package event

import (
	"reflect"
	"testing"

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
		expected := []Event{}

		result := GameStateDiffToEvents(oldData, newData)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Game just started, AxisScore == AlliesScore should return no events", func(t *testing.T) {
		oldData := defaultData
		newData := defaultData
		newData.GameScore.Axis = 2
		newData.GameScore.Allies = 2
		expected := []Event{}

		result := GameStateDiffToEvents(oldData, newData)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Game score changes should return ObjectiveCaptureEvent", func(t *testing.T) {
		oldData := defaultData
		newData := defaultData
		newData.GameScore.Axis = 3
		newData.GameScore.Allies = 2

		events := GameStateDiffToEvents(oldData, newData)
		if len(events) != 1 {
			t.Fatalf("Expected 1 event, but got %d", len(events))
		}

		event, ok := events[0].(ObjectiveCaptureEvent)
		if !ok {
			t.Fatalf("Expected ObjectiveCaptureEvent, but got %T", events[0])
		}

		if event.OldScore != oldData.GameScore || event.NewScore != newData.GameScore {
			t.Errorf("Expected old score %v and new score %v, but got old score %v and new score %v",
				oldData.GameScore, newData.GameScore, event.OldScore, event.NewScore)
		}

		if event.Type() != EVENT_OBJECTIVE_CAPPED {
			t.Errorf("Expected event type %s, but got %s", EVENT_OBJECTIVE_CAPPED, event.Type())
		}
	})

	t.Run("No score change should return no events", func(t *testing.T) {
		oldData := defaultData
		newData := defaultData
		expected := []Event{}

		result := GameStateDiffToEvents(oldData, newData)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
