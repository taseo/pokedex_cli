package cache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu     sync.RWMutex
	data   map[string]cacheEntry
	stopCh chan struct{}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		data:   make(map[string]cacheEntry),
		stopCh: make(chan struct{}),
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Delete() {
	close(c.stopCh)
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()

	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()

	defer c.mu.RUnlock()

	v, ok := c.data[key]

	if ok {
		return v.val, ok
	}

	return nil, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()

			now := time.Now()

			for k := range c.data {
				if now.Sub(c.data[k].createdAt) >= interval {
					delete(c.data, k)
				}
			}

			c.mu.Unlock()

		case <-c.stopCh:
			fmt.Println("Stopping reapLoop goroutine")
			return
		}
	}
}
