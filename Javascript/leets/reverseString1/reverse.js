// @ts-check

function reverseString1(s) {
  let x = s.split("");
  let l = x.length;
  let result = [];
  for (let i = 0; i < l; i++) {
    result.push(x.pop());
  }
  return result.join("");
}

function reverseString1b(s) {
  let chars = s.split("");
  let results = [];
  for (let c of chars) {
    results.unshift(c);
  }
  return results.join("");
}

function reverseString2(s) {
  // Convert to array since strings are immutable
  let chars = s.split("");

  // In-place reversal using two pointers
  let left = 0;
  let right = chars.length - 1;

  while (left < right) {
    // Swap characters
    let temp = chars[left];
    chars[left] = chars[right];
    chars[right] = temp;

    left++;
    right--;
  }

  return chars.join("");
}

// Test cases

console.log(reverseString1("1234567"), "Expected: 7654321");
console.log(reverseString1b("1234567"), "Expected: 7654321");


console.log(reverseString2("1234567"), "Expected: 7654321");
console.log(reverseString2("hello"), "Expected: olleh");
console.log(reverseString2("a"), "Expected: a");
console.log(reverseString2(""), "Expected: ");
