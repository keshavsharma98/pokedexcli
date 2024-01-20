package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()
		timeIntervalAgo := time.Now().Add(-interval)

		for key, value := range c.cache {
			if value.createdAt.Before(timeIntervalAgo) {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {

	cache := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cache, ok := c.cache[key]

	if !ok {
		return nil, false
	}

	return cache.val, true
}
