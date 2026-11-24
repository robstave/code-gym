# Answer: Saved snapshot

**Difficulty: Medium**

Assigning a slice copies its descriptor, not its backing array. `saved` and
`scores` therefore refer to the same elements, so the sample prints `99`
instead of the original score `10`.

Create independent storage, for example:

```go
saved := append([]int(nil), scores...)
```

Alternatively, allocate a slice of the same length and use `copy`.
