# Sagas

A saga coordinates a business operation across services as a sequence of local transactions. If a later step fails, compensating actions may undo or mitigate earlier steps.

## How it is used

For an order, reserve inventory, take payment, and confirm the order. If payment fails, release the reservation. An orchestrator can direct the steps, or services can react to events in a choreographed flow.

## When applicable

Use a saga when a workflow spans service-owned databases and cannot reasonably use a single database transaction.

## Design tradeoffs

Sagas avoid a distributed database transaction but allow intermediate states to be visible. Compensation may fail or be impossible to fully reverse, so steps need idempotency, durable progress tracking, retries, and operational recovery.

[Back to concepts](../../readme.md)
