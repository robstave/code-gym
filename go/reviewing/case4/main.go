// Case 4: Completion rate. Difficulty: Easy.
package main

import "fmt"

func completionRate(done, total int) float64 {
	return float64(done / total)
}

func main() {
	fmt.Printf("Rate: %.2f\n", completionRate(3, 4))
}
