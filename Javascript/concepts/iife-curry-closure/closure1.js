/**
 * Closures Study File
 * -------------------
 * Complements iife1.js (scope isolation) and curry1.js (function factories).
 * A closure = function + its lexical environment (captured variables).
 *
 * SECTIONS:
 *  1. Basic closure capture
 *  2. Function factory (makeAdder)
 *  3. Private state via closure (counter) vs object mutation
 *  4. Loop capture gotcha (var) & fixes (let / IIFE)
 *  5. Memoization with closure cache
 *  6. once(fn) utility (idempotent initializer)
 *  7. Encapsulating async state (rate limiter skeleton)
 *  8. Pitfalls: retaining more than needed / leaks
 *  9. this vs lexical variables
 * 10. Exercises
 */

console.log("================ 1. Basic Closure =================");
function outer(msg) {
  const stamp = Date.now(); // captured
  return function inner() { // inner closes over msg + stamp
    console.log(`Inner sees msg='${msg}' stamp=${stamp}`);
  };
}
const fnA = outer("hello");
setTimeout(fnA, 10); // still prints original values later

console.log("\n================ 2. Function Factory (makeAdder) =================");
function makeAdder(base) {
  return function add(n) { return base + n; }; // base captured
}
const add10 = makeAdder(10);
console.log("add10(5):", add10(5)); // 15

console.log("\n================ 3. Private State (Counter) =================");
function makeCounter() {
  let count = 0; // private
  return {
    inc(by = 1) { count += by; },
    dec(by = 1) { count -= by; },
    value() { return count; }
  };
}
const c = makeCounter();
c.inc(); c.inc(2); c.dec();
console.log("Counter value (should be 2):", c.value());

console.log("\n================ 4. Loop Capture Gotcha =================");
// Classic pitfall with var (function scoped)
var funcsBad = [];
for (var i = 0; i < 3; i++) {
  funcsBad.push(function() { console.log("bad i=", i); });
}
funcsBad[0](); // 3  (all 3)
funcsBad[1](); // 3

// Fix #1: let (block scoped)
var funcsGood = [];
for (let j = 0; j < 3; j++) {
  funcsGood.push(function() { console.log("good j=", j); });
}
funcsGood[0](); // 0
funcsGood[1](); // 1

// Fix #2: IIFE capture (pre-ES6)
var funcsIIFE = [];
for (var k = 0; k < 3; k++) {
  (function(capturedK) {
    funcsIIFE.push(function() { console.log("iife k=", capturedK); });
  })(k);
}
funcsIIFE[0](); // 0
funcsIIFE[1](); // 1

console.log("\n================ 5. Memoization =================");
function memoize(fn) {
  const cache = new Map(); // closed over
  return function(...args) {
    const key = JSON.stringify(args);
    if (cache.has(key)) return cache.get(key);
    const val = fn.apply(this, args);
    cache.set(key, val);
    return val;
  };
}
const slowFib = n => (n < 2 ? n : slowFib(n-1) + slowFib(n-2));
const fastFib = memoize(function fib(n){ return n < 2 ? n : fib(n-1) + fib(n-2); });
console.log("fastFib(10):", fastFib(10));

console.log("\n================ 6. once(fn) Utility =================");
function once(fn) {
  let called = false, result;
  return function(...args) {
    if (!called) {
      called = true;
      result = fn.apply(this, args);
    }
    return result;
  };
}
const init = once(() => { console.log("Initializing expensive resource"); return { ready: true }; });
init(); // runs
init(); // reuses

console.log("\n================ 7. Async State (Rate Limit Skeleton) =================");
function makeRateLimiter(limit, intervalMs) {
  let tokens = limit;
  let queue = [];
  setInterval(() => {
    tokens = limit;
    while (tokens > 0 && queue.length) {
      tokens--;
      queue.shift()();
    }
  }, intervalMs);
  return function schedule(fn) {
    if (tokens > 0) { tokens--; fn(); }
    else queue.push(fn);
  };
}
const limiter = makeRateLimiter(2, 100);
limiter(() => console.log("job1"));
limiter(() => console.log("job2"));
limiter(() => console.log("job3 delayed"));

console.log("\n================ 8. Pitfalls / Leaks =================");
// Capturing large objects keeps them alive.
function makeListener(bigObj) {
  return () => console.log(bigObj.id);
}
// If you add listeners to long-lived structures, ensure to remove them so closures can be GC'd.

console.log("\n================ 9. this vs Lexical Variables =================");
const obj = {
  val: 42,
  makeFn() {
    return () => this.val; // arrow closes over lexical this (obj)
  }
};
console.log("obj.makeFn()():", obj.makeFn()());

console.log("\n================ 10. Exercises =================");
console.log("1. Implement a makeOncePerInterval(fn, ms) wrapper.");
console.log("2. Build a memoizeWithLimit(fn, maxEntries).");
console.log("3. Create a sequence generator makeSeq(start=0, step=1).");
console.log("4. Implement an event emitter using closures for internal maps.");
console.log("5. Show a closure causing a leak; then refactor to break the reference.");

/*
CHEAT SHEET
-----------
Closure = Function + lexical environment.
Uses: factory functions, private state, callbacks, memoization, partial application, once, module emulation.
Gotchas: loop var capture with var, capturing too much state, accidental retention via listeners.
Not about: dynamic scope, runtime copying of variables (engine optimizes environments lazily).
*/

console.log("\n================ Done (Closures) =================");
