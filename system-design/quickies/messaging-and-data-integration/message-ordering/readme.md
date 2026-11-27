# Message ordering

Message ordering determines whether consumers see related events in the order they were produced. For example, an account's update should not be applied before its creation.

## How it is used

Route related events to the same partition or ordered queue using a key such as account ID. Process that key serially where needed, and use sequence numbers or versions to detect missing or out-of-order events.

## When applicable

Ordering matters when operations depend on earlier events, such as updating an order's status.

## Design tradeoffs

Ordering is often guaranteed only within a partition, not across the whole system. Serial processing can limit throughput, and retries or replay can disrupt order unless the consumer and broker are configured accordingly.

[Back to concepts](../../readme.md)
