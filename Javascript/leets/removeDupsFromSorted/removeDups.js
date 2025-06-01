/**
 * Remove duplicates from a sorted array in-place and return the new length.
 * This function modifies the array in place to contain only unique elements
 * in their original order, and returns the new length.
 * 
 * Time Complexity: O(n) where n is the length of the array
 * Space Complexity: O(1) as we're modifying the array in-place
 * 
 * @param {number[]} list - A sorted array that may contain duplicates
 * @return {number} - Length of the array after removing duplicates
 */
function removeDups(list) {
    // Handle empty array case
    if (list.length === 0) return 0;
    
    // i maintains the position where the next unique element should be placed
    let i = 0;
    
    // Use the first element as our initial reference
    let last = list[i];
    
    // Traverse the array starting from the second element
    for (let j = 1; j < list.length; j++) {
        // If current element is different from the last unique element
        if (list[j] > last) {
            // Update our reference to the new unique element
            last = list[j];
            
            // Move i to the next position
            i++;
            
            // Place the unique element at position i
            list[i] = last;
        }
    }
    
    // The new length is i + 1 (since i is 0-indexed)
    const newLength = i + 1;
    
    console.log(`Array after removing duplicates: [${list.slice(0, newLength)}], New length: ${newLength}`);
    
    return newLength;
}

// Test cases
console.log("\nTest Case 1: Basic array with duplicates");
removeDups([0, 0, 0, 2, 3]);

console.log("\nTest Case 2: Array with multiple duplicates");
removeDups([0, 0, 1, 1, 2, 3, 4, 4, 5]);

console.log("\nTest Case 3: Array with no duplicates");
removeDups([1, 2, 3, 4, 5]);

console.log("\nTest Case 4: Empty array");
removeDups([]);

console.log("\nTest Case 5: Array with all duplicates");
removeDups([2, 2, 2, 2, 2]);

console.log("\nTest Case 6: Array with duplicates at the end");
removeDups([0, 1, 2, 3, 3, 3]);

console.log("\nTest Case 7: Larger numbers");
removeDups([100, 100, 101, 102, 102, 200, 200]);