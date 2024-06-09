# 33. Search in Rotated Sorted Array


## Level - medium


## Task
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). 
For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.


## Объяснение
Задача предполагает поиск заданного элемента в отсортированном массиве, который был циклически сдвинут. 
Это значит, что массив был отсортирован, а затем циклически сдвинут на некоторое количество позиций.

Например, если мы имеем массив [0,1,2,4,5,6,7] и циклически сдвинули его на одну позицию вправо, то получим [7,0,1,2,4,5,6].

Задача состоит в том, чтобы найти индекс заданного элемента в таком массиве, если он там есть, или вернуть -1, если его нет.

Один из возможных подходов к решению этой задачи - это использование бинарного поиска. 
В отсортированном массиве бинарный поиск работает за логарифмическое время, что является очень эффективным. 
Однако в этом случае нам нужно будет изменить бинарный поиск, чтобы он работал с отсортированным и циклически сдвинутым массивом.


## Example 1:
````
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
````


## Example 2:
````
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
````


## Example 3:
````
Input: nums = [1], target = 0
Output: -1
````


## Constraints:
- 1 <= nums.length <= 5000
- -10^4 <= nums[i] <= 10^4
- All values of nums are unique.
- nums is an ascending array that is possibly rotated.
- -10^4 <= target <= 10^4
