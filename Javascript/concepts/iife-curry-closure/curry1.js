/**
 * Currying & Partial Application Study File
 * ----------------------------------------
 * DEFINITIONS (concise):
 *  - Currying: Transforming f(a,b,c) into f(a)(b)(c). One argument at a time.
 *  - Partial Application: Pre-filling some arguments of a function to get
 *    another function expecting the rest (e.g., pa = f.bind(null, a)).
 *
 * WHY CARE?
 *  - Reusability: specialize generic functions (e.g., makeTax(0.07)).
 *  - Composition: works nicely in functional pipelines.
 *  - Deferred evaluation: supply args as they become available.
 *  - Declarative readability in some domains (validation, math, predicates).
 *
 * COVERED SECTIONS:
 *  1. Uncurried vs simple manual curried
 *  2. Manual curry for 2+ args (classic add example)
 *  3. Generic fixed-arity curry utility
 *  4. Variadic "gather until no args" style (infinite currying)
 *  5. Partial application vs currying (contrast)
 *  6. Practical uses (config/logging/predicate factory)
 *  7. Composition example (pipe) with curried helpers
 *  8. Placeholders (simple lightweight implementation)
 *  9. Edge cases & gotchas
 * 10. Exercises
 */

console.log("================ 1. Uncurried vs Curried =================");
// Uncurried function needing all args now.
function volume(l, w, h) { return l * w * h; }
console.log("Uncurried volume(2,3,4):", volume(2, 3, 4));

// Curried style: one arg at a time (simplest explicit form).
function volumeCurried(l) {
	return function (w) {
		return function (h) {
			return l * w * h; // closure captures l, w
		};
	};
}
console.log("Curried volumeCurried(2)(3)(4):", volumeCurried(2)(3)(4));

console.log("\n================ 2. Manual Curry (add3) =================");
// Classic add3 example. Each function returns the next.
const add3 = a => b => c => a + b + c; // nested arrow functions
console.log("add3(1)(2)(3):", add3(1)(2)(3));

// You can partially apply along the way:
const add1 = add3(1);       // now a=1 stored in closure
const add1And2 = add1(2);   // now a=1, b=2 stored
console.log("add1And2(5):", add1And2(5)); // 1+2+5 = 8

console.log("\n================ 3. Generic Fixed-Arity curry(fn) =================");
// A utility that curries a function of known arity (length) until all args supplied.
function curry(fn, expected = fn.length) {
	return function curried(...args) {
		if (args.length >= expected) {
			return fn.apply(this, args); // enough args -> invoke
		}
		return (...rest) => curried(...args, ...rest); // collect more
	};
}

function sum4(a, b, c, d) { return a + b + c + d; }
const curriedSum4 = curry(sum4);
console.log("curriedSum4(1)(2)(3)(4):", curriedSum4(1)(2)(3)(4));
console.log("curriedSum4(1,2)(3,4):", curriedSum4(1, 2)(3, 4));
console.log("curriedSum4(1)(2,3,4):", curriedSum4(1)(2, 3, 4));

console.log("\n================ 4. Variadic / Infinite Currying =================");
// Allow chaining until an empty call or no argument call triggers evaluation.
function addChain(x) {
	let total = x;
	function next(n) {
		if (n === undefined) return total; // termination when called with no arg
		total += n;
		return next; // keep chaining
	}
	return next;
}
console.log("addChain(1)(2)(3)(4)():", addChain(1)(2)(3)(4)());

console.log("\n================ 5. Partial Application vs Currying =================");
// Partial application: pre-filling some arguments at once.
function multiply(a, b, c) { return a * b * c; }
const times2And3 = multiply.bind(null, 2, 3); // partially apply first two
console.log("times2And3(5):", times2And3(5)); // 2*3*5 = 30

// Curried version for comparison:
const multiplyCurried = a => b => c => a * b * c;
console.log("multiplyCurried(2)(3)(5):", multiplyCurried(2)(3)(5));

console.log("\n================ 6. Practical Uses =================");
// a) Prefixed logger
const makeLogger = tag => msg => console.log(`[${tag}]`, msg);
const errorLog = makeLogger("ERROR");
errorLog("Disk not found");

// b) Configured tax calculator
const makeTax = rate => amount => amount * (1 + rate);
const caTax = makeTax(0.0725);
console.log("CA tax on 100:", caTax(100));

// c) Predicate factory (e.g., min length validator)
const minLen = n => s => (s?.length ?? 0) >= n;
console.log("minLen(5)(\"hello\"):", minLen(5)("hello"));
console.log("minLen(5)(\"hi\"):", minLen(5)("hi"));

console.log("\n================ 7. Composition (pipe) =================");
// Simple left-to-right pipe using variadic reduce.
const pipe = (...fns) => x => fns.reduce((v, fn) => fn(v), x);
// Curried helpers
const add = a => b => a + b;
const multiplyBy = m => x => x * m;
const toStr = x => String(x);
const pipeline = pipe(add(2), multiplyBy(5), toStr);
console.log("pipeline(3):", pipeline(3)); // ((3+2)*5) -> 25 -> "25"

console.log("\n================ 8. Placeholders (Lightweight) =================");
// Support skipping positions using a placeholder symbol.
const _ = Symbol('placeholder');
function curryWithPlaceholders(fn, expected = fn.length) {
	function resolver(collected) {
		return function (...newArgs) {
			const merged = [];
			let newIdx = 0;
			for (const c of collected) {
				if (c === _) {
					if (newIdx < newArgs.length) {
						merged.push(newArgs[newIdx++]);
					} else {
						merged.push(_);
					}
				} else {
					merged.push(c);
				}
			}
			// Append remaining new args
			while (newIdx < newArgs.length) merged.push(newArgs[newIdx++]);
			const readyCount = merged.filter(a => a !== _).length;
			if (readyCount >= expected && !merged.includes(_)) {
				return fn(...merged);
			}
			return resolver(merged);
		};
	}
	return resolver(Array(expected).fill(_));
}

function quad(a, b, c, d) { return `${a}-${b}-${c}-${d}`; }
const cq = curryWithPlaceholders(quad);
const step1 = cq(_, 'B');
const step2 = step1('A', _);
const resultQuad = step2('C')('D');
console.log("Placeholder curry result:", resultQuad); // A-B-C-D

console.log("\n================ 9. Edge Cases & Gotchas =================");
console.log("Arity of curriedSum4:", curriedSum4.length, "(original length preserved in fn.length but wrapper hides internal chain)");
console.log("Functions with rest (...args) don't curry cleanly via fn.length based approach.");
console.log("Beware of over-currying: sometimes direct functions are clearer.");

console.log("\n================ 10. Exercises =================");
console.log("1. Implement a memoized curried factorial.");
console.log("2. Build a curryRight utility (apply args from right to left).");
console.log("3. Create a curried between(min)(max)(value) predicate.");
console.log("4. Implement a sum() infinite currying version that also supports valueOf/ toString coercion.");
console.log("5. Extend placeholder curry to allow providing more args than remaining slots.");

/*
CHEAT SHEET
-----------
Currying:
	f(a,b,c) -> f(a)(b)(c)

Partial Application:
	f(a,b,c) with some args preset -> g(b,c)

Generic curry algorithm (fixed arity):
	- Collect args until length >= fn.length
	- Then invoke

Infinite currying pattern:
	function f(x){ let total=x; return function g(y){ if(y===undefined) return total; total+=y; return g; }}

Use Cases:
	- Configuration (makeTax(rate)(amount))
	- Logging (makeLogger(tag)(msg))
	- Composition pipelines (pipe / compose)
	- Predicate building (minLen(n)(str))

Avoid Overuse:
	- When it obscures intent
	- When argument order isn't naturally progressive

Placeholders:
	Let you skip positions; fill later in any order.

Exercise Hint (between):
	const between = min => max => x => x >= min && x <= max;
*/

console.log("\n================ Done (Currying) =================");
