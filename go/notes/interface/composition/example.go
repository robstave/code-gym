package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
INTERFACE & COMPOSITION EXAMPLE
================================

This file demonstrates two related (but distinct) kinds of composition in Go:

1. Interface composition (an interface embedding other interfaces)
2. Struct composition via embedding (a struct anonymously embedding another struct)

We will model a tiny document processing system.

Core ideas showcased:
---------------------
Interface composition:
  - Reader:    something that can Read() string
  - Writer:    something that can Write(string) error
  - Flusher:   something that can Flush() error
  - ReadWriterFlusher: composed interface embedding Reader, Writer, Flusher

Struct embedding:
  - BaseBuffer: provides shared fields + methods
  - MemoryBuffer: embeds BaseBuffer and implements the composed interface
  - LoggerBuffer: wraps/embeds another ReadWriterFlusher to add logging (decorator style)

Additional patterns included:
  - Interface satisfaction by method set (implicit, no "implements" keyword)
  - Returning interface types vs concrete types
  - Using zero values and constructor helpers
  - Error wrapping

Read from top to bottom; main() runs illustrative scenarios.
*/

// ----------- Interface Composition -----------

// Reader defines the ability to read the current buffer contents.
type Reader interface {
	Read() (string, error)
}

// Writer defines the ability to append data to the buffer.
type Writer interface {
	Write(data string) error
}

// Flusher defines the ability to clear underlying state.
type Flusher interface {
	Flush() error
}

// ReadWriterFlusher composes the three smaller interfaces. Any type
// implementing all embedded methods satisfies this larger interface.
type ReadWriterFlusher interface {
	Reader
	Writer
	Flusher
}

// ----------- Struct Composition (Embedding) -----------

// BaseBuffer holds shared state. It's not exported (lowercase) because we
// don't intend users to rely on it directly; they should use higher-level types.
type BaseBuffer struct {
	data []string
	cap  int // simple capacity limit for demonstration
}

// newBaseBuffer constructs a BaseBuffer with a capacity.
func newBaseBuffer(capacity int) BaseBuffer {
	buffer := BaseBuffer{
		data: make([]string, 0, capacity),
		cap:  capacity,
	}
	return buffer
}

// size returns number of elements currently stored.
func (b *BaseBuffer) size() int {
	return len(b.data)
}

// remaining returns remaining capacity.
func (b *BaseBuffer) remaining() int {
	return b.cap - len(b.data)
}

// MemoryBuffer composes behavior by embedding BaseBuffer anonymously.
// Its method set now includes promoted methods/fields (size, remaining, etc.).
// MemoryBuffer will implement Read, Write, Flush, so it satisfies ReadWriterFlusher.
type MemoryBuffer struct {
	BaseBuffer
	name string
}

// NewMemoryBuffer is a public constructor returning the interface, not the concrete
// type. This hides implementation details and allows swapping later.
func NewMemoryBuffer(name string, capacity int) ReadWriterFlusher {
	mem := &MemoryBuffer{
		BaseBuffer: newBaseBuffer(capacity),
		name:       name,
	}
	return mem
}

// Write implements Writer.
func (mb *MemoryBuffer) Write(data string) error {
	if mb.remaining() <= 0 {
		return fmt.Errorf("buffer %s full: cannot write '%s'", mb.name, data)
	}
	mb.data = append(mb.data, data)
	return nil
}

// Read implements Reader. It joins stored strings.
func (mb *MemoryBuffer) Read() (string, error) {
	if mb.size() == 0 {
		return "", errors.New("buffer empty")
	}
	return strings.Join(mb.data, " "), nil
}

// Flush implements Flusher. We re-slice to zero length while keeping capacity.
func (mb *MemoryBuffer) Flush() error {
	mb.data = mb.data[:0]
	return nil
}

// ----------- Decorator via Embedding / Wrapping -----------

// LoggerBuffer wraps another ReadWriterFlusher to augment behavior (composition over inheritance).
// We embed the interface so its methods are promoted; we can override selectively.
type LoggerBuffer struct {
	ReadWriterFlusher
	id   string
	logs []string
}

// NewLoggerBuffer creates a logging decorator around an existing buffer.
func NewLoggerBuffer(id string, inner ReadWriterFlusher) *LoggerBuffer {
	lb := LoggerBuffer{
		ReadWriterFlusher: inner,
		id:                id,
		logs:              make([]string, 0, 16),
	}
	return &lb
}

// Write intercepts the call, logs, then delegates to the wrapped implementation.
func (lb *LoggerBuffer) Write(data string) error {
	lb.logs = append(lb.logs, fmt.Sprintf("WRITE(%s): %s", lb.id, data))
	return lb.ReadWriterFlusher.Write(data)
}

// Flush intercepts flush to add a log entry.
func (lb *LoggerBuffer) Flush() error {
	lb.logs = append(lb.logs, fmt.Sprintf("FLUSH(%s)", lb.id))
	return lb.ReadWriterFlusher.Flush()
}

// LogDump returns concatenated logs. (Not part of composed interface—extra behavior.)
func (lb *LoggerBuffer) LogDump() string {
	return strings.Join(lb.logs, " | ")
}

// ----------- Utility Functions Consuming Interfaces -----------

// writeAll attempts to write a list of strings. Demonstrates accepting the *composed* interface.
func writeAll(buf ReadWriterFlusher, items ...string) error {
	for _, it := range items {
		if err := buf.Write(it); err != nil {
			return fmt.Errorf("writeAll failed after '%s': %w", it, err)
		}
	}
	return nil
}

// resetAndFill demonstrates reusing the same interface value through lifecycle events.
func resetAndFill(buf ReadWriterFlusher, fill string) error {
	if err := buf.Flush(); err != nil {
		return err
	}
	return buf.Write(fill)
}

// printBuffer tries a read and prints either data or the error.
func printBuffer(label string, buf ReadWriterFlusher) {
	data, err := buf.Read()
	if err != nil {
		fmt.Printf("[%s] read error: %v\n", label, err)
		return
	}
	fmt.Printf("[%s] contents: %q\n", label, data)
}

// ----------- main: walkthrough -----------

func main() {
	// Create a base memory buffer (through constructor returning interface).
	buf := NewMemoryBuffer("primary", 4)

	// Wrap it with a logging decorator. Note: logger itself still satisfies ReadWriterFlusher.
	logged := NewLoggerBuffer("L1", buf)

	// Write multiple items.
	if err := writeAll(logged, "hello", "composed", "interfaces"); err != nil {
		fmt.Println("writeAll error:", err)
	}
	printBuffer("after writes", logged)

	// Demonstrate capacity handling by not exceeding capacity.
	if err := logged.Write("overflow? no "); err != nil {
		fmt.Println("Not expected error:", err)
	}

	// Demonstrate capacity handling by exceeding capacity.
	if err := logged.Write("overflow2? yes "); err != nil {
		fmt.Println("expected error:", err)
	}

	// Reset and put a single value back.
	if err := resetAndFill(logged, "reset-complete"); err != nil {
		fmt.Println("reset error:", err)
	}
	printBuffer("after reset", logged)

	// Show logs from decorator (extra functionality not in interface) by using concrete type.
	fmt.Println("logs:", logged.LogDump())

	// Show that we can pass either the original buf or the logged version interchangeably
	// because both satisfy the composed interface.
	_ = writeAll(buf, "direct", "write")
	printBuffer("original buf (no logs)", buf)
}
