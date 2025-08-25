// simple triangle

var triangle = function (size) {


  for (i = 0; i < size; i++) {
    let c = "*";
    for (j = 0; j < i; j++) {
      c = c + "*";
    }
    console.log(c);
  }
};

// bit slicker...some checking
var triangleb = function (size) {
  //kinda dealt with in the loops.  Maybe add an upperbound
  if (size < 1) {
    return;
  }
  for (i = 0; i < size; i++) {
    let line = "*".repeat(i + 1);

    console.log(line);
  }
};


// lets add an optional char
var trianglec = function (size, ch) {
  //kinda dealt with in the loops.  Maybe add an upperbound
  if (size < 1) {
    return;
  }
  let c = "*"

  if (ch != null) {
    c = ch
  }
  for (i = 0; i < size; i++) {
    let line = c.repeat(i + 1);

    console.log(line);
  }
};


console.log("== triange");
triangle(5);
console.log("== triangeb");
triangleb(5);
console.log("== triangec default");
trianglec(5);
console.log("== triangec");
trianglec(5,"#");

