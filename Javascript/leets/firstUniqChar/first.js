/* LeetCode: First Unique Character in a String
 *
 * Problem:
 * Given a string s, find the first non-repeating character in it and return its index.
 * If it does not exist, return -1.
 *
 * Constraints:
 * - 1 <= s.length <= 10^5
 * - s consists of only lowercase English letters.
 *
 * Example 1:
 * Input: s = "leetcode"
 * Output: 0
 *
 * Example 2:
 * Input: s = "concussion"
 * Output: 4
 */
 
let firstUnique = function (word) {
  let index = 0;

  charMap = new Map();

  chars = word.split("");

  for (let c of chars) {
    if (charMap.get(c) != null) {
      charMap.set(c, charMap.get(c) + 1);
    } else {
      charMap.set(c, 1);
    }
  }

  for (let i = 0; i < chars.length; i++) {
    let count = charMap.get(chars[i]);
    if (count == 1) {
      return i;
    }
  }

  return index;
};

console.log(firstUnique("concussion"));
