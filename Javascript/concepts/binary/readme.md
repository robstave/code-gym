
https://www.hello-algo.com/en/chapter_tree/binary_tree_traversal/#2-complexity-analysis



# Binary Tree Traversal
Binary tree traversal is the process of visiting each node in a binary tree in a specific order.


There are two kinds of binary tree traversal: depth-first search (DFS) and breadth-first search (BFS). DFS explores as far as possible along each branch before backtracking, while BFS explores all the nodes at the present depth level before moving on to the nodes at the next depth level.


## Depth-first search (DFS)


Depth-first search (DFS) is a tree traversal method that explores as far as possible along each branch before backtracking. In this case, we will use DFS to compare the left and right subtrees of the binary tree.

In DFS, we can traverse the binary tree in three different orders: pre-order, in-order, and post-order.


this image is from https://www.hello-algo.com/en/chapter_tree/binary_tree_traversal/#2-complexity-analysis

its awesome.   memorize this

//link to image
![image](https://www.hello-algo.com/en/chapter_tree/binary_tree_traversal.assets/binary_tree_dfs.png)

basically....in the resurive calls,  we have 3 steps:
1. visit the current node (pre-order)
2. visit the left subtree (in-order)
3. visit the right subtree (post-order)

it boils down to an evaluation that happens
first, before the recursive calls, we have pre-order traversal

second, we have the recursive call to the left subtree, which is in-order traversal

third, we have the recursive call to the right subtree, which is post-order traversal




## Breadth-first search (BFS)

BFS is a tree traversal method that explores all the nodes at the present depth level before moving on to the nodes at the next depth level. In this case, we will use BFS to compare the left and right subtrees of the binary tree.

Note that the BFS approach is not as efficient as the DFS approach for this problem, because it requires more memory to store the nodes at each level of the tree. However, it can be useful in certain cases where we want to explore the tree level by level.


