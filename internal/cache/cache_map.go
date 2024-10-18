package cache

import (
    "time"
    "sync"
)

type Cache struct {
    storage map[string]cacheEntry
    mu *sync.Mutex
}

type cacheEntry struct {
    createdAt time.Time
    val []byte
}

func NewCache(interval time.Duration) Cache {
    cache := Cache{ storage: make(map[string]cacheEntry), mu: &sync.Mutex{} }
    ticker := time.NewTicker(interval)
    go func() {
        for {
            <- ticker.C
            cache.reapLoop(interval)
        }
    }()
    return cache
}

func (c Cache) Add(url string, val []byte) {
    c.mu.Lock()
    c.storage[url] = cacheEntry{ createdAt: time.Now().UTC(), val: val }
    c.mu.Unlock()
}

func (c Cache) Get(url string) ([]byte, bool) {
    c.mu.Lock()
    val, ok := c.storage[url]
    c.mu.Unlock()
    return val.val, ok
}

func (c Cache) reapLoop(interval time.Duration) {
    c.mu.Lock()
    for key, val := range c.storage {
        if val.createdAt.Add(interval).Before(time.Now().UTC()) {
            delete(c.storage, key)
        }
    }
    c.mu.Unlock()
}
