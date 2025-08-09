/*
Example 1:

Input: n = 3
Output: ["1","2","Fizz"]
Example 2:

Input: n = 5
Output: ["1","2","Fizz","4","Buzz"]
Example 3:

Input: n = 15
Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
*/

function fizzBuzz(count) {
  let result = [];

  for (let i = 1; i <= count; i++) {
    let ans = "";

    if (i%15==0) {
      ans = "FizzBuzz";
    } else if (i%5==0) {
      ans = "Buzz";
    } else if (i%3==0) {
      ans = "Fizz";
    } else {
      ans = "" + i;
    }
    result.push(ans);
  }

  return result;
}

console.log("test")

 
console.log(fizzBuzz(33));