function fizz(num: number): string[] {
  let answer: Array<string> = [];

  for (let i = 1; i <= num; i++) {
    var item: string = "";
    switch (i % 15) {
      case 3:
      case 6:
      case 9:
      case 12:
        item = "Fizz";
        break;
      case 5:
      case 10:
        item = "Buzz";
        break;
      case 0:
        item = "FizzBuzz";
        break;
      default:
        item = item + i;
    }

    answer.push(item)
  }

  return answer;
}

console.log(fizz(15));
