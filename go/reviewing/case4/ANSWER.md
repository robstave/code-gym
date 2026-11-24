# Answer: Completion rate

**Difficulty: Easy**

Both operands are integers, so `done / total` performs integer division before
conversion to `float64`. The sample prints `0.00`, not `0.75`.

Convert before dividing: `float64(done) / float64(total)`. If the function
later accepts zero totals, its API should define how to handle that separately.
