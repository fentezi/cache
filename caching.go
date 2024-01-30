package cache

import (
	"fmt"
)

type Cache struct {
	items map[string]Item
}

type Item struct {
	Value any
}

func New() *Cache {
	return &Cache{
		items: make(map[string]Item),
	}
}

func (c *Cache) Set(key string, value any) {
	c.items[key] = Item{Value: value}
}

func (c *Cache) Get(key string) any {
	item, _ := c.items[key]
	return item.Value
}

func (c *Cache) Delete(key string) {
	_, ok := c.items[key]
	if !ok {
		fmt.Println("Key not found")
		return
	}
	delete(c.items, key)
}
