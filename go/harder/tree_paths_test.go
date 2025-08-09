// tree_paths_test.go
package util

import (
	"reflect"
	"sort"
	"testing"
)

// Helper function to create a simple binary tree
func createTree(values []int, nullIndex int) *TreeNode {
	if len(values) == 0 || values[0] == nullIndex {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != nullIndex {
			node.Left = &TreeNode{Val: values[i]}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nullIndex {
			node.Right = &TreeNode{Val: values[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// Helper function to normalize path ordering for comparison
func normalizePaths(paths [][]int) [][]int {
	sort.Slice(paths, func(i, j int) bool {
		// Compare path elements
		for k := 0; k < len(paths[i]) && k < len(paths[j]); k++ {
			if paths[i][k] != paths[j][k] {
				return paths[i][k] < paths[j][k]
			}
		}
		return len(paths[i]) < len(paths[j])
	})
	return paths
}

func TestAllRootToLeafPaths(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		null   int // Value representing null in the values array
		want   [][]int
	}{
		{
			name:   "simple tree",
			values: []int{1, 2, 3},
			null:   -1,
			want:   [][]int{{1, 2}, {1, 3}},
		},
		{
			name:   "balanced tree with depth 3",
			values: []int{1, 2, 3, 4, 5, 6, 7},
			null:   -1,
			want:   [][]int{{1, 2, 4}, {1, 2, 5}, {1, 3, 6}, {1, 3, 7}},
		},
		{
			name:   "unbalanced tree",
			values: []int{1, 2, 3, 4, -1, -1, 5},
			null:   -1,
			want:   [][]int{{1, 2, 4}, {1, 3, 5}},
		},
		{
			name:   "single node",
			values: []int{1},
			null:   -1,
			want:   [][]int{{1}},
		},
		{
			name:   "null root",
			values: []int{},
			null:   -1,
			want:   [][]int{},
		},
		{
			name:   "deep left path",
			values: []int{1, 2, -1, 3, -1, -1, -1, 4},
			null:   -1,
			want:   [][]int{{1, 2, 3, 4}},
		},
		{
			name:   "complex tree",
			values: []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, -1, 5, 1},
			null:   -1,
			want:   [][]int{{5, 4, 11, 7}, {5, 4, 11, 2}, {5, 8, 13}, {5, 8, 4, 5}, {5, 8, 4, 1}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := createTree(tc.values, tc.null)
			got := AllRootToLeafPaths(root)

			// Normalize paths for comparison
			normalizedGot := normalizePaths(got)
			normalizedWant := normalizePaths(tc.want)

			if !reflect.DeepEqual(normalizedGot, normalizedWant) {
				t.Errorf("AllRootToLeafPaths() = %v, want %v", normalizedGot, normalizedWant)
			}
		})
	}
}

func TestSumRootToLeafPaths(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		null   int
		want   int
	}{
		{
			name:   "simple tree",
			values: []int{1, 2, 3},
			null:   -1,
			want:   12 + 13, // 25
		},
		{
			name:   "example from description",
			values: []int{4, 9, 0, 5, 1},
			null:   -1,
			want:   495 + 491 + 40, // 1026
		},
		{
			name:   "single node",
			values: []int{7},
			null:   -1,
			want:   7,
		},
		{
			name:   "null root",
			values: []int{},
			null:   -1,
			want:   0,
		},
		{
			name:   "large numbers",
			values: []int{9, 8, 7, 6, 5, 4, 3},
			null:   -1,
			want:   986 + 985 + 974 + 973, // 3918
		},
		{
			name:   "zeros in tree",
			values: []int{0, 1, 0},
			null:   -1,
			want:   1 + 0, // 1
		},
		{
			name:   "complex tree",
			values: []int{6, 3, 5, 2, 5, 0, 8, -1, -1, 7, 1, -1, -1, -1, 9},
			null:   -1,
			want:   632 + 635 + 637 + 631 + 650 + 658 + 659, // 3902
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := createTree(tc.values, tc.null)
			got := SumRootToLeafPaths(root)

			if got != tc.want {
				t.Errorf("SumRootToLeafPaths() = %v, want %v", got, tc.want)
			}
		})
	}
}
