# Service communication

Services communicate synchronously through requests or asynchronously through messages and events. Each approach changes how services depend on one another.

## How it is used

Use synchronous HTTP or RPC when a caller needs an immediate response. Publish an event or queue a command when work can happen later, and use correlation IDs to trace the workflow.

## When applicable

Synchronous calls suit immediate lookups and decisions; asynchronous messaging suits background work and independently processed updates.

## Design tradeoffs

Synchronous calls are straightforward but couple availability and latency across services. Messaging reduces direct timing dependencies but adds eventual consistency, duplicate handling, and operational complexity. Both need clear contracts, versioning, and failure handling.

[Back to concepts](../../readme.md)
