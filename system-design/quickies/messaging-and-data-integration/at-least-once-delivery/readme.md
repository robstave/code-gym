# At-least-once delivery

At-least-once delivery means a message can be delivered more than once, typically because the sender or broker retries when an acknowledgment is missing.

## How it is used

Consumers acknowledge messages after successful processing and make side effects idempotent, for example by recording processed message IDs atomically with the business update.

## When applicable

This model is common when losing messages is less acceptable than occasionally processing duplicates.

## Design tradeoffs

Retries improve delivery resilience but increase work and require deduplication or idempotent handlers. Delivery guarantees depend on broker configuration, retention, and acknowledgment behavior.

[Back to concepts](../../readme.md)
