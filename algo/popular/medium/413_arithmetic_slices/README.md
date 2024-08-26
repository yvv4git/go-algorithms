# 413. Arithmetic Slices


## Level - medium


## Task
An integer array is called arithmetic if it consists of at least three elements and if the difference between any two consecutive elements is the same.
- For example, [1,3,5,7,9], [7,7,7,7], and [3,-1,-5,-9] are arithmetic sequences.

Given an integer array nums, return the number of arithmetic subarrays of nums.

A subarray is a contiguous subsequence of the array.


## Объяснение
Арифметический срез — это последовательность элементов, в которой разница между соседними элементами постоянна. 
Например, последовательность [1, 3, 5, 7] является арифметическим срезом, так как разница между каждыми двумя соседними элементами равна 2.

Задача 413 заключается в нахождении количества арифметических срезов определенной длины в заданной последовательности чисел. 
Обычно в таких задачах требуется найти все арифметические срезы длиной не менее 3.

Например, для последовательности [1, 2, 3, 4] арифметическими срезами длиной не менее 3 будут:
- [1, 2, 3]
- [2, 3, 4]
- [1, 2, 3, 4]

Таким образом, в данном примере есть 3 арифметических среза длиной не менее 3.

Для решения задачи 413 нам нужно будет разработать алгоритм, который будет перебирать все возможные подмножества заданной последовательности 
и проверять, являются ли они арифметическими срезами требуемой длины.

## Example 1:
````
Input: nums = [1,2,3,4]
Output: 3
Explanation: We have 3 arithmetic slices in nums: [1, 2, 3], [2, 3, 4] and [1,2,3,4] itself.
````


## Example 2:
````
Input: nums = [1]
Output: 0
````


## Constraints:
- 1 <= nums.length <= 5000
- -1000 <= nums[i] <= 1000