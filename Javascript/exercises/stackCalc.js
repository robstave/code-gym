function doRervsePolish(exp) {
  
  let stack = [];

  let tokens = exp.split(" ");

  for (let token of tokens) {
    if (!isNaN(Number(token))) {
      console.log("push:" + token);
      stack.push(token);
      continue;
    }

    let b = Number(stack.pop());
    let a = Number(stack.pop());
    switch (token) {
      case "+":
        c = a + b;

        console.log(a + " + " + b + " = " + c);
        stack.push(c);
        break;
      case "*":
        c = a * b;
        console.log(a + " * " + b + " = " + c);

        stack.push(c);
        break;

      case "-":
        c = a - b;
        console.log(a + " - " + b + " = " + c);
        stack.push(c);
        break;

      default:
    }
  }

  return stack.pop();
}

console.log(doRervsePolish("2 1 + 3 * 2 -"));
