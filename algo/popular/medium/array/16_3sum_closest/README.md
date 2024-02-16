# 16. 3Sum Closest


## Level - Medium


## Task
Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.


## Объяснение
Задача заключается в том, чтобы найти три числа в заданном массиве, чтобы их сумма была ближе всего к заданному числу.

Например, если у нас есть массив [-1, 2, 1, -4] и число 1, то ближайшая к 1 сумма будет -1 + 2 + 1 = 2.

Решение этой задачи может быть реализовано с помощью алгоритма, 
который будет перебирать все возможные комбинации чисел в массиве и искать такую, 
сумма которой будет ближе всего к заданному числу.



## Example 1:
````
Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
````


## Example 2:
````
Input: nums = [0,0,0], target = 1
Output: 0
Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).
````


## Constraints:
- 3 <= nums.length <= 500
- -1000 <= nums[i] <= 1000
- -10^4 <= target <= 10^4