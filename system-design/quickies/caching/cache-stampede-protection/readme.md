# Cache stampede protection

A cache stampede occurs when many requests miss the same key at once and all try to rebuild it, overloading the source.

## How it is used

Coalesce concurrent loads for the same key, use a short-lived distributed lock when needed across instances, or serve stale data while one worker refreshes it. Add jitter to expiration times.

## When applicable

Use protection for popular keys with expensive rebuilds, especially around cache expiration or a cache outage.

## Design tradeoffs

Locks and request coalescing add coordination and failure-handling complexity. Serving stale data improves availability but reduces freshness; expiration jitter alone does not prevent every stampede.

[Back to concepts](../../readme.md)
