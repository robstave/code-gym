package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
)

// Greet writes one newline-terminated message to any io.Writer. The newline
// lets the receiver recognize the end of the message on the TCP stream.
func Greet(writer io.Writer, name string) error {
	_, err := fmt.Fprintf(writer, "Hello, %s\n", name)
	return err
}

// sendGreeting opens one connection, sends one greeting, and closes it.
// Each job calls this independently, so the jobs do not share a connection.
func sendGreeting(ctx context.Context, address string) error {
	// DialContext abandons a pending connection attempt if ctx is canceled.
	conn, err := (&net.Dialer{}).DialContext(ctx, "tcp", address)
	if err != nil {
		return err
	}
	// Close the connection when this function returns. A cancellation callback
	// may already have closed it, so there is no useful close error to handle.
	defer func() {
		_ = conn.Close()
	}()

	// DialContext only applies the context while connecting. Once connected,
	// closing the connection on cancellation also interrupts a blocked write.
	stop := context.AfterFunc(ctx, func() {
		_ = conn.Close()
	})
	// Prevent the callback from running later if the job finishes first.
	// If it has already started, stop does not wait for it to finish.
	defer stop()

	if err := Greet(conn, "world"); err != nil {
		// If cancellation caused (or coincided with) the write error, report the
		// context error so the caller can recognize cancellation.
		if ctx.Err() != nil {
			return ctx.Err()
		}
		return err
	}
	// A successful write means the bytes were written to the connection; it
	// does not confirm that the receiver processed the message.
	return nil
}

// run starts three independent jobs and waits for every job before returning.
// It collects all job errors rather than stopping at the first failure.
func run(ctx context.Context, address string) error {
	const jobs = 3
	var wg sync.WaitGroup
	// Each job sends at most one error. Buffering space for every job means
	// none can block on sending while the main goroutine waits for completion.
	errs := make(chan error, jobs)

	for i := 0; i < jobs; i++ {
		// Register the job before starting it so Wait cannot miss it.
		wg.Add(1)
		go func(job int) {
			// Done runs on every return path, including a failed send.
			defer wg.Done()
			if err := sendGreeting(ctx, address); err != nil {
				// Wrap the error with its job number while preserving the
				// underlying error for errors.Is and errors.As.
				errs <- fmt.Errorf("job %d: %w", job, err)
			}
			// Pass the number explicitly to give each goroutine its own value.
		}(i + 1)
	}

	// Only close the channel after all senders have finished; sending on a
	// closed channel would panic. Ranging below then drains the buffered errors.
	wg.Wait()
	close(errs)

	var failures []error
	for err := range errs {
		failures = append(failures, err)
	}
	// Join returns nil when there are no failures, or one error containing all
	// failures otherwise. The order can vary because the jobs run concurrently.
	return errors.Join(failures...)
}

func main() {
	// NotifyContext cancels ctx when the process receives an interrupt, such as
	// Ctrl+C. The shared context lets all three jobs react to the same signal.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	// Release the signal resources when main returns.
	defer stop()

	// run waits for the jobs even after cancellation. A failure is logged and
	// gives the process a nonzero exit status; success exits normally.
	if err := run(ctx, "localhost:5002"); err != nil {
		log.Fatal(err)
	}
}
