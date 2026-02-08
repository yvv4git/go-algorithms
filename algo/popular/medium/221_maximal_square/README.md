# 221. Maximal Square


## Level - medium


## Task
Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.


## Объяснение

Задача: найти наибольший квадрат из единиц в бинарной матрице и вернуть его площадь.

Подходы:
- Динамическое программирование: dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1. O(m*n).
- Brute Force: перебор всех квадратов. O(m*n*min(m,n)²). Медленно.
- Префиксные суммы: быстрая проверка квадратов. O(m*n*min(m,n)).

Рекомендуется: DP.
![alt text](image.png)

```
Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output: 4
```


## Example 2:
![alt text](image-1.png)

```
Input: matrix = [["0","1"],["1","0"]]
Output: 1
```


## Example 3:
```
Input: matrix = [["0"]]
Output: 0
```


## Constraints:
- m == matrix.length
- n == matrix[i].length
- 1 <= m, n <= 300
- matrix[i][j] is '0' or '1'.