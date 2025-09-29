package main

import (
	"fmt"
	"slices"
)

// 2. Slicing and full slice expression
func slicing() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	a := data[2:5]
	b := data[:4]
	c := data[5:]

	fmt.Println("data:", data)

	// capacity is from start of slice to end of underlying array
	// that's why cap(a)=7 (from index 2 to end)
	fmt.Println("a:", a, "len", len(a), "cap", cap(a))

	fmt.Println("b:", b, "len", len(b), "cap", cap(b))
	fmt.Println("c:", c, "len", len(c), "cap", cap(c))

	a[0] = 99
	c[0] = 88
	fmt.Println("after a[0]=99 data:", data)
}

// 17. Clone
func cloneExample() {
	a := []int{1, 2, 3}
	b := slices.Clone(a)
	a[0] = 99
	fmt.Println("a:", a, "b:", b)
}

func main() {

	fmt.Println("\n== slicing ==")
	slicing()

	fmt.Println("\n== clone ==")
	cloneExample()

}
