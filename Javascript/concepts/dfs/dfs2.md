# DFS 2 — Graph (Mermaid)

The graph used in `dfs2.js` as an adjacency list:
- A → B, C
- B → D, E
- C → F
- D → C
- E → F
- F → B (creates a cycle)

```mermaid
graph TD
  A[A] --> B[B]
  A --> C[C]
  B --> D[D]
  B --> E[E]
  C --> F[F]
  D --> C
  E --> F
  F --> B
```

Notes
- There is a cycle: A → B → D → C → F → B.
- DFS with a `visited` set prevents infinite loops on such cycles.
