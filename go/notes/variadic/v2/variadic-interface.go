package main

import (
	"fmt"
)

// printAll demonstrates a variadic function that accepts values of
// any type by using the empty interface type: interface{}.
//
//   - items ...interface{} means callers can pass zero or more values
//     of any type. Inside the function, items is a slice of interface{}.
//   - We iterate the slice and print each value and its concrete type.
func printAll(items ...interface{}) {
	for i, item := range items {
		// %v prints the value, %T prints the concrete type held by the interface
		fmt.Printf("Item %d: %v (type: %T)\n", i+1, item, item)
	}
}

// quackAttack accepts a variadic list of values that satisfy the quack
// interface. This demonstrates two things:
// 1) Defining behavior with interfaces (doQuack method)
// 2) Passing a slice to a variadic parameter using the '...' spread operator
func quackAttack(animals ...quack) {
	for _, animal := range animals {
		// Call the interface method; the concrete type provides the implementation
		fmt.Printf("%s\n", animal.doQuack())
	}
}

// quack is a small interface that requires a single method: doQuack.
// Any type that implements doQuack() string implicitly satisfies this interface.
type quack interface {
	doQuack() string
}

// Two concrete types that implement the quack interface.
type mallard struct{}
type dog struct{}
type decoy struct{}

// mallard implements doQuack. The receiver is a value receiver because
// mallard has no fields and there's no need for a pointer.
func (m mallard) doQuack() string {
	return "Quack!"
}

// dog implements doQuack; it uses the same method signature so dog also
// satisfies the quack interface. Note the playful "Qoof!" return.
func (d dog) doQuack() string {
	return "Qoof!"
}

func (d *decoy) doQuack() string {
	return "silent!"
}

func main() {
	// Demonstrate the variadic printAll function with mixed types
	fmt.Println("\n=== Print All Types ===")
	printAll(42, "hello", 3.14, true, []int{1, 2, 3})

	// Create concrete values
	dog1 := dog{}
	mallard1 := mallard{}
	decoy1 := &decoy{}

	// Build a slice of the interface type. Both dog and mallard satisfy quack
	animals := []quack{dog1, mallard1, decoy1}

	// Use the '...' operator to expand the slice into individual arguments
	// for the variadic function quackAttack.
	quackAttack(animals...)
}
