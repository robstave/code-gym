/**
 * Longest Common Prefix (LCP)
 * Given an array of strings, return the longest common prefix shared by all strings.
 * If there is no common prefix, return an empty string "".
 *
 * Examples:
 * - ["flow","flower","flop","flowu"] -> "flo" (length 3)
 * - ["dog","racecar","car"] -> ""
 *
 * Do not implement the solution yet. This is a stub for interview practice.
 */

/**
 * TODO: Implement this.
 * @param {string[]} words - Array of strings (assume ASCII/case-sensitive).
 * @returns {string} Longest common prefix string, or "" if none.
 */
function longestCommonPrefix(words) {
  let prefix = "";

  let l = words.length;

  if (l == 0) {
    return prefix;
  }

  if (l == 1) {
    return words[0];
  }

  let firstWordLength = words[0].length;

  for (let i = 0; i < firstWordLength; i++) {
    const ch = words[0][i];
    for (let j = 0;  j< l; j++) {
      if (words[j][i] != ch) {
        return prefix;
      }
    }
    prefix = prefix + ch;
  }

  return prefix; // placeholder so tests run and show FAIL until implemented
}

// --- Test harness ---
function expectEqual(name, actual, expected) {
  const pass = actual === expected;
  console.log(
    `${pass ? "PASS" : "FAIL"} - ${name}: actual=${JSON.stringify(
      actual
    )} expected=${JSON.stringify(expected)}`
  );
}

function runCase(name, words, expectedPrefix) {
  const actualPrefix = longestCommonPrefix(words);
  expectEqual(`${name} (prefix)`, actualPrefix, expectedPrefix);

  // Also check the expected length for clarity
  const actualLen = (actualPrefix || "").length;
  const expectedLen = (expectedPrefix || "").length;
  expectEqual(`${name} (length)`, actualLen, expectedLen);
}

// --- Test cases ---
console.log("=== Longest Common Prefix Tests ===");

// Provided example: common length is 3 ("flo")
runCase(
  "Sample: flow/flower/flop/flowu",
  ["flow", "flower", "flop", "flowu"],
  "flo"
);
 
// No common prefix
runCase("No common", ["dog", "racecar", "car"], "");

// One is a full prefix of others
runCase("One word is prefix of others", ["flo", "flop", "flower"], "flo");

// All identical
runCase("All identical", ["same", "same", "same"], "same");

// Single word
runCase("Single word", ["alone"], "alone");

// Duplicate entries
runCase("Duplicates", ["aa", "aa", "aa"], "aa");


// Empty array
runCase("Empty array", [], "");

// Contains empty string
runCase("Contains empty string", ["", "abc", "ab"], "");

// Case sensitivity (no common due to capital F)
runCase("Case sensitive", ["Flow", "flower"], "");

 


// Export for reuse if needed
module.exports = { longestCommonPrefix };
