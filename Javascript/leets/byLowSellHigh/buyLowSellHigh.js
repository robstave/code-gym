// @ts-check	


var maxProfit = function(prices) {


    let profit = 0;
    let min = prices[0]

    for ( let i=1;i<prices.length;i++) {


        // Track minimum price seen so far
        if (prices[i] < min) {
            min = prices[i];
        }
        
        // Calculate maximum profit by comparing current profit with 
        // potential profit from buying at minimum price and selling at current price
        profit = Math.max(profit, prices[i] - min);
    }
    
    return profit;
};

// Test cases with expected results
console.log("Test Case 1:", maxProfit([5, 3, 1, 4, 6, 8, 3]), "Expected: 7");
console.log("Test Case 2:", maxProfit([3, 2, 1]), "Expected: 0");
console.log("Test Case 3:", maxProfit([7, 1, 5, 3, 6, 4]), "Expected: 5");
console.log("Test Case 4:", maxProfit([2, 4, 1]), "Expected: 2");
console.log("Test Case 5:", maxProfit([7, 6, 5, 4, 3]), "Expected: 0");
console.log("Test Case 6:", maxProfit([1, 2]), "Expected: 1");
console.log("Test Case 7:", maxProfit([1]), "Expected: 0");
console.log("Test Case 8:", maxProfit([]), "Expected: 0");
console.log("Test Case 9:", maxProfit([2, 1, 2, 1, 0, 1, 2]), "Expected: 2");
console.log("Test Case 10:", maxProfit([3, 3, 5, 0, 0, 3, 1, 4]), "Expected: 4");
