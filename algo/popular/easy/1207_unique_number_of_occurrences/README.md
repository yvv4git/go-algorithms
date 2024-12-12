# 1207. Unique Number of Occurrences


## Level - easy


## Task
Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.


## Объяснение
Дан массив целых чисел. Нужно определить, являются ли количества вхождений каждого числа в массиве уникальными.
Дан массив arr, состоящий из целых чисел. Верните true, если количество вхождений каждого числа в массиве уникально, и false в противном случае.

Пример 1:
- Вход: arr = [1, 2, 2, 1, 1, 3]
- Выход: true
- Объяснение: Число 1 встречается 3 раза, число 2 — 2 раза, число 3 — 1 раз. Все количества вхождений уникальны.


## Example 1:
```
Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
```


## Example 2:
```
Input: arr = [1,2]
Output: false
```


## Example 3:
```
Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
Output: true
```


## Constraints:
- 1 <= arr.length <= 1000
- -1000 <= arr[i] <= 1000