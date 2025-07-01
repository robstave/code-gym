// util.go
package util

// UniqueInts returns a slice with duplicate integers removed,
// preserving the original order of first appearances.
func UniqueInts1(nums []int) []int {

	seen := make(map[int]bool)
	results := make([]int, 0, len(nums))

	for _, val := range nums {

		if !seen[val] {
			seen[val] = true

			results = append(results, val)
		}
	}
	return results
}
