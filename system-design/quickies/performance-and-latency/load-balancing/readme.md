# Load balancing

Load balancing distributes requests across multiple service instances instead of sending all traffic to one.

## How it is used

A load balancer selects a healthy backend using a policy such as round-robin, least connections, or hashing by a key. Health checks remove unsuitable backends from routing.

## When applicable

Use it when a service runs on multiple instances for capacity, availability, or both.

## Design tradeoffs

Round-robin is simple but ignores differing request costs; least-connections needs accurate state and may not reflect actual load. Hash-based routing preserves affinity but can create hotspots. The balancer itself must be scaled and made resilient.

[Back to concepts](../../readme.md)
