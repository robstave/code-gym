# Database migrations

Database migrations are versioned changes to a database schema or data, such as adding a column or creating an index.

## How it is used

Apply migrations in a controlled order using a migration tool. For a live schema change, use an expand-and-contract sequence: add a compatible structure, migrate application traffic and data, then remove the old structure after verification.

## When applicable

Use migrations whenever a shared database schema evolves, particularly when old and new application versions run simultaneously.

## Design tradeoffs

Compatibility steps make deployments safer but temporarily increase complexity and storage. Large backfills and indexes can affect production traffic; test them with representative data, monitor progress, and plan recovery before destructive changes.

[Back to concepts](../../readme.md)
