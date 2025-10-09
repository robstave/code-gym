package main

import (
	"fmt"
)

type Stack struct {
	stack []string
}

func (s *Stack) Push(value string) {
	s.stack = append(s.stack, value)
}

func (s *Stack) Pop() string {
	l := len(s.stack)
	if l == 0 {
		return ""
	}
	val := s.stack[l-1]
	s.stack = s.stack[:l-1]
	return val
}

func (s *Stack) Peek() string {
	l := len(s.stack)
	if l == 0 {
		return ""
	}
	return s.stack[l-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func main() {
	fmt.Println("-------")

	s := Stack{}

	s.Push("sss")
	s.Push("aaa")

	val := s.Pop()
	fmt.Printf("%s\n", val)

	val = s.Pop()
	fmt.Printf("%s\n", val)

	fmt.Printf("Is stack empty? %v\n", s.IsEmpty())

	s.Push("new")
	fmt.Printf("Peek: %s\n", s.Peek())
}
