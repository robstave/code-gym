// Go's `slices` package implements sorting for builtins
// and user-defined types. We'll look at sorting for
// builtins first.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Sorting functions are generic, and work for any
	// _ordered_ built-in type. For a list of ordered
	// types, see [cmp.Ordered](https://pkg.go.dev/cmp#Ordered).
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// We can also use the `slices` package to check if
	// a slice is already in sorted order.
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)

	// Demonstration: sorting a subslice.
	// A slice header (len/cap + pointer) refers to an underlying array. When you
	// take a subslice (strs2a := strs2[1:5]) both slices share that same backing array.
	// Sorting the subslice mutates the elements in the underlying array for the
	// indices covered by the subslice range [1:5). It does NOT globally sort the
	// entire original slice; elements outside that window keep their relative order.
	fmt.Println("-------- ")

	strs2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	strs2a := strs2[1:5]

	fmt.Println(strs2)
	fmt.Println(strs2a)

	strs2a[0] = 99

	fmt.Println(strs2)
	fmt.Println("-----sort subslice--- ")
	slices.Sort(strs2a)
	fmt.Println(strs2a)

	fmt.Println("-----full slice--- ")
	fmt.Println(strs2)

}
