// use-calculator.mjs (ES module consumer)

import calc, { PI } from './calculator.mjs';

console.log('calc.add(2,3)=', calc.add(2,3));
console.log('calc.multiply(4,5)=', calc.multiply(4,5));
console.log('PI ~', PI);
