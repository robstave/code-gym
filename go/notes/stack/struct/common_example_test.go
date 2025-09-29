package common

import (
	"fmt"
	"testing"
)

// ExampleStack_basic demonstrates basic usage of the generic Stack.
func ExampleStack_basic() {
	var s Stack[int] // zero value ready
	fmt.Println("empty?", s.IsEmpty())

	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println("len after pushes:", s.Len())

	if top, ok := s.Peek(); ok {
		fmt.Println("peek:", top)
	}

	// Iterate from top to bottom
	s.ForEach(func(v int) { fmt.Printf("visit %d ", v) })
	fmt.Println()

	// Pop all
	for !s.IsEmpty() {
		v, _ := s.Pop()
		fmt.Printf("pop %d ", v)
	}
	fmt.Println()

	// Output:
	// empty? true
	// len after pushes: 3
	// peek: 30
	// visit 30 visit 20 visit 10 
	// pop 30 pop 20 pop 10 
}

// Basic sanity test.
func TestStackPushPop(t *testing.T) {
	var s Stack[string]
	s.Push("a"); s.Push("b")
	if s.Len() != 2 { t.Fatalf("expected len 2 got %d", s.Len()) }
	if v, _ := s.Peek(); v != "b" { t.Fatalf("expected top b got %s", v) }
	if v, ok := s.Pop(); !ok || v != "b" { t.Fatalf("expected pop b got %v %v", v, ok) }
	if v, ok := s.Pop(); !ok || v != "a" { t.Fatalf("expected pop a got %v %v", v, ok) }
	if _, ok := s.Pop(); ok { t.Fatalf("expected empty pop to fail") }
}
