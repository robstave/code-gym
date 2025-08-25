/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

 

Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]
 
*/
/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var moveZeroes = function (nums) {
  let l = nums.length;
  let pointer = 0;

  if (l < 2) {
    return nums;
  }

  for (let i = 0; i < l; i++) {
    if (nums[i] !== 0) {
      if (pointer !== i) {
        [nums[pointer], nums[i]] = [nums[i], nums[pointer]];
      }
      pointer++;
    }
  }

  return nums;
};

console.log(moveZeroes([1, 0, 0, 3, 12])); // [1,3,12,0,0]
console.log(moveZeroes([0, 1, 0, 3, 12])); // [1,3,12,0,0]
console.log(moveZeroes([1, 0])); // [1,0]
console.log(moveZeroes([0, 1])); // [1,0]
