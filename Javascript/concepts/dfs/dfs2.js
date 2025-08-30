// DFS 2: DFS on a general graph (with a cycle) using a visited set
// Idea: mark visited to avoid infinite loops when cycles exist

// Graph as adjacency list
const graph = {
  A: ['B', 'C'],
  B: ['D', 'E'],
  C: ['F'],
  D: ['C'], // creates a cycle D -> C -> F (end) but also C could point back in other graphs
  E: ['F'],
  F: ['B'], // back edge to B makes a cycle A->B->D->C->F->B
};

// ascii version of the graph
/*
   A
  / \
 B   C
/ \   \
D   E   F
 \ /
  C
*/

function dfsGraph(start) {
  const visited = new Set();
  const order = [];

  function visit(node) {
    if (visited.has(node)) return;
    visited.add(node);
    order.push(node);
    for (const nei of graph[node] || []) {
      visit(nei);
    }
  }

  visit(start);
  return order;
}

console.log('DFS graph from A:', dfsGraph('A').join(' -> '));
// Expected order is DFS-dependent but should include each node once, e.g.:
// A -> B -> D -> C -> F -> E
