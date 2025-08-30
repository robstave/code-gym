/**
 * Find all ordered combinations of dice rolls that sum to a target.
 * Backtracking template: choose -> explore -> un-choose, with pruning.
 *
 * Inputs
 * - diceCount: number of dice to roll
 * - sides: number of faces per die (values 1..sides)
 * - targetSum: desired total sum across all dice
 *
 * Output
 * - Array of arrays, each inner array is a length `diceCount` combination
 *   like [2, 3, 5] whose sum is targetSum. Order matters.
 *
 * Complexity
 * - Worst-case O(sides^diceCount); pruning reduces branches when remainingSum is small.
 */
function findDiceSumCombinations(diceCount, sides, targetSum) {
  const results = [];

  // Helper: builds combinations in `path`.
  // Invariants:
  // - remainingDice + path.length === diceCount
  // - remainingSum is the sum we still need to hit from the remaining dice
  function backtrack(path, remainingDice, remainingSum) {
    // Base case: all dice chosen; accept only exact sum
    if (remainingDice === 0) {
      if (remainingSum === 0) {
        // console.log(`Found combination: ${path}`);
        results.push([...path]); // store a copy of the current solution
      }
      return; // cut branch (backtrack)
    }

    // Try all possible die faces for the next position
    for (let value = 1; value <= sides; value++) {
      // Prune: if this value already exceeds what's left, skip
      if (value > remainingSum) break;

      // Choose
      path.push(value);

      // console.log(`Explore ${path}: ${remainingDice - 1}, ${remainingSum - value}`, path);
      // Explore
      backtrack(path, remainingDice - 1, remainingSum - value);
      // Un-choose (the literal "backtrack")
      path.pop();
    }
  }

  backtrack([], diceCount, targetSum);
  return results;
}

// Demo
console.log(
  "Combinations for 3 dice, 6 sides, sum 10:"
  //  findDiceSumCombinations(3, 6, 10).length
);
console.log(findDiceSumCombinations(3, 6, 10));

module.exports = { findDiceSumCombinations };
