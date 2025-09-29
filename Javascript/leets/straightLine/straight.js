/*

Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
Output: true
Example 2:



Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
Output: false
*/

var isLine = function (p1, p2, p3) {

   return (p2[1] - p1[1]) * (p3[0] - p2[0]) == (p3[1] - p2[1]) * (p2[0] - p1[0]);
};

var checkStraightLine = function (coordinates) {
  count = coordinates.length;
  if (count == 0) {
    return false;
  }
  if (count == 1) {
    return true;
  }

  for (i = 0; i < count - 2; i++) {
    if (!isLine(coordinates[i], coordinates[i+1], coordinates[i+2])) {
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
  ]),
  true
);

console.log(
  checkStraightLine([
    [1, 1],
    [2, 2],
    [3, 4],
    [4, 5],
    [5, 6],
    [7, 7],
  ]),
  false
);

console.log(
  checkStraightLine([
    [0, 0],
    [0, 1],
    [0, -1],
  ]),
  true
);

console.log(
  checkStraightLine([
    [1, -8],
    [2, -3],
    [1, 2],
  ]),
  false
);

console.log(
  checkStraightLine([
    [1, 1],
    [2, 2],
    [2, 0],
  ]),
  false
);
