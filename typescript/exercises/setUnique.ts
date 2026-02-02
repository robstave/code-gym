/*
You are given an array of numbers.
Return a new array containing only the unique numbers, in the order of first appearance.

Requirements:

Do not modify the original array

Preserve original order

Handle empty input

Example:
input:  [1, 2, 2, 3, 1, 4]
output: [1, 2, 3, 4]
*/

 
export function uniqueNumbers(nums: number[]): number[] {

  let mySet: Set<number> = new Set<number>()
  let results: Array<number> = new Array<number>()

  for (let value of nums) {
    if (!mySet.has(value)) {
        mySet.add(value);
        results.push(value)
    }
  }
  return  results;
}

console.log("----------")
console.log(uniqueNumbers([5,1,2,2,3,1,4]));
console.log(uniqueNumbers([0,0,5,1,2,2,3,1,4]))
console.log(uniqueNumbers([]))
console.log(uniqueNumbers([7,11]))