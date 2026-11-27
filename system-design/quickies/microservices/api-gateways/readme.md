# API gateways

An API gateway is an entry point that routes client requests to backend services and can provide shared cross-cutting features.

## How it is used

The gateway can route requests, authenticate clients, enforce rate limits, and collect request metrics. For example, a mobile app may call one gateway that routes order requests to the order service.

## When applicable

Use a gateway when many clients need a consistent external interface to multiple backend services.

## Design tradeoffs

A gateway simplifies client access and centralizes policies but adds latency and another critical component to operate. Avoid putting too much business logic in it, and ensure backend services are protected if they can be reached directly.

[Back to concepts](../../readme.md)
