// @ts-check	

/**
 * Checks if a number is a palindrome by reversing it and comparing with the original.
 * A palindrome reads the same forward and backward (e.g., 121, 1221, 12321).
 * 
 * This implementation uses a stack to store each digit and then reconstructs
 * the reversed number.
 * 
 * @param {number} val - The number to check
 * @return {boolean} - True if the number is a palindrome, false otherwise
 */
function isPalindrome(val) {
    let stack = []
    
    // Store original value for comparison
    let value = val
    
    // Extract each digit and push to stack
    while (value > 0) {
        let digit = value % 10
        stack.push(digit)
        value = Math.floor(value/10)
    }
    
    // Reconstruct the reversed number
    let rev = 0;
    for (let i = 0; i < stack.length; i++) {
        rev = rev * 10 + stack[i]
    }
 
    // Compare the reversed number with the original
    return rev === val;
}

/**
 * Checks if a number is a palindrome using string conversion.
 * This is more efficient as it avoids stack operations and arithmetic.
 * 
 * @param {number} val - The number to check
 * @return {boolean} - True if the number is a palindrome, false otherwise
 */
function isPalindrome2(val) {
    // Convert to string
    const str = val.toString();
    
    // Compare characters from both ends moving inward
    for (let i = 0, j = str.length - 1; i < j; i++, j--) {
        if (str[i] !== str[j]) {
            return false;
        }
    }
    
    return true;
}

/**
 * Another efficient palindrome checker that doesn't convert to string.
 * Instead, it directly reverses the entire number and compares.
 * 
 * @param {number} val - The number to check
 * @return {boolean} - True if the number is a palindrome, false otherwise
 */
function isPalindrome3(val) {
    // Handle edge cases: negative numbers are not palindromes
    if (val < 0) return false;
    
    // Handle single digit numbers (always palindromes)
    if (val >= 0 && val < 10) return true;
    
    // Handle numbers that end with 0 (not palindromes unless the number is 0)
    if (val % 10 === 0 && val !== 0) return false;
    
    let originalVal = val;
    let reversed = 0;
    
    // Reverse the entire number
    while (val > 0) {
        reversed = reversed * 10 + (val % 10);
        val = Math.floor(val / 10);
    }
    
    return reversed === originalVal;
}

// Test cases
console.log(isPalindrome(1234321), "true");
console.log(isPalindrome(123), "false");
console.log(isPalindrome(1221), "true");

console.log("Testing isPalindrome2:");
console.log(isPalindrome2(1234321), "true");
console.log(isPalindrome2(123), "false");
console.log(isPalindrome2(1221), "true");

console.log("Testing isPalindrome3:");
console.log(isPalindrome3(1234321), "true");
console.log(isPalindrome3(123), "false");
console.log(isPalindrome3(1221), "true");