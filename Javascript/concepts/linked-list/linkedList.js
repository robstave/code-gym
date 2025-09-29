const { ListNode } = require("./node");

// Singly linked list. Methods mutate the list in-place.
class LinkedList {
  // Create a list with an optional head node
  constructor(head = null) {
    this.head = head;
  }

  // Convert list to array (head -> tail)
  toArray() {
    let arr = [];

    let node = this.head;

    while (node != null) {
      arr.push(node.data);
      node = node.next;
    }

    return arr;
  }

  // Count nodes
  size() {
    let sz = 0;

    let node = this.head;

    while (node != null) {
      sz++;
      node = node.next;
    }

    return sz;
  }

  // Return the head node (or null)
  getFirst() {
    return this.head;
  }

  // Insert at head
  addAtHead(val) {
    let newHead = new ListNode(val);
    newHead.next = this.head;
    this.head = newHead;
  }

  // Append at tail
  addAtTail(val) {
    let tail = new ListNode(val);

    if (this.head == null) {
      this.head = tail;
      return;
    }
    let node = this.head;

    while (node.next != null) {
      node = node.next;
    }
    node.next = tail;
  }

  // Return node at zero-based index (or null if out of range)
  getAtIndex(index) {
    let ptr = 0;

    let node = this.head;

    while (node != null) {
      if (ptr == index) {
        return node;
      }
      ptr++;
      node = node.next;
    }

    return null;
  }

  // Return node at zero-based index (or null if out of range)
  getMiddle() {
    let ptr = 0;

    if (this.head == null) {
      return null;
    }
    let node = this.head;
    let middle = this.head;

    let count = 0;

    while (node.next != null) {
      node = node.next;
      count++;
      if (count % 2 == 0) {
        middle = middle.next;
      }
    }

    return middle;
  }

  getTail() {
    let node = this.head;

    while (node != null && node.next != null) {
      node = node.next;
    }

    return node;
  }

  // Insert value before the current node at index; no-op if out of range
  addAtIndex(index, val) {
    if (index == 0) {
      this.addAtHead(val);
      return;
    }

    let newNode = new ListNode(val);
    if (this.head == null) {
      this.head = newNode;
      return;
    }
    let ptr = 0;

    let node = this.head;
    let lastNode = this.head;

    while (node != null) {
      lastNode = node;
      node = node.next;
      ptr++;
      if (ptr == index) {
        newNode.next = node;
        lastNode.next = newNode;
      }
    }
  }

  // TODO: Delete node at index
  deleteAtIndex(index) {
    if (index == 0) {
      this.head = this.head.next;
      return;
    }

    let ptr = 0;
    let node = this.head;
    let lastNode = this.head;

    while (node != null) {
      lastNode = node;
      node = node.next;
      ptr++;
      if (ptr == index) {
        if (node == null) {
          lastNode.next = null;
        } else {
          lastNode.next = node.next;
        }
        return;
      }
    }
  }

  reverse() {
    let node = this.head;
    let prev = null;

    while (node != null) {
      let temp = node.next; // store off node.next

      node.next = prev; // point node backwards
      prev = node; // push prev pointer
      node = temp; //advance pointer
    }
    this.head = prev;
  }

  kthFromEnd(k) {
    // trick is to just search the list, but after
    //  the pointer hits K, add a trailing pointer
    let node = this.head;
    let trailing = this.head;

    // span the list until to you get to K
    while (k-- > 0) {
      if (!node) return null;
      node = node.next;
    }

    // span the list with trailing pointer running behind
    while (node) {
      node = node.next;
      trailing = trailing.next;
    }
    return trailing;
  }

  hasCycle() {
    let slow = this.head;
    let fast = this.head;

    while (fast && fast.next) {
      slow = slow.next;
      fast = fast.next.next;
      if (slow === fast) return true;
    }
    return false;
  }

breakCycle() {
    let slow = this.head;
    let slowBehind = null;
    
    let fast = this.head;

    while (fast && fast.next) {
    
        slowBehind = slow
    
      slow = slow.next;
      fast = fast.next.next;
      if (slow === fast) {
        slowBehind.next = null
        return true;
      }
    }
    return false;
  }

  arrayToList(arr) {
    let ptr = null;
    this.head = null;

    arr.forEach((element) => {
      let node = new ListNode(element);

      if (this.head == null) {
        this.head = node;
        ptr = node;
      } else {
        ptr.next = node;
        ptr = node;
      }
    });
  }
}

module.exports = { LinkedList };
