// Brute-force (bitmask) version: generate all subsets using bits 0..(2^n - 1)
// This is an iterative approach, not recursive, but it enumerates the same power set.

function powerSetBitmask(arr) {
  const n = arr.length;
  const result = [];
  const total = 1 << n; // 2^n

  for (let mask = 0; mask < total; mask++) {
    const subset = [];
    for (let i = 0; i < n; i++) {
      if (mask & (1 << i)) subset.push(arr[i]);
    }
    result.push(subset);
  }
  return result;
}

// Demo
const demo = ['A', 'B', 'C'];
const sets = powerSetBitmask(demo);
console.log('Bitmask power set size (expected 8):', sets.length);
console.log('Subsets:', sets.map(s => '[' + s.join('') + ']').join(' '));

module.exports = { powerSetBitmask };
