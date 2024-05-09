# 697. Degree of an Array


## Level - easy


## Task
Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.

Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.


## Объяснение
Задача состоит в том, чтобы найти наименьшую длину подмассива в массиве, который содержит все элементы максимальной частоты

Например, если у нас есть массив [1, 2, 2, 3, 1], то наименьшая длина подмассива, 
который содержит все элементы максимальной частоты (в данном случае это число 2, которое встречается два раза), 
будет равна 2, так как подмассив [2, 2] является наименьшим, который содержит все элементы максимальной частоты.

Решение задачи должно быть эффективным, так как оно должно работать за линейное время O(n), 
где n - количество элементов в массиве.


## Example 1:
````
Input: nums = [1,2,2,3,1]
Output: 2
Explanation:
The input array has a degree of 2 because both elements 1 and 2 appear twice.
Of the subarrays that have the same degree:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
The shortest length is 2. So return 2.
````


## Example 2:
````
Input: nums = [1,2,2,3,1,4,2]
Output: 6
Explanation:
The degree is 3 because the element 2 is repeated 3 times.
So [2,2,3,1,4,2] is the shortest subarray, therefore returning 6.
````


## Constraints:
- nums.length will be between 1 and 50,000.
- nums[i] will be an integer between 0 and 49,999.