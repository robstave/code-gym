# Delivery semantics

Delivery semantics describe what a messaging system promises about how often a message is delivered or processed after failures.

## How it is used

- **At-most-once:** A message is delivered no more than once, but may be lost. For example, acknowledging before processing avoids redelivery if processing fails.
- **At-least-once:** Messages are retried until acknowledged, so processing can happen more than once. Use idempotent handlers or deduplication to make retries safe.
- **Exactly-once processing:** Within a specified system and transaction boundary, the outcome is committed once despite retries. Some brokers provide transactional features for this; external side effects need their own coordination or idempotency.
- **Effectively-once effects:** A common application goal: accept duplicate deliveries but ensure the business effect happens once, such as by atomically recording a message ID with the update.

## When applicable

Choose based on whether missed messages, duplicate effects, or additional coordination are acceptable. Metrics may tolerate loss; payments generally require stronger safeguards.

## Design tradeoffs

At-most-once is simpler but can lose work. At-least-once improves delivery resilience but requires duplicate handling. Exactly-once guarantees have a defined scope and may add coordination and cost; they do not automatically cover external APIs. See [at-least-once delivery](../at-least-once-delivery/readme.md) for more detail.

[Back to concepts](../../readme.md)
