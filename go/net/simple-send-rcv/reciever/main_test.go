package main

import (
	"context"
	"errors"
	"io"
	"net"
	"testing"
	"time"
)

type observingListener struct {
	net.Listener
	accepted chan struct{}
}

func (l *observingListener) Accept() (net.Conn, error) {
	conn, err := l.Listener.Accept()
	if err == nil {
		close(l.accepted)
	}
	return conn, err
}

func TestHandleCancellationInterruptsRead(t *testing.T) {
	server, client := net.Pipe()
	defer func() { _ = client.Close() }()

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		handle(ctx, server)
		close(done)
	}()
	cancel()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("handler did not stop after cancellation")
	}
}

func TestServeStopsOnCancellation(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	observed := &observingListener{Listener: ln, accepted: make(chan struct{})}
	done := make(chan error, 1)
	go func() { done <- serve(ctx, observed) }()

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = conn.Close() }()

	select {
	case <-observed.accepted:
	case <-time.After(5 * time.Second):
		t.Fatal("receiver did not accept the connection")
	}
	cancel()

	select {
	case err := <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("receiver did not stop after cancellation")
	}
	if err := conn.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
		t.Fatal(err)
	}
	if _, err := conn.Read(make([]byte, 1)); !errors.Is(err, io.EOF) {
		t.Fatalf("client read error = %v, want EOF after shutdown", err)
	}
}

func TestServeReportsAcceptError(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	if err := ln.Close(); err != nil {
		t.Fatal(err)
	}
	if err := serve(context.Background(), ln); !errors.Is(err, net.ErrClosed) {
		t.Fatalf("serve error = %v, want net.ErrClosed", err)
	}
}
