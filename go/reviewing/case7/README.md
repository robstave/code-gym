# Case 7: Request timeout

**Difficulty: Medium**

The helper should return a context that a caller can use for up to one second
to perform work, unless the parent is canceled first. The caller should release
the timer when finished. Review the context's lifetime.

Run: `go run .\case7\main.go` from the reviewing directory.
