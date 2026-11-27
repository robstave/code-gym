# Rate limiting

Rate limiting restricts how many requests a client or operation can make in a period, helping protect shared capacity.

## How it is used

Choose an algorithm and apply a limit per user, API key, service, or other relevant identity. Reject excess requests, commonly with HTTP 429 and, where possible, a `Retry-After` header.

### Common algorithms

- **Fixed window:** Count requests in fixed periods, such as 100 per minute. It is simple and cheap, but clients can send almost twice the limit around a window boundary.
- **Sliding window log:** Record request timestamps and count those within the preceding window. This enforces a precise rolling limit, but storing and removing timestamps costs memory and work.
- **Sliding window counter:** Combine counts from the current and previous fixed windows, typically weighting the previous count by the overlapping fraction. It uses less memory than a log but approximates the rolling count and can allow or reject requests differently from an exact log.
- **Token bucket:** Refill tokens at a set rate up to a capacity; each accepted request consumes a token. It allows bursts up to the available capacity while limiting the sustained rate. Choose both refill rate and capacity carefully.
- **Leaky bucket:** Admit work into a bounded queue and process it at a steady rate, rejecting arrivals when the queue is full. It smooths traffic but adds queueing latency and can reject bursts even when clients could tolerate them. Some variants enforce a rate without actually queueing requests.

## When applicable

Use it for public APIs, expensive operations, abuse prevention, and fair sharing among tenants.

## Design tradeoffs

Strict limits protect the service but may reject legitimate bursts. Distributed enforcement requires shared coordination or accepts approximate limits; limits must be chosen for actual capacity and usage patterns.

Rate limiting controls a client's request rate; [load shedding](../load-shedding/readme.md) rejects work when the service itself is overloaded. Both can be used together.

[Back to concepts](../../readme.md)
