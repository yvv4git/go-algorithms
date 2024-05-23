# 674. Longest Continuous Increasing Subsequence


## Level - easy


## Task
Given an unsorted array of integers nums, return the length of the longest continuous increasing subsequence (i.e. subarray). 
The subsequence must be strictly increasing.

A continuous increasing subsequence is defined by two indices l and r (l < r) such that it is [nums[l], 
nums[l + 1], ..., nums[r - 1], nums[r]] and for each l <= i < r, nums[i] < nums[i + 1].


## Объяснение
Задача заключается в нахождении самой длинной подпоследовательности возрастающих чисел в заданном массиве чисел. 
Подпоследовательность должна быть непрерывной, то есть числа в ней должны идти в исходном порядке и не должны пропускаться.

Например, если дан массив чисел [1, 3, 5, 4, 7], то самой длинной подпоследовательностью будет [1, 3, 5], 
так как она имеет наибольшую длину и является возрастающей.

Задача требует реализации алгоритма, который будет находить эту подпоследовательность и возвращать ее длину.

## Example 1:
````
Input: nums = [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5] with length 3.
Even though [1,3,5,7] is an increasing subsequence, it is not continuous as elements 5 and 7 are separated by element
4.
````


## Example 2:
````
Input: nums = [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2] with length 1. Note that it must be strictly
increasing.
````


## Constraints:
- 1 <= nums.length <= 10^4
- -10^9 <= nums[i] <= 10^9