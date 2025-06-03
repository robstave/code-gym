

- https://leetcode.com/problems/remove-element/description/
- https://leetcode.com/problems/remove-element/submissions/1521286115/ - js


# Problem

Given an integer array `nums` and an integer `val`, remove all occurrences of `val` in `nums` [**in-place**](https://en.wikipedia.org/wiki/In-place_algorithm). The order of the elements may be changed. Then return _the number of elements in_ `nums` _which are not equal to_ `val`.

Consider the number of elements in `nums` which are not equal to `val` be `k`, to get accepted, you need to do the following things:

- Change the array `nums` such that the first `k` elements of `nums` contain the elements which are not equal to `val`. The remaining elements of `nums` are not important as well as the size of `nums`.
- Return `k`.

**Custom Judge:**

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int val = ...; // Value to remove
int[] expectedNums = [...]; // The expected answer with correct length.
                            // It is sorted with no values equaling val.

int k = removeElement(nums, val); // Calls your implementation

assert k == expectedNums.length;
sort(nums, 0, k); // Sort the first k elements of nums
for (int i = 0; i < actualLength; i++) {
    assert nums[i] == expectedNums[i];
}

If all assertions pass, then your solution will be **accepted**.

**Example 1:**

**Input:** nums = [3,2,2,3], val = 3
**Output:** 2, nums = [2,2,_,_]
**Explanation:** Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).

**Example 2:**

**Input:** nums = [0,1,2,2,3,0,4,2], val = 2
**Output:** 5, nums = [0,1,4,0,3,_,_,_]
**Explanation:** Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
Note that the five elements can be returned in any order.
It does not matter what you leave beyond the returned k (hence they are underscores).



# Solution




## Code mine

well.  I totally did not do this right.
It should be in place.  This works

- push it all to stack that is not the value
- write from stack to nums

its fast but not 'in place'


```js

var removeElement = function(nums, val) {

  l = nums.length
  var stack = []
  var k = 0

  for(i=0;i<l;i++){
    if(nums[i] != val){
      k++
      stack.push(nums[i])
    }
  }

  for(j=0;j<k;j++){
    nums[j] = stack[j]
  }

    return k
};
```


## better solution

so this is really the i trailing k thing   two pointer

scan list.   if its NOT the value, write to k and k++
otherwise ( a match...continue looping)

k = size of good data.

```js

var removeElement2 = function(nums, val) {
    var k = 0;
    for(let i=0;i<nums.length;i++){
        if(nums[i]!==val){
            nums[k]=nums[i];
            k++
        }
    }
    return k; 
};
```


## Go version

https://go.dev/play/p/lgARb2M8UmA

```go
func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

```


 