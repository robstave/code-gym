// Brute-force version for exactly 3 dice with n sides targeting a sum
// Uses 3 nested loops to enumerate all possibilities

function findCombinations3Dice(n, target) {
  const results = [];

  for (let a = 1; a <= n; a++) {
    for (let b = 1; b <= n; b++) {
      for (let c = 1; c <= n; c++) {
        if (a + b + c === target) {
          results.push([a, b, c]);
        }
      }
    }
  }

  return results;
}

// Demo
console.log('3 dice, 6 sides, sum 10 (brute-force) count:', findCombinations3Dice(6, 10).length);
console.log(findCombinations3Dice(6, 10));
