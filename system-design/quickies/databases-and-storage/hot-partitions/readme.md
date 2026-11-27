# Hot partitions

A hot partition receives a disproportionate share of traffic or data, making it a bottleneck while other partitions have spare capacity.

## How it is used

Monitor traffic, latency, and resource usage per partition. Depending on the access pattern, spread load with a different key, key salting, caching, or workload-specific replication.

## When applicable

Look for hot partitions in sharded databases, event streams, and caches, especially when popular users or keys dominate traffic.

## Design tradeoffs

Splitting or salting keys spreads load but can make reads and aggregation more complex. Caching or replicas help some read-heavy cases but do not necessarily solve write hotspots.

[Back to concepts](../../readme.md)
