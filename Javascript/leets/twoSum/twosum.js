// @ts-check

function twoSum(arr, target) {
  // Removed unused variables
  
  // Use a proper object/Map instead of an array for mapping values to indices
  const mapping = {};

  // Create a map of all values to their indices
  for (let i = 0; i < arr.length; i++) {
    mapping[arr[i]] = i;
  }
   
  // Loop through array to find the complement that would sum to target
  for (let j = 0; j < arr.length; j++) {
    let complement = target - arr[j];
    
    // Check if the complement exists in our mapping and is not the same element
    if (complement in mapping && mapping[complement] !== j) {
      return [j, mapping[complement]]; // Return indices of the two numbers
    }
  }
  
  // If no solution is found
  return null;
}

/**
 * Tests the twoSum function with various test cases
 * @param {Array} testCase - Array containing [array, target, expected result]
 * @param {number} caseNumber - The test case number
 */
function runTest(testCase, caseNumber) {
  const [array, target, expected] = testCase;
  const result = twoSum(array, target);
  
  // Compare results (considering order doesn't matter)
  let success = false;
  if (result === null && expected === null) {
    success = true;
  } else if (result && expected) {
    // Sort both arrays for comparison since order of indices doesn't matter
    success = JSON.stringify(result.sort()) === JSON.stringify(expected.sort());
  }
  
  console.log(`Test Case ${caseNumber}: ${success ? 'PASS' : 'FAIL'}`);
  console.log(`  Input: [${array}], Target: ${target}`);
  console.log(`  Expected: ${JSON.stringify(expected)}, Got: ${JSON.stringify(result)}\n`);
}

// Test cases
const testCases = [
  [[1, 4, 5, 7, 18], 12, [1, 3]], // Original test case (4+7=11)
  [[2, 7, 11, 15], 9, [0, 1]],    // LeetCode example (2+7=9)
  [[3, 2, 4], 6, [1, 2]],         // Numbers not in order (2+4=6)
  [[3, 3], 6, [0, 1]],            // Same number twice (3+3=6)
  [[1, 2, 3, 4, 5], 10, null],    // No solution
  [[-1, -2, -3, -4], -7, [2, 3]], // Negative numbers (-3+-4=-7)
  [[0, 4, 3, 0], 0, [0, 3]]       // Zeros in the array (0+0=0)
];

// Run each test case
testCases.forEach((testCase, index) => {
  runTest(testCase, index + 1);
});
