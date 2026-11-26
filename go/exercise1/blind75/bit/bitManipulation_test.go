package utils

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		value int
		want  int
	}{
		{value: 0, want: 0},
		{value: 1, want: 1},
		{value: 2, want: 1},
		{value: 7, want: 3},
		{value: 11, want: 3},
		{value: 128, want: 1},
		{value: 255, want: 8},
	}

	for _, tt := range tests {
		if got := hammingWeight(tt.value); got != tt.want {
			t.Errorf("hammingWeight(%d) = %d, want %d", tt.value, got, tt.want)
		}
		if got := hammingWeightIdiomatic(tt.value); got != tt.want {
			t.Errorf("hammingWeightIdiomatic(%d) = %d, want %d", tt.value, got, tt.want)
		}

		if got := hammingWeightBit(tt.value); got != tt.want {
			t.Errorf("hammingWeightBit(%d) = %d, want %d", tt.value, got, tt.want)
		}
	}
}
