# Read amplification

Read amplification occurs when satisfying one logical read requires multiple physical reads or lookups, such as reading several storage files or querying multiple shards.

## How it is used

Inspect query plans and storage metrics to identify extra lookups. Indexes, compaction, denormalization, or caching can reduce the work required for common reads.

## When applicable

This is useful when read latency or disk usage is high despite modest application-level query volume.

## Design tradeoffs

Indexes, compaction, and precomputed views can speed reads but add storage, write work, or maintenance. Caches reduce reads at the cost of possible staleness.

[Back to concepts](../../readme.md)
