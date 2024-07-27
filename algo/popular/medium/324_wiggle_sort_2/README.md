# 324. Wiggle Sort II


## Level - medium


## Task
Given an integer array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....

You may assume the input array always has a valid answer.


## Объяснение
Задача требует от нас переставить элементы массива таким образом, чтобы они удовлетворяли условию "wiggle sort". 
Формально, для массива nums длины n должно выполняться следующее условие:  
nums[0] < nums[1] > nums[2] < nums[3] > ...

Это означает, что элементы должны чередоваться так, чтобы каждый второй элемент был больше своих соседей.

Примеры:
- Для nums = [1, 5, 1, 1, 6, 4], одно из возможных решений может быть [1, 6, 1, 5, 1, 4].
- Для nums = [1, 3, 2, 2, 3, 1], одно из возможных решений может быть [2, 3, 1, 3, 1, 2].

Основные шаги для решения:
1. Сортировка массива.  
Сначала отсортируем массив.
2. Разделение на две части.  
Разделим отсортированный массив на две части. Если длина массива нечетная, первая часть будет на один элемент больше.
3. Создание результирующего массива.  
Заполним результирующий массив, чередуя элементы из первой и второй частей.


## Example 1:
````
Input: nums = [1,5,1,1,6,4]
Output: [1,6,1,5,1,4]
Explanation: [1,4,1,5,1,6] is also accepted.
````


## Example 2:
````
Input: nums = [1,3,2,2,3,1]
Output: [2,3,1,3,1,2]
````


## Constraints:
- 1 <= nums.length <= 5 * 10^4
- 0 <= nums[i] <= 5000
- It is guaranteed that there will be an answer for the given input nums.


## Follow Up
Can you do it in O(n) time and/or in-place with O(1) extra space?