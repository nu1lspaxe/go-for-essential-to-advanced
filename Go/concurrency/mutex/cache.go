package mutex

import (
	"sync"
	"testing"
)

// A Cache Structure for Safely Handling Concurrent Reads and Writes in
// High-Concurrency Scenarios.

// Lock-free
// [sync.Map] provides a thread-safe map. sync.Map.Store, sync.Map.Load, etc.
// Compare-And-Swap (CAS), e.g., `sync/atomic` to safely update without locks.

// Scinerio comparison:
// 	1. sync.Map: read-heavy, write-light
// 	2. sync.RWMutex: read-heavy, write-heavy
// 	3. sync.Mutex: write-heavy, read-light
// 	4. shard locking: write-heavy, read-heavy
// 	5. CAS (atomic): ultra-high performance

type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex // thread-safe
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.data[key]
	return value, exists
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func TestCache(t *testing.T) {
	cache := NewCache()

	cache.Set("name", "Golang")
	cache.Set("version", "1.20")

	if value, exists := cache.Get("name"); !exists || value != "Golang" {
		t.Errorf("Expected key 'name' to have value 'Golang', got '%v'", value)
	}

	if value, exists := cache.Get("version"); !exists || value != "1.20" {
		t.Errorf("Expected key 'version' to have value '1.20', got '%v'", value)
	}

	cache.Delete("version")

	if _, exists := cache.Get("version"); exists {
		t.Error("Expected key 'version' to be deleted")
	}

	if value, exists := cache.Get("name"); !exists || value != "Golang" {
		t.Errorf("Expected key 'name' to remain with value 'Golang', got '%v'", value)
	}
}
