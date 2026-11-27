# Sharding keys

A sharding key determines which partition stores a record. For example, a system might place each customer's data based on a customer ID.

## How it is used

Choose a key based on common queries, data distribution, and expected growth. Route reads and writes using the key, and plan how to handle queries that span shards.

## When applicable

Sharding becomes useful when one database cannot meet data size or throughput needs and the operational cost is justified.

## Design tradeoffs

A key that keeps related data together simplifies queries but can create hot shards. A more evenly distributed key may require cross-shard queries, and changing a key later can be costly.

[Back to concepts](../../readme.md)
