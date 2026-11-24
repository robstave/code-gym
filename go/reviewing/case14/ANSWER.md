# Answer: Retry delay

**Difficulty: Easy**

`time.Duration` represents nanoseconds. Converting the number directly gives
`500ns` rather than the intended `500ms`. The program only prints the duration;
using it for an actual retry would make the wait far too short.

Convert the number and multiply by the unit:

```go
return time.Duration(milliseconds) * time.Millisecond
```
