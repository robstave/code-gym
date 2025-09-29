package main

import (
	"fmt"
	"slices"
)

// Extensive slice operations exercise playground.
// Read, run, experiment. Implement the TODOs at the bottom.

// 1. Creation examples
func basics() {
	var nilSlice []int
	emptyLiteral := []int{}
	withValues := []int{1, 2, 3}
	made := make([]int, 5)
	madeCap := make([]int, 0, 8)

	fmt.Printf("nilSlice: %v len=%d cap=%d isNil=%v\n", nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("emptyLiteral: %v len=%d cap=%d isNil=%v\n", emptyLiteral, len(emptyLiteral), cap(emptyLiteral), emptyLiteral == nil)
	fmt.Printf("withValues: %v len=%d cap=%d\n", withValues, len(withValues), cap(withValues))
	fmt.Printf("made: %v len=%d cap=%d\n", made, len(made), cap(made))
	fmt.Printf("madeCap: %v len=%d cap=%d\n", madeCap, len(madeCap), cap(madeCap))
}

// 2. Slicing and full slice expression
func slicing() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	a := data[2:5]
	b := data[:4]
	c := data[5:]
	full := data[:]

	// the 3 fields indicate start, end, capacity limit
	limitCap := data[2:5:5]

	fmt.Println("data:", data)

	// capacity is from start of slice to end of underlying array
	// that's why cap(a)=7 (from index 2 to end)
	fmt.Println("a:", a, "len", len(a), "cap", cap(a))

	fmt.Println("b:", b, "len", len(b), "cap", cap(b))
	fmt.Println("c:", c, "len", len(c), "cap", cap(c))
	fmt.Println("full:", full, "len", len(full), "cap", cap(full))
	fmt.Println("limitCap:", limitCap, "len", len(limitCap), "cap", cap(limitCap))

	a[0] = 99
	fmt.Println("after a[0]=99 data:", data)
}

// 3. Capacity growth illustration
func capacityGrowth() {
	s := []int{}
	lastCap := cap(s)
	for i := 0; i < 20; i++ {
		s = append(s, i)
		if cap(s) != lastCap {
			fmt.Printf("append %2d -> len=%2d cap=%2d (cap changed)\n", i, len(s), cap(s))
			lastCap = cap(s)
		}
	}
}

// 4. Removals
func removeAtPreserve[T any](s []T, idx int) []T {
	copy(s[idx:], s[idx+1:])
	return s[:len(s)-1]
}
func removeAtSwap[T any](s []T, idx int) []T {
	s[idx] = s[len(s)-1]
	return s[:len(s)-1]
}

// 5. Insert
func insertAt[T any](s []T, idx int, v T) []T {
	if idx < 0 || idx > len(s) {
		panic("index out of range")
	}
	s = append(s, v)
	copy(s[idx+1:], s[idx:])
	s[idx] = v
	return s
}

// 6. Detach underlying array
func detach() {
	src := []int{1, 2, 3, 4}
	view := src[1:3]
	copied := append([]int(nil), view...)
	view[0] = 99
	fmt.Println("src:", src, "view:", view, "copied:", copied)
}

// 7. Leak pitfall & safe extract
func leakPitfall() []byte { large := make([]byte, 1<<20); header := large[:10]; return header }
func safeExtract() []byte {
	large := make([]byte, 1<<20)
	header := large[:10]
	return append([]byte(nil), header...)
}

// 8. Multi-dimensional irregular slice
func multiDim() {
	grid := make([][]int, 3)
	for r := range grid {
		grid[r] = make([]int, r+1)
		for c := range grid[r] {
			grid[r][c] = r + c
		}
	}
	fmt.Println("grid:", grid)
}

// 9. Filter in place
func filterEven(nums []int) []int {
	out := nums[:0]
	for _, n := range nums {
		if n%2 == 0 {
			out = append(out, n)
		}
	}
	return out
}

// 10. Reverse
func reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 11. Map
func mapSlice[T, R any](in []T, f func(T) R) []R {
	out := make([]R, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}

// 12. Sorting
func sorting() {
	s := []int{9, 3, 6, 1, 2}
	slices.Sort(s)
	fmt.Println("sorted:", s)
	desc := []int{9, 3, 6, 1, 2}
	slices.SortFunc(desc, func(a, b int) int { return b - a })
	fmt.Println("desc:", desc)
}

// 13. Append vs copy merge
func extendExample() {
	a := []int{1, 2}
	b := []int{3, 4}
	a = append(a, b...)
	fmt.Println("merged:", a)
	dst := make([]int, len(a)+len(b))
	n := copy(dst, a)
	copy(dst[n:], b)
	fmt.Println("copy merged:", dst)
}

// 14. Preallocate
func prealloc(n int) []int {
	out := make([]int, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, i*i)
	}
	return out
}

// 15. Struct vs pointer
type Item struct {
	ID   int
	Name string
}

func structVsPointer() {
	structs := []Item{{1, "one"}, {2, "two"}}
	pointers := []*Item{{3, "three"}, {4, "four"}}
	structs[0].Name = "ONE"
	pointers[0].Name = "THREE"
	fmt.Println("structs:", structs)
	fmt.Println("pointers:", pointers)
}

// 16. Capacity trap
func capacityTrap() {
	base := []int{1, 2, 3, 4}
	head := base[:2]
	tail := base[2:]
	head = append(head, 99) // overwrites base[2]
	fmt.Println("base:", base)
	fmt.Println("head:", head)
	fmt.Println("tail:", tail)
}

// 17. Clone
func cloneExample() {
	a := []int{1, 2, 3}
	b := slices.Clone(a)
	a[0] = 99
	fmt.Println("a:", a, "b:", b)
}

// ------------------ TODO EXERCISES ------------------

// TODO 1: rotateLeft rotates elements left by k (in-place preferred).
// Example: [1 2 3 4 5], k=2 -> [3 4 5 1 2]
func rotateLeft[T any](s []T, k int) { panic("implement me") }

// TODO 2: dedupeSorted removes consecutive duplicates from sorted slice in-place.
func dedupeSorted[T comparable](s []T) []T { panic("implement me") }

// TODO 3: chunk splits into size-sized chunks (last may be smaller) using slicing only.
func chunk[T any](s []T, size int) [][]T { panic("implement me") }

// TODO 4: stablePartition stable in-place partition keeping order where keep(x)==true first.
func stablePartition[T any](s []T, keep func(T) bool) []T { panic("implement me") }

// TODO 5: growZero extends slice by n zero-value elements with a single append.
func growZero[T any](s []T, n int) []T { panic("implement me") }

func main() {
	fmt.Println("== basics ==")
	basics()
	fmt.Println("\n== slicing ==")
	slicing()
	fmt.Println("\n== capacity growth ==")
	capacityGrowth()
	fmt.Println("\n== remove / insert ==")
	r := []int{10, 11, 12, 13, 14}
	fmt.Println("orig:", r)
	r = removeAtPreserve(r, 2)
	fmt.Println("remove preserve idx2:", r)
	r = removeAtSwap(r, 1)
	fmt.Println("remove swap idx1:", r)
	r = insertAt(r, 1, 99)
	fmt.Println("insert 99 at idx1:", r)
	fmt.Println("\n== detach ==")
	detach()
	fmt.Println("\n== multi-dim ==")
	multiDim()
	fmt.Println("\n== filter even in-place ==")
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("before:", nums)
	nums = filterEven(nums)
	fmt.Println("after:", nums)
	fmt.Println("\n== reverse & mapSlice ==")
	arr := []string{"a", "b", "c"}
	reverse(arr)
	fmt.Println("reversed:", arr)
	lengths := mapSlice(arr, func(s string) int { return len(s) })
	fmt.Println("mapped lengths:", lengths)
	fmt.Println("\n== sorting ==")
	sorting()
	fmt.Println("\n== extendExample ==")
	extendExample()
	fmt.Println("\n== prealloc ==")
	fmt.Println("prealloc squares:", prealloc(6))
	fmt.Println("\n== struct vs pointer ==")
	structVsPointer()
	fmt.Println("\n== capacity trap ==")
	capacityTrap()
	fmt.Println("\n== clone ==")
	cloneExample()
	// Uncomment to test TODOs after implementing
	// rot := []int{1,2,3,4,5}; rotateLeft(rot,2); fmt.Println("rotated:", rot)
	// d := []int{1,1,2,2,2,5}; d = dedupeSorted(d); fmt.Println("deduped:", d)
	// ch := chunk([]int{1,2,3,4,5,6,7}, 3); fmt.Println("chunks:", ch)
}
