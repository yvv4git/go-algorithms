# 962. Maximum Width Ramp


## Level - medium


## Task
A ramp in an integer array nums is a pair (i, j) for which i < j and nums[i] <= nums[j]. The width of such a ramp is j - i.

Given an integer array nums, return the maximum width of a ramp in nums. If there is no ramp in nums, return 0.


## Объяснение
Задача состоит в том, чтобы найти максимальную ширину "рампы" в массиве целых чисел nums. 
"Рампа" - это пара индексов (i, j) такая, что i < j и nums[i] <= nums[j]. 
Ширина такой "рампы" определяется как j - i.

В других словами, нам нужно найти две позиции в массиве, 
где значение в первой позиции меньше или равно значению во второй позиции, 
и расстояние между этими позициями максимально.

Например, в массиве [6, 0, 8, 2, 1, 5] максимальная ширина "рампы" равна 4, 
и она достигается при (i, j) = (1, 5), где nums[1] = 0 и nums[5] = 5.

Если в массиве нет "рампы", то функция должна вернуть 0.


## Example 1:
```
Input: nums = [6,0,8,2,1,5]
Output: 4
Explanation: The maximum width ramp is achieved at (i, j) = (1, 5): nums[1] = 0 and nums[5] = 5.
```


## Example 2:
```
Input: nums = [9,8,1,0,1,9,4,0,4,1]
Output: 7
Explanation: The maximum width ramp is achieved at (i, j) = (2, 9): nums[2] = 1 and nums[9] = 1.
```


## Constraints:
- 2 <= nums.length <= 5 * 10^4
- 0 <= nums[i] <= 5 * 10^4