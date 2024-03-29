# 400. Nth Digit


## Level - medium


## Task
Given an integer n, return the nth digit of the infinite integer sequence [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...].


## Объяснение
Задача 400. Nth Digit - это задача, в которой нужно найти n-ю цифру в последовательности 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ..., 
где каждое число представлено отдельно.

Решение этой задачи можно найти, просматривая каждое число и его цифры, пока не найдем n-ю цифру. 
Однако, это не эффективно, поскольку мы будем проходить через все числа, пока не найдем n-ю цифру.

Эффективнее будет использовать математику. Мы можем найти длину и начальное число для блока, в котором находится n-я цифра. 
Затем мы можем найти n-ю цифру в этом блоке.


## Example 1:
````
Input: n = 3
Output: 3
````


## Example 2:
````
Input: n = 11
Output: 0
Explanation: The 11th digit of the sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... is a 0, which is part of the number 10.
````


## Constraints:
- 1 <= n <= 2^31 - 1