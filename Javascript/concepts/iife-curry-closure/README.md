## Scope Tools: IIFE, Closures, Currying

Three related concepts:
* IIFE: Immediately run a function to isolate scope / do setup.
* Closure: A function paired with the lexical variables it references.
* Currying: Using closures to feed arguments one at a time, producing new functions.

Use them together: an IIFE can set up a private environment, returning curried / closure‑based utilities.

---
## IIFE (Immediately Invoked Function Expression) Study Notes

An IIFE is a function expression you execute immediately after you create it. Historically it was the main pattern for isolating scope before ES6 (2015) gave us block scope (`let` / `const`) and native ES modules. Still useful today for quick, one‑off isolated execution or encapsulating temporary setup code in environments without modules.

---
### Core Idea
Wrap a function so JS parses it as an expression, then invoke it right away:
```js
(function () {
	// private work here
})();
```
Nothing inside leaks unless you explicitly return or assign it.

---
### Why Historically?
Before modules:
* `var` was function‑scoped only; all top-level `var` polluted `window` (browser) or the global object.
* Library authors wanted private helpers & minimal public surface.
* Minifiers benefited from local variable renaming inside the IIFE.

---
### Patterns Demonstrated (see `iife1.js`)
1. Plain function expression vs manual call (NOT an IIFE)
2. Basic IIFE (runs once immediately)
3. IIFE returning a computed value (encapsulation of temporaries)
4. Module / revealing pattern (private state, returned API)
5. IIFE with parameters (dependency injection / aliasing globals)
6. Named IIFE (better stack traces, recursion)
7. Arrow function IIFE variant
8. Modern perspective & contrast with block scope / modules

---
### Syntax Variants
```js
(function(){ /* code */ })();     // common
(function(){ /* code */ }());     // equivalent
(() => { /* code */ })();         // arrow IIFE
!function(){ /* code */ }();       // leading operator trick (legacy)
~function(){ /* code */ }();       // other unary operator forms
void function(){ /* code */ }();   // avoids value result
```
All force the function into “expression position” so the following `()` invokes it.

---
### When to Use Today
* Quick local sandbox in a non‑module script tag
* One‑time setup that produces a value or an API object
* Avoid name collisions when mixing multiple inline scripts
* Wrapping demo or scratch code without leaving debris in globals

### Prefer Alternatives When
| Need | Prefer |
|------|--------|
| File‑level encapsulation in modern build / ESM | Native ES modules (default) |
| Simple temporary variable isolation | Just use a block `{ ... }` |
| Reusable abstraction w/ state | Factory function or class |
| Async top-level logic in ESM | Use top-level `await` (in ESM) |

---
### Common Pitfalls
* Forgetting the final `()` – then you just defined a function expression but never ran it.
* Returning large mutable objects thinking they are still “private” – once returned, they are external.
* Overusing IIFEs inside already modular code (noise without benefit).
* Mixing arrow IIFEs where `this` binding matters (arrows have lexical `this`).

---
### Minimal Module Pattern Example
```js
const counter = (function () {
	let count = 0;
	return {
		inc(by = 1) { count += by; },
		value() { return count; }
	};
})();

counter.inc();
console.log(counter.value()); // 1
```

---
### Exercises
Try adding these to `iife1.js` (or a scratch file):
1. Memoizer: Build an IIFE that returns a `fib(n)` function with a private cache.
2. Feature detect: Use an IIFE to detect if `Promise` exists, then export a tiny shim if not.
3. Timer stats: An IIFE that runs, records start time, returns an API with `elapsed()`.
4. Config injector: Pass an object into an IIFE and expose only a validated subset.
5. Recursion practice: Named IIFE computing gcd(a, b).

---
### Quick Comparison: Block vs IIFE
```js
// Block scope (simpler when you just need temp vars)
{
	const temp = 42;
	console.log(temp);
}
// temp is gone here

// IIFE (if you want an expression result or older style)
const answer = (() => {
	const temp = 42;
	return temp * 2;
})();
console.log(answer); // 84
```

---
### Further Reading
* MDN: https://developer.mozilla.org/docs/Glossary/IIFE
* Jake Archibald on module patterns
* Exploring JS – historical patterns section

---
### TL;DR (IIFE)
Legacy-but-still-useful pattern for immediate execution + scope isolation. In modern ES module code, use sparingly for quick sandboxes, polyfills, or one-off initialization.

---
## Closures (see `closure1.js`)

Definition: A closure is a function plus the lexical environment (variable bindings) present when it was created. The function “remembers” those bindings even if invoked later elsewhere.

Core pattern:
```js
function outer(x) {
	return function inner(y) { return x + y; }; // inner closes over x
}
const add5 = outer(5);
console.log(add5(3)); // 8
```

Why it matters:
* Private state (counters, caches)
* Factories (configure then reuse)
* Memoization & performance
* Asynchronous callbacks that need stable references
* Implementing once, debounce, throttle utilities

Common pitfalls:
* Loop capture with `var` -> all callbacks share the final value.
* Retaining large objects unintentionally (memory leaks) via long-lived closures.
* Confusing closure (lexical) with dynamic `this` binding (separate concepts).

Loop fix example:
```js
for (let i = 0; i < 3; i++) {
	setTimeout(() => console.log(i), 0); // 0 1 2
}
```
or (pre-ES6):
```js
for (var i = 0; i < 3; i++) {
	(function(c){ setTimeout(() => console.log(c), 0); })(i);
}
```

When to refactor away:
* If a closure grows to manage complex mutable state -> consider a class or module.
* If memory profiling shows retained objects only due to unused closures -> break references.

Exercises (closure1.js lists more):
1. makeOncePerInterval(fn, ms)
2. memoizeWithLimit(fn, n)
3. makeSeq(start, step)
4. Event emitter
5. Leak demo + fix

---
## Currying & Partial Application (see `curry1.js`)

Definitions:
* Currying: f(a,b,c) -> f(a)(b)(c), each call supplies the next argument.
* Partial application: Pre-filling some args to get a new function expecting the rest.

Motivation:
* Specialize generic functions (makeTax(rate)(amount))
* Build reusable pipelines (pipe/compose)
* Defer providing values until later
* Improve readability in functional style code

Basic example:
```js
const add3 = a => b => c => a + b + c;
add3(1)(2)(3); // 6
```

Generic curry outline:
```js
function curry(fn, expected = fn.length) {
	return function curried(...args) {
		return args.length >= expected
			? fn(...args)
			: (...rest) => curried(...args, ...rest);
	};
}
```

Infinite currying pattern:
```js
const sum = a => b => (b === undefined ? a : sum(a + b));
sum(1)(2)(3)(); // 6 when final empty call used as terminator
```

Placeholders (advanced): allow skipping positions until later fill.

When NOT to curry:
* Argument order unclear / unnatural
* Overhead in hot paths (extra closures)
* Team unfamiliar -> readability drop

Compare:
```js
// Partial application
const mult = (a,b,c) => a*b*c;
const times6 = mult.bind(null, 2,3); // partial
times6(5); // 30
// Curried variant
const multC = a=>b=>c=>a*b*c;
multC(2)(3)(5); // 30
```

Composition synergy:
```js
const add = a=>b=>a+b;
const double = x=>x*2;
const pipe = (...fns) => x => fns.reduce((v,f)=>f(v), x);
const incAndDouble = pipe(add(1), double);
console.log(incAndDouble(3)); // 8
```

Exercises (see file for more):
1. curryRight(fn)
2. between(min)(max)(x)
3. sum infinite currying with valueOf override
4. Placeholder enhancements
5. Memoized curried factorial

---
## Putting It Together
Pattern progression:
1. IIFE isolates and returns a factory.
2. Factory uses closures for private configuration/state.
3. Returned functions are curried for composability.

Example fusion:
```js
const makeMath = (() => {
	const log = (...m) => console.log('[MATH]', ...m); // private helper
	const add = a => b => (log('add', a, b), a + b);
	const mul = a => b => a * b;
	return { add, mul };
})();
console.log(makeMath.add(2)(3)); // 5
```

---
## Glossary
| Term | Summary |
|------|---------|
| IIFE | Immediately executed function expression for scope isolation |
| Closure | Function plus captured lexical variables |
| Currying | Transform f(a,b,...) -> f(a)(b)... one argument per call |
| Partial Application | Pre-filling some arguments to produce a new function |
| Arity | Number of parameters a function expects (fn.length) |

---
## Next Ideas
* Add a debounce/throttle closure example
* Add a point-free style section for currying
* Show memory profiling tip for closures (Chrome dev tools)
* Explore async/await IIFE for top-level await simulation in CJS

