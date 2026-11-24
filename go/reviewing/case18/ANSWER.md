# Answer: Importing lines

**Difficulty: Medium**

`Scanner.Scan` returns false both at normal end-of-input and on an error.
The function ignores `scanner.Err()`, so the sample reports one line and no
error even though the reader failed. Scanner errors also include tokens that
exceed its configured maximum size.

Return the scanner error after the loop:

```go
return count, scanner.Err()
```

The count remains available to the caller even when the error is non-nil.
