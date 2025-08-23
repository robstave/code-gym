/**
 * @param {string} s
 * @return {number}
 */
var countGoodSubstrings = function (s) {
  let l = s.length;
  let count = 0;
  if (l <3){
    return 0;
  }
  for (i = 0; i < l - 2; i++) {
    ss = s.substring(i, i + 3);

    if (ss[0] != ss[1] && ss[0] != ss[2] && ss[1] != ss[2]) {
      count++;
    }
  }

  return count;
};

console.log(countGoodSubstrings("xyzzaz"));
