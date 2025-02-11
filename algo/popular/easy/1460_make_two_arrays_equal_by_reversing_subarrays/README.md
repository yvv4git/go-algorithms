# 1460. Make Two Arrays Equal by Reversing Subarrays


## Level - easy


## Task
You are given two integer arrays of equal length target and arr. In one step, you can select any non-empty subarray of arr and reverse it. You are allowed to make any number of steps.

Return true if you can make arr equal to target or false otherwise.


## Объяснение
Задача состоит в том, чтобы выяснить, можно ли из одного массива сделать второй, переворачивая его подмассивы.

Условие задачи: Даны два массива одинаковой длины — target и arr. 
Нужно определить, можно ли преобразовать массив arr в массив target, применяя любое количество операций, 
где в каждой операции можно выбрать подмассив и перевернуть его.

Алгоритм:
1. Сортировать оба массива.
2. Сравнить отсортированные массивы. 
Если они одинаковые, то можно сделать оба массива равными, переворачивая подмассивы. Если нет — нельзя.

Например:
- Массив target = [1, 2, 3, 4]

Массив arr = [2, 4, 3, 1]

После сортировки обоих массивов:

    target = [1, 2, 3, 4]
    arr = [1, 2, 3, 4]


## Example 1:
```
Input: target = [1,2,3,4], arr = [2,4,1,3]
Output: true
Explanation: You can follow the next steps to convert arr to target:
1- Reverse subarray [2,4,1], arr becomes [1,4,2,3]
2- Reverse subarray [4,2], arr becomes [1,2,4,3]
3- Reverse subarray [4,3], arr becomes [1,2,3,4]
There are multiple ways to convert arr to target, this is not the only way to do so.
```


## Example 2:
```
Input: target = [7], arr = [7]
Output: true
Explanation: arr is equal to target without any reverses.
```


## Example 3:
```
Input: target = [3,7,9], arr = [3,7,11]
Output: false
Explanation: arr does not have value 9 and it can never be converted to target.
```


## Constraints:
- target.length == arr.length
- 1 <= target.length <= 1000
- 1 <= target[i] <= 1000
- 1 <= arr[i] <= 1000