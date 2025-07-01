// utils2_test.go
package util

import "testing"

func TestSumInts(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{}, 0},
		{[]int{-1, -2, -3}, -6},
		{[]int{5, -5, 10}, 10},
	}
	for _, tc := range tests {
		got := SumInts(tc.input)
		if got != tc.want {
			t.Errorf("SumInts(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		slice   []int
		element int
		want    bool
	}{
		{[]int{1, 2, 3, 4}, 3, true},
		{[]int{1, 2, 3, 4}, 5, false},
		{[]int{}, 1, false},
		{[]int{-1, -2, -3}, -2, true},
	}
	for _, tc := range tests {
		got := Contains(tc.slice, tc.element)
		if got != tc.want {
			t.Errorf("Contains(%v, %v) = %v; want %v", tc.slice, tc.element, got, tc.want)
		}
	}
}

func TestFilterEvenInts(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 4}},
		{[]int{}, []int{}},
		{[]int{1, 3, 5}, []int{}},
		{[]int{2, 4, 6}, []int{2, 4, 6}},
	}
	for _, tc := range tests {
		got := FilterEvenInts(tc.input)
		if !equalSlices(got, tc.want) {
			t.Errorf("FilterEvenInts(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

func TestReverseInts(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{}, []int{}},
		{[]int{5}, []int{5}},
	}
	for _, tc := range tests {
		got := ReverseInts(tc.input)
		if !equalSlices(got, tc.want) {
			t.Errorf("ReverseInts(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

func TestFindMax(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{-1, -2, -3}, -1},
		{[]int{5}, 5},
	}
	for _, tc := range tests {
		got := FindMax(tc.input)
		if got != tc.want {
			t.Errorf("FindMax(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

func TestCountOccurrences(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 4, 3}, 3, 2},
		{[]int{1, 2, 3, 4}, 5, 0},
		{[]int{}, 1, 0},
	}
	for _, tc := range tests {
		got := CountOccurrences(tc.input, tc.target)
		if got != tc.want {
			t.Errorf("CountOccurrences(%v, %v) = %v; want %v", tc.input, tc.target, got, tc.want)
		}
	}
}

func TestMergeSlices(t *testing.T) {
	tests := []struct {
		slice1 []int
		slice2 []int
		want   []int
	}{
		{[]int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
		{[]int{}, []int{3, 4}, []int{3, 4}},
		{[]int{1, 2}, []int{}, []int{1, 2}},
	}
	for _, tc := range tests {
		got := MergeSlices(tc.slice1, tc.slice2)
		if !equalSlices(got, tc.want) {
			t.Errorf("MergeSlices(%v, %v) = %v; want %v", tc.slice1, tc.slice2, got, tc.want)
		}
	}
}
