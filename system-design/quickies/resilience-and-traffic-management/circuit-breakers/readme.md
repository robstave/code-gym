# Circuit breakers

A circuit breaker temporarily stops calls to a failing dependency so failures do not repeatedly consume resources.

## How it is used

Track failures and open the circuit after a configured threshold. After a cooldown, allow limited trial calls in a half-open state; close the circuit if the dependency recovers.

## When applicable

Use it when a remote dependency fails or stalls often enough to threaten the calling service.

## Design tradeoffs

Fast failures protect resources but can reject requests that might have succeeded. Poor thresholds can cause unnecessary outages or delay recovery; fallbacks and monitoring remain important.

[Back to concepts](../../readme.md)
