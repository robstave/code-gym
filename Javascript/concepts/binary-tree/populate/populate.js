//https://www.geeksforgeeks.org/dsa/construct-complete-binary-tree-given-array/
//
// other notes 
// https://www.hello-algo.com/en/chapter_tree/binary_tree_traversal/#__tabbed_1_5

// Javascript program to construct binary tree from
// given array in level order fashion


class Node {
  constructor(data) {
    this.left = null;
    this.right = null;
    this.data = data;
  }
}

// build the tree recursively
//
// start with the parent node.

// Function to insert nodes in level order
function insertLevelOrder(arr, i) {
  let root = null;
  // Base case for recursion
  if (i < arr.length) {
    root = new Node(arr[i]);

    // insert left child
    root.left = insertLevelOrder(arr, 2 * i + 1);

    // insert right child
    root.right = insertLevelOrder(arr, 2 * i + 2);
  }
  return root;
}

// Function to print tree nodes in InOrder fashion
// this is a breadth first traversal
function inOrder(root) {
  if (root != null) {
  console.log(root.data + " ");
    inOrder(root.left);
  
    inOrder(root.right);
  }
}

/*
Input  :  arr[] = {1, 2, 3, 4, 5, 6}
Output : Root of the following tree
                  1
                 / \
                2   3
               / \ /
              4  5 6
              */



let arr1 = [1, 2, 3, 4, 5, 6 ];
let root1 = insertLevelOrder(arr1, 0);
console.log("----- tree ( condensed )")
console.log(root1);
console.log("----- ")
inOrder(root1);

/*
Input: arr[] = {1, 2, 3, 4, 5, 6, 6, 6, 6, 6}
Output: Root of the following tree
                   1
                  / \
                 2   3
                / \ / \
               4  5 6  6
              / \ /
             6  6 6
*/

let arr2 = [1, 2, 3, 4, 5, 6, 6, 6, 6];
let root2 = insertLevelOrder(arr2, 0);
console.log("----- tree ( condensed )")
console.log(root2);
console.log("----- ")
inOrder(root2);


