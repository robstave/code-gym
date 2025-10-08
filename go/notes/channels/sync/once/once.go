package main

import (
	"fmt"
	"sync"
	"time"
)

// Singleton represents the single shared instance we want to lazily and safely initialize.
// Fields:
//
//	data  - immutable string set at construction
//	value - mutable counter incremented by goroutines
//	mu    - protects concurrent access to 'value'
//
// NOTE: The struct itself does NOT enforce the singleton property; the package‑level
// variables and sync.Once below do.
type Singleton struct {
	data  string
	value int
	mu    sync.Mutex
}

// instance holds the single *Singleton once created. It starts as nil.
// once is a sync.Once that guarantees the init function runs at most one time.
// After once.Do finishes successfully, subsequent calls are fast and skip the function.
var instance *Singleton
var once sync.Once

// GetInstance returns the shared *Singleton.
// The anonymous function passed to once.Do is executed exactly once in a
// thread‑safe manner, even if many goroutines call GetInstance concurrently.
// Memory visibility: All writes performed during the first execution are
// visible to every goroutine that observes the returned pointer afterward.
func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating Singleton instance")
		instance = &Singleton{
			data:  "I'm the only one!",
			value: 0,
		}
	})
	return instance
}

// mutate safely increments the singleton's value; demonstrates internal locking.
// We use a pointer receiver so we mutate the underlying shared object rather than a copy.
func (s *Singleton) mutate(delta int) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.value += delta
	return s.value
}

func main() {
	// Launch several goroutines all racing to obtain the singleton. Only the first
	// caller will execute the initialization closure; others block until it completes.
	for i := range 5 {
		go func(id int) {
			s := GetInstance()

			// Show that every goroutine gets the same pointer.
			fmt.Printf("goroutine %d sees pointer %p\n", id, s)
			// Increment the shared counter under lock.
			newVal := s.mutate(1)
			// Direct read of s.value here is safe ONLY because mutate already
			// completed under the lock; we read after it releases (benign data race
			// avoided since we don't read concurrently inside mutate). For a strict
			// race‑free pattern you'd add a read method guarded by mu.
			fmt.Printf("%s value=%d (after increment)\n", s.data, newVal)
		}(i)
	}

	// Sleep to allow goroutines to finish (simplest demo approach). In real code
	// prefer sync.WaitGroup instead of sleep for coordination.
	time.Sleep(500 * time.Millisecond)
}
