package utils

import (
	"github.com/FishGoddess/cachego"
)

type Cache struct {
	cache *cachego.Cache
}

func NewCache() *Cache {
	cache := cachego.NewCache()
	return &Cache{
		cache: cache,
	}
}

func (c *Cache) Set(key string, value interface{}, ttl int64) {
	c.cache.SetWithTTL(key, value, ttl)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.cache.Get(key)
}

func (c *Cache) Close() {
	c.cache.Gc()
}

func (c *Cache) Del(key string) {
	c.cache.Remove(key)
}
