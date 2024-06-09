# 34. Find First and Last Position of Element in Sorted Array


## Level - medium


## Task
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.


## Объяснение
Задача предлагает найти первую и последнюю позиции заданного элемента в отсортированном массиве.

Для решения этой задачи можно использовать бинарный поиск, так как массив отсортирован. 
Бинарный поиск эффективен для поиска элемента в отсортированном массиве, 
так как каждый раз делит массив на две равные части и исключает половину, где не может быть искомого элемента.


## Example 1:
````
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
````


## Example 2:
````
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
````


## Example 3:
````
Input: nums = [], target = 0
Output: [-1,-1]
````


## Constraints:
- 0 <= nums.length <= 10^5
- -10^9 <= nums[i] <= 10^9
- nums is a non-decreasing array.
- -10^9 <= target <= 10^9