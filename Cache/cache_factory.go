package cache

func New() *Cache {
	return &Cache{
		make(map[string]any),
	}
}
