



/*
given a number, find an array that sums up to zero
*/

var sumZero = function (n) {
  let arr = [];
  k = Math.floor(n / 2);

  for (i = 1; i <= k; i++) {
    arr.push(i, -i);
  }

  if (n % 2 == 1) {
    arr.push(0);
  }

  return arr;
};

console.log(sumZero(5));
