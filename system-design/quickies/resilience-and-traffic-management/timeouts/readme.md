# Timeouts

A timeout limits how long an operation may wait before it is treated as failed. A deadline sets an overall time limit for the work.

## How it is used

Set appropriate connection and request timeouts, propagate deadlines to downstream calls, and cancel work when possible. Leave room in the overall budget for retries or fallback.

## When applicable

Use timeouts for network calls, database queries, and other operations that might stall.

## Design tradeoffs

Short timeouts free resources quickly but can fail slow, valid requests. Long timeouts tolerate latency but tie up resources; a timeout does not guarantee that the remote operation stopped or had no effect.

[Back to concepts](../../readme.md)
