package event

import (
	"errors"
	"sync"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

type Cache struct {
	gameState hll.GameState
	gsLock    *sync.RWMutex

	serverView hll.ServerView
	svLock     *sync.RWMutex

	players *ttlcache.Cache[string, hll.DetailedPlayerInfo]
}

func NewCache() *Cache {
	cache := Cache{
		gameState:  hll.GameState{},
		gsLock:     &sync.RWMutex{},
		serverView: hll.ServerView{},
		svLock:     &sync.RWMutex{},
		players: ttlcache.New(
			ttlcache.WithTTL[string, hll.DetailedPlayerInfo](2*time.Minute),
			ttlcache.WithDisableTouchOnHit[string, hll.DetailedPlayerInfo](),
		),
	}

	go cache.players.Start()

	return &cache
}

func (c *Cache) GetGameState() hll.GameState {
	c.gsLock.RLock()
	defer c.gsLock.RUnlock()
	return c.gameState
}

func (c *Cache) setGameState(gs hll.GameState) {
	c.gsLock.Lock()
	c.gameState = gs
	c.gsLock.Unlock()
}

func (c *Cache) GetServerView() hll.ServerView {
	c.svLock.RLock()
	defer c.svLock.RUnlock()
	return c.serverView
}

func (c *Cache) setServerView(sv hll.ServerView) {
	c.svLock.Lock()
	c.serverView = sv
	c.svLock.Unlock()
}

func (c *Cache) GetPlayerInfo(playerID string) (hll.DetailedPlayerInfo, error) {
	pd := c.players.Get(playerID)
	if pd == nil {
		return hll.DetailedPlayerInfo{}, errors.New("no information available")
	}
	return pd.Value(), nil
}

func (c *Cache) GetOnlinePlayerIDs() []string {
	return c.players.Keys()
}

func (c *Cache) setPlayerInfo(pd hll.DetailedPlayerInfo) {
	c.players.Set(pd.ID, pd, ttlcache.DefaultTTL)
}
