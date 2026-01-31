// You are tasked with safeguarding a web service or an
// API endpoint from potential abuse, such as denial-of-service
// (DOS) attacks, by implementing a rate limiter. Build a function called
//  <code>isAllowed</code>

//  This function should deny any request from a
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

  /**
   * Flush (remove) rate limit data for a specific client key.
   * Useful for resetting a client's rate limit manually.
   * @param {string} key - The client key to flush
   * @returns {boolean} - True if key existed and was removed, false otherwise
   */
  flush(key) {
    if (this.rateMap.has(key)) {
      this.rateMap.delete(key);
      console.log(`Flushed rate limit for: ${key}`);
      return true;
    }
    console.log(`No rate limit found for: ${key}`);
    return false;
  }

  /**
   * Flush all rate limit data for all clients.
   * Useful for clearing the entire rate limiter state.
   * @returns {number} - Number of clients that were flushed
   */
  flushAll() {
    const count = this.rateMap.size;
    this.rateMap.clear();
    console.log(`Flushed all rate limits (${count} clients cleared)`);
    return count;
  }

  /**
   * Flush stale entries older than a given age (in milliseconds).
   * Useful for periodic cleanup to prevent memory buildup.
   * @param {number} maxAge - Age in milliseconds (default: 60000 = 1 minute)
   * @returns {number} - Number of stale entries removed
   */
  flushStale(maxAge = 60000) {
    const now = new Date().getTime();
    let flushedCount = 0;

    for (const [key, rate] of this.rateMap.entries()) {
      if (now - rate.ts.getTime() > maxAge) {
        this.rateMap.delete(key);
        flushedCount++;
      }
    }

    console.log(`Flushed ${flushedCount} stale entries (older than ${maxAge}ms)`);
    return flushedCount;
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

// === Flush examples ===
console.log("\n=== Flush Examples ===");

const rl4 = new RateLimiter();

// Make some requests
for (let i = 0; i < 50; i++) {
  rl4.isAllowed("ClientD");
}
console.log("ClientD should be at 50 requests");

// Flush a specific client
rl4.flush("ClientD");

// Now ClientD should be able to make requests again without hitting the limit
console.log("After flush, ClientD allowed?", rl4.isAllowed("ClientD"));

// Add multiple clients
for (let i = 0; i < 30; i++) {
  rl4.isAllowed("ClientE");
  rl4.isAllowed("ClientF");
}

// Flush all clients
rl4.flushAll();

console.log("After flushAll, ClientE allowed?", rl4.isAllowed("ClientE"));

// Demonstrate flushStale
const rl5 = new RateLimiter();
rl5.isAllowed("ClientG");
rl5.isAllowed("ClientH");

setTimeout(() => {
  console.log("\nAfter 2 seconds, flush stale entries older than 1 second:");
  rl5.flushStale(1000); // Flush entries older than 1 second
}, 2000);
