// Binary search on a sorted (ascending) array.
// Returns the index of target, or null if not found.
function binarySearch(arr, target) {
  let left = 0;
  let right = arr.length - 1;

  while (left <= right) {
    // calc mid.  check

    let mid = left + Math.floor((right - left) / 2);

    value = arr[mid];
    if (value == target) {
      return mid;
    }
    if (value > target) {
      // search lower
      right = mid -1;
    } else {
      // search higher
      left = mid +1;
    }
  }

  return null;
}

// Export for reuse if needed
module.exports = { binarySearch };

// --- Test harness ---
function expectEqual(name, actual, expected) {
  const pass = actual === expected;
  console.log(
    `${pass ? "PASS" : "FAIL"} - ${name}: actual=${actual} expected=${expected}`
  );
}

function expectFound(name, arr, target) {
  const idx = binarySearch(arr, target);
  const pass = idx !== null && arr[idx] === target;
  console.log(
    `${pass ? "PASS" : "FAIL"} - ${name}: idx=${idx} value=${
      idx !== null ? arr[idx] : null
    }`
  );
}

console.log("\n=== binarySearch tests ===");

expectEqual("found mid", binarySearch([1, 3, 5, 7, 8, 9, 11, 34, 100], 5), 2);

// Basic
expectEqual("found mid", binarySearch([1, 3, 5, 7, 9], 7), 3);
expectEqual("not found", binarySearch([1, 3, 5, 7, 9], 8), null);

// Edges
expectEqual("first element", binarySearch([1, 3, 5, 7, 9], 1), 0);
expectEqual("last element", binarySearch([1, 3, 5, 7, 9], 9), 4);
expectEqual("empty array", binarySearch([], 10), null);
expectEqual("single found", binarySearch([42], 42), 0);
expectEqual("single not found", binarySearch([42], 7), null);

// Mixed values
expectEqual("negatives", binarySearch([-10, -3, 0, 5, 12], -3), 1);

// Duplicates: any index pointing to the target is acceptable
expectFound("duplicates any index", [1, 2, 2, 2, 3], 2);
