# Answer: Remove a prefix

**Difficulty: Easy**

`strings.TrimLeft` removes any leading characters in the provided cutset,
not the exact prefix. For `log-golf`, it removes `log-gol` and returns `f`;
the intended result is `golf`.

Use `strings.TrimPrefix(label, "log-")`. It removes one exact prefix if
present and otherwise returns the original string.
