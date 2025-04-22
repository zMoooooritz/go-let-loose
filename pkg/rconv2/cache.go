package rconv2

import (
	"errors"

	"github.com/jellydator/ttlcache/v3"
	"github.com/zMoooooritz/go-let-loose/internal/socketv2/api"
)

type rconCache struct {
	data    *ttlcache.Cache[string, any]
	enabled bool
}

type RconOption func(*Rcon)

func WithCache() RconOption {
	return func(r *Rcon) {
		r.cache.data = ttlcache.New[string, any](
			ttlcache.WithDisableTouchOnHit[string, any](),
		)
		r.cache.enabled = true
		go r.cache.data.Start()
	}
}

func WithoutCache() RconOption {
	return func(r *Rcon) {
		r.cache.data = nil
		r.cache.enabled = false
	}
}

func (c *rconCache) get(key string) (any, error) {
	if !c.enabled {
		return nil, errors.New("cache is disabled")
	}

	value := c.data.Get(key)
	if value == nil {
		return nil, errors.New("cache miss")
	}

	return value.Value(), nil
}

func (c *rconCache) set(key string, value any) {
	if !c.enabled {
		return
	}

	if ttlGetter, ok := any(value).(api.Cacheable); ok {
		ttl := ttlGetter.CacheTTL()
		if ttl > 0 {
			c.data.Set(key, value, ttl)
		}
	}
}
