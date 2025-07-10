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
	cache.entries = make(map[string]cacheEntry)
	return &cache
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.mu.Unlock()
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
		currentTime := <-ticker.C
		c.mu.Lock()
		for key, entry := range c.entries {
			if currentTime.Sub(entry.createdAt) > t {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
