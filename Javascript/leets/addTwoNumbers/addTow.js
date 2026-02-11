//todo...this is a class?
function ListNode(val, next) {
  this.val = val === undefined ? 0 : val;
  this.next = next === undefined ? null : next;
}

/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function (l1, l2) {};

//Input: l1 = [2,4,3], l2 = [5,6,4]
//Output: [7,0,8]
//Explanation: 342 + 465 = 807.

var number1a = new ListNode(2);
var number1b = new ListNode(4);
var number1c = new ListNode(3);

var number2a = new ListNode(5);
var number2b = new ListNode(6);
var number2c = new ListNode(4);

number1a.next = number1b;
number1b.next = number1c;

number2a.next = number2b;
number2b.next = number2c;

var sum = function (l1, l2) {
  let carry = 0;
  let solutionHead = null;
  let solutionPtr = null;
  let n1Ptr = l1;
  let n2Ptr = l2;

  while (n1Ptr != null || n2Ptr != null) {
    let x = 0;
    let y = 0;
 
    if (n1Ptr != null) {
      x = n1Ptr.val;
      n1Ptr = n1Ptr.next;
    }
    if (n2Ptr != null) {
      y = n2Ptr.val;
      n2Ptr = n2Ptr.next;
    }

    let sum = x + y + carry;

    if (sum >= 10) {
      sum = sum - 10;
      carry = 1;
    } else {
      carry = 0;
    }

    console.log(sum, carry);

    let ptr = new ListNode(sum);
    if (solutionHead == null) {
      solutionHead = ptr;
      solutionPtr = solutionHead;
    } else {
      solutionPtr.next = ptr;
      solutionPtr = solutionPtr.next;
    }
  }

  if (carry == 1) {
 
    solutionPtr.next = new ListNode(1);
  }

  return solutionHead;
};

console.log(sum(number1a, number2a));
