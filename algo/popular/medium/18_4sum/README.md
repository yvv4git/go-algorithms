# 18. 4Sum


## Level - medium


## Task
Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
- 0 <= a, b, c, d < n
- a, b, c, and d are distinct.
- nums[a] + nums[b] + nums[c] + nums[d] == target

You may return the answer in any order.


## Объяснение
Это задача на поиск всех уникальных четверок чисел из массива, которые в сумме дают определенное значение.

Например, если у нас есть массив чисел [1, 0, -1, 0, -2, 2] и мы ищем сумму 0, то один из возможных ответов может быть четверкой [-1, 0, 0, 1].

Основная суть задачи заключается в том, чтобы найти все такие четверки, которые удовлетворяют условию суммы. 
Это может быть сделано с помощью различных алгоритмов, таких как перебор, хеширование или двоичный поиск.


## Example 1:
````
Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
````


## Example 2:
````
Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]
````


## Constraints:
- 1 <= nums.length <= 200
- -10^9 <= nums[i] <= 10^9
- -10^9 <= target <= 10^9