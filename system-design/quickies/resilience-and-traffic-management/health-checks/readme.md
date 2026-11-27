# Health checks

Health checks report whether an instance is running and whether it is ready to receive traffic.

## How it is used

Expose separate liveness and readiness checks. A scheduler can restart an instance that fails liveness, while a load balancer can stop sending traffic to an instance that fails readiness.

## When applicable

Use them for services managed by orchestrators or load balancers, especially during startup, deployment, and recovery.

## Design tradeoffs

Checks that are too shallow miss problems; expensive or overly strict dependency checks can remove healthy instances during a shared outage. Liveness should generally not fail solely because a remote dependency is down, since restarts may worsen the outage.

[Back to concepts](../../readme.md)
