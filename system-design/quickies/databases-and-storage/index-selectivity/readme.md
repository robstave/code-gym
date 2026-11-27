# Index selectivity

Selectivity describes how effectively an indexed value narrows down rows. A lookup matching a small fraction of a table is generally more selective than one matching most rows.

## How it is used

Consider actual value distributions and query predicates when selecting indexes. Check query plans and statistics to see whether the database uses the index.

## When applicable

This matters when designing indexes for filters, joins, and composite keys on large tables.

## Design tradeoffs

High selectivity often makes an index useful, but a scan can be faster when most rows match. Composite index order and workload matter, and each extra index costs storage and slows writes.

[Back to concepts](../../readme.md)
