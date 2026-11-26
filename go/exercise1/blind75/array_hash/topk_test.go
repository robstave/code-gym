package utils

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{name: "first example", nums: []int{1, 2, 2, 3, 3, 3}, k: 2, want: []int{2, 3}},
		{name: "second example", nums: []int{7, 7}, k: 1, want: []int{7}},
		{name: "single element", nums: []int{5}, k: 1, want: []int{5}},
		{name: "negative values and zero", nums: []int{-1, -1, -1, 0, 0, 2}, k: 2, want: []int{-1, 0}},
		{name: "all distinct values selected", nums: []int{1, 2, 3}, k: 3, want: []int{1, 2, 3}},
		{name: "ties within selected elements", nums: []int{1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
		{name: "highest frequency appears last", nums: []int{1, 2, 3, 3, 3}, k: 1, want: []int{3}},
		{name: "constraint endpoints", nums: []int{-1000, 1000, 1000, 0, 0, 0}, k: 2, want: []int{0, 1000}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			slices.Sort(got)
			slices.Sort(tt.want)
			if !slices.Equal(got, tt.want) {
				t.Errorf("topKFrequent(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
