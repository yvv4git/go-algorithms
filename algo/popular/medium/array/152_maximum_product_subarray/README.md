# 152. Maximum Product Subarray


## Level - medium


## Task
Given an integer array nums, find a subarray that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.


## Объяснение
Задача заключается в нахождении подмассива с максимальным произведением элементов в заданном массиве.

Например, если у нас есть массив [2,3,-2,4], то максимальное произведение будет 6 (получается, когда мы берем подмассив [2,3]).

Однако, в этой задаче есть особенность: поскольку произведение может давать отрицательные результаты, 
нам нужно также отслеживать минимальное произведение, так как это может стать максимальным, 
если умножить его на отрицательное число.

Таким образом, нам нужно решить задачу на программировании, которая будет искать максимальное и минимальное произведения подмассивов, 
а затем вернуть максимальное из этих двух значений.


## Example 1:
````
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
````


## Example 2:
````
Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
````


## Constraints:
- 1 <= nums.length <= 2 * 10^4
- -10 <= nums[i] <= 10
- The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.