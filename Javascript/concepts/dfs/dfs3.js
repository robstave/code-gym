// DFS 3: Backtracking (subset/power-set generation)
// Idea: at each index, choose to include or exclude an element.

function powerSet(arr) {
  const result = [];
  const curr = [];

  function dfs(i) {
    if (i === arr.length) {
      result.push(curr.slice());
      return;
    }
    // Exclude arr[i]
    dfs(i + 1);
    // Include arr[i]
    curr.push(arr[i]);
    dfs(i + 1);
    curr.pop(); // backtrack
  }

  dfs(0);
  return result;
}

const sets = powerSet(['A', 'B', 'C']);
console.log('Power set size:', sets.length);
console.log('Some subsets:', sets.map(s => '[' + s.join('') + ']').join(' '));
// Expected size for 3 items: 2^3 = 8
