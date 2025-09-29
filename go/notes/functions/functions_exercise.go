package main

import (
	"errors"
	"fmt"
	"time"
)

/*
FUNCTIONS EXERCISE PLAYGROUND
=============================

This file collects many different patterns and features of Go functions.
Read the examples and complete the TODOs at the bottom (search for 'TODO(').

Concepts covered / to explore:
  1. Basic functions & multiple return values
  2. Named return values & naked returns (when to avoid)
  3. Variadic parameters (and forwarding)
  4. Closures capturing variables
  5. Higher-order functions (functions as values)
  6. Defer execution order & common idioms
  7. Error handling patterns
  8. Panic / recover for controlled failure boundaries
  9. Recursion & tail recursion note
 10. Functional option pattern
 11. Method vs free function (receiver choice)
 12. Time-based throttling with closure state
 13. Memoization with map + closure
 14. Generic function (simple demonstration)

After implementing TODOs you can run:
  go run .
(from this directory) or wire pieces into tests.
*/

// 1. Basic function with multiple returns
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

// 2. Named returns (use sparingly; can help with defer modifying result)
func areaRect(w, h int) (area int) { // 'area' named
	area = w * h
	return // naked return (acceptable here but be cautious in longer funcs)
}

// 3. Variadic + forwarding
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
func sumTwice(nums ...int) int { return sum(nums...) * 2 } // forwarding

// 4. Closure capturing variable
func counter(start int) func() int {
	n := start
	return func() int { n++; return n }
}

// 5. Higher-order function example
func apply[T any](vals []T, f func(T) T) []T {
	out := make([]T, len(vals))
	for i, v := range vals {
		out[i] = f(v)
	}
	return out
}

// 6. Defer order & resource pattern
func withTiming(label string, work func()) {
	start := time.Now()
	defer func() {
		fmt.Printf("%s took %v\n", label, time.Since(start))
	}()
	work()
}

// 7. Error wrapping pattern
func loadThing(id string) (string, error) {
	if id == "" {
		return "", fmt.Errorf("loadThing: %w", errors.New("empty id"))
	}
	return "thing:" + id, nil
}

// 8. Panic / recover boundary
func safeRun(fn func()) (recovered interface{}) {
	defer func() {
		if r := recover(); r != nil {
			recovered = r
		}
	}()
	fn()
	return nil
}

// 9. Recursion (factorial example) - NOTE: large n will overflow int
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 10. Functional option example

type server struct {
	addr    string
	timeout time.Duration
	secure  bool
}

type ServerOption func(*server)

func WithTimeout(d time.Duration) ServerOption { return func(s *server) { s.timeout = d } }
func WithSecure() ServerOption                 { return func(s *server) { s.secure = true } }

func NewServer(addr string, opts ...ServerOption) *server {
	s := &server{addr: addr, timeout: 5 * time.Second}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// 11. Method vs function
// (Here: attach a method to server)
func (s *server) URL() string {
	proto := "http"
	if s.secure {
		proto = "https"
	}
	return fmt.Sprintf("%s://%s", proto, s.addr)
}

// 12. Throttle closure: returns a function that only allows execution every interval
func throttle(interval time.Duration, fn func()) func() bool {
	var last time.Time
	return func() bool {
		now := time.Now()
		if now.Sub(last) >= interval {
			last = now
			fn()
			return true
		}
		return false
	}
}

// 13. Memoization
func memoize[K comparable, V any](f func(K) V) func(K) V {
	cache := make(map[K]V)
	return func(k K) V {
		if v, ok := cache[k]; ok {
			return v
		}
		v := f(k)
		cache[k] = v
		return v
	}
}

// 14. Generic function example (map of slice) already shown in apply; another:
func firstOrZero[T any](s []T) (zero T, _ T) {
	if len(s) == 0 {
		return zero, zero
	}
	return s[0], zero
}

// ------------------- TODO EXERCISES -------------------

// TODO(1): Implement fib(n int) int using iterative approach (no recursion) with O(1) space.
func fib(n int) int { panic("implement me") }

// TODO(2): Write a function timeCall(label string, fn func()) time.Duration that runs fn and returns duration.
func timeCall(label string, fn func()) time.Duration { panic("implement me") }

// TODO(3): Implement debounce(interval time.Duration, fn func()) func() that delays execution until silence period.
// Successive calls reset the timer; use a goroutine + channel + timer.
func debounce(interval time.Duration, fn func()) func() { panic("implement me") }

// TODO(4): Implement retry(max int, delay time.Duration, fn func() error) error with exponential backoff (delay * 2 each try).
func retry(max int, delay time.Duration, fn func() error) error { panic("implement me") }

// TODO(5): Implement captureLoopVars() []func() int that returns a slice of funcs each returning its own loop index (0..n-1).
// Classic closure capture pitfall exercise; ensure each closure returns distinct value.
func captureLoopVars() []func() int { panic("implement me") }

// TODO(6): Implement parallelMap[T,R any](in []T, workerCount int, fn func(T) R) []R preserving input order.
// Use goroutines + channel or WaitGroup; ensure deterministic order in result.
func parallelMap[T, R any](in []T, workerCount int, fn func(T) R) []R { panic("implement me") }

// TODO(7): Implement once(fn func()) func() that returns a function executing fn only the first time (thread-safe).
func once(fn func()) func() { panic("implement me") }

// TODO(8): Implement chain(funcs ...func()) func() that returns a function calling all in order.
func chain(funcs ...func()) func() { panic("implement me") }

// TODO(9): Implement compose[A,B,C any](f func(B) C, g func(A) B) func(A) C (classic functional compose).
func compose[A, B, C any](f func(B) C, g func(A) B) func(A) C { panic("implement me") }

// TODO(10): Implement safeGo(fn func()) that starts a goroutine and recovers panic logging it (print to stdout) without crashing main.
func safeGo(fn func()) { panic("implement me") }

// TODO(11): Implement timeoutCall(d time.Duration, fn func() error) error -> err if fn not finished in d (use channel + select).
func timeoutCall(d time.Duration, fn func() error) error { panic("implement me") }

// TODO(12): Implement pipeline(in []int) []int that: (a) squares numbers concurrently, (b) filters evens, (c) collects to slice.
// Use channels and close them properly.
func pipeline(in []int) []int { panic("implement me") }

// Optional: add tests to validate the above.

// ------------------- DEMO MAIN -------------------
func main() {
	q, err := div(10, 2)
	fmt.Println("div:", q, err)
	fmt.Println("areaRect:", areaRect(3, 4))
	fmt.Println("sum:", sum(1, 2, 3))
	c := counter(5)
	fmt.Println("counter:", c(), c(), c())
	fmt.Println("factorial 5:", factorial(5))
	fmt.Println("server URL:", NewServer("localhost:8080", WithSecure()).URL())

	th := throttle(50*time.Millisecond, func() { fmt.Println("throttle fired") })
	th()
	th()
	time.Sleep(60 * time.Millisecond)
	th()

	memo := memoize(func(n int) int { fmt.Println("compute", n); return n * n })
	fmt.Println("memo first:", memo(4))
	fmt.Println("memo second:", memo(4))

	fmt.Println("TODOs pending - implement exercises to extend demo")
}
