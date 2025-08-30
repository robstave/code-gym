const { ListNode } = require("./node");
const { LinkedList } = require("./linkedList");
/*
Add two non-negative integers represented as singly linked lists.

Assumptions:
- Each node holds a single digit 0-9.
- The head is the least-significant digit (reverse order).
- Lists may have different lengths.
- Result is returned as a new linked list (also reverse order).

Example:
Input:  l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
*/

let l1 = new LinkedList();
let l2 = new LinkedList();

//l1.arrayToList([9,9,9,9,9,9,9])
//l2.arrayToList([9,9,9,9])

// Simple case: 353 + 474 = 827  -> list: [7,2,8]
l1.arrayToList([3, 5, 3]);
l2.arrayToList([4, 7, 4]);

// Add two lists given their head nodes; returns the head node of the sum list
// Time: O(max(n, m)), Space: O(max(n, m)) for the output
var sum = function (l1, l2) {
  let ptr1 = l1;
  let ptr2 = l2;
  let carry = 0;
  let sumHead = null;
  let sumPtr = null;

  while (ptr1 != null || ptr2 != null) {
    let x = 0;
    let y = 0;
    // Read digits from current nodes (default to 0 if one list is shorter)
    if (ptr1 != null) {
      x = ptr1.data;
      ptr1 = ptr1.next;
    }
    if (ptr2 != null) {
      y = ptr2.data;
      ptr2 = ptr2.next;
    }

    // Current digit = (x + y + carry) % 10, update carry accordingly
    let nodeData = x + y + carry;


    // Trace logging; comment out if not needed
    //console.log(`x=${x}, y=${y} , carry=${carry}, nodeData=${nodeData}`);

    // clear the carry
    carry = 0;
    if (nodeData >= 10) {
      nodeData = nodeData % 10;
      carry = 1;
    }  

    // Append new digit node to result list
    let sumNode = new ListNode();
    sumNode.data = nodeData;

    if (sumHead == null) {
      sumHead = sumNode;
    } else {
      sumPtr.next = sumNode;
    }
    sumPtr = sumNode;
  }

  // Append final carry if present
  if (carry > 0) {
    let sumNode = new ListNode(carry);
    sumPtr.next = sumNode;
  }

  return sumHead;
};

console.log("Sum1:");
let result = sum(l1.head, l2.head);
resultLinkedList = new LinkedList(result);
console.log(resultLinkedList.toArray());

l1.arrayToList([9, 9, 9, 9, 9, 9, 9]);
l2.arrayToList([9, 9, 9, 9]);

console.log("Sum2:");
result = sum(l1.head, l2.head);
resultLinkedList = new LinkedList(result);
console.log(resultLinkedList.toArray());
//Output: [8,9,9,9,0,0,0,1]

// Additional test cases
function runCase(label, a, b, expected) {
  const A = new LinkedList();
  const B = new LinkedList();
  A.arrayToList(a);
  B.arrayToList(b);
  const r = sum(A.head, B.head);
  const R = new LinkedList(r).toArray();
  console.log(label + ":");
  console.log("  A:", a, "B:", b);
  console.log("  ->", R, expected ? "expected:" : "");
  if (expected) console.log("     ", expected);
}

// Different lengths, carry propagates
runCase("Sum3", [9, 9, 9], [1], [0, 0, 0, 1]);

// No carry anywhere
runCase("Sum4", [2, 2], [5, 5], [7, 7]);

// One empty list
runCase("Sum5", [], [7, 3], [7, 3]);

// Zeros
runCase("Sum6", [0], [0], [0]);

