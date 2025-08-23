# JavaScript Modules: ESM and CommonJS

This folder shows both modern ES Modules (ESM) and legacy/CommonJS (CJS) patterns, with runnable examples.

Contents:
- operations.js (ESM named exports)
- ex1.js (ESM consumer of operations.js)
- calculator.mjs (ESM default + named export)
- use-calculator.mjs (ESM consumer of calculator.mjs)
- operations.cjs (CJS module.exports)
- ex2.cjs (CJS consumer using require)

## How to run (Windows PowerShell)

ESM example using .js with explicit imports:
```powershell
node ex1.js
```

ESM example using .mjs (Node treats .mjs as ESM by default):
```powershell
node use-calculator.mjs
```

CommonJS example using require/module.exports:
```powershell
node ex2.cjs
```

Notes:
- When using .js files as ESM, Node infers module type via nearest package.json "type": "module" or by using .mjs. Here, ex1.js imports another .js file and is run directly; if your global/default is CommonJS, prefer .mjs for clarity.
- In browsers, use `<script type="module">` and relative URLs; Node-specific `require()` won’t work there.

## ESM Cheatsheet
```js
// Named exports
export const PI = Math.PI;
export function add(a,b){ return a+b; }

// Default export
export default { add };

// Importing
import calculator, { PI as circlePI } from './calculator.mjs';
```

## CJS Cheatsheet
```js
// exports
module.exports = { add, subtract };

// import
const { add } = require('./operations.cjs');
```

## Dynamic Import (ESM)
```js
const math = await import('./operations.js');
console.log(math.add(1,2));
```

## When to use which
- Use ESM for modern Node and browsers; it’s the standard going forward.
- Use CJS for older Node codebases or when a dependency only supports require/module.exports.

