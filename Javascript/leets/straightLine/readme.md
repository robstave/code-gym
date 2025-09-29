
https://leetcode.com/problems/check-if-it-is-a-straight-line/description/

validate that points are in a line

so...the way I did it was to check the slope between each pair of points.


this will work,

(y2 - y1) / (x2 - x1) = (y3 - y2) / (x3 - x2)

but its not stellar

and the infinities are a mess.  Rebalance the equation to avoid division by zero errors.

(y2 - y1) * (x3 - x2) = (y3 - y2) * (x2 - x1)