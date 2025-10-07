package rcon

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/zMoooooritz/go-let-loose/internal/util"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
)

var players *ttlcache.Cache[string, hll.DetailedPlayerInfo] = ttlcache.New(
	ttlcache.WithTTL[string, hll.DetailedPlayerInfo](2*time.Minute),
	ttlcache.WithDisableTouchOnHit[string, hll.DetailedPlayerInfo](),
)

func eventHandlerRoutine(events <-chan hll.Event, eventNotifier *eventNotifier, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-events:
			if !ok {
				return
			}
			eventNotifier.notify(event)
		}
	}
}

func logsFetcherRoutine(rcn *Rcon, events chan<- hll.Event, ctx context.Context, wg *sync.WaitGroup) {
	initialRun := true
	lastSeenTime := int64(0)
	processedLogs := make(map[string]bool)

	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			logsEntries, err := rcn.GetLogEntries(30, "")

			if err != nil {
				logger.Error("fetching log entries failed", err)
				time.Sleep(time.Second)
				continue
			}

			for _, entry := range logsEntries {
				match := logPattern.FindStringSubmatch(entry.Message)
				if len(match) < 3 {
					continue
				}
				timestamp := util.ToInt64(match[1])
				currentLine := match[2]

				if timestamp < lastSeenTime {
					continue
				}

				if timestamp == lastSeenTime && processedLogs[currentLine] {
					continue
				}

				if !initialRun { // ignore past events on startup
					for _, event := range logToEvents(entry.Message) {
						events <- event
					}
				}

				if timestamp > lastSeenTime {
					lastSeenTime = timestamp
					processedLogs = make(map[string]bool)
				}

				processedLogs[currentLine] = true
			}

			initialRun = false

			time.Sleep(400 * time.Millisecond)
		}
	}
}

func serverInfoFetcherRoutine(rcn *Rcon, events chan<- hll.Event, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	var oldGameState hll.GameState

	for {
		select {
		case <-ctx.Done():
			return
		default:

			gameState, err := rcn.GetGameState()
			if err == nil {
				if oldGameState != (hll.GameState{}) {
					stateEvents := gameStateDiffToEvents(oldGameState, gameState)
					for _, event := range stateEvents {
						events <- event
					}
				}

				oldGameState = gameState
			}

			players, err := rcn.GetPlayersInfo()
			if err == nil {
				for _, player := range players {
					oldPlayerData, err := getPlayerInfo(player.ID)
					if err == nil {
						playerEvents := playerInfoDiffToEvents(oldPlayerData, player)
						for _, event := range playerEvents {
							events <- event
						}
					}
					setPlayerInfo(player)
				}
			}

			time.Sleep(time.Second)
		}
	}
}

func gameStateDiffToEvents(oldData hll.GameState, newData hll.GameState) []hll.Event {
	events := []hll.Event{}

	emptyGameState := hll.GameState{}
	if oldData == emptyGameState {
		return events
	}

	if oldData.GameScore != newData.GameScore {
		if newData.GameScore.Axis == newData.GameScore.Allies { // game just started
			return events
		}
		events = append(events, hll.ObjectiveCaptureEvent{
			GenericEvent: hll.GenericEvent{
				EventType: hll.EVENT_OBJECTIVE_CAPPED,
				EventTime: time.Now(),
			},
			OldScore: oldData.GameScore,
			NewScore: newData.GameScore,
		})
	}
	return events
}

func playerInfoDiffToEvents(oldData hll.DetailedPlayerInfo, newData hll.DetailedPlayerInfo) []hll.Event {
	events := []hll.Event{}

	emptyPlayerInfo := hll.DetailedPlayerInfo{}
	if oldData == emptyPlayerInfo {
		return events
	}

	if oldData.Team != newData.Team {
		events = append(events, hll.PlayerSwitchTeamEvent{
			GenericEvent: hll.GenericEvent{
				EventType: hll.EVENT_TEAM_SWITCHED,
				EventTime: time.Now(),
			},
			Player:  newData.PlayerInfo,
			OldTeam: oldData.Team,
			NewTeam: newData.Team,
		})
	}
	if oldData.Unit != newData.Unit {
		events = append(events, hll.PlayerSwitchSquadEvent{
			GenericEvent: hll.GenericEvent{
				EventType: hll.EVENT_SQUAD_SWITCHED,
				EventTime: time.Now(),
			},
			Player:   newData.PlayerInfo,
			OldSquad: oldData.Unit,
			NewSquad: newData.Unit,
		})
	}
	if oldData.Role != newData.Role {
		events = append(events, hll.PlayerChangeRoleEvent{
			GenericEvent: hll.GenericEvent{
				EventType: hll.EVENT_ROLE_CHANGED,
				EventTime: time.Now(),
			},
			Player:  newData.PlayerInfo,
			OldRole: oldData.Role,
			NewRole: newData.Role,
		})
	}
	if oldData.Loadout != newData.Loadout {
		events = append(events, hll.PlayerChangeLoadoutEvent{
			GenericEvent: hll.GenericEvent{
				EventType: hll.EVENT_LOADOUT_CHANGED,
				EventTime: time.Now(),
			},
			Player:     newData.PlayerInfo,
			OldLoadout: oldData.Loadout,
			NewLoadout: newData.Loadout,
		})
	}
	ds := hll.Score{}
	ds.Combat = max(newData.Score.Combat-oldData.Score.Combat, 0)
	ds.Defense = max(newData.Score.Defense-oldData.Score.Defense, 0)
	ds.Offense = max(newData.Score.Offense-oldData.Score.Offense, 0)
	ds.Support = max(newData.Score.Support-oldData.Score.Support, 0)
	if ds.Combat != 0 || ds.Support != 0 || ds.Offense != 0 || ds.Defense != 0 {
		events = append(events, hll.PlayerScoreUpdateEvent{
			GenericEvent: hll.GenericEvent{
				EventType: hll.EVENT_SCORE_UPDATE,
				EventTime: time.Now(),
			},
			Player:   newData.PlayerInfo,
			OldScore: oldData.Score,
			NewScore: newData.Score,
		})
	}
	if oldData.Position.IsActive() && newData.Position.IsActive() {
		if oldData.Position != newData.Position {
			events = append(events, hll.PlayerPositionChangedEvent{
				GenericEvent: hll.GenericEvent{
					EventType: hll.EVENT_POSITION_CHANGED,
					EventTime: time.Now(),
				},
				Player: newData.PlayerInfo,
				OldPos: oldData.Position,
				NewPos: newData.Position,
			})
		}
	}
	if oldData.ClanTag != newData.ClanTag {
		events = append(events, hll.PlayerClanTagChangedEvent{
			GenericEvent: hll.GenericEvent{
				EventType: hll.EVENT_CLAN_TAG_CHANGED,
				EventTime: time.Now(),
			},
			Player:     newData.PlayerInfo,
			OldClanTag: oldData.ClanTag,
			NewClanTag: newData.ClanTag,
		})
	}
	return events
}

func getPlayerInfo(playerID string) (hll.DetailedPlayerInfo, error) {
	pd := players.Get(playerID)
	if pd == nil {
		return hll.DetailedPlayerInfo{}, errors.New("no information available")
	}
	return pd.Value(), nil
}

func setPlayerInfo(pd hll.DetailedPlayerInfo) {
	players.Set(pd.ID, pd, ttlcache.DefaultTTL)
}
