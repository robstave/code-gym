# p95 and p99 latency

Percentiles describe how long requests take across a population. A p95 of 200 ms means approximately 95% of observed requests completed in 200 ms or less; p99 describes the slowest 1%.

## How it is used

Record request durations in a histogram and monitor percentiles by endpoint and time window. Compare them with latency objectives to find slow experiences that an average can hide.

## When applicable

Use percentiles for user-facing services, dependency calls, and workloads where occasional slow requests matter.

## Design tradeoffs

High percentiles need enough samples to be meaningful and depend on the measurement window. Histogram buckets trade precision for storage; do not average per-instance percentiles to obtain a global percentile.

[Back to concepts](../../readme.md)
