# 169. Majority Element


## Level - easy


## Task
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.


## Объяснение
В данной задаче требуется найти такой элемент массива, который встречается в нём более половины раз (то есть более ⌊n / 2⌋ раз, где n — длина массива). Гарантируется, что такой элемент существует всегда. Необходимо вернуть этот элемент.

 

## Example 1:
```
Input: nums = [3,2,3]
Output: 3
```

## Example 2:
```
Input: nums = [2,2,1,1,1,2,2]
Output: 2
```
 

## Constraints:
- n == nums.length
- 1 <= n <= 5 * 104
- 109 <= nums[i] <= 109


## Follow-up
Could you solve the problem in linear time and in O(1) space?