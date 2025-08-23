


/**
 * IIFE (Immediately Invoked Function Expression) Study File
 * ---------------------------------------------------------
 * Shows common historical + still-useful patterns for isolating scope in JS.
 * Before ES6 modules & block scoping (let/const), IIFEs were the primary way
 * to avoid polluting the global namespace. They still help for: one‑time
 * setup, creating private state, and wrapping demo / scratch code.
 *
 * PATTERNS COVERED:
 *  1. Function expression (not an IIFE) vs direct invocation
 *  2. Basic IIFE
 *  3. IIFE returning a value
 *  4. Module / revealing pattern (private state)
 *  5. IIFE with parameters (dependency injection / aliasing)
 *  6. Named IIFE (better stack traces / recursion)
 *  7. Arrow function IIFE variant
 *  8. Modern perspective / when NOT to use
 */

console.log("================ 1. Function Expression =================");
// A function expression assigned to a variable (NOT immediately invoked)
// Contrast this with an IIFE below.
let f1 = function () {
    console.log("Normal function expression – invoked later");
};

f1(); // explicit call

console.log("\n================ 2. Basic IIFE =================");
// Wrap the function in parens so JS parses it as an expression, then invoke.
(function () {
    console.log("Basic IIFE runs immediately once");
})(); // Semicolon guards against ASI issues if concatenated with previous code.


 

console.log("\n================ 3. IIFE Returning a Value =================");
// Use an IIFE to compute a value while keeping intermediate vars private.
var result = (function () {
    var x = 10; // private
    var y = 20; // private
    return x + y; // only the sum escapes
})();

console.log("Result (should be 30):", result);

console.log("\n================ 4. Module Pattern / Private State =================");
// Classic: return an object exposing only what you want public.
// (Precursor to ES modules / classes for many simple cases.)
var counter = (function () {
    var count = 0; // private, not accessible directly outside
    function normalize(n) { return Number(n) || 0; } // also private helper
    return {
        increment(by = 1) { count += normalize(by); },
        decrement(by = 1) { count -= normalize(by); },
        getCount() { return count; }
    };
})();

counter.increment();
counter.increment(2);
counter.decrement();
console.log("Counter (should be 2):", counter.getCount());

console.log("\n================ 5. IIFE With Parameters =================");
// Pass in globals (dependency injection) so inside code can reference a local alias.
// Common in older libs to allow minification & avoid repeated global lookups.
(function (g, log) {
    const start = Date.now();
    log("Inside param IIFE. globalThis === g?", g === globalThis);
    log("Elapsed ms (quick demo):", Date.now() - start);
})(globalThis, console.log);

console.log("\n================ 6. Named IIFE =================");
// Naming the function improves stack traces & allows self-recursion.
// We now CAPTURE the returned value so you can see 120.
const factorialOf5 = (function factorialDemo(n) {
    if (n === 0) {
        // Base case: factorial(0) = 1
        return 1;
    }
    if (n === 5) console.log("Computing factorial(5) via named IIFE recursion...");
    return n * factorialDemo(n - 1);
})(5);
console.log("Factorial result (should be 120):", factorialOf5);

console.log("\n================ 7. Arrow Function IIFE =================");
// Arrow IIFE - concise but can be less readable for multi-line logic.
const doubled = (() => {
    const arr = [1, 2, 3];
    return arr.map(x => x * 2);
})();
console.log("Arrow IIFE result:", doubled);

console.log("\n================ 8. Modern Perspective =================");
// Today, prefer:
//  - Block scope: { let temp = ... }
//  - ES modules: each file has its own top-level scope
//  - Classes or factory functions for reusable abstractions
// IIFEs remain handy for quick one-off isolated execution or legacy scripts.

// QUICK CONTRAST: block scope alternative (no IIFE needed)
{
    let temp = 42;
    console.log("Block scope temp:", temp);
}
// console.log(temp); // Would throw ReferenceError – temp is block-scoped

console.log("\n================ Done =================");

/*
CHEAT SHEET
-----------
IIFE Syntax Variants:
    (function(){ /* code */ /* })();
    (function(){ /* code */ /* }());  // equivalent
    (() => { /* code */ /* })();      // arrow
    !function(){ /* code */ /* }();    // using a leading operator (historical)

Use Cases:
    - Isolate scope
    - One-time setup / precomputations
    - Avoid naming collisions in global scripts
    - Emulate module pattern before ES modules

Avoid When:
    - Writing modern module-based code (ESM) where file scope suffices
    - Overusing for trivial blocks where { } with let/const is clearer
*/
