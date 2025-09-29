// Exercise demo: build a list and try core operations
const { ListNode } = require("./node");
const { LinkedList } = require("./linkedList");

let myList = new LinkedList();

function runCase(data) {
	// Accept null/undefined by treating as an empty list
	const input = data ?? [];
	myList.arrayToList(input);
	const arr = myList.toArray();
	const mid = myList.getMiddle();
	console.log(arr);
	console.log(mid ? mid.data : null);
}

// Original cases
runCase([1, 2, 3, 4, 5]);
runCase([1, 2, 3, 4]);
runCase([1, 2, 3]);
runCase([1, 2]);
runCase([1]);
runCase([]);

// Also handle null input explicitly
runCase(null);
