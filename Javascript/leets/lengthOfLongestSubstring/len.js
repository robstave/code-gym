/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
  let l = s.length;
  let max = 0;

  let start = 0;
  let end = 0;

  while (end < l ) {
    ss = s.substring(start, end+1);
  console.log(start, end, ss);
    if (hasDup(ss)) {
      start++;
      end=start+max;
      continue
    }

    max = Math.max(max, ss.length);
 
    end++;
  }

  return max;
};

var hasDup = function (s) {
  let chars = s.split("");
  let l = chars.length;
  let m = new Set();

  for (let i = 0; i < l; i++) {
    if (m.has(s[i])) {
      return true;
    }
    m.add(s[i]);
  }

  return false;
};


let l = lengthOfLongestSubstring("pwwkew");
if (l != 3) {
    console.log("fail",l);
} else {
    console.log("pass");
}

l =lengthOfLongestSubstring("dvdh");
if (l != 3) {
    console.log("fail",l);
} else {
    console.log("pass");
}

l = lengthOfLongestSubstring("au");
if (l != 2) {
    console.log("fail",l,l);
} else {
    console.log("pass");

}

l = lengthOfLongestSubstring("qwertyuasdfguzxcvbnm7463");
if (l != 17) {
    console.log("fail",l,l);
} else {
    console.log("pass");

}
//""
