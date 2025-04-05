package eventv2

import (
	"context"
	"sync"
	"time"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
)

func ServerInfoFetcher(rcn *rconv2.Rcon, cache *Cache, events chan<- Event, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			players, err := rcn.GetPlayersInfo()
			if err != nil {
				continue
			}

			for _, player := range players {
				oldPlayerData, err := cache.GetPlayerInfo(player.ID)
				if err == nil {
					playerEvents := PlayerInfoDiffToEvents(oldPlayerData, player)
					for _, event := range playerEvents {
						events <- event
					}
				}
				cache.setPlayerInfo(player)
			}

			serverView := hll.PlayersToServerView(players)
			cache.setServerView(serverView)

			time.Sleep(time.Second)
		}
	}
}
