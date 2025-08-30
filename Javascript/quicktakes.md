# JavaScript quick takes (pre‑interview)

Short, memorable snippets for common patterns and gotchas.

---

## Count frequencies quickly

```js
const count = {};
for (const ch of s) {
    count[ch] = (count[ch] || 0) + 1;
}
// or
const count2 = [...s].reduce((m, ch) => (m[ch] = (m[ch] || 0) + 1, m), {});
```

## substring vs slice

```js
// substring swaps args if start > end, negative -> 0
s.substring(i, i + 3);

// slice keeps order and supports negatives
s.slice(i, i + 3);
s.slice(-3); // last 3 chars
```

## Numeric sort (don’t forget the comparator)

```js
nums.sort((a, b) => a - b);
```

## Unique values and dedupe

```js
const unique = [...new Set(arr)];
```

## Remove falsy values

```js
const compact = arr.filter(Boolean); // removes 0, '', null, undefined, NaN, false
```

## Map/filter/reduce one‑liners

```js
const squares = arr.map(x => x * x);
const evens = arr.filter(x => x % 2 === 0);
const sum = arr.reduce((a, x) => a + x, 0);
```

## Group by with reduce

```js
const byKey = arr.reduce((m, x) => {
    const k = keyFn(x);
    (m[k] ||= []).push(x);
    return m;
}, {});
```

## Object.entries / fromEntries

```js
const obj = { a: 1, b: 2 };
Object.keys(obj);   // ['a','b']
Object.values(obj); // [1,2]
Object.entries(obj);// [['a',1],['b',2]]
Object.fromEntries([['x', 3], ['y', 4]]); // { x:3, y:4 }
```

## Optional chaining and nullish coalescing

```js
const city = user?.address?.city; // safe access
const n = value ?? 0; // only uses 0 if value is null/undefined (not for '' or 0)
```

## Destructure with defaults and renaming

```js
const { a: alpha = 1, b = 2 } = obj;
const [first, , third = 0] = arr;
```

## Spread/rest; shallow copy

```js
const copy = { ...obj }; // shallow
const merged = { ...a, ...b };
const arrCopy = [...arr];
function f(x, ...rest) { /* rest is an array */ }
```

## Deep clone caveats

```js
// Simple but loses functions/Dates/Maps/Sets and breaks on cycles
const deep1 = JSON.parse(JSON.stringify(data));

// Better when available (Node 17+/modern browsers), handles cycles/most types
const deep2 = structuredClone(data);
```

## Map/Set quick ops

```js
const m = new Map([["a", 1]]);
m.set("b", 2); 
m.get("a");
 m.has("b");

const A = new Set([1,2,3]);
const B = new Set([2,3,4]);
const inter = new Set([...A].filter(x => B.has(x)));
const diff = new Set([...A].filter(x => !B.has(x)));
```

## Promises: all/settled/race/any

```js
await Promise.all([p1, p2]);         // fail fast
await Promise.allSettled([p1, p2]);  // always resolves
await Promise.race([p1, p2]);        // first to settle
await Promise.any([p1, p2]);         // first to fulfill (AggregateError if all reject)
```

## Timeout/Abort for fetch

```js
const c = new AbortController();
const t = setTimeout(() => c.abort(), 5000);
try {
    const res = await fetch(url, { signal: c.signal });
    const data = await res.json();
} finally {
    clearTimeout(t);
}
```

## Event loop: microtasks before macrotasks

```js
setTimeout(() => console.log('timeout'));
Promise.resolve().then(() => console.log('micro'));
// logs: 'micro' then 'timeout'
```

## Simple debounce and throttle

```js
const debounce = (fn, ms) => {
    let id; 
    return (...args) => { 
        clearTimeout(id); 
        id = setTimeout(() => fn(...args), ms); };
};
const throttle = (fn, ms) => {
    let last = 0;
     return (...args) => { const now = Date.now(); if (now - last >= ms) { last = now; fn(...args); } };
};
```

## Truthy/falsy gotcha vs ??

```js
// || uses truthiness; '' and 0 are treated as falsey
const label1 = maybeStr || 'default';
// ?? uses nullish; '' and 0 are kept
const label2 = maybeStr ?? 'default';
```

## In‑place array updates without mutation (copy first)

```js
const next = arr.map((x, i) => i === idx ? newVal : x);
const without = arr.filter((_, i) => i !== idx);
```

---

Use these as building blocks; combine them to solve most whiteboard tasks fast.


## Check anagram (two ways)

```js
// 1) Sort-based (O(n log n))
const isAnagramSort = (a, b) =>
    a.length === b.length && a.split('').sort().join('') === b.split('').sort().join('');

// 2) Count-based (O(n))
const isAnagramCount = (a, b) => {
    if (a.length !== b.length) return false;
    const freq = {};
    for (const ch of a) freq[ch] = (freq[ch] || 0) + 1;
    for (const ch of b) { if (!freq[ch]) return false; freq[ch]--; }
    return true;
};
```

## Check palindrome (two-pointer)

```js
const isPalindrome = (val) => {
    const s = String(val);
    for (let i = 0, j = s.length - 1; i < j; i++, j--) { // two inits, two updates
        if (s[i] !== s[j]) return false;
    }
    return true;
};
// Variant ignoring case/non-alphanumerics:
// const t = s.toLowerCase().replace(/[^a-z0-9]/g, ''); return t === [...t].reverse().join('');
```

## Add to beginning (Array.prototype.unshift)

```js
const fruits = ['Banana', 'Orange'];

// Add one (returns new length)
fruits.unshift('Apple');
// ['Apple', 'Banana', 'Orange']

// Add multiple
fruits.unshift('Kiwi', 'Mango');
// ['Kiwi', 'Mango', 'Apple', 'Banana', 'Orange']

// Note: unshift is O(n) (shifts existing elements). For immutable updates:
const next = ['Grape', ...fruits];
```



## Flatten an array

```js
// 1) Deep flatten (recursive)
const flattenDeep = (arr) =>
    arr.reduce((acc, v) => acc.concat(Array.isArray(v) ? flattenDeep(v) : v), []);

// 2) One-level (shallow) flatten
const flatten1 = (arr) => [].concat(...arr);

// 3) Built-in flat(depth)
const arr1 = [1, 2, [3, 4, [5, 6]]];
arr1.flat();          // [1, 2, 3, 4, [5, 6]]
arr1.flat(2);         // [1, 2, 3, 4, 5, 6]
arr1.flat(Infinity);  // deep flatten
```