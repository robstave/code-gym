# Reading query plans

A query plan describes how a database intends to execute a query, including scans, index access, joins, and estimated costs.

## How it is used

Use the database's explain tool to inspect a slow query. Compare estimated and actual rows when available, and look for expensive scans, joins, or sorts before changing indexes or queries.

## When applicable

Use query plans when diagnosing slow queries or evaluating whether an index helps a particular workload.

## Design tradeoffs

Plans depend on data, statistics, parameters, and database version. An index that speeds one query adds write and storage costs; tools that execute the query for actual measurements may be expensive or have side effects.

[Back to concepts](../../readme.md)
