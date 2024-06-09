# 334. Increasing Triplet Subsequence


## Level - medium


## Task
Given an integer array nums, 
return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. 
If no such indices exists, return false.


## Объяснение
Задача заключается в нахождении подпоследовательности из трех элементов в массиве, которая строго возрастает. 
Это означает, что каждый следующий элемент должен быть больше предыдущего.

Например, в массиве [1, 2, 3, 4, 5] 
есть подпоследовательность [1, 2, 3], [1, 2, 4], [1, 2, 5], [1, 3, 4], [1, 3, 5], [1, 4, 5], [2, 3, 4], [2, 3, 5], [2, 4, 5], [3, 4, 5].

Задача требует написать функцию, которая будет проверять, существует ли такая подпоследовательность в заданном массиве. 
Функция должна возвращать true, если такая подпоследовательность существует, и false в противном случае.


## Example 1:
````
Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.
````

## Example 2:
````
Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.
````


## Example 3:
````
Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
````


## Constraints:
- 1 <= nums.length <= 5 * 10^5
- -2^31 <= nums[i] <= 2^31 - 1