# Replication lag

Replication lag is the delay between a write being applied on a primary and becoming visible on a replica.

## How it is used

Monitor lag and route freshness-sensitive reads to the primary or use a database-supported consistency mechanism. Less sensitive reads can use replicas to spread load.

## When applicable

Lag matters whenever replicas serve reads, especially after a user writes data and immediately expects to see the update.

## Design tradeoffs

Replica reads can improve capacity but may be stale. Primary reads improve freshness at the cost of primary load; waiting for replicas can add latency and reduce availability.

[Back to concepts](../../readme.md)
