# Answer: First character

**Difficulty: Medium**

String slicing uses byte offsets, not Unicode code points. `text[:1]` returns
only the first byte of the multibyte character `é`, producing invalid UTF-8.
The sample prints `Valid UTF-8: false`; it should print `true`.

For valid UTF-8 input, decode the first rune using
`utf8.DecodeRuneInString` and slice using the returned byte width, or range over
the string and return `string(r)` for the first rune. Define behavior for
empty or invalid UTF-8 inputs if those are allowed. A Unicode code point is not
necessarily a whole user-perceived character (grapheme cluster).
