# Retry with backoff and jitter

Retries repeat an operation after a transient failure. Backoff increases the delay between attempts, while jitter randomizes it so clients do not all retry at once.

## How it is used

Retry only suitable transient errors with a bounded number of attempts and an overall deadline. Use exponential backoff with jitter and make side-effecting operations idempotent.

## When applicable

Use this for temporary network failures, throttling, or briefly unavailable dependencies.

## Design tradeoffs

Retries can improve success rates but increase latency and load during outages. Retrying non-idempotent operations can duplicate effects; retrying at multiple layers can multiply traffic.

[Back to concepts](../../readme.md)
