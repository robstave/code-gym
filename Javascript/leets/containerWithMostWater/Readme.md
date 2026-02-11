
# Container With Most Water

## Problem Statement
https://leetcode.com/problems/container-with-most-water/description/

Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.


```javascript

/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function(height) {
    
};

```




## solution






--
--

its a two pointer problem, we can start with the left and right pointer at the beginning and end of the array respectively. We can calculate the area formed by the lines at these two pointers and update our maximum area if necessary. Then, we can move the pointer that has the shorter line towards the other pointer, as this is the only way to potentially increase the area. We repeat this process until the two pointers meet.

the logic behind moving the pointer with the shorter line is that the area is limited by the shorter line, so moving it towards the other pointer may increase the height of the container and thus increase the area. On the other hand, moving the pointer with the taller line will not increase the area since it is limited by the shorter line.

```javascript

