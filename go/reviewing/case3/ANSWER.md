# Answer: Reading a quantity

**Difficulty: Easy**

The function discards the error from `strconv.Atoi` and returns success even
when parsing fails. The sample prints `Quantity: 0`; invalid input should
produce an error instead. Ignoring the error can also hide out-of-range input.

Return the parsing error, optionally wrapping it with context:

```go
quantity, err := strconv.Atoi(input)
if err != nil {
    return 0, fmt.Errorf("invalid quantity: %w", err)
}
return quantity, nil
```
