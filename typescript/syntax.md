

 
## initialization

```typescript

// initialize an empty map of string to number in ts
let myMap: Map<string, number> = new Map<string, number>();
//init the map as a string key with set of numbers as value
let mapSet: Map<string, Set<number>> = new Map<string, Set<number>>();

// initialize an empty array of strings in ts
let myArray: string[] = [];
// initialize an empty set of numbers in ts
let numbers: Set<number> = new Set<number>();


```

## types

### Sets

```typescript
// loop through a set of numbers
let numSet: Set<number> = new Set([1, 2, 3, 4, 5]); 
for (let num of numSet) {
    console.log(num);
}

// loop using forEach
numSet.forEach((num) => {
    console.log(num);
});

// loop using values()
for (let num of numSet.values()) {
    console.log(num);
}   

// set to array
let numArray: number[] = Array.from(numSet);

// set to map (value as key, index as value)
let numMap: Map<number, number> = new Map();    
let index = 0;
for (let num of numSet) {
    numMap.set(num, index++);
}
// map to set (keys as values)
let keysSet: Set<number> = new Set(numMap.keys());

```

### Maps

```typescript
// loop through a map of string to number
let myMap: Map<string, number> = new Map([
    ["a", 1],
    ["b", 2],
    ["c", 3]
]);
for (let [key, value] of myMap) {
    console.log(`${key}: ${value}`);
}
// loop using forEach
myMap.forEach((value, key) => {
    console.log(`${key}: ${value}`);
});

// loop using entries()
for (let [key, value] of myMap.entries()) {
    console.log(`${key}: ${value}`);
}   
```


###
```typescript
export function uniqueNumbers(nums: number[]): number[] {


  for (let value of nums) {
    console.log(value)
  }
  return  nums;
}
```


## Sort

```typescript
items.sort((a, b) => {
  if (a.score !== b.score) return b.score - a.score;
  return a.name.localeCompare(b.name);
});

```
Hidden gotcha:

localeCompare vs <


localeCompare is used for string comparison to handle different character sets and locales correctly. Using `<` or `>` for string comparison may lead to unexpected results, especially with non-ASCII characters.

For example:

```typescript   
const names = ["Zoe", "Ana", "Émile", "Éric"];
names.sort((a, b) => a.localeCompare(b));
console.log(names); // Correctly sorted: ["Ana", "Émile", "Éric

", "Zoe"]
```

## STrings

### Substrings

```typescript

let str: string = "Hello, TypeScript!";
let substr1: string = str.substring(0, 5); // "Hello"
let substr2: string = str.slice(7, 17); // "TypeScript"
let substr3: string = str.substr(7, 10); // "TypeScript"
