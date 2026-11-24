# Answer: Request timeout

**Difficulty: Medium**

`defer cancel()` runs when `requestContext` returns, immediately canceling the
returned context. The caller sees `context canceled` instead of an active
context; the intended one-second timeout is never usable.

Return both the context and its cancel function, and have the caller defer
the cancel function after creating the context:

```go
func requestContext(parent context.Context) (context.Context, context.CancelFunc) {
    return context.WithTimeout(parent, time.Second)
}

ctx, cancel := requestContext(context.Background())
defer cancel()
```
