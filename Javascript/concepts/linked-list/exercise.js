// Exercise demo: build a list and try core operations
const { ListNode } = require("./node");
const { LinkedList } = require("./linkedList");

var node1 = new ListNode(1);
var node2 = new ListNode(2);
var node3 = new ListNode(4);
var node4 = new ListNode(6);

node1.next = node2;
node2.next = node3;
node3.next = node4;

const myList = new LinkedList(node1);

console.log("Linked List:");
console.log("----");
console.log("initial array:");
console.log(myList.toArray());
console.log("Add 55 to head--");
myList.addAtHead(55);
console.log(myList.toArray());

console.log("----");
console.log("Print element at index 1");
console.log(myList.getAtIndex(1).data);
console.log("Print element at index 3");
console.log(myList.getAtIndex(3).data);

console.log("----");
console.log(myList.toArray());
console.log("Add 88 at index 3");
myList.addAtIndex(3, 88);
console.log(myList.toArray());

console.log("Add 999 at index 55 ( will fail)");
myList.addAtIndex(555, 999);
console.log(myList.toArray());
console.log(myList.toArray());

myList.addAtIndex(6, 62);
console.log(myList.toArray());
console.log("delete element at index 3");
myList.deleteAtIndex(3);
console.log(myList.toArray());

console.log("print middle");
console.log(myList.getMiddle().data);

console.log("delete element at index 0");
myList.deleteAtIndex(0);
console.log(myList.toArray());

console.log("print middle");
console.log(myList.getMiddle().data);


console.log("delete element at index 4");
myList.deleteAtIndex(4);
console.log(myList.toArray());

console.log("reverse");
myList.reverse();
console.log(myList.toArray());

console.log("Cycle test");
tail = myList.getTail();
node = myList.getAtIndex(1);
console.log(`connecting tail ${tail.data} to node ${node.data}`);
tail.next = node;
console.log("Cycle created");
console.log("has cycle?", myList.hasCycle());

console.log("Break Cycle");
myList.breakCycle()
console.log("has cycle?", myList.hasCycle());


console.log("reset with array");
myList.arrayToList([1, 2, 3, 4, 5]);

console.log(myList.toArray());
