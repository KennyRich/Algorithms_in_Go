## Day 6

###Question
Write a function that takes in `special` array and returns its `product sum`.
A special array is a non-empty array that contains either integers or the other special arrays.
The product sum of `special array` is the sum of its element where special array inside it should be summed themselves
and then multiplied by their level of depth. 

For example, the product sum of `[x,y]` is `x+y`. The product sum for `[x, [y,z]]` is `x+2y+2z`

`Sample Input:` [5,2,[7,-1], 3, [6,[-13,8], 4]]

`Output`: 5+2+2(7-1)+3+2(6+3(-13+8)+4)  = 12