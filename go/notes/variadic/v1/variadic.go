package main

import (
	"fmt"
	"strings"
)

// Basic variadic function - sum numbers
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total = total + v
	}
	return total
}

// Variadic with fixed parameters first
func greetUsers(greeting string, names ...string) {

	for _, n := range names {
		fmt.Printf("%s: %s\n", greeting, n)
	}

}

func buildSentence(words ...string) string {
	if len(words) == 0 {
		return ""
	}
	return strings.Join(words, " ") + "."
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0 // or panic, depending on requirements
	}

	// since this is a slice, you can start with this too
	maximum := nums[0]
	for _, num := range nums[1:] {
		if num > maximum {
			maximum = num
		}
	}
	return maximum
}

// Variadic function that takes other functions
func applyToAll(fn func(int) int, nums ...int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = fn(num)
	}
	return result
}

// Helper function for the example above
func square(x int) int {
	return x * x
}

func main() {

	sumResult := sum(2, 3, 4)
	fmt.Printf("%d\n", sumResult)

	items := []int{2, 4, 6}
	// you can pass slices into this, but you need the ... to do it
	sumResult = sum(items...)
	fmt.Printf("%d\n", sumResult)

	greetUsers("Greeting", "rob", "ken", "mike")
	sliceONames := []string{"rob", "ken", "mike"}
	greetUsers("Greeting2", sliceONames...)

	// note: sliceONames := [...]string{"rob", "ken", "mike"}
	// will not work as its an array and not a slice

	fmt.Printf("-------\n")

	fmt.Printf("%s\n", buildSentence("How", "now", "brown", "cow"))

	fmt.Printf("Max %d\n", max(4, 3, 5, 1, 5, 6))

	fmt.Println("\n=== Apply Function to All ===")
	squared := applyToAll(square, 1, 2, 3, 4, 5)
	fmt.Println("Squared:", squared)

	// Using anonymous function
	doubled := applyToAll(func(x int) int { return x * 2 }, 10, 20, 30)
	fmt.Println("Doubled:", doubled)

}
