// @ts-check

/**
 * Removes all instances of a value from an array in-place and returns the new length.
 * @param {number[]} nums - The input array to modify
 * @param {number} val - The value to remove
 * @return {number} - The new length of the array after removing elements
 */
var removeElement = function (nums, val) {
  let index = 0;

  for (let i = 0; i < nums.length; i++) {
    // If current element is not the value to remove,
    // copy it to the index position
    if (nums[i] !== val) {
      nums[index] = nums[i];
      index++;
    }
  }

  return index;
};

/**
 * Test helper function
 * @param {number[]} nums - Input array
 * @param {number} val - Value to remove
 * @param {number[]} expected - Expected array after removal
 */
function testRemoveElement(nums, val, expected) {
  const originalNums = [...nums]; // Copy for display
  const newLength = removeElement(nums, val);
  const result = nums.slice(0, newLength);
  const passed = JSON.stringify(result) === JSON.stringify(expected);

  console.log(`Test: Remove ${val} from [${originalNums}]`);
  console.log(`Expected: [${expected}], Got: [${result}]`);
  console.log(`Length: ${newLength}`);
  console.log(`Status: ${passed ? 'PASS' : 'FAIL'}\n`);
}

// Test cases
testRemoveElement([1, 2, 3, 4, 5, 2, 3, 4], 2, [1, 3, 4, 5, 3, 4]);
testRemoveElement([3, 2, 2, 3], 3, [2, 2]);
testRemoveElement([], 1, []);
testRemoveElement([1], 1, []);
testRemoveElement([1, 1, 1], 1, []);
testRemoveElement([1, 2, 3, 4], 5, [1, 2, 3, 4]);
testRemoveElement([1, 1, 2, 2, 3, 3], 2, [1, 1, 3, 3]);
