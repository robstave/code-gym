# Answer: Launching jobs

**Difficulty: Medium**

The goroutines close over the same `job` variable, which is declared outside
the loop. The `start` channel holds them until after the loop completes, at
which point `job` is 3. Consequently, the sample prints `[3 3 3]` instead of
`[0 1 2]`.

Pass the current value as an argument when starting the goroutine:

```go
go func(id int) {
    <-start
    results <- id
}(job)
```

Go 1.22 changed variables declared by a loop to be per-iteration, but this
example declares `job` outside the loop, so that change does not fix it.
