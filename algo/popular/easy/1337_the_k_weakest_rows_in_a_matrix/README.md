# 1337. The K Weakest Rows in a Matrix


## Level - easy


## Task
You are given an m x n binary matrix mat of 1's (representing soldiers) and 0's (representing civilians). The soldiers are positioned in front of the civilians. That is, all the 1's will appear to the left of all the 0's in each row.

A row i is weaker than a row j if one of the following is true:
- The number of soldiers in row i is less than the number of soldiers in row j.
- Both rows have the same number of soldiers and i < j.

Return the indices of the k weakest rows in the matrix ordered from weakest to strongest.


## Объяснение
Дана матрица (двумерный массив) размером m x n, состоящая из нулей и единиц. 
В каждой строке матрицы все единицы стоят перед нулями. 
Это означает, что каждая строка представляет собой отсортированный массив, 
где сначала идут все единицы, а затем все нули.

Необходимо определить k самых "слабых" строк в матрице. Строка считается слабее другой, если:
- Количество единиц в ней меньше.
- Если количество единиц одинаково, то слабее считается та строка, у которой индекс меньше.

Пример:
```
matrix = [
  [1, 1, 0, 0, 0],
  [1, 1, 1, 0, 0],
  [1, 0, 0, 0, 0],
  [1, 1, 0, 0, 0],
  [1, 1, 1, 1, 0]
]
k = 3
```

Как решать:
1. Подсчет единиц в каждой строке.
Для каждой строки можно посчитать количество единиц. 
Поскольку строки отсортированы, можно использовать бинарный поиск для эффективного подсчета.

2. Сортировка строк. 
Отсортировать строки сначала по количеству единиц, а затем по индексу.

3. Выбор первых k строк. 
После сортировки выбрать первые k строк и вернуть их индексы.



## Example 1:
```
Input: mat = 
[[1,1,0,0,0],
 [1,1,1,1,0],
 [1,0,0,0,0],
 [1,1,0,0,0],
 [1,1,1,1,1]], 
k = 3
Output: [2,0,3]
Explanation: 
The number of soldiers in each row is: 
- Row 0: 2 
- Row 1: 4 
- Row 2: 1 
- Row 3: 2 
- Row 4: 5 
The rows ordered from weakest to strongest are [2,0,3,1,4].
```


## Example 2:
```
Input: mat = 
[[1,0,0,0],
 [1,1,1,1],
 [1,0,0,0],
 [1,0,0,0]], 
k = 2
Output: [0,2]
Explanation: 
The number of soldiers in each row is: 
- Row 0: 1 
- Row 1: 4 
- Row 2: 1 
- Row 3: 1 
The rows ordered from weakest to strongest are [0,2,3,1].
```


## Constraints:
- m == mat.length
- n == mat[i].length
- 2 <= n, m <= 100
- 1 <= k <= m
- matrix[i][j] is either 0 or 1.