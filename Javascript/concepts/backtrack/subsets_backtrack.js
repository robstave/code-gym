// Backtracking: Generate all subsets (power set)
// At each index, choose to exclude or include the element, then recurse.

function powerSetBacktrack(arr) {
  const result = [];
  const path = [];

  function dfs(i) {
    if (i === arr.length) {
      result.push([...path]);
      return; // backtrack
    }
    // Exclude arr[i]
    dfs(i + 1);
    // Include arr[i]
    path.push(arr[i]);
    dfs(i + 1);
    path.pop(); // un-choose (backtrack)
  }

  dfs(0);
  return result;
}

// Demo
const demo = ['A', 'B', 'C'];
const sets = powerSetBacktrack(demo);
console.log('Backtracking power set size (expected 8):', sets.length);
console.log('Subsets:', sets.map(s => '[' + s.join('') + ']').join(' '));

module.exports = { powerSetBacktrack };
