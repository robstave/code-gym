

Loop

## initialization

```typescript

// initialize an empty map of string to number in ts
let myMap: Map<string, number> = new Map<string, number>();

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



###
```typescript
export function uniqueNumbers(nums: number[]): number[] {


  for (let value of nums) {
    console.log(value)
  }
  return  nums;
}
```

