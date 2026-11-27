# Bulkheads

Bulkheads isolate resources so a problem in one workload or dependency does not exhaust capacity needed by others.

## How it is used

Give separate workloads their own connection pools, worker pools, queues, or concurrency limits. For example, isolate slow report generation from interactive requests.

## When applicable

Use bulkheads when services share infrastructure but have different priorities, failure modes, or traffic patterns.

## Design tradeoffs

Isolation limits the blast radius but reserves capacity that may sit idle. Too many partitions increase configuration and operational complexity.

[Back to concepts](../../readme.md)
