// ex2.cjs (CommonJS consumer)

const { add, subtract, multiply, divide } = require('./operations.cjs');

console.log('CJS add 10+5 =', add(10,5));
console.log('CJS subtract 10-5 =', subtract(10,5));
console.log('CJS multiply 10*5 =', multiply(10,5));
console.log('CJS divide 10/5 =', divide(10,5));
