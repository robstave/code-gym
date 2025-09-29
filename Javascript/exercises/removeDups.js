// @ts-check

/**
 * Remove duplicates from a sorted array in-place and return the new length.
 * This function modifies the array in place to contain only unique elements
 * in their original order, and returns the new length.
 * 
 
 * 
 * @param {number[]} list - A sorted array that may contain duplicates
 * @return {number} - Length of the array after removing duplicates
 */
function removeDups(list) {
  let pointer = 0;

  if (list.length <= 1) {
    return list.length;
  }

  for (let i = 1; i < list.length; i++) {
    if (list[i] != list[pointer]) {
      pointer++;
      list[pointer] = list[i];
    }
  }

  const newLength = pointer + 1;

  console.log(
    `Array after removing duplicates: [${list.slice(
      0,
      newLength
    )}], New length: ${newLength}`
  );

  return newLength;
}

// Test cases
console.log("\nTest Case 1: Basic array with duplicates");
removeDups([0, 0, 0, 2, 3]);

console.log("\nTest Case 2: Array with multiple duplicates");
removeDups([0, 0, 1, 1, 2, 3, 4, 4, 5]);

console.log("\nTest Case 3: Array with no duplicates");
removeDups([1, 2, 3, 4, 5]);

console.log("\nTest Case 4: Empty array");
removeDups([]);

console.log("\nTest Case 5: Array with all duplicates");
removeDups([2, 2, 2, 2, 2]);

console.log("\nTest Case 6: Array with duplicates at the end");
removeDups([0, 1, 2, 3, 3, 3]);

console.log("\nTest Case 7: Larger numbers");
removeDups([100, 100, 101, 102, 102, 200, 200]);
