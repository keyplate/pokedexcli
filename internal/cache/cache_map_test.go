package cache 

import (
    "testing"
    "fmt"
    "time"
)

func TestCacheAdded(t *testing.T) {
    interval := 5 * time.Second
    testCases := []struct {
        key string
        val []byte
    }{
        {
            key: "https://google.com",
            val: []byte("hello"),
        },
        {
            key: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
            val: []byte("my fav video"),
        },
    }
    for i, c := range testCases {
        t.Run(fmt.Sprintf("Test case #%d", i), func(t *testing.T) {
            cache := NewCache(interval)
            cache.Add(c.key, c.val)
            valActual, ok := cache.Get(c.key)
            if !ok {
                t.Errorf("Expected to find key %s", c.key)
                return
            }
            if string(c.val) != string(valActual) {
                t.Errorf("Expected to find value %v", c.val)
                return
            }
        })
    }
}

func TestCacheCleaned(t *testing.T) {
    const cleanInterval = 5 * time.Millisecond
    const waitInterval = cleanInterval + 5 * time.Millisecond 

    cache := NewCache(cleanInterval)
    key := "abc.com"
    cache.Add(key, []byte("la-la-la"))
    
    _, ok := cache.Get(key)
    if !ok {
        t.Errorf("Expected to find key")
    }

    time.Sleep(waitInterval)
    _, ok = cache.Get(key)
    if ok {
        t.Errorf("Expected to not find key")
    }
}
