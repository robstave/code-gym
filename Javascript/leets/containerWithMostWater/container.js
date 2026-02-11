var calcWater = function (rightHeight, leftHeight, width) {
  return Math.min(rightHeight, leftHeight) * width;
};

/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function (height) {
  // set up two pointers

  let p1 = 0;
  let p2 = height.length - 1;

  let max = 0;
  while (p1 < p2) {
    max = Math.max(max,calcWater(height[p1], height[p2], p2 - p1));

    if (height[p1] < height[p2]) {
      //advance the shorter pointer
      p1++;
    } else {
      p2--;
    }
  }

  return max
};

let height = [1, 8, 6, 2, 5, 4, 8, 3, 7];

console.log("----",calcWater(2,8,9))

console.log(maxArea(height));
