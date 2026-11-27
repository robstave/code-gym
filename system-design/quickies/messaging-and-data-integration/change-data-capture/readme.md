# Change Data Capture

Change Data Capture (CDC) records database changes and makes them available to other systems, often by reading the database's transaction log.

## How it is used

Publish inserts, updates, and deletes to a stream so consumers can update search indexes, analytics stores, or other services.

## When applicable

Use CDC to propagate database changes without having every application explicitly publish each change.

## Design tradeoffs

CDC introduces operational complexity, replication lag, and possible duplicate events. Consumers must handle ordering, schema changes, and recovery; access to logs and data also needs appropriate security.

[Back to concepts](../../readme.md)
