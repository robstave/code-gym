# Pessimistic locking

Pessimistic locking acquires a lock before modifying data, making conflicting operations wait or fail until the lock is released.

## How it is used

Within a short database transaction, lock the relevant rows, validate the operation, update them, and commit. Configure lock timeouts and handle deadlocks.

## When applicable

It can help when conflicts are frequent or when an operation must coordinate access to specific records.

## Design tradeoffs

Locks can prevent conflicting updates but reduce concurrency and can cause waits or deadlocks. Transactions should remain short; locks alone do not automatically enforce every business invariant.

[Back to concepts](../../readme.md)
