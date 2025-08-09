// You are tasked with safeguarding a web service or an
// API endpoint from potential abuse, such as denial-of-service
// (DOS) attacks, by implementing a rate limiter. Build a function called
//  <code>isAllowed</code>

//This function should deny any request from a
//  unique client that exceeds the threshold of 100
//  requests within the last second.

// tests:  99 requests from clientA in a second is pass
// 101 from ClientB is a fail
// 50 requests followed by 80 a second later is a pass

class RateLimiter {
  // Write your code here
  constructor() {
    this.rateMap = new Map();
  }

  isAllowed(key) {
    var now = new Date();

    if (this.rateMap.has(key)) {
      let rate = this.rateMap.get(key);

      if (now.getTime() - 1000 > rate.ts) {
        console.log("reset");
        rate.count = 1;
        rate.ts = now;
        return true;
      } else {
        if (rate.count >= 100) {
          console.log("bzzzz");
          return false;
        }
      }
      //console.log(rate.count);

      rate.count++;
    } else {
      let rateValue = { ts: now, count: 1 };
      this.rateMap.set(key, rateValue);
    }

    return true;
  }
}

//test 1.
// 99 problems...its a pass
const rl1 = new RateLimiter();

let allowed = true;

for (let i = 0; i < 100; i++) {
  if (!rl1.isAllowed("ClientA")) {
    allowed = false;
  }
}
console.log("Test1: is true:", allowed);

//test 2.
// 99 problems...its a pass
const rl2 = new RateLimiter();

allowed = true;

for (let i = 0; i < 101; i++) {
  if (!rl2.isAllowed("ClientB")) {
    allowed = false;
  }
}
console.log("Test2: is false:", allowed);

//test 2.
// 99 problems...its a pass
const rl3 = new RateLimiter();

allowed = true;

for (let i = 0; i < 101; i++) {
  if (!rl2.isAllowed("ClientC")) {
    allowed = false;
  }
}
console.log("Test3a: is false:", allowed);

setTimeout(() => {
  console.log("Test3b: is true:", rl2.isAllowed("ClientC"));
}, 1100);
