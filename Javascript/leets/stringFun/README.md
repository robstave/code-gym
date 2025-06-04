# String Manipulation Functions

Set of utility functions for manipulating strings and words.

## Functions

### sortWords(s: string)
Sorts words in a string alphabetically.
```javascript
sortWords("happy now brownish le cow") 
// Returns: "brownish cow happy le now"
```

### reverseWords(s: string)
Reverses the order of words in a string using two-pointer technique.
```javascript
reverseWords("happy now brownish le cow")
// Returns: "cow le brownish now happy"
```

### wordSize(s: string)
Sorts words first by length, then alphabetically for same-length words.
```javascript
wordSize("happy now brownish le cow")
// Returns: "le cow now happy brownish"
```

## Features
- Handles empty strings and edge cases
- Removes extra spaces
- Filters out empty words
- Case-sensitive sorting using localeCompare
- In-place reversal using two-pointer technique

## Implementation Details
- All functions use string trimming and splitting
- Input validation for null/undefined and non-string inputs
- Efficient array manipulation techniques
- See [sortWords.js](sortWords.js) for the implementation
