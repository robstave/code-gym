/**
 * Demonstrates default parameters and object destructuring with defaults.
 * 
 * Key concepts:
 * 1. Default parameter: options = {} prevents undefined if called with no args
 * 2. Destructuring with defaults: { align = "left", char = "*", name = "foo" }
 *    extracts properties from the options object, using defaults if properties are missing
 * 
 * @param {Object} options - Configuration object (optional)
 * @param {string} options.align - Alignment (default: "left")
 * @param {string} options.char - Character to use (default: "*")
 * @param {string} options.name - Name value (default: "foo")
 */
function options(options = {}) {
  // Destructure the options object, applying defaults for any missing properties
  const { align = "left", char = "*", name = "foo" } = options;

  console.log(options);
  console.log("align:", align);
  console.log("char:", char);
  console.log("name:", name);
}

console.log("--------");
// Called with no arguments → options = {}, all values use defaults
options();
console.log("--------");

// Called with partial options → char="d", align and name use defaults
options({ char: "d" });
console.log("--------");

// Called with multiple options → char="d", name="duggg", align uses default
options({ char: "d", name: "duggg" });

console.log("\n======== options2 (direct destructuring in parameters) ========\n");

/**
 * Equivalent to options() but using direct destructuring in the function parameter.
 * This is a more modern, concise pattern that combines both techniques in one line.
 * 
 * Key difference:
 * - options():  function options(options = {}) { const { align = "left", ... } = options; }
 * - options2(): function options2({ align = "left", ... } = {}) { ... }
 * 
 * Both achieve the same result, but options2 is more compact.
 * 
 * @param {Object} params - Destructured directly in the parameter list
 * @param {string} params.align - Alignment (default: "left")
 * @param {string} params.char - Character to use (default: "*")
 * @param {string} params.name - Name value (default: "foo")
 */
function options2({ align = "left", char = "*", name = "foo" } = {}) {
  // No need for separate destructuring — parameters are already extracted
  console.log("align:", align);
  console.log("char:", char);
  console.log("name:", name);
}

console.log("--------");
// Called with no arguments → empty object {}, all values use defaults
options2();
console.log("--------");

// Called with partial options → char="d", align and name use defaults
options2({ char: "d" });
console.log("--------");

// Called with multiple options → char="d", name="duggg", align uses default
options2({ char: "d", name: "duggg" });
