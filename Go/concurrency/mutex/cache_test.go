package mutex_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/nu1lspaxe/go-for-essential-to-advanced/Go/concurrency/mutex"
	"golang.org/x/exp/rand"
)

func TestCache(t *testing.T) {
	cache := mutex.NewCache()

	// Test Set and Get operations
	cache.Set("name", "Golang")
	cache.Set("version", "1.20")

	// Verify the "name" key exists and its value is correct
	if value, exists := cache.Get("name"); !exists || value != "Golang" {
		t.Errorf("Expected key 'name' to have value 'Golang', got '%v'", value)
	}

	// Verify the "version" key exists and its value is correct
	if value, exists := cache.Get("version"); !exists || value != "1.20" {
		t.Errorf("Expected key 'version' to have value '1.20', got '%v'", value)
	}

	// Test Delete operation
	cache.Delete("version")

	// Verify the "version" key no longer exists
	if _, exists := cache.Get("version"); exists {
		t.Error("Expected key 'version' to be deleted")
	}

	// Verify the "name" key still exists after deleting "version"
	if value, exists := cache.Get("name"); !exists || value != "Golang" {
		t.Errorf("Expected key 'name' to remain with value 'Golang', got '%v'", value)
	}
}

func TestCacheConcurrency(t *testing.T) {
	cache := mutex.NewCache()
	var wg sync.WaitGroup

	// Define expected key-value pairs to validate results
	expected := make(map[string]int)
	var mu sync.Mutex // To synchronize writes to the `expected` map

	// Start multiple writer goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", id)
			value := rand.Intn(100)

			cache.Set(key, value)

			mu.Lock()
			expected[key] = value
			mu.Unlock()

			t.Logf("Writer %d: Set %s = %d", id, key, value)
		}(i)
	}

	// Start multiple reader goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", rand.Intn(5))
			if value, exists := cache.Get(key); exists {
				mu.Lock()
				expectedValue, ok := expected[key]
				mu.Unlock()

				if !ok || expectedValue != value {
					t.Errorf("Reader %d: Mismatched value for %s. Expected %v, Got %v", id, key, expectedValue, value)
				} else {
					t.Logf("Reader %d: Get %s = %v", id, key, value)
				}
			} else {
				t.Logf("Reader %d: %s not found", id, key)
			}
		}(i)
	}

	wg.Wait()
	t.Log("All goroutines finished")

	// Final validation of expected values
	for key, expectedValue := range expected {
		if value, exists := cache.Get(key); !exists || value != expectedValue {
			t.Errorf("Final validation failed for %s: Expected %v, Got %v", key, expectedValue, value)
		}
	}
}
