# 1512. Number of Good Pairs


## Level - Easy


## Task
Given an array of integers nums, return the number of good pairs.

A pair (i, j) is called good if nums[i] == nums[j] and i < j.


## Объяснение
Нам дается массив целых чисел nums.
Задача - определить количество пар (i, j), где nums[i] == nums[j] и i < j.
То есть, мы ищем пары элементов в массиве, где элементы равны и индекс первого элемента меньше индекса второго элемента.

Пример:  
Для массива [1,2,3,1,1,3]:
- Найдутся 4 хорошие пары: (0,3), (0,4), (3,4), (2,5)
- Где (0,3) означает, что nums[0]=nums[3]=1



## Example 1:
```
Input: nums = [1,2,3,1,1,3]
Output: 4
Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.
```


## Example 2:
```
Input: nums = [1,1,1,1]
Output: 6
Explanation: Each pair in the array are good.
```


## Example 3:
```
Input: nums = [1,2,3]
Output: 0
```


## Constraints:
- 1 <= nums.length <= 100
- 1 <= nums[i] <= 100