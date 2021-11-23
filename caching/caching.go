package caching

import (
	"fmt"
	"go-weather/search"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type Memorycache struct {
	cache *memcache.Client
}

var cache *Memorycache

func init() {
	cache = &Memorycache{cache: memcache.New("localhost:11211")}
}

func FetchFromCache(date time.Time, location search.Coordinate) (string, error) {
	item, error := cache.cache.Get(fmt.Sprintf("%d:%d:%d:%f:%f", date.Day(), date.Month(), date.Year(), location.Longitude, location.Latitude))
	if error != nil {
		fmt.Println("Not found in cache")
		return "", error
	}
	return string(item.Value), nil
}

func AddToCache(value string, location search.Coordinate) {
	date := time.Now()
	item := memcache.Item{Key: fmt.Sprintf("%d:%d:%d:%f:%f", date.Day(), date.Month(), date.Year(), location.Longitude, location.Latitude), Value: []byte(value)}
	if err := cache.cache.Set(&item); err != nil {
		fmt.Println("Could not save to cache")
	}
}
