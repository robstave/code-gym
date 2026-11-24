# Answer: Applying a raise

**Difficulty: Easy**

The range value `employee` is a copy of each struct. Updating the copy does not
update the slice. The sample prints `100 200`, not the intended `110 210`.

Update the elements by index instead:

```go
for i := range employees {
    employees[i].salary += amount
}
```
