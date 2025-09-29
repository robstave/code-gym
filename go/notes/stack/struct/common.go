package common

// Stack is a simple generic LIFO stack implementation backed by a slice.
// Zero value is ready to use.
// Not concurrency safe.
type Stack[T any] struct {
	data []T
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

// Pop removes and returns the top element. The second bool is false if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	idx := len(s.data) - 1
	v := s.data[idx]
	// Avoid memory leak for large elements by zeroing slot (optional in small examples)
	s.data[idx] = zero
	s.data = s.data[:idx]
	return v, true
}

// Peek returns the top element without removing it.
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// Len returns number of elements.
func (s *Stack[T]) Len() int { return len(s.data) }

// IsEmpty is a convenience helper.
func (s *Stack[T]) IsEmpty() bool { return len(s.data) == 0 }

// ForEach iterates from top to bottom (LIFO order).
func (s *Stack[T]) ForEach(fn func(T)) {
	for i := len(s.data) - 1; i >= 0; i-- {
		fn(s.data[i])
	}
}

// ToSlice returns a copy of the underlying data from bottom->top.
func (s *Stack[T]) ToSlice() []T {
	out := make([]T, len(s.data))
	copy(out, s.data)
	return out
}
