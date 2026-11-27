# Load shedding

Load shedding deliberately rejects or drops some work when a service is overloaded, preserving capacity for work it can still handle.

## How it is used

Set bounded queues and concurrency limits, then reject excess requests quickly or prioritize essential work. Return an explicit overload response, such as HTTP 503 where appropriate, and monitor rejected requests.

## When applicable

Use it when traffic spikes or a slow dependency could otherwise exhaust resources and make all requests fail.

## Design tradeoffs

Shedding protects the service but some requests fail immediately. Priority policies can starve less important work, and uncoordinated retries can worsen overload. Unlike [rate limiting](../rate-limiting/readme.md), shedding responds to service capacity rather than only a client's quota.

[Back to concepts](../../readme.md)
