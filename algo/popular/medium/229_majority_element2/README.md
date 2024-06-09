# 229. Majority Element II


## Level - medium


## Task
Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.


## Объяснение
Задача находит все элементы, которые встречаются более n/3 раз в массиве.
Это означает, что если в массиве есть элемент, который встречается более n/3 раз, где n - размер массива, 
то он должен быть включен в результирующий список.

Например, если у нас есть массив [3, 2, 3], то элемент 3 встречается два раза, 
что больше 3/3, поэтому он является "majority element".

Также, если у нас есть массив [1, 1, 1, 3, 3, 2, 2, 2], то элементы 1 и 2 встречаются три раза, что больше 8/3, 
поэтому они являются "majority elements".

Решение этой задачи может быть реализовано с использованием алгоритма подсчета элементов 
или алгоритма "Boyer-Moore Voting Algorithm".


## Example 1:
````
Input: nums = [3,2,3]
Output: [3]
````


## Example 2:
````
Input: nums = [1]
Output: [1]
````


## Example 3:
````
Input: nums = [1,2]
Output: [1,2]
````


## Constraints:
- 1 <= nums.length <= 5 * 10^4
- -10^9 <= nums[i] <= 10^9


Follow up: Could you solve the problem in linear time and in O(1) space?