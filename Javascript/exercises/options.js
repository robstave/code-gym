function options(options = {}) {
  const { align = "left", char = "*", name = "foo" } = options;

  console.log(options);
  console.log("align:", align);
  console.log("char:", char);
  console.log("name:", name);
}

console.log("--------");
options();
console.log("--------");

options({ char: "d" });
console.log("--------");
options({ char: "d", name: "duggg" });
