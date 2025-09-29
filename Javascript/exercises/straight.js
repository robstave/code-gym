/*

Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
Output: true
Example 2:



Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
Output: false
*/

var getSlope = function (p1, p2) {
  if (p2[0] - p1[0] == 0) {
    return Infinity;
  }
  return  ((p2[1] - p1[1]) / (p2[0] - p1[0]));
};

var checkStraightLine = function (coordinates) {
  count = coordinates.length;
  if (count == 0) {
    return false;
  }
  if (count == 1) {
    return true;
  }

  p1 = coordinates[0];
  p2 = coordinates[1];

  slope = getSlope(p1, p2);
  console.log(slope);

  for (i = 1; i < count; i++) {
    if (getSlope(coordinates[0], coordinates[i]) != slope) {
      return false;
    }
  }
  return true;
};

console.log(
  checkStraightLine([
    [1, 2],
    [2, 3],
    [3, 4],
    [4, 5],
    [5, 6],
    [6, 7],
  ]), true
);

console.log(
  checkStraightLine([
    [1, 1],
    [2, 2],
    [3, 4],
    [4, 5],
    [5, 6],
    [7, 7],
  ]), false
);

console.log(
  checkStraightLine([
    [0, 0],
    [0, 1],
    [0, -1],
  ]), true
);

console.log(
  checkStraightLine([[1,-8],[2,-3],[1,2]]), false
);

console.log(
  checkStraightLine([[1,1],[2,2],[2,0]]), false
);


