# SLOs and error budgets

A service level objective (SLO) is a target for a measured service level indicator (SLI), such as the percentage of successful requests over 30 days. The error budget is the amount of unreliability allowed by that target.

## How it is used

For a 99.9% availability SLO, at most 0.1% of eligible requests may fail within the measurement window. Monitor budget consumption and use it to balance feature delivery with reliability work.

## When applicable

Use SLOs for important user-facing services where teams need a shared, measurable reliability target.

## Design tradeoffs

Strict objectives can require substantial engineering investment and slow releases. Poorly chosen indicators or excluded traffic can hide user problems; objectives should reflect real user needs and be reviewed regularly.

[Back to concepts](../../readme.md)
