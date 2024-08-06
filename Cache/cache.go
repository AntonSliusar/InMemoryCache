package cache

type Cache struct {
	items map[string]any
}

func (c *Cache) Set(key string, value any) {
	c.items[key] = value
}

func (c *Cache) Get(key string) any {
	value, _ := c.items[key]
	return value
}

func (c *Cache) Delete(key string) {
	delete(c.items, key)
}
