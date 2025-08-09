// lru_cache.go
package util

// LRUCache is a fixed-size cache that evicts the least recently used item
// when the capacity is reached. It supports:
// - Put: Add or update a key-value pair
// - Get: Retrieve a value by key (and mark it as recently used)
// - Contains: Check if a key exists in the cache
type LRUCache struct {
	// Your fields here
}

// NewLRUCache creates a new LRUCache with the given capacity
func NewLRUCache(capacity int) *LRUCache {
	// Your implementation here
	return nil
}

// Get retrieves the value for a key and marks it as recently used
// Returns the value and true if found, or 0 and false if not found
func (c *LRUCache) Get(key int) (int, bool) {
	// Your implementation here
	return 0, false
}

// Put adds a new key-value pair or updates an existing one
// If the cache is at capacity, it should evict the least recently used item
func (c *LRUCache) Put(key, value int) {
	// Your implementation here
}

// Contains checks if a key exists in the cache without updating its "recently used" status
func (c *LRUCache) Contains(key int) bool {
	// Your implementation here
	return false
}

// Len returns the current number of items in the cache
func (c *LRUCache) Len() int {
	// Your implementation here
	return 0
}

// Clear removes all items from the cache
func (c *LRUCache) Clear() {
	// Your implementation here
}
