# Service discovery

Service discovery lets a service find the network addresses of other service instances as those instances start, move, or stop.

## How it is used

Instances register with a discovery system or are exposed through platform-managed DNS. Clients or a proxy use that information to route requests to available instances, often alongside health checks.

## When applicable

Use discovery when service addresses change dynamically, such as in an autoscaled container environment.

## Design tradeoffs

Discovery reduces hard-coded configuration but adds infrastructure and another component to monitor. Stale registrations or DNS caches can route traffic to unavailable instances, so timeouts and health handling remain necessary.

[Back to concepts](../../readme.md)
