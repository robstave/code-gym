
# Slices in Go


## Notes

### what is the difference between arrays and slices in Go?

- Arrays have a fixed size, while slices are dynamically sized.
- Arrays are value types, meaning that assigning one array to another copies all elements. Slices are reference types, meaning that assigning one slice to another shares the underlying array.
- Slices provide more flexibility and are generally preferred over arrays in Go.


### how to create a slice in Go?

You can create a slice in several ways:

1. Using a slice literal:

```go
s := []int{1, 2, 3}
```

2. Using the `make` function:

```go
s := make([]int, 3) // creates a slice of length 3
```

3. Slicing an existing array or slice:

```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4] // creates a slice containing elements 2, 3, 4
```
### how to append elements to a slice

You can use the `append` function to add elements to a slice:

```go
s := []int{1, 2, 3}
s = append(s, 4)       // adds a single element
s = append(s, 5, 6, 7) // adds multiple elements
```

### how to get the length and capacity of a slice

You can use the `len` and `cap` functions to get the length and capacity of a slice:

```go
s := []int{1, 2, 3}
fmt.Println(len(s)) // prints 3
fmt.Println(cap(s)) // prints 3
```
### how to copy a slice

You can use the `copy` function to copy elements from one slice to another:

```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
fmt.Println(dst) // prints [1 2 3]
```

### how to iterate over a slice

You can use a `for` loop to iterate over the elements of a slice:

```go
s := []int{1, 2, 3}
for i, v := range s {
    fmt.Println(i, v)
}
```
### how to delete an element from a slice

You can delete an element from a slice by using slicing and the `append` function:

```go
s := []int{1, 2, 3, 4}
i := 2 // index of the element to delete
s = append(s[:i], s[i+1:]...)
fmt.Println(s) // prints [1 2 4]
``` 
### how to check if a slice is empty

You can check if a slice is empty by comparing its length to 0:

```go
s := []int{}
if len(s) == 0 {
    fmt.Println("slice is empty")
}
``` 

### how to clear a slice

You can clear a slice by assigning it to `nil` or an empty slice:

```go
s := []int{1, 2, 3}
s = nil // clears the slice
// or
s = []int{} // clears the slice
```

### how to check if a slice contains a specific element

You can check if a slice contains a specific element by iterating over the slice:

```go
s := []int{1, 2, 3}
target := 2
found := false
for _, v := range s {
    if v == target {
        found = true
        break
    }
}
fmt.Println(found) // prints true
```

### how to find the index of an element in a slice

You can find the index of an element in a slice by iterating over the slice:

```go
s := []int{1, 2, 3}
target := 2
index := -1
for i, v := range s {
    if v == target {
        index = i
        break
    }
}
fmt.Println(index) // prints 1
```

### how to reverse a slice

You can reverse a slice by swapping elements from both ends towards the center:

```go
s := []int{1, 2, 3, 4}
for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
}
fmt.Println(s) // prints [4 3 2 1]
```

### how to sort a slice

You can sort a slice using the `sort` package:

```go
import "sort"

s := []int{3, 1, 4, 2}
sort.Ints(s)
fmt.Println(s) // prints [1 2 3 4]
```

