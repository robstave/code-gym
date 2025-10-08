package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			// Context was cancelled - clean up and exit
			fmt.Printf("Worker %d: stopping (reason: %v)\n", id, ctx.Err())
			return
		default:
			// Do work
			fmt.Printf("Worker %d: working...\n", id)

		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// Start multiple workers, all sharing the same context
	go worker(ctx, 1)
	go worker(ctx, 2)
	go worker(ctx, 3)

	// Let them work for 2 seconds
	time.Sleep(2 * time.Second)

	// Cancel all workers at once with a single call
	fmt.Println("Main: cancelling all workers...")
	cancel()

	// Give workers time to clean up
	time.Sleep(1 * time.Second)
	fmt.Println("Main: done")
}
