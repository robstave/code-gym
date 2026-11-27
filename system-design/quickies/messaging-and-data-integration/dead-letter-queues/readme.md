# Dead-letter queues

A dead-letter queue (DLQ) holds messages that could not be processed successfully after a defined policy, keeping them available for investigation.

## How it is used

After bounded retries, move or route failed messages to a DLQ with failure details. Alert operators and provide a safe way to inspect, fix, and replay them.

## When applicable

Use a DLQ when individual bad messages should not indefinitely block a consumer or partition.

## Design tradeoffs

Moving a message aside can break processing order or leave downstream state incomplete. DLQs require retention, monitoring, secure handling, and replay procedures; they do not fix the underlying error.

[Back to concepts](../../readme.md)
