// @ts-check

/**
 * Sorts words in a string alphabetically
 * @param {string} s - Input string
 * @return {string} - Space-separated sorted words
 */
function sortWords(s) {
  if (!s || typeof s !== "string") return "";
  let words = s
    .trim()
    .split(" ")
    .filter((word) => word.length > 0);
  words.sort((a, b) => a.localeCompare(b));
  return words.join(" ");
}

/**
 * Reverses the order of words in a string
 * @param {string} s - Input string
 * @return {string} - Space-separated reversed words
 */
function reverseWords(s) {
  // edge cases
  if (!s || typeof s !== "string") return "";
  // split ( fancy)
  let words = s
    .trim()
    .split(" ")
    .filter((word) => word.length > 0);

  // two pointers
  let left = 0;
  let right = words.length - 1;

  while (left < right) {
    // or use a temp
    [words[left], words[right]] = [words[right], words[left]];
    left++;
    right--;
  }

  return words.join(" ");
}

/**
 * Sorts words by length, then alphabetically for same-length words
 * @param {string} s - Input string
 * @return {string} - Space-separated sorted words
 */
function wordSize(s) {
  if (!s || typeof s !== "string") return "";
  let words = s
    .trim()
    .split(" ")
    .filter((word) => word.length > 0);

  words.sort((a, b) => {
    let result = a.length - b.length;
    if (result === 0) {
      return a.localeCompare(b);
    }
    return result;
  });

  return words.join(" ");
}

function sortUnique(s) {
  if (!s || typeof s !== "string") return "";
  let words = s
    .trim()
    .split(" ")
    .filter((word) => word.length > 0);

  let aSet = new Set(words);
  let filtered = [...aSet];
  filtered.sort((a, b) => a.localeCompare(b));
  return filtered.join(" ");
}

// Test cases
function runTests() {
  console.log("Testing sortWords:");
  console.log(
    ":",
    sortWords("happy now brownish le cow"),
    "\n:",
    "brownish cow happy le now"
  );
  console.log(":", sortWords(""), "\n:", "");
  console.log(":", sortWords("  hello   world  "), "\n:", "hello world");

  console.log("\nTesting reverseWords:");
  console.log(
    ":",
    reverseWords("happy now brownish le cow"),
    "\n:",
    "cow le brownish now happy"
  );
  console.log(":", reverseWords("hello"), "\n:", "hello");
  console.log(":", reverseWords(""), "\n:", "");

  console.log("\nTesting wordSize:");
  console.log(
    ":",
    wordSize("happy now brownish le cow"),
    "\n:",
    "le cow now happy brownish"
  );
  console.log(":", wordSize("catty cat cats "), "\n:", "cat cats catty");
  console.log(":", wordSize(""), "\n:", "");

  console.log("\nTesting sortUnique:");
  console.log(
    ":",
    sortUnique("gamma gamma alpha beta alpha beta"),
    "\n:",
    "alpha beta gamma"
  );
}

runTests();
