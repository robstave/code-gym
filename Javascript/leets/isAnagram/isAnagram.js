var isAnagram = function (s, t) {
  l = s.length;

  if (l != t.length) {
    return false;
  }

  check1 = t.split("").sort().join();
  check2 = s.split("").sort().join();
  return check1 == check2;
};

console.log(isAnagram("hello", "ollhe"));
