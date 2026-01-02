package aggregator

import (
	"github.com/patrickmn/go-cache"
	"time"
	"fmt"
)

func NewFlightCache() *cache.Cache {
    return cache.New(
        5*time.Minute,   // default expiration
        10*time.Minute,  // cleanup interval
    )
}

func cacheKey(origin, destination, date string) string {
    return fmt.Sprintf("%s:%s:%s", origin, destination, date)
}

var cacheInstance = NewFlightCache()
