# Idempotency keys

An idempotency key identifies a logical request so repeating it does not repeat its side effect, such as charging a customer twice.

## How it is used

The client sends a unique key with the request. The server stores the key and outcome, and returns the prior outcome for a matching retry; it rejects reuse with a different payload.

## When applicable

Use keys for operations that clients may retry after a timeout, especially payments, order creation, and other side-effecting requests.

## Design tradeoffs

Keys require durable storage, an appropriate retention period, and concurrency-safe handling of simultaneous requests. After a key expires, a late retry may be treated as new; scope keys to the relevant client and operation.

[Back to concepts](../../readme.md)
