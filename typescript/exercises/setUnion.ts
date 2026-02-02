/*
Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must be unique and you may return the result in any order.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4] or [4,9]
*/

export function intersection1(nums1: number[], nums2: number[]): number[] {
  let set1: Set<number> = new Set<number>(nums1);
  let set2: Set<number> = new Set<number>(nums2);

  let results: number[] = [];
  for (const check of set1) {
    if (set2.has(check)) {
      results.push(check);
    }
  }

  // Implement here
  return results;
}

export function intersection2(nums1: number[], nums2: number[]): number[] {
  // Check lengths and create set from larger array
  if (nums1.length > nums2.length) {
    [nums1, nums2] = [nums2, nums1]; // Swap so nums1 is smaller
  }

 
  let set2: Set<number> = new Set<number>(nums2);

  let results: number[] = [];
  for (const check of nums1) {
    if (set2.has(check)) {
      results.push(check);
    }
  }

 
  return results;
}

console.log(intersection1([4, 9, 5], [9, 4, 9, 8, 4]));
console.log(intersection1([1, 2, 3, 4], [0, 0, 1, 1, 1, 2, 2, 2, 2, 1, 0]));

console.log(intersection2([4, 9, 5], [9, 4, 9, 8, 4]));
console.log(intersection2([1, 2, 3, 4], [0, 0, 1, 1, 1, 2, 2, 2, 2, 1, 0]));


// Determine which array is smaller to optimize iteration (optional but good practice)
// Convert one array to a Set for O(1) lookups
// Iterate through the other array and check for existence in the Set
