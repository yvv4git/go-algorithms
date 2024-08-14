# 378. Kth Smallest Element in a Sorted Matrix


## Level - medium


## Task
Given an n x n matrix where each of the rows and columns is sorted in ascending order, return the kth smallest element in the matrix.

Note that it is the kth smallest element in the sorted order, not the kth distinct element.

You must find a solution with a memory complexity better than O(n^2).


## Объяснение
Задача заключается в нахождении k-го наименьшего элемента в матрице, где каждая строка и каждый столбец отсортированы в неубывающем порядке.
Дана матрица matrix размером n x n, где каждая строка и каждый столбец отсортированы по возрастанию. 
Необходимо найти k-й наименьший элемент в этой матрице.

Вход: matrix =
````
[[1, 5, 9],
[10, 11, 13],
[12, 13, 15]], 
k = 8
````
Выход: 13

Объяснение: В данной матрице элементы в порядке возрастания: 1, 5, 9, 10, 11, 12, 13, 13, 15. 
Восьмой наименьший элемент — 13.


## Example 1:
````
Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
Output: 13
Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest number is 13
````


## Example 2:
````
Input: matrix = [[-5]], k = 1
Output: -5
````


## Constraints:
- n == matrix.length == matrix[i].length
- 1 <= n <= 300
- -10^9 <= matrix[i][j] <= 10^9
- All the rows and columns of matrix are guaranteed to be sorted in non-decreasing order.
- 1 <= k <= n^2
