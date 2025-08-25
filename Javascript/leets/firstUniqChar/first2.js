

/**
 * @param {string} s
 * @return {number}
 */
var firstUniqChar = function(s) {

    hash = new Map()
    arr = s.split('')
 
    for(  let i = 0; i < arr.length; i++) {
        let l = arr[i];
         if (hash.has(l)) {
            hash.set(l, hash.get(l) + 1);
        } else {
            hash.set(l, 1);
        }
    }

    for (let j = 0; j < arr.length; j++) {
      let l = arr[j];
        if (hash.get(l) == 1) {
            return j;
        }
    }

    return -1;
}
