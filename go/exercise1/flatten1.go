// utils-flatten.go
package util

// FlattenInts takes a slice of int-slices and returns a single
// flat slice with all elements in order.
func FlattenInts(matrix [][]int) []int {

	results := make([]int, 0, 100)

	for _, row := range matrix {
		for _, val := range row {
			results = append(results, val)
		}

	}
	return results
}

// this is a little nicer
func FlattenInts2(matrix [][]int) []int {

	results := make([]int, 0, 100)

	for _, row := range matrix {
		results = append(results, row...)

	}
	return results
}
