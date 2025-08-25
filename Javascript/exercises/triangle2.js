// lets add an optional char
var triangle2b = function (size, ch, left) {
  //kinda dealt with in the loops.  Maybe add an upperbound
  if (size < 1) {
    return;
  }
  let c = "*";

  if (ch != null) {
    c = ch;
  }

  let isLeft = false;
  if (left == true) {
     isLeft = true;
  }

  for (i = 0; i < size; i++) {
    let line = c.repeat(i + 1);

    if (isLeft) {
      line = " ".repeat(size - i - 1) + line;
    }

    console.log(line);
  }
};

triangle2b(5);

triangle2b(5, "4", true);
