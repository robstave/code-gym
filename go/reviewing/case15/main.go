// Case 15: Rank scores. Difficulty: Easy.
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func rankScores(scores []int) []string {
	result := make([]string, 0, len(scores))
	for _, score := range scores {
		result = append(result, strconv.Itoa(score))
	}
	sort.Strings(result)
	return result
}

func main() {
	fmt.Println("Ranked:", rankScores([]int{2, 10, 3}))
}
