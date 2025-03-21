package pokecache

import (
	"sync"
	"time"
)

type PokeCache interface {
	Get(key string) []byte
	Add(key string, value []byte)
}

type CacheItem struct {
	value     []byte
	createdAt time.Time
}
type Cache struct {
	mutex sync.RWMutex
	ttl   time.Duration
	items map[string]CacheItem
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	item, ok := c.items[key]
	if !ok {
		return nil, false
	}
	if time.Since(item.createdAt) > c.ttl {
		delete(c.items, key)
		return nil, false
	}
	return item.value, true
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items[key] = CacheItem{
		value:     value,
		createdAt: time.Now(),
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.items {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.items, k)
		}
	}
}

func NewCache(ttl time.Duration) *Cache {
	c := &Cache{
		ttl:   ttl,
		items: make(map[string]CacheItem),
	}
	go c.reapLoop(ttl)
	return c

}
