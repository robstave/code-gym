# Strong vs. eventual consistency

Strong consistency commonly means a read observes the latest completed write under the system's stated consistency model. Eventual consistency means replicas converge if updates stop, but reads may temporarily differ.

## How it is used

Choose read and write settings based on the operation. For example, require a current balance for a transfer while allowing a product catalog display to briefly lag.

## When applicable

This choice matters when data is replicated or accessed across distributed services and regions.

## Design tradeoffs

Stronger guarantees can require coordination, increasing latency or reducing availability during network partitions. Eventual consistency can improve responsiveness and availability but requires handling stale data and conflicts.

[Back to concepts](../../readme.md)
