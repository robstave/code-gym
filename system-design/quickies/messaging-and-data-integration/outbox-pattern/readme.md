# Outbox pattern

The outbox pattern stores a business change and an event describing it in the same database transaction. A separate process publishes pending events.

## How it is used

When placing an order, write both the order and an outbox row atomically. A relay reads the outbox and sends events to the message broker, then records progress.

## When applicable

Use it when database updates and messages must not get out of sync because one succeeds and the other fails.

## Design tradeoffs

The relay adds infrastructure and delivery delay. It may publish duplicates after failures, so consumers should be idempotent; outbox cleanup, ordering, and monitoring also need attention.

[Back to concepts](../../readme.md)
