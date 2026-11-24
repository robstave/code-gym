// Case 19: Launching jobs. Difficulty: Medium.
package main

import (
	"fmt"
	"sort"
)

func runJobs() []int {
	start := make(chan struct{})
	results := make(chan int, 3)
	var job int
	for job = 0; job < 3; job++ {
		go func() {
			<-start
			results <- job
		}()
	}
	close(start)

	ids := make([]int, 0, 3)
	for i := 0; i < 3; i++ {
		ids = append(ids, <-results)
	}
	sort.Ints(ids)
	return ids
}

func main() {
	fmt.Println("Job IDs:", runJobs())
}
