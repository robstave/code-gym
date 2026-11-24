# Simple TCP sender and receiver

This example sends three greetings over TCP. It requires Go 1.21 or later.

## Run

In a terminal, start the receiver from this directory:

```powershell
go run .\reciever\main.go
```

In a second terminal, start the sender:

```powershell
go run .\sender\main.go
```

The receiver listens on port 5002. It prints a line for each greeting and
continues running after the sender exits. Stop it with Ctrl+C.

## How it works

- The sender starts three goroutines. Each opens a **separate TCP connection**
  to `localhost:5002`, sends `Hello, world` followed by a newline, and closes
  its connection.
- The receiver accepts each connection and starts a goroutine that reads and
  prints one newline-terminated message. The order of the messages is not
  guaranteed; their remote addresses normally have different source ports.
- The sender uses a wait group to wait for all three jobs, then exits. It
  reports any connection or write errors. A successful write does not confirm
  that the receiver processed the message: there is no acknowledgement.
- The sender's context is canceled by Ctrl+C. Cancellation interrupts pending
  connection attempts and closes established connections to interrupt writes.
  The sender still waits for the jobs to finish before exiting.
- The receiver also uses a signal context. On Ctrl+C, it closes the listener
  and active connections to interrupt blocked accepts and reads, then waits for
  its handlers to finish before exiting.

## Separate or shared connections?

Both approaches are normal. Separate connections keep jobs independent and let
the receiver handle them concurrently, but opening and closing connections adds
overhead. This is reasonable for three occasional greetings.

A persistent shared connection can be more efficient for a steady stream of
messages. It requires a protocol to delimit messages and handle any responses,
and concurrent senders may need coordination to prevent messages from
interfering. The current receiver reads only one line per connection and then
closes it, so it would need to change to read multiple lines before the sender
could reuse a connection.

To run the sender tests from this directory without a Go module:

```powershell
$env:GO111MODULE = "off"
go test .\sender\...
```
