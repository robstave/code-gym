# Answer: JSON response

**Difficulty: Easy**

The struct fields are unexported because their names begin with lowercase
letters. `encoding/json` ignores them even though they have JSON tags. The
sample prints `{}` instead of `{"name":"Ada","age":37}`.

Export the fields while keeping the tags for the desired JSON names:

```go
type person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

Update the struct literal to use `Name` and `Age` as well.
