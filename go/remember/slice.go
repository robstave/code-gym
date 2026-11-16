package main

import "fmt"

func change(s []int) {
	// The slice header is copied, but s and the caller still share the same
	// backing array, so changing an element is visible to the caller.
	s[0] = 99
}

func add(s []int) {
	// append returns a slice with a possibly new length (and backing array).
	// Assigning it to s changes only this function's copy of the slice header.
	s = append(s, 4) // only the copy's length changed
}

func addMutate(s []int) []int {
	// Return the updated slice so the caller can keep its new length/header.
	s = append(s, 4)
	return s
}

func main() {
	a := []int{1, 2, 3}
	change(a)
	fmt.Printf("Length: %v array: %v\n", len(a), a) // element change is visible [99 2 3]

	add(a)
	fmt.Printf("Length: %v array: %v\n", len(a), a) // caller's length is unchanged [99 2 3]

	a = addMutate(a)
	fmt.Printf("Length: %v array: %v\n", len(a), a) // caller keeps the returned length  [99 2 3 4]

}
