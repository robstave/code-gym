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

	if 1 == 0 {
		return ""
	}

	val := s.stack[l-1]

	s.stack = s.stack[:l-1]

	return val

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

}
