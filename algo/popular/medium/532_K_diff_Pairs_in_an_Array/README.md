# 532. K-diff Pairs in an Array


## Level - medium


## Task
Given an array of integers nums and an integer k, return the number of unique k-diff pairs in the array.

A k-diff pair is an integer pair (nums[i], nums[j]), where the following are true:

- 0 <= i, j < nums.length
- i != j
- |nums[i] - nums[j]| == k

Notice that |val| denotes the absolute value of val.


## Объяснение

Дана задача поиска количества уникальных пар чисел в массиве, которые отличаются друг от друга ровно на заданное значение k.

**Суть задачи:**
- Необходимо найти все пары чисел (nums[i], nums[j]), где |nums[i] - nums[j]| = k
- Пары должны быть уникальными (не учитывать порядок элементов в паре)
- Индексы i и j должны быть разными (i != j)
- Нужно вернуть количество таких уникальных пар

**Особенности:**
- Если k = 0, то ищем пары одинаковых чисел (дубликаты)
- Учитываем только уникальные пары, даже если числа повторяются в массиве

## Подходы к решению

### 1. Сортировка + два указателя
- Отсортировать массив
- Использовать два указателя для поиска пар с разницей k
- Сложность: O(n log n) из-за сортировки

### 2. Хэш-таблица
- Использовать map для хранения встреченных чисел
- Для каждого числа nums[i] проверить, есть ли nums[i] + k или nums[i] - k в map
- Сложность: O(n) в среднем

### 3. Брутфорс (не рекомендуется)
- Проверить все возможные пары
- Сложность: O(n²) - слишком медленно для больших массивов

## Example 1:
```
Input: nums = [3,1,4,1,5], k = 2
Output: 2
Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
Although we have two 1s in the input, we should only return the number of unique pairs.
```

## Example 2:
```
Input: nums = [1,2,3,4,5], k = 1
Output: 4
Explanation: There are four 1-diff pairs in the array, (1, 2), (2, 3), (3, 4) and (4, 5).
```

## Example 3:
```
Input: nums = [1,3,1,5,4], k = 0
Output: 1
Explanation: There is one 0-diff pair in the array, (1, 1).
```