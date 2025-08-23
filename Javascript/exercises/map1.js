var topKFrequent = function (fruits, k) {
  let myMap = new Map();
  for (let fruit of fruits) {
    if (myMap.has(fruit)) {
      myMap.set(fruit, myMap.get(fruit) + 1);
    } else {
      myMap.set(fruit, 1);
    }
  }

  console.log("----");

  for (let [fruit, count] of myMap.entries()) {
    console.log(`${fruit}: ${count}`);
  }

  // we can convert the map to an array of entries and sort it
  let sortedEntries = Array.from(myMap.entries()).sort((a, b) => b[1] - a[1]);

  console.log("--sorted--");

  for (let [fruit, count] of sortedEntries.entries()) {
    console.log(`${fruit}: ${count}`);
  }

  console.log("----");
  // then we can take the top k entries
  let topK = sortedEntries.slice(0, k);

  result = [];

  for (let f of topK) {
    result.push(f[0]);
  }

  return result;
};

let fruits = [
  "apple",
  "banana",
  "orange",
  "kiwi",
  "grape",
  "pear",
  "apple",
  "apple",
  "apple",
  "banana",
  "orange",
  "banana",
  "grape",
  "grape",
  "grape",
  "pear",

  "pear",
  "pear",
  "pear",
  "pear",
];

console.log(topKFrequent(fruits, 4));
