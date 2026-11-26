package utils

import "testing"

func TestHasDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{name: "empty slice", nums: []int{}, want: false},
		{name: "single value", nums: []int{1}, want: false},
		{name: "all values unique", nums: []int{1, 2, 3, 4}, want: false},
		{name: "duplicate adjacent values", nums: []int{1, 2, 2, 3}, want: true},
		{name: "duplicate non-adjacent values", nums: []int{1, 2, 3, 1}, want: true},
		{name: "duplicate negative values", nums: []int{-1, 0, -1}, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasDuplicate(tt.nums); got != tt.want {
				t.Errorf("hasDuplicate(%v) = %t, want %t", tt.nums, got, tt.want)
			}
			if got := hasDuplicate2(tt.nums); got != tt.want {
				t.Errorf("hasDuplicate2(%v) = %t, want %t", tt.nums, got, tt.want)
			}
		})
	}
}
