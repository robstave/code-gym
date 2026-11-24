# Answer: Missing account

**Difficulty: Medium**

`fmt.Errorf` with `%w` wraps the original error in a new error value. Direct
comparison with `==` does not follow the wrapping chain. The sample prints
`Lookup failed`; the intended branch prints `Account not found`.

Use `errors.Is(err, errNotFound)` to check for the sentinel through wrapping.
Keep `%w` in `lookupAccount` so the original error remains discoverable.
