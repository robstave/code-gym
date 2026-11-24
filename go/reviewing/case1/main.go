// Case 1: Order total. Difficulty: Easy.
package main

import "fmt"

func totalPrices(prices []int) int {
	total := 0
	for i := 0; i < len(prices)-1; i++ {
		total += prices[i]
	}
	return total
}

func main() {
	fmt.Println("Total:", totalPrices([]int{2, 5, 3}))
}
