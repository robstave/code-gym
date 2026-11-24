# Answer: Order total

**Difficulty: Easy**

The loop stops at `len(prices)-1`, so it never adds the last item. The sample
prints `Total: 7`; the intended total is `10`. A one-item order incorrectly
totals zero. An empty order happens to return zero.

Fix the loop condition to `i < len(prices)`, or range over the values:

```go
for _, price := range prices {
    total += price
}
```