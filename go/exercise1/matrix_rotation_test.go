// matrix_rotation_test.go
package util

import (
	"reflect"
	"testing"
)

func TestRotateMatrix90(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			name: "3x3 matrix",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},

		{
			name: "2x2 matrix",
			input: [][]int{
				{1, 2},
				{3, 4},
			},
			want: [][]int{
				{3, 1},
				{4, 2},
			},
		},

		{
			name:  "1x1 matrix",
			input: [][]int{{5}},
			want:  [][]int{{5}},
		},

		{
			name:  "empty matrix",
			input: [][]int{},
			want:  [][]int{},
		},
		{
			name: "4x4 matrix",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: [][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of the input to ensure it's not modified
			inputCopy := make([][]int, len(tc.input))
			for i, row := range tc.input {
				inputCopy[i] = make([]int, len(row))
				copy(inputCopy[i], row)
			}

			got := RotateMatrix90(inputCopy)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("RotateMatrix90() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRotateMatrixInPlace90(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			name: "3x3 matrix",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},

		{
			name: "2x2 matrix",
			input: [][]int{
				{1, 2},
				{3, 4},
			},
			want: [][]int{
				{3, 1},
				{4, 2},
			},
		},
		{
			name:  "1x1 matrix",
			input: [][]int{{5}},
			want:  [][]int{{5}},
		},
		{
			name: "4x4 matrix",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: [][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of the input since it will be modified in place
			input := make([][]int, len(tc.input))
			for i, row := range tc.input {
				input[i] = make([]int, len(row))
				copy(input[i], row)
			}

			RotateMatrixInPlace90(input)
			if !reflect.DeepEqual(input, tc.want) {
				t.Errorf("RotateMatrixInPlace90() modified input to %v, want %v", input, tc.want)
			}
		})
	}
}
