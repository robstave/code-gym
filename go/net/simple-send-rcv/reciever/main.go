package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
)

// handle reads one newline-terminated message from a single client.
func handle(ctx context.Context, conn net.Conn) {
	// Closing the connection on cancellation interrupts a blocked read, so an
	// idle client cannot prevent the receiver from shutting down.
	stop := context.AfterFunc(ctx, func() {
		_ = conn.Close()
	})
	defer func() {
		stop()
		_ = conn.Close()
	}()

	line, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		// A read interrupted by shutdown is expected, not a client error.
		if ctx.Err() == nil {
			log.Println("read error:", err)
		}
		return
	}
	fmt.Printf("received from %s: %q\n", conn.RemoteAddr(), line)
}

// serve accepts clients until the context is canceled or accepting fails.
// It owns the listener and waits for all handlers before returning.
func serve(ctx context.Context, ln net.Listener) error {
	// A child context also lets an unexpected accept error stop active handlers.
	ctx, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup

	// Closing the listener interrupts Accept when the context is canceled.
	stop := context.AfterFunc(ctx, func() {
		_ = ln.Close()
	})
	defer func() {
		// Stop all handlers before waiting, including after an accept error.
		cancel()
		stop()
		_ = ln.Close()
		wg.Wait()
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			if ctx.Err() != nil {
				// Closing the listener during shutdown is expected.
				return nil
			}
			return fmt.Errorf("accept connection: %w", err)
		}
		// Register each handler before starting it so shutdown waits for it.
		wg.Add(1)
		go func() {
			defer wg.Done()
			handle(ctx, conn)
		}()
	}
}

func main() {
	// Ctrl+C cancels the context shared by the listener and all handlers.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	ln, err := net.Listen("tcp", ":5002")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listening on :5002")

	if err := serve(ctx, ln); err != nil {
		log.Fatal(err)
	}
}
