const { TreeNode } = require("../node/treeNode");

function getExample1() {
  let ll = new TreeNode(3);
  let lr = new TreeNode(4);
  let l = new TreeNode(2, ll, lr);

  let rl = new TreeNode(4);
  let rr = new TreeNode(3);
  let r = new TreeNode(2, rl, rr);

  return new TreeNode(1, l, r);
}

function getExample2() {
  let ll = new TreeNode(3);
  
  let lr = new TreeNode(4);
  let l = new TreeNode(2, ll, lr);

  let rl = new TreeNode(4);
  let rr = new TreeNode(5);
  let r = new TreeNode(2, rl, rr);

  return new TreeNode(1, l, r);
}

var isSymmetric = function (root) {
  if (root.left === null || root.right === null) return false;
  return inOrderCheck(root.left, root.right);
};

function inOrderCheck(roota, rootb) {
  if (roota === null && rootb === null) return true;
  if (roota === null || rootb === null) return false;

  console.log(roota.val + " " + rootb.val);
  if (roota.val !== rootb.val) {
    return false;
  }

  // Traverse the left subtree
  if (!inOrderCheck(roota.left, rootb.right)) {
    return false;
  }

  // Visit the root node

  // Traverse the right subtree
  if (!inOrderCheck(roota.right, rootb.left)) {
    return false;
  }

  return true;
}

console.log(getExample1())
console.log("--");
console.log(isSymmetric(getExample1()));
console.log("--x");

console.log(isSymmetric(getExample2()));
