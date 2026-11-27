# Correlation IDs and tracing

A correlation ID links logs related to a request or workflow. Distributed tracing adds spans that show the timing and relationships between operations across services.

## How it is used

Propagate context across service boundaries, include IDs in structured logs, and instrument key operations as spans. Use the trace to follow a slow request through its dependencies.

## When applicable

These tools are useful for debugging failures and latency in systems where a request crosses multiple components.

## Design tradeoffs

Instrumentation, storage, and context propagation add overhead. Sampling reduces cost but can miss individual traces; avoid recording secrets or unnecessary personal data in logs and spans.

[Back to concepts](../../readme.md)
