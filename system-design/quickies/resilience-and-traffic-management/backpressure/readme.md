# Backpressure

Backpressure slows or limits producers when a consumer cannot keep up, preventing unbounded queues and resource exhaustion.

## How it is used

Use bounded queues, concurrency limits, or flow-control signals. When capacity is exhausted, producers can wait, slow down, or receive a rejection they can handle.

## When applicable

Use it between services, in streaming pipelines, and anywhere producers can send work faster than it can be processed.

## Design tradeoffs

Waiting increases latency and can propagate congestion upstream; rejection requires clients to handle errors. Queue limits protect resources but require careful sizing and monitoring.

[Back to concepts](../../readme.md)
