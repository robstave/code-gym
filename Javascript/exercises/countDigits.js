// Helper: check if single character is a digit 0-9
function isDigit(ch) {
  return ch >= '0' && ch <= '9';
}

// Count individual digit characters (0-9)
function countDigits(s) {
  let count = 0;
  for (let i = 0; i < s.length; i++) {
    if (isDigit(s[i])) {
      count++;
    }
  }
  return count;
}

// Count runs of consecutive digits (0-9) as numbers
// Examples: "0" => 1, "003" => 1, "9 8 7" => 3
function countDigits2(s) {
  let count = 0;
  let startDigit = false;
  for (let i = 0; i < s.length; i++) {
    if (isDigit(s[i])) {
      startDigit = true;
    } else {
      if (startDigit) {
        count++;
      }
      startDigit = false;
    }
  }

  if (startDigit) {
    count++;
  }
  return count;
}

// --- Test harness and grouped tests ---
function expectEqual(name, actual, expected) {
  const pass = actual === expected;
  console.log(`${pass ? 'PASS' : 'FAIL'} - ${name}: actual=${actual} expected=${expected}`);
}

console.log("\n=== countDigits (single characters) ===");
expectEqual("Digits: 23hh5h6", countDigits("23hh5h6"), 4);
expectEqual("Digits: mixed spaces", countDigits("23h3ki 47 9h5h6"), 8);
expectEqual("Digits: empty", countDigits(""), 0);
expectEqual("Digits: no digits", countDigits("abc def"), 0);
expectEqual("Digits: leading", countDigits("12abc"), 2);
expectEqual("Digits: trailing", countDigits("abc34"), 2);
expectEqual("Digits: zeros mixed", countDigits("10a0b05"), 5); // 1,0,0,0,5
expectEqual("Digits: leading zeros", countDigits("007bond"), 3); // 0,0,7
expectEqual("Digits: punctuation", countDigits("a-12.3"), 3); // 1,2,3
expectEqual("Digits: spaces", countDigits("9 8 7"), 3);
expectEqual("Digits: single zero", countDigits("0"), 1);
expectEqual("Digits: 101", countDigits("101"), 3);

console.log("\n=== countDigits2 (digit runs) ===");
expectEqual("Digits2: 23hh5h6", countDigits2("23hh5h6"), 3); // 23,5,6
expectEqual("Digits2: mixed spaces", countDigits2("23h3ki 47 9h5h6"), 6); // 23,3,47,9,5,6
expectEqual("Digits2: empty", countDigits2(""), 0);
expectEqual("Digits2: no digits", countDigits2("abc def"), 0);
expectEqual("Digits2: leading", countDigits2("12abc"), 1); // 12
expectEqual("Digits2: trailing", countDigits2("abc34"), 1); // 34
expectEqual("Digits2: zeros mixed", countDigits2("10a0b05"), 3); // 10,0,05
expectEqual("Digits2: leading zeros", countDigits2("007bond"), 1); // 007
expectEqual("Digits2: punctuation", countDigits2("a-12.3"), 2); // 12,3
expectEqual("Digits2: spaces", countDigits2("9 8 7"), 3); // 9,8,7
expectEqual("Digits2: single zero", countDigits2("0"), 1); // 0
expectEqual("Digits2: 101", countDigits2("101"), 1); // 101
