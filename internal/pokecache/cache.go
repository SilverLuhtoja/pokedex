package pokecache

import (
	"sync"
	"time"

	"github.com/fatih/color"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheData map[string]cacheEntry
	mutex     sync.Mutex
	timeout   time.Duration
}

func NewCache(interval time.Duration) Cache {
	return Cache{timeout: interval}
}

func (c *Cache) Add(key string, val []byte) {
	color.Blue("ADDING TO CACHE")
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.cacheData == nil {
		c.cacheData = make(map[string]cacheEntry)
	}
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cacheData[key] = entry
}

func (c *Cache) Get(key *string) ([]byte, bool) {
	color.Blue("READING FROM CACHE")
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if key == nil {
		return nil, false
	}

	entry, ok := c.cacheData[*key]
	if !ok {
		return nil, false
	}
	return entry.val, ok
}
