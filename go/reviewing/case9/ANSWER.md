# Answer: Event timestamp

**Difficulty: Medium**

`time.Unix` takes seconds as its first argument, not milliseconds. The sample
prints a date in the year 55840; `1700000000000` milliseconds corresponds to
`2023-11-14T22:13:20Z`.

Use `time.UnixMilli(milliseconds).UTC()` (Go 1.17 or newer), or divide into
seconds and the remaining nanoseconds. Do not divide by 1000 alone if
millisecond precision must be retained.
