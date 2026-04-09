package main

type Cache struct {
	Data map[string]string
}

func NewCache() *Cache {
	return &Cache{Data: make(map[string]string)}
}

func (c *Cache) Set(key, val string) {
	c.Data[key] = val
}

func (c *Cache) Get(key string) string {
	return c.Data[key]
}
