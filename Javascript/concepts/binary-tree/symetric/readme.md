## Symmetric Tree (LeetCode 101)

Given the root node of a binary tree, determine whether the tree is symmetric around its center. A tree is symmetric if the left subtree is a mirror reflection of the right subtree, meaning each corresponding node pair holds the same value and the children match in reverse order.

https://leetcode.com/problems/symmetric-tree/description/

### Example diagrams
- ![Example 1](https://assets.leetcode.com/uploads/2021/02/19/symtree1.jpg)
- ![Example 2](https://assets.leetcode.com/uploads/2021/02/19/symtree2.jpg)


https://leetcode.com/problems/symmetric-tree/description/




# Solution
To determine if a binary tree is symmetric, we can use a recursive approach. We will create a helper function that takes two nodes as arguments and checks if they are mirror images of each other. The main function will call this helper function with the left and right children of the root node.


use a depth-first search (DFS) approach to compare the left and right subtrees of the binary tree. The helper function will check if the values of the current nodes are equal and if their children are mirror images of each other.

recall that a depth-first search (DFS) is a tree traversal method that explores as far as possible along each branch before backtracking. In this case, we will use DFS to compare the left and right subtrees of the binary tree.

the pattern is 
1. Check if both nodes are null. If they are, return true (they are symmetric).
2. Check if one of the nodes is null and the other is not. If so, return false (they are not symmetric).
3. Check if the values of the current nodes are equal. If they are not, return false (they are not symmetric).
4. Recursively check if the left child of the first node is a mirror image of the right child of the second node, and if the right child of the first node is a mirror image of the left child of the second node.  If both recursive calls return true, then the current nodes are symmetric.      