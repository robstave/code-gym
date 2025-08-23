/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var topKFrequent = function (nums, k) {
  let m = new Map();

  for (i = 0; i < nums.length; i++) {
    let v = nums[i];

    if (m.has(v)) {
      let count = m.get(v);
      m.set(v, count++);
    } else {
      m.set(v, 1);
    }
  }

  m.values()
};

console.log(topKFrequent([1, 1, 1, 2, 2, 3], 2));
