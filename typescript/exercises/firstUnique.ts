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

export function firstUnique(s: string): number {
  let first = -1;

  let hash: Map<string, number> = new Map<string, number>();

  for (let value of s) {
    if (!hash.has(value)) {
      hash.set(value, 1);
    } else {
      let count = hash.get(value);
      count++;
      hash.set(value, count);
    }
  }

  console.log(hash);

  for (let index = 0; index < s.length; index++) {
    if (hash.get(s[index]) == 1) {
      return index;
    }
  }

  return -1;
}

console.log(firstUnique("concussion"));
