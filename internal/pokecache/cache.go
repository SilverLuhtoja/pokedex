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
	cacheData map[string]cacheEntry
	mutex     *sync.Mutex
	timeout   time.Duration
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cacheData: make(map[string]cacheEntry),
		timeout:   interval,
		mutex:     &sync.Mutex{},
	}

	go cache.readLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cacheData[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, ok := c.cacheData[key]
	if !ok {
		return nil, false
	}
	return entry.val, ok
}

func (c *Cache) readLoop() {
	ticker := time.NewTicker(c.timeout)
	defer ticker.Stop()

	for range ticker.C {
		c.deleteFromMap()
	}

}

func (c *Cache) deleteFromMap() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, value := range c.cacheData {
		if time.Since(value.createdAt) > c.timeout {
			delete(c.cacheData, key)
		}
	}
}
