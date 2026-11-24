package main

import (
	"bufio"
	"context"
	"errors"
	"net"
	"testing"
	"time"
)

func TestRunSendsThreeGreetings(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = listener.Close()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	done := make(chan error, 1)
	go func() {
		done <- run(ctx, listener.Addr().String())
	}()

	if err := listener.(*net.TCPListener).SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 3; i++ {
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}
		if err := conn.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
			_ = conn.Close()
			t.Fatal(err)
		}
		line, readErr := bufio.NewReader(conn).ReadString('\n')
		_ = conn.Close()
		if readErr != nil {
			t.Fatal(readErr)
		}
		if line != "Hello, world\n" {
			t.Errorf("greeting = %q, want %q", line, "Hello, world\n")
		}
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
}

func TestRunWithCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if err := run(ctx, "127.0.0.1:1"); !errors.Is(err, context.Canceled) {
		t.Fatalf("run error = %v, want context.Canceled", err)
	}
}

func TestRunReportsConnectionErrors(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	address := listener.Addr().String()
	if err := listener.Close(); err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := run(ctx, address); err == nil {
		t.Fatal("run succeeded with no listener")
	}
}
