# Graceful degradation

Graceful degradation keeps essential functionality available when a dependency or feature fails, even if the experience is reduced.

## How it is used

Provide a deliberate fallback, such as showing cached product details when recommendations are unavailable. Make the degraded state visible where users need to understand it.

## When applicable

Use it when nonessential components can fail independently of the primary user workflow.

## Design tradeoffs

Fallbacks improve availability but may show stale or incomplete data. They add paths to test and monitor; never use a fallback that falsely reports an important operation as successful.

[Back to concepts](../../readme.md)
