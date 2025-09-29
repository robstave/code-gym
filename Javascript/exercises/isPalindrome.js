

let isPalindrome = function(x){

    let str = x.toString();

    arr = str.split("")

    for ( let i = 0; i < arr.length; i++){
        if ( arr[i] !== arr[arr.length-1-i]){
            return false
        }
    }

    return true
    
}

console.log(isPalindrome(2112))

