package utils

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "first example", nums: []int{2, 20, 4, 10, 3, 4, 5}, want: 4},
		{name: "second example", nums: []int{0, 3, 2, 5, 4, 6, 1, 1}, want: 7},
		{name: "empty", nums: []int{}, want: 0},
		{name: "single value", nums: []int{42}, want: 1},
		{name: "duplicates only", nums: []int{7, 7, 7}, want: 1},
		{name: "no consecutive values", nums: []int{10, 4, 1}, want: 1},
		{name: "negative values across zero", nums: []int{2, -2, 0, -1, 1}, want: 5},
		{name: "separate sequences", nums: []int{10, 11, 12, 1, 2, 3, 4}, want: 4},
		{name: "constraint endpoints", nums: []int{-1000000000, -999999999, 999999999, 1000000000}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.nums); got != tt.want {
				t.Errorf("longestConsecutive(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutiveNotRight(tt.nums); got != tt.want {
				t.Errorf("longestConsecutive(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}

}

func TestLongestConsecutiveLargeInput(t *testing.T) {
	const sequenceLength = 50000
	nums := make([]int, 2*sequenceLength)
	for i := 0; i < sequenceLength; i++ {
		nums[i] = i
	}

	if got := longestConsecutive(nums); got != sequenceLength {
		t.Errorf("longestConsecutive(large input) = %d, want %d", got, sequenceLength)
	}
}
