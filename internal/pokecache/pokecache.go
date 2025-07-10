package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

func NewCache(t time.Duration) *Cache {
	cache := Cache{}
	go cache.reapLoop(t)
	return &cache
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return []byte{}, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(t time.Duration) {
	ticker := time.NewTicker(t)
	defer ticker.Stop()

	for {
		for key, entry := range c.entries {
			c.mu.Lock()
			currentTime := <-ticker.C
			if currentTime.Sub(entry.createdAt) > t {
				delete(c.entries, key)
			}
		}
	}
}
