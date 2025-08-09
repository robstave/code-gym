// tree_paths.go
package util

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// AllRootToLeafPaths returns all possible paths from the root node to all leaf nodes.
// Each path is represented as a slice of integers, starting with the root value
// and ending with the leaf value.
// A leaf node is a node with no children (both Left and Right are nil).
// The order of returned paths does not matter.
func AllRootToLeafPaths(root *TreeNode) [][]int {
	// Your implementation here
	return nil
}

// SumRootToLeafPaths calculates all root-to-leaf path sums in a binary tree
// where each path sum is obtained by concatenating all node values from root to leaf.
//
// For example, in this tree:
//     1
//    / \
//   2   3
// There are two paths: 1->2 and 1->3, which represent 12 and 13, and sum to 25.
//
// If the tree is:
//      4
//     / \
//    9   0
//   / \
//  5   1
// The paths are 4->9->5 (495), 4->9->1 (491), and 4->0 (40), summing to 1026.
func SumRootToLeafPaths(root *TreeNode) int {
	// Your implementation here
	return 0
}
