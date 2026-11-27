# Write amplification

Write amplification occurs when one logical write causes multiple physical writes. For example, updating a row may also update indexes, logs, replicas, and storage files.

## How it is used

Measure write throughput and disk activity while considering indexes, replication, and storage-engine behavior. Remove unnecessary indexes or batch writes where the workload allows.

## When applicable

Investigate it when write-heavy workloads encounter storage bottlenecks, high costs, or excessive flash wear.

## Design tradeoffs

Indexes and replicas increase write work but can improve reads, availability, and durability. Batching may improve throughput at the cost of latency or more complex failure handling.

[Back to concepts](../../readme.md)
