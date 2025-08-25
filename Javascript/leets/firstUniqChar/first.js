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
