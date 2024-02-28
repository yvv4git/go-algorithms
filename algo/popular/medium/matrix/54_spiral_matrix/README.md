# 54. Spiral Matrix


## Level - medium


## Task
Given an m x n matrix, return all elements of the matrix in spiral order.


## Объяснение
Задача предполагает написание функции, которая будет выводить элементы двумерного массива (матрицы) по спирали.

Для решения этой задачи, вам нужно будет использовать идею, чтобы двигаться по спирали, пока не пройдете все элементы матрицы.


## Example 1:
![img.png](img.png)
````
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
````


## Example 2:
![img_1.png](img_1.png)
````
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
````


## Constraints:
- m == matrix.length
- n == matrix[i].length
- 1 <= m, n <= 10
- -100 <= matrix[i][j] <= 100