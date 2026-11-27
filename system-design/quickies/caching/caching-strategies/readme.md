# Caching strategies

A cache stores copies of data so future requests can avoid a slower source, such as a database. Different strategies decide who loads the cache, how writes reach the source, and when cached values are refreshed.

## How they are used

### Cache-aside (lazy loading)

The application checks the cache first. On a miss, it reads the source, stores the result in the cache, and returns it. On a write, it typically updates the source and invalidates the cached entry.

- **When applicable:** General-purpose workloads where only some data is frequently read.
- **Tradeoffs:** Simple and loads only requested data, but the application manages cache logic. Misses add latency, and concurrent reads and writes can leave stale entries.

### Read-through

The application reads through a cache service or library. On a miss, the cache loads the value from the source and returns it.

- **When applicable:** Repeated reads where a cache abstraction can consistently load data from the source.
- **Tradeoffs:** Simplifies application read logic, but requires a cache integration that supports loading. The first read is still slow, and freshness needs a separate policy.

### Write-through

The application writes through the cache, which synchronously writes to the source as part of the operation and updates the cached value.

- **When applicable:** Read-heavy data where writes should leave the cache populated with the latest successful update.
- **Tradeoffs:** Writes have extra latency and may populate values that are never read. Failure handling and concurrent writes still require care; write-through alone does not guarantee consistency for every reader.

### Write-behind (write-back)

The application writes to the cache first. The cache later writes changes to the source asynchronously, often batching them.

- **When applicable:** Write-heavy workloads that can tolerate delayed persistence and possible data loss.
- **Tradeoffs:** Can reduce write latency and batch source updates, but a failure before persistence may lose data. Ordering, retries, durability, and shutdown behavior need careful design.

### Write-around

Writes go directly to the source without populating the cache. Reads can then load data using cache-aside or read-through.

- **When applicable:** Data that is written often but read infrequently, where caching every write would waste space.
- **Tradeoffs:** Avoids filling the cache with unused data, but the next read may be a miss. Existing cached entries must still be invalidated or updated to avoid stale reads.

### Refresh-ahead

The cache refreshes popular entries before they expire, often based on their access patterns.

- **When applicable:** Predictable, frequently accessed data where a miss would be expensive.
- **Tradeoffs:** Reduces expiry-related latency, but refreshes may waste work on data that is no longer needed. Refresh failures and concurrent refreshes still need handling.

## Choosing a strategy

Strategies can be combined: for example, use cache-aside reads with write-around writes and a time-to-live (TTL). A TTL bounds how long an entry stays cached, but it does not by itself ensure that reads are current or prevent simultaneous reloads.

Consider freshness, acceptable data loss, read and write patterns, and the behavior during cache or source failures. See [cache invalidation](../cache-invalidation/readme.md) and [cache stampede protection](../cache-stampede-protection/readme.md) for related concerns.

[Back to concepts](../../readme.md)
