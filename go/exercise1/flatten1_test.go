// utils-flatten_test.go
// This file should contain the test cases for the flatten functions.
// utils-flatten_test.go
package util

import "testing"

func TestFlattenInts(t *testing.T) {
	tests := []struct {
		input [][]int
		want  []int
	}{
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{}, []int{}},
		{[][]int{{}, {5, 6}, {}}, []int{5, 6}},
		{[][]int{{7}, {8, 9}, {10, 11, 12}}, []int{7, 8, 9, 10, 11, 12}},
	}
	for _, tc := range tests {
		got := FlattenInts(tc.input)
		if !equalSlices(got, tc.want) {
			t.Errorf("FlattenInts(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

func TestFlattenInts2(t *testing.T) {
	tests := []struct {
		input [][]int
		want  []int
	}{
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{}, []int{}},
		{[][]int{{}, {5, 6}, {}}, []int{5, 6}},
		{[][]int{{7}, {8, 9}, {10, 11, 12}}, []int{7, 8, 9, 10, 11, 12}},
	}
	for _, tc := range tests {
		got := FlattenInts2(tc.input)
		if !equalSlices(got, tc.want) {
			t.Errorf("FlattenInts(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}
