function twoSum(arr, target) {
  // Removed unused variables
  
 // arr[i]   = target - arr[j]
 if (arr.length == 0) {
    return null
 }

  const mapping = {};

  for ( let i=0;i<arr.length;i++){
    mapping[arr[i]] = i
  }

  for ( let j=0;j<arr.length;j++){
    check = target - arr[j]

    if ( check in mapping && mapping[check] != j){
        return [j,mapping[check]]
    }
  }

  return null // no solution
}

 // 
console.log(twoSum([2, 7, 11, 15], 18));
console.log(twoSum([1, 4, 5, 7, 18], 12));
console.log(twoSum([3, 2, 4], 6));
console.log(twoSum([3, 3], 6));
console.log("--");


