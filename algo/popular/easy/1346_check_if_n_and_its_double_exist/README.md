# 1346. Check If N and Its Double Exist


## Level - easy


## Task
Given an array arr of integers, check if there exist two indices i and j such that :
- i != j
- 0 <= i, j < arr.length
- arr[i] == 2 * arr[j]


## Объяснение
Задача требует проверить, существует ли в данном массиве arr два различных индекса i и j, 
такие что элемент под индексом i равен удвоенному элементу под индексом j. 
Другими словами, нужно найти пару чисел в массиве, где одно из них является двойным значением другого.

Детали задачи:
- Условия индексов: Индексы i и j должны быть разными (i != j) и находиться в пределах от 0 до длины массива минус один (0 <= i, j < arr.length).
- Условие значений: Значение элемента массива с индексом i должно быть равно двойному значению элемента массива с индексом j (arr[i] == 2 * arr[j]).


## Example 1:
```
Input: arr = [10,2,5,3]
Output: true
Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]
```

## Example 2:
```
Input: arr = [3,1,7,11]
Output: false
Explanation: There is no i and j that satisfy the conditions.
```
 

## Constraints:
- 2 <= arr.length <= 500
- -103 <= arr[i] <= 103