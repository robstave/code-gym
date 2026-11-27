# Optimistic locking

Optimistic locking assumes concurrent updates are uncommon. It checks whether data changed before saving and rejects conflicting updates instead of holding a lock during the read.

## How it is used

Read a record and its version, then update it only if the stored version still matches. If no row is updated, report or resolve the conflict, or retry when safe.

## When applicable

Use it for low-contention edits, such as a user updating a profile or configuration.

## Design tradeoffs

It avoids long-held locks but conflicts require application handling and may cause repeated work. High contention can make retries costly, and all relevant writers must honor the version check.

[Back to concepts](../../readme.md)
