package event

import (
	"context"
	"sync"
	"time"

	"github.com/zMoooooritz/go-let-loose/pkg/config"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
	"github.com/zMoooooritz/go-let-loose/pkg/rcon"
)

func ServerInfoFetcher(rcn *rcon.Rcon, cache *Cache, events chan<- Event, ctx context.Context, wg *sync.WaitGroup) {
	piRecvChannel := make(chan []string, 100)
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			basicPlayers, err := rcn.GetPlayers()
			if err != nil {
				continue
			}

			gameState, err := rcn.GetGameState()
			if err == nil {
				stateEvents := GameStateDiffToEvents(cache.GetGameState(), gameState)
				for _, event := range stateEvents {
					events <- event
				}
				cache.setGameState(gameState)
			}

			for _, player := range basicPlayers {
				if hll.IsNameProblematic(player.Name) {
					detailedPlayer := hll.EmptyDetailedPlayerInfo()
					detailedPlayer.PlayerInfo = player
					cache.setPlayerInfo(detailedPlayer)
				} else {
					rcn.QueueJob(
						rcon.RconJob{
							Data: rcon.CommandData{
								Command: "playerinfo " + player.Name,
								Format:  config.RF_DIRECT,
							},
							Response: piRecvChannel,
						},
					)
				}
			}

			players := []hll.DetailedPlayerInfo{}
			for range basicPlayers {
				select {
				case <-ctx.Done():
					return
				case data := <-piRecvChannel:
					if len(data) == 0 {
						continue
					}
					playerData := data[0]
					detailedPlayer, err := rcon.ParsePlayerInfo(playerData)
					if err != nil {
						logger.Error("parsing player data failed", err, playerData)
						continue
					}
					players = append(players, detailedPlayer)

					oldPlayerData, err := cache.GetPlayerInfo(detailedPlayer.ID)
					if err == nil {
						playerEvents := PlayerInfoDiffToEvents(oldPlayerData, detailedPlayer)
						for _, event := range playerEvents {
							events <- event
						}
					}
					cache.setPlayerInfo(detailedPlayer)
				case <-time.After(time.Second):
					break
				}
			}

			serverView := hll.PlayerstoServerView(players)
			cache.setServerView(serverView)

			time.Sleep(2 * time.Second)
		}
	}
}
