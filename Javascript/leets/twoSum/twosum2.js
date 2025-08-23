/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum2 = function (nums, target) {
  let m = new Map();

  for (let i = 1; i < nums.length; i++) {
    let nj = target - nums[i];
    m.set(nj, i);
  }

  for (let j = 0; j < nums.length; j++) {
    let nj = nums[j];
    if (m.has(nj)) {
      let i = m.get(nj);
      if (i != j) {
        return [j, i];
      }
    }
  }
};

console.log(twoSum([2, 7, 11, 15], 18));
console.log(twoSum([3, 2, 4], 6));
console.log(twoSum([3, 3], 6));
