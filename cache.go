package myutils

import "sync"

type Cache[T interface{}] struct {
	cache map[string]T
	*sync.RWMutex
}

// NewCache creates a simple cache data structure. Data is stored in a thread safe map.
func NewCache[T interface{}]() *Cache[T] {
	return &Cache[T]{
		cache:   make(map[string]T),
		RWMutex: &sync.RWMutex{},
	}
}

// Set adds a value v by key k to the underlying map.
func (c *Cache[T]) Set(k string, v T) {
  c.Lock()
  defer c.Unlock()

  c.cache[k] = v
}

// Get retrieves a value by key k from the underlying map. If there is no value
//associated with the given key the second returned value will be false, 
//like in a basic map data structure.
func (c *Cache[T]) Get(k string) (T, bool) {
  c.RLock()
  defer c.RUnlock()

  v, ok := c.cache[k]
  return v, ok
}
