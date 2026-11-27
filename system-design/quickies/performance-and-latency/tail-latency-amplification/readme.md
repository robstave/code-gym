# Tail latency amplification

When a request waits for several dependencies, a single slow dependency can delay the whole response. As the number of parallel calls increases, the chance that at least one is slow increases too.

## How it is used

Measure individual dependency latency and the end-to-end request. Reduce unnecessary fan-out, set deadlines, and consider returning partial results or hedging safe requests when appropriate.

## When applicable

This matters for fan-out APIs, search aggregation, and other requests that wait for multiple services or shards.

## Design tradeoffs

Hedged requests can reduce latency but add load, especially during an outage. Partial responses improve responsiveness but may omit useful data; deadlines must account for the overall request budget.

[Back to concepts](../../readme.md)
