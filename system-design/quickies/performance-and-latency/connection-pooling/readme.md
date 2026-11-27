# Connection pooling

A connection pool keeps reusable connections to a database or service instead of establishing a new connection for every request.

## How it is used

Configure a shared client with limits on open and idle connections. Requests borrow a connection and return it when finished; monitor pool waits and connection usage.

## When applicable

Pooling helps when connection setup is expensive and many requests repeatedly contact the same dependency.

## Design tradeoffs

Pools reduce setup overhead but consume resources even when idle. Too small a pool causes waiting; too large a pool can overwhelm the dependency, especially across many application instances.

[Back to concepts](../../readme.md)
