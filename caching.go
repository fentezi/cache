package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	items map[string]Item
	mu    sync.Mutex
}

type Item struct {
	Value   any
	TTL     time.Duration
	TimeNow time.Time
}

func New() *Cache {
	return &Cache{
		items: make(map[string]Item),
	}
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = Item{
		Value:   value,
		TTL:     ttl,
		TimeNow: time.Now(),
	}
}

func (c *Cache) Get(key string) (any, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	t := time.Since(c.items[key].TimeNow)
	if t > c.items[key].TTL {
		c.Delete(key)
		return nil, errors.New("key not found")
	}
	item, _ := c.items[key]
	return item.Value, nil
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.items[key]
	if !ok {
		fmt.Println("Key not found")
		return
	}
	delete(c.items, key)
}
