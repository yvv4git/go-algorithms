# 78. Subsets


## Level - medium


## Task
Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.


## Объяснение
Задача заключается в том, чтобы найти все возможные подмножества заданного набора чисел.

Например, если у нас есть набор чисел [1, 2, 3], то подмножества будут выглядеть так:  
[] (пустое подмножество)
[1]
[2]
[3]
[1, 2]
[1, 3]
[2, 3]
[1, 2, 3]

Это задача на использование рекурсии или комбинаторики. Вы можете решить ее, используя рекурсивный подход, 
где вы на каждом шаге добавляете или не добавляете текущий элемент в подмножество.


## Example 1:
````
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
````


## Example 2:
````
Input: nums = [0]
Output: [[],[0]]
````

## Constraints:
- 1 <= nums.length <= 10
- -10 <= nums[i] <= 10
- All the numbers of nums are unique.