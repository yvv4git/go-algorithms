# 1030. Matrix Cells in Distance Order


## Level - easy


## Task
You are given four integers row, cols, rCenter, and cCenter. 
There is a rows x cols matrix and you are on the cell with the coordinates (rCenter, cCenter).

Return the coordinates of all cells in the matrix, sorted by their distance from (rCenter, cCenter) from the smallest distance to the largest distance. 
You may return the answer in any order that satisfies this condition.

The distance between two cells (r1, c1) and (r2, c2) is |r1 - r2| + |c1 - c2|.


## Объяснение
Задача предполагает, что вы работаете с матрицей, представляющей собой сетку, где каждая ячейка имеет координаты (r, c), 
где r - это номер строки, а c - номер столбца. Ваша задача - отсортировать все ячейки в матрице в порядке их расстояния от центральной ячейки (r0, c0). 
Расстояние между двумя ячейками вычисляется как Манхэттенское расстояние, которое определяется как сумма абсолютных разностей их координат.

Если две ячейки имеют одинаковое расстояние, они должны быть отсортированы по их координатам. 
В случае, если две ячейки имеют одинаковые координаты, они считаются эквивалентными и могут быть в любом порядке.

Например, если у вас есть матрица 2x2 с центральной ячейкой в (0, 0), 
то отсортированный порядок ячеек будет следующим: [(0,0), (0,1), (1,0), (1,1)].

Ваша задача - реализовать функцию, которая принимает размеры матрицы (rows, cols) и координаты центральной ячейки (rCenter, cCenter), 
а затем возвращает отсортированный список всех ячеек в матрице в соответствии с их расстоянием до центральной ячейки.


## Example 1:
````
Input: rows = 1, cols = 2, rCenter = 0, cCenter = 0
Output: [[0,0],[0,1]]
Explanation: The distances from (0, 0) to other cells are: [0,1]
````


## Example 2:
````
Input: rows = 2, cols = 2, rCenter = 0, cCenter = 1
Output: [[0,1],[0,0],[1,1],[1,0]]
Explanation: The distances from (0, 1) to other cells are: [0,1,1,2]
The answer [[0,1],[1,1],[0,0],[1,0]] would also be accepted as correct.
````


## Example 3:
````
Input: rows = 2, cols = 3, rCenter = 1, cCenter = 2
Output: [[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
Explanation: The distances from (1, 2) to other cells are: [0,1,1,2,2,3]
There are other answers that would also be accepted as correct, such as [[1,2],[1,1],[0,2],[1,0],[0,1],[0,0]].
````


## Constraints:
- 1 <= rows, cols <= 100
- 0 <= rCenter < rows
- 0 <= cCenter < cols