# Service boundaries

Microservices split an application into independently deployable services. A service boundary defines which business capabilities and data a service owns.

## How it is used

Group closely related behavior around a business capability, such as orders or inventory. Expose a clear API or events and keep each service responsible for its own data and rules.

## When applicable

Consider this approach when a system has distinct domains or teams that benefit from independent deployment and scaling.

## Design tradeoffs

Independent services allow focused ownership but add network calls, operational work, and distributed data consistency challenges. Boundaries that are too fine create excessive coordination; a modular monolith may be simpler for smaller systems.

[Back to concepts](../../readme.md)
