package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func main() {
	// Create a new stack
	s := stack.New()

	fmt.Println("=== Stack Operations Demo ===")

	// Push elements onto the stack
	fmt.Println("\n1. Pushing elements...")
	s.Push("first")
	s.Push("second")
	s.Push(42)
	s.Push(3.14)
	s.Push([]string{"hello", "world"})

	fmt.Printf("Stack length after pushes: %d\n", s.Len())

	// Peek at the top element without removing it
	fmt.Println("\n2. Peeking at top element...")
	if s.Len() > 0 {
		top := s.Peek()
		fmt.Printf("Top element: %v\n", top)
		fmt.Printf("Stack length after peek: %d\n", s.Len())
	}

	// Pop elements from the stack (LIFO - Last In, First Out)
	fmt.Println("\n3. Popping elements...")
	for s.Len() > 0 {
		element := s.Pop()
		fmt.Printf("Popped: %v (type: %T), remaining: %d\n",
			element, element, s.Len())
	}

	// Demonstrate stack is empty
	fmt.Println("\n4. Stack operations on empty stack...")
	fmt.Printf("Stack length: %d\n", s.Len())

	// Peek on empty stack returns nil
	emptyPeek := s.Peek()
	fmt.Printf("Peek on empty stack: %v\n", emptyPeek)

	// Pop on empty stack returns nil
	emptyPop := s.Pop()
	fmt.Printf("Pop on empty stack: %v\n", emptyPop)

	// Practical example: Expression evaluation helper
	fmt.Println("\n=== Practical Example: Matching Parentheses ===")
	expressions := []string{
		"(())",
		"(()",
		"())",
		"(()())",
		"",
	}

	for _, expr := range expressions {
		result := checkParentheses(expr)
		fmt.Printf("Expression '%s': %s\n", expr, result)
	}
}

// checkParentheses demonstrates a practical use of stack
// for checking balanced parentheses in an expression
func checkParentheses(expression string) string {
	s := stack.New()

	for _, char := range expression {
		switch char {
		case '(':
			s.Push(char)
		case ')':
			if s.Len() == 0 {
				return "Unbalanced - extra closing parenthesis"
			}
			s.Pop()
		}
	}

	if s.Len() == 0 {
		return "Balanced"
	}
	return "Unbalanced - unclosed parentheses"
}
