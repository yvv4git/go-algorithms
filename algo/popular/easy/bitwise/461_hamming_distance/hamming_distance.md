# 461. Hamming Distance


## Level - easy


## Task
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
Given two integers x and y, return the Hamming distance between them.


## Объяснение
В этой задаче нужно написать функцию hammingDistance, которая будет вычислять расстояние Хэмминга между двумя целыми числами.

Расстояние Хэмминга между двумя целыми числами - это количество позиций, в которых биты различны. 
Например, расстояние Хэмминга между числами 1 (0001 в двоичной системе) и 4 (0100 в двоичной системе) равно 2, 
потому что биты на первой и второй позиции различаются.


## Example 1:
````
Input: x = 1, y = 4
Output: 2
Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
↑   ↑
The above arrows point to positions where the corresponding bits are different.
````

## Example 2:
````
Input: x = 3, y = 1
Output: 1
````

## Constraints:
- 0 <= x, y <= 231 - 1