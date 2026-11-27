# Consistent hashing

Consistent hashing maps keys and nodes onto a hash space so adding or removing a node usually moves only a subset of keys, rather than remapping most of them.

## How it is used

Hash each key and assign it to a node on a ring or use another consistent-hashing scheme. Virtual nodes can spread load more evenly across physical nodes.

## When applicable

Use it for distributed caches or storage where membership changes and minimizing key movement matters.

## Design tradeoffs

The scheme reduces remapping but does not automatically solve hot keys or guarantee even load. Virtual nodes and membership changes add coordination, rebalancing, and operational complexity.

[Back to concepts](../../readme.md)
