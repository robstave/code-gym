# Answer: Removing canceled orders

**Difficulty: Medium**

After an element is removed, the next element shifts into its index. The loop
increments `i` anyway, so it skips the shifted element. The sample prints
`[active canceled active]` rather than `[active active]`.

Increment only when keeping an element, or build a new slice by appending each
non-canceled order. If filtering in place, document the mutation of the input's
backing array; use a new slice if callers require the original to remain intact.
