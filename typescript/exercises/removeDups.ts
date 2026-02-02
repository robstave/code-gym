/**
 * Remove Duplicates from Sorted Array — Exercise
 *
 * Given an integer array `nums` sorted in non-decreasing order, remove the duplicates
 * in-place such that each unique element appears only once. The relative order of the
 * elements should be kept the same.
 *
 * After removing the duplicates, return the number of unique elements `k`.
 * The first `k` elements of `nums` should contain the unique numbers in sorted order.
 * The remaining elements beyond index `k - 1` can be ignored.
 *
 * Custom Judge:
 * The judge will call your implementation like:
 *   const k = removeDuplicates(nums);
 *   // assert k === expected.length
 *   // assert nums.slice(0, k) matches expected
 *
 * Constraints / Requirements:
 * - Modify the array in-place.
 * - Use O(1) extra memory.
 * - Do not allocate a new array for the result.
 *
 * Example:
 *   const nums = [0,0,1,1,1,2,2,3,3,4];
 *   const k = removeDuplicates(nums);
 *   // After call: k === 5, nums.slice(0, k) === [0,1,2,3,4]
 *
 * NOTE: This file is an exercise stub. Do NOT provide a full implementation here.
 */

/**
 * Implement this function to remove duplicates in-place and return k.
 * Do not change the function signature.
 */
export function removeDuplicates(nums: number[]): number {
  let pointer: number = 1;

  if (nums.length < 2) {
    return nums.length;
  }
  for (let index = 1; index < nums.length; index++) {
    if (nums[index] !== nums[pointer]) {
      
      nums[pointer] = nums[index];
      pointer++;
    }
  }
  return pointer;
}

let array = [0, 0, 1, 2, 3, 4, 4, 4, 6, 8, 8, 9];

const result = removeDuplicates(array);
console.log(result + ":" + array);

// Replace single ad-hoc example with a few simple testcases.
const tests: number[][] = [
	[],                          // empty
	[42],                        // single element
	[5, 5, 5, 5],                // all duplicates
	[0, 0, 1, 2, 3, 3, 4],       // mixed duplicates
];

for (const t of tests) {
	const arr = t.slice(); // clone because function modifies in-place
	const k = removeDuplicates(arr);
	console.log(JSON.stringify(t) + " -> k=" + k + ", result=" + JSON.stringify(arr.slice(0, k)));
}
