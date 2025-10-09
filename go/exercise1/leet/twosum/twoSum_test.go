package utils

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"case1", []int{3, 2, 4}, 6, []int{1, 2}},
		{"case2", []int{2, 7, 11, 15}, 18, []int{1, 2}},
		{"case3", []int{2, 4, 11, 16}, 27, []int{2, 3}},
		{"case4", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
