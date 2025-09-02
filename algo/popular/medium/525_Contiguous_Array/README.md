# 525. Contiguous Array


## Level - medium


## Task
Given a binary array nums, return the maximum length of a contiguous subarray with an equal number of 0 and 1.


## Объяснение
Задача требует найти самый длинный непрерывный подмассив (contiguous subarray) в бинарном массиве (состоящем из 0 и 1), 
в котором количество 0 равно количеству 1. Если такого подмассива нет, вернуть 0.

- Входные данные: Массив nums типа int[], содержащий только 0 и 1.
- Задача: Найти длину самого длинного подмассива, где количество элементов, равных 0, равно количеству элементов, равных 1.
- Выходные данные: Целое число — длина самого длинного подходящего подмассива.
- Ограничения:
    - Длина массива 1 <= nums.length <= 10^5.
    - Элементы массива — только 0 или 1.

## Подходы к решению
- Подход с префиксными суммами и хэш-таблицей (оптимальный)
- Брутфорс (перебор всех подмассивов)
- Подход с префиксными суммами без преобразования 0 в -1


## Example 1:
```
Input: nums = [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.
```


## Example 2:
```
Input: nums = [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
```

## Example 3:
```
Input: nums = [0,1,1,1,1,1,0,0,0]
Output: 6
Explanation: [1,1,1,0,0,0] is the longest contiguous subarray with equal number of 0 and 1.
```

## Constraints:
- 1 <= nums.length <= 10^5
- nums[i] is either 0 or 1.