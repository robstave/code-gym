// DFS 1: The simplest DFS — pre-order traversal on a tiny binary tree
// Idea: visit node, then left, then right (recursion uses the call stack)

const tree = {
  val: "A",
  left: {
    val: "B",
    left: { val: "D", left: null, right: null },
    right: { val: "E", left: null, right: null },
  },
  right: {
    val: "C",
    left: null,
    right: { val: "F", left: null, right: null },
  },
};

// ascii version of above
/*
        A
       / \
      B   C
     / \   \
    D   E   F
*/

// DFS (Depth-First Search) — Pre-order variant: visit, then go deep left, then deep right.
// Why DFS? Each recursive call dives down one child before any siblings.
// "Cutting a branch" is just returning early when there's no node (base case),
// which is a tiny example of pruning/backtracking.
function dfsPreOrder(node, visit) {
  // Base case (prune this branch): nothing to do
  if (!node) return; // early return cuts the branch

  // 1) Visit this node first (pre-order)
  visit(node.val);

  // 2) Explore the left subtree deeply before the right subtree
  dfsPreOrder(node.left, visit);

  // 3) After finishing left, explore right subtree
  dfsPreOrder(node.right, visit);
}

// DFS — Post-order variant (right, left, node)
// Still DFS: fully explore children before visiting the node.
function dfsPostOrder(node, visit) {
  // Base case (prune this branch): nothing to do
  if (!node) return; // early return cuts the branch

  // Explore right subtree deeply first (order choice doesn't change DFS nature)
  dfsPostOrder(node.right, visit);

  // Then explore left subtree deeply
  dfsPostOrder(node.left, visit);

  // Visit this node last (post-order relative to chosen child order)
  visit(node.val);
}

// DFS — In-order variant (left, node, right), common for binary search trees.
// Still DFS: we dive left fully, come back to node, then dive right fully.
function dfsInOrder(node, visit) {
  // Base case (prune this branch): nothing to do
  if (!node) return; // early return cuts the branch

  // Explore left subtree deeply first
  dfsInOrder(node.left, visit);

  // Visit node between left and right (in-order)
  visit(node.val);

  // Then explore right subtree deeply
  dfsInOrder(node.right, visit);
}

const order = [];
dfsPreOrder(tree, (v) => order.push(v));
console.log("DFS pre-order:", order.join(" -> "));
// Expected: A -> B -> D -> E -> C -> F

const order2 = [];
dfsPostOrder(tree, (v) => order2.push(v));
console.log("DFS post-order (R,L,node):", order2.join(" -> "));
// Expected (with right-then-left child order): F -> C -> E -> D -> B -> A


const order3 = [];
dfsInOrder(tree, (v) => order3.push(v));
console.log("DFS in-order:", order3.join(" -> "));
// Expected: D -> B -> E -> A -> C -> F
