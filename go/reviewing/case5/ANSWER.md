# Answer: Counter

**Difficulty: Easy**

`increment` has a value receiver, so it changes a copy. The sample prints `0`
instead of `1`.

Use a pointer receiver: `func (c *counter) increment()`. Calling
`c.increment()` on the addressable variable in `main` then updates `c`.
