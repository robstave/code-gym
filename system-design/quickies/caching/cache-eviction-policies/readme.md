# Cache eviction policies

An eviction policy chooses which cached entries to remove when the cache needs space. Eviction is different from expiration: a time-to-live removes or invalidates an entry because of age.

## How it is used

- **Least recently used (LRU):** Evict the entry that has not been used for the longest time; useful when recent access predicts future access.
- **Least frequently used (LFU):** Evict less frequently accessed entries; useful for consistently popular items, often with aging so old popularity does not dominate forever.
- **First in, first out (FIFO):** Evict the oldest inserted entry; simple, but ignores how often it is used.
- **Random:** Evict a randomly selected entry; inexpensive but may remove a popular item.

## When applicable

Choose a policy when the cache has a bounded memory budget and cannot hold all possible entries.

## Design tradeoffs

LRU and LFU track access information, adding overhead; FIFO and random are simpler but may have lower hit rates for some workloads. The best choice depends on access patterns and entry sizes. A TTL and an eviction policy can be used together.

[Back to concepts](../../readme.md)
