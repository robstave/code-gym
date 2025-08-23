// calculator.mjs (ESM with a default export)

const calculator = {
  add(a, b) { return a + b; },
  subtract(a, b) { return a - b; },
  multiply(a, b) { return a * b; },
  divide(a, b) {
    if (b === 0) throw new Error('Division by zero');
    return a / b;
  }
};

export const PI = Math.PI;
export default calculator;
