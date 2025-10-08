package main

import (
	"fmt"
	"time"
)

// task is a simple unit of work.
// NOTE: We pass tasks BY VALUE through the channels below.
// In Go, a struct value is copied on send and receive. That means:
//   - The worker receives its own copy of the task (j below)
//   - Mutating fields (j.value = "done") only affects that local copy
//   - When we send the modified j on the results channel, THAT copy (with value mutated) is copied again
//
// For small structs this is fine; for large structs or to share mutable state, you'd usually
// send *task (a pointer) instead, so all goroutines refer to the same underlying data.
type task struct {
	value    string
	jobCount int
}

// worker is a goroutine that continually reads from the jobs channel until it is closed.
// Channel direction annotations:
//
//	jobs    <-chan task   (receive-only for this function)
//	results chan<- task   (send-only for this function)
//
// This documents intent and the compiler enforces directional usage inside the function body.
func worker(id int, jobs <-chan task, results chan<- task) {
	for j := range jobs { // range ends when jobs is closed AND drained
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second) // simulate work

		// Because j is a COPY of the task struct, this mutation only changes the copy
		// we're about to forward on the results channel (not the original value sent).
		j.value = "done"
		fmt.Println("worker", id, "finished job", j)

		results <- j // send modified copy
	}
}

func main() {
	// Number of jobs to produce
	const numJobs = 5

	// Buffered channels sized to the total number of jobs so sends won't block
	// once the workers are running. (For illustration; unbuffered also works.)
	jobs := make(chan task, numJobs)
	results := make(chan task, numJobs)

	// Launch a small worker pool (3 workers). They immediately begin waiting on jobs.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Produce jobs. Each send copies the task value into the channel buffer.
	for j := 1; j <= numJobs; j++ {
		jobs <- task{value: "Started", jobCount: j}
	}
	close(jobs) // signals no more jobs; workers finish after draining existing jobs.

	// Collect results. Each receive gives us the copy sent by the worker (with value mutated to "done").
	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-results)
	}
}

// POINTER VS VALUE SUMMARY
// ------------------------
// Do you "need" pointers to mutate? It depends on what you want to mutate and who should observe it.
// In this example, we mutate j.value inside the worker, then send the modified struct on the results channel.
// That works fine with values because the mutation is reflected in the value we send out.
// You only need a pointer if:
//   1. Multiple goroutines must observe changes to the SAME instance (shared state), or
//   2. The struct is large and copying it frequently would be a performance issue, or
//   3. You want methods with pointer receivers (e.g., to avoid copying or to mutate in methods), or
//   4. You store the value in an interface that requires pointer receiver methods to satisfy it.
// Using *task would change channel types to chan *task. Then each send/receive moves a pointer (cheap)
// and all mutations happen on the single underlying task object.
// For small immutable-ish job descriptors, passing by value is idiomatic and clear.
