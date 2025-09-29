var reverseStack = function (s) {
  vals = s.split("");
  resultArray = [];

  while (vals.length > 0) {
    resultArray.push(vals.pop());
  }
  return resultArray.join("");
};

var reverseUnshift = function (s) {
  vals = s.split("");
  resultArray = [];

  for (c of vals) {
    resultArray.unshift(c);
  }
  return resultArray.join("");
};

var reverseTwoPointer = function (s) {
  vals = s.split("");
  let l = vals.length;

  let start = 0;
  let end = l - 1;

  while (start <= end) {
    const temp = vals[start];
    vals[start] = vals[end];
    vals[end] = temp;

    start++;
    end--;
  }
  return vals.join("");
};

var reverseArrayInPlace = function (arr) {
  let len = arr.length;
  if (len < 2) {
    return arr;
  }

  for (let i = 0; i < Math.floor((len - 1) / 2); i++) {
    temp = arr[i];
    arr[i] = arr[len - 1 - i];
    arr[len - 1 - i] = temp;
  }

  return arr;
};

function reverseRecursive(str) {
  if (str === "") {
    return "";
  } else {
    return reverseRecursive(str.substr(1)) + str.charAt(0);
  }
}

console.log("-----reverseStack--------");
console.log(reverseStack("1234567"));
console.log(reverseStack("hello"));
console.log("\n");

console.log("-----reverseUnshift------");
console.log(reverseUnshift("1234567"));
console.log(reverseUnshift("hello"));
console.log("\n");

console.log("-----reverseTwoPointer------");
console.log(reverseTwoPointer("1234567"));
console.log(reverseTwoPointer("hello"));
console.log("\n");

console.log("-----reverseArrayInPlace-----");
console.log(reverseArrayInPlace([1, 2, 3, 4, 5]));
console.log(reverseArrayInPlace(["h", "e", "l", "l", "o"]));
console.log("\n");

console.log("-----reverseRecursive------");
console.log(reverseRecursive("1234567"));
console.log(reverseRecursive("hello"));
console.log("\n");
