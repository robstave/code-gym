# Backtracking examples

This folder contains small, runnable examples that show backtracking vs. brute force.

## Files

- `dice1.js` — Backtracking to find all ordered dice rolls that sum to a target (choose → explore → un-choose). Includes pruning.
- `dicebrute.js` — Brute-force 3-nested-loops version for exactly 3 dice.
- `subsets_backtrack.js` — Backtracking to generate all subsets (power set).
- `subsets_bruteforce.js` — Bitmask/iterative brute-force for the same power set.

## Run

From repo root or this directory:

```powershell
# Dice sum backtracking
node .\Javascript\concepts\backtrack\dice1.js

# Dice sum 3-dice brute-force
node .\Javascript\concepts\backtrack\dicebrute.js

# Subsets (backtracking)
node .\Javascript\concepts\backtrack\subsets_backtrack.js

# Subsets (bitmask brute-force)
node .\Javascript\concepts\backtrack\subsets_bruteforce.js
```

## Notes

- Backtracking pattern: push (choose) → recurse (explore) → pop (un-choose). Add pruning to cut impossible branches early.
- Brute force enumerates everything without recursion or pruning. It’s simpler but can be slower for large inputs.
