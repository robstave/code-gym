# Answer: Rank scores

**Difficulty: Easy**

The function converts the scores to strings and then sorts them
lexicographically. As a result, `10` sorts before `2`. The sample prints
`[10 2 3]`, not the intended numeric order `[2 3 10]`.

Copy the integer slice, sort the copy with `sort.Ints`, and then convert the
sorted values to strings. Copying matters because the specification says not
to change the input.
