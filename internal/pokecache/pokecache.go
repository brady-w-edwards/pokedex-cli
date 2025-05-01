package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	entries  map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := new(Cache)
	c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.entries[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if c.entries[key].val != nil {
		return c.entries[key].val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	for _, entry := range c.entries {
		if time.Since(entry.createdAt) > c.interval {

		}
	}
}
