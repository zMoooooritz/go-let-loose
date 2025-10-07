package api

import "time"

type Cacheable interface {
	CacheTTL() time.Duration
}
