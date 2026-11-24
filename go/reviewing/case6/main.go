// Case 6: Saved snapshot. Difficulty: Medium.
package main

import "fmt"

func snapshot(scores []int) []int {
	saved := scores
	return saved
}

func main() {
	scores := []int{10, 20, 30}
	saved := snapshot(scores)
	scores[0] = 99
	fmt.Println("First saved score:", saved[0])
}
