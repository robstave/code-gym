# Cache invalidation

Cache invalidation removes or updates cached data when the source changes, reducing the chance of serving stale results.

## How it is used

For a cache-aside design, update the source of truth and invalidate the affected key. A later read reloads it; a time-to-live can limit how long stale entries persist if invalidation is missed.

## When applicable

Use it whenever cached data changes and consumers need a defined freshness level.

## Design tradeoffs

Frequent invalidation improves freshness but increases source load. Concurrent reads and writes can repopulate stale values, so ordering, versioning, or other coordination may be needed for stricter guarantees.

[Back to concepts](../../readme.md)
