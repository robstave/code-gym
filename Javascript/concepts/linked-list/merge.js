const { ListNode } = require("./node");
const { LinkedList } = require("./linkedList");

/**
 * Merge two sorted singly linked lists into one sorted list.
 * - Inputs are heads of lists sorted in non-decreasing order.
 * - Runs in O(n + m) time and O(1) extra space (relinks existing nodes).
 * @param {ListNode|null} l1
 * @param {ListNode|null} l2
 * @returns {ListNode|null} head of merged sorted list
 */
function merge(l1, l2) {
  // Pointers to traverse the input lists
  let ptr1 = l1;
  let ptr2 = l2;

  // Dummy head to simplify tail handling; tail always points to last node in result
  const dummy = new ListNode();
  let tail = dummy;

  while (ptr1 != null || ptr2 != null) {
    if (ptr1 != null && ptr2 != null) {
      if (ptr1.data < ptr2.data) {
        // trace: taking from list 1
        console.log(`L1:${ptr1.data}`);
        tail.next = ptr1;
        tail = tail.next;
        ptr1 = ptr1.next;
      } else {
        // trace: taking from list 2
        console.log(`L2:${ptr2.data}`);
        tail.next = ptr2;
        tail = tail.next;
        ptr2 = ptr2.next;
      }
    } else {
      // One list exhausted: append the remainder of the other and finish
      if (ptr1 != null) {
        tail.next = ptr1;
        break;
      }
      if (ptr2 != null) {
        tail.next = ptr2;
        break;
      }
    }
  }
  return dummy.next;
}

let l1 = new LinkedList();
let l2 = new LinkedList();

l1.arrayToList([3, 5, 9, 19]);
l2.arrayToList([4, 7, 8, 12, 14, 15, 21, 22]);

console.log(l1.toArray());
console.log(l2.toArray());
let result = merge(l1.head, l2.head);

const r = new LinkedList(result);

console.log(r.toArray());
 