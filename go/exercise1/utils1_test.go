// util_test.go
package util

import "testing"

func TestUniqueInts(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 2, 1}, []int{1, 2, 3}},
		{[]int{}, []int{}},
		{[]int{5, 5, 5}, []int{5}},
		{[]int{3, 1, 2, 3, 2}, []int{3, 1, 2}},
	}
	for _, tc := range tests {
		got := UniqueInts1(tc.input)
		if !equalSlices(got, tc.want) {
			t.Errorf("UniqueInts(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

// helper to compare two int slices for equality
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
