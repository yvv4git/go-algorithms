# 1331. Rank Transform of an Array


## Level - easy


## Task
Given an array of integers arr, replace each element with its rank.

The rank represents how large the element is. The rank has the following rules:
- Rank is an integer starting from 1.
- The larger the element, the larger the rank. If two elements are equal, their rank must be the same.
- Rank should be as small as possible.


## Объяснение
Задача требует преобразовать массив чисел в массив их рангов. 
Ранг числа — это его позиция в отсортированном массиве, где числа с одинаковыми значениями имеют одинаковый ранг.

Например:
- Вход: [40, 10, 20, 30]
- Выход: [4, 1, 2, 3]

Объяснение:
- 10 — первое число в отсортированном массиве, поэтому его ранг 1.
- 20 — второе число в отсортированном массиве, поэтому его ранг 2.
- 30 — третье число в отсортированном массиве, поэтому его ранг 3.
- 40 — четвертое число в отсортированном массиве, поэтому его ранг 4.

## Example 1:
```
Input: arr = [40,10,20,30]
Output: [4,1,2,3]
Explanation: 40 is the largest element. 10 is the smallest. 20 is the second smallest. 30 is the third smallest.
```


## Example 2:
```
Input: arr = [100,100,100]
Output: [1,1,1]
Explanation: Same elements share the same rank.
```


## Example 3:
```
Input: arr = [37,12,28,9,100,56,80,5,12]
Output: [5,3,4,2,8,6,7,1,3]
```


## Constraints:
- 0 <= arr.length <= 10^5
- -109 <= arr[i] <= 10^9