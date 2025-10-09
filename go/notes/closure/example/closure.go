package main

import "fmt"

// multiplier returns a closure that multiplies its input by a fixed factor.
// The closure captures the value of the 'factor' variable when it is created.
func multiplier(factor int) func(int) int {
	return func(x int) int {
		// 'factor' is captured by value, so changes to the original variable
		// outside this function will not affect the closure.
		return x * factor
	}
}

// multiplier2 returns a closure that multiplies its input by a factor.
// The closure captures a pointer to the 'factor' variable, allowing it
// to reflect changes to the variable outside the closure.
func multiplier2(factor *int) func(int) int {
	return func(x int) int {
		// 'factor' is captured by reference (pointer), so changes to the
		// original variable outside this function will affect the closure.
		return x * (*factor)
	}
}

func main() {
	// Create a closure that doubles its input.
	double := multiplier(2)
	// Create a closure that triples its input.
	triple := multiplier(3)

	fmt.Println(double(5)) // 10 (5 * 2)
	fmt.Println(triple(5)) // 15 (5 * 3)

	d := 10                      // plain int
	multipass := multiplier2(&d) // capture pointer to d

	fmt.Println(multipass(5)) // 50 (5 * 10)

	// Changing the value of 'd' updates the factor used by the closure.
	d = 7
	fmt.Println(multipass(5)) // 35 (5 * 7)
}
