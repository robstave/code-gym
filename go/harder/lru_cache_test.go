// lru_cache_test.go
package util

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	t.Run("basic put and get", func(t *testing.T) {
		cache := NewLRUCache(2)

		cache.Put(1, 1)
		cache.Put(2, 2)

		// Test retrieval
		v1, exists1 := cache.Get(1)
		if !exists1 || v1 != 1 {
			t.Errorf("Expected to get 1 from key 1, got %d, exists=%v", v1, exists1)
		}

		v2, exists2 := cache.Get(2)
		if !exists2 || v2 != 2 {
			t.Errorf("Expected to get 2 from key 2, got %d, exists=%v", v2, exists2)
		}
	})

	t.Run("eviction policy", func(t *testing.T) {
		cache := NewLRUCache(2)

		// Fill cache
		cache.Put(1, 1)
		cache.Put(2, 2)

		// Access key 1 to make it the most recently used
		cache.Get(1)

		// Add a new item which should evict key 2
		cache.Put(3, 3)

		// Key 2 should be gone
		_, exists2 := cache.Get(2)
		if exists2 {
			t.Errorf("Expected key 2 to be evicted")
		}

		// Keys 1 and 3 should exist
		_, exists1 := cache.Get(1)
		_, exists3 := cache.Get(3)
		if !exists1 || !exists3 {
			t.Errorf("Expected keys 1 and 3 to exist, got exists1=%v, exists3=%v", exists1, exists3)
		}
	})

	t.Run("update existing key", func(t *testing.T) {
		cache := NewLRUCache(2)

		cache.Put(1, 1)
		cache.Put(1, 10) // Update value

		v1, _ := cache.Get(1)
		if v1 != 10 {
			t.Errorf("Expected updated value 10 for key 1, got %d", v1)
		}
	})

	t.Run("contains without updating recency", func(t *testing.T) {
		cache := NewLRUCache(2)

		cache.Put(1, 1)
		cache.Put(2, 2)

		// Check key 1 exists without marking it as recently used
		if !cache.Contains(1) {
			t.Errorf("Expected Contains(1) to return true")
		}

		// Add key 3 which should evict the least recently used (key 1)
		cache.Put(3, 3)

		// Key 1 should be gone
		if cache.Contains(1) {
			t.Errorf("Expected key 1 to be evicted")
		}

		// Keys 2 and 3 should exist
		if !cache.Contains(2) || !cache.Contains(3) {
			t.Errorf("Expected keys 2 and 3 to exist")
		}
	})

	t.Run("empty cache", func(t *testing.T) {
		cache := NewLRUCache(0)

		cache.Put(1, 1) // Should be a no-op

		if cache.Len() != 0 {
			t.Errorf("Expected empty cache, got length %d", cache.Len())
		}

		if cache.Contains(1) {
			t.Errorf("Expected key 1 to not exist in empty cache")
		}
	})

	t.Run("clear cache", func(t *testing.T) {
		cache := NewLRUCache(5)

		for i := 0; i < 5; i++ {
			cache.Put(i, i*10)
		}

		if cache.Len() != 5 {
			t.Errorf("Expected cache length to be 5, got %d", cache.Len())
		}

		cache.Clear()

		if cache.Len() != 0 {
			t.Errorf("Expected cache to be empty after Clear(), got length %d", cache.Len())
		}

		for i := 0; i < 5; i++ {
			if cache.Contains(i) {
				t.Errorf("Expected key %d to not exist after clear", i)
			}
		}
	})

	t.Run("complex eviction sequence", func(t *testing.T) {
		cache := NewLRUCache(3)

		// Fill cache
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Put(3, 3)

		// Access in reverse order to change recency
		cache.Get(3) // Most recent
		cache.Get(2)
		cache.Get(1) // Least recent

		// Add new item, should evict 1
		cache.Put(4, 4)

		// Check eviction
		if cache.Contains(1) {
			t.Errorf("Expected key 1 to be evicted")
		}

		// Change access order again
		cache.Get(2) // Most recent
		cache.Get(4)
		cache.Get(3) // Least recent

		// Add new item, should evict 3
		cache.Put(5, 5)

		// Check final state
		expected := map[int]bool{
			1: false,
			2: true,
			3: false,
			4: true,
			5: true,
		}

		for k, expectedExists := range expected {
			if cache.Contains(k) != expectedExists {
				t.Errorf("Key %d: expected exists=%v, got %v", k, expectedExists, cache.Contains(k))
			}
		}
	})
}
