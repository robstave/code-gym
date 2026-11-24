# Answer: Applying a raise

**Difficulty: Easy**

The range value `employee` is a copy of each struct. Updating the copy does not
update the slice. The sample prints `100 200`, not the intended `110 210`.

The distinction is between the slice and an element of the slice. Passing
`employees` to the function copies the slice descriptor, but both the caller's
slice and the function's copy still refer to the same underlying array.
Accessing `employees[i]` refers to an actual struct in that array, so changing
its salary is visible to the caller.

By contrast, in `for _, employee := range employees`, each `employee` is a
copy of an array element. Changing that copy does not change the original
struct in the array.

Update the elements by index instead:

```go
for i := range employees {
    employees[i].salary += amount
}
```
