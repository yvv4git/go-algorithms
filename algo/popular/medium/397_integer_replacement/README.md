# 397. Integer Replacement


## Level - medium


## Task
Given a positive integer n, you can apply one of the following operations:
1. If n is even, replace n with n / 2.
2. If n is odd, replace n with either n + 1 or n - 1.

Return the minimum number of operations needed for n to become 1.


## Объяснение
Задача относится к категории динамического программирования и рекурсии. 
В этой задаче требуется найти минимальное количество операций, необходимых для преобразования заданного целого числа n в 1. 
Допустимые операции следующие:
- Если n четное, то n можно заменить на n / 2.
- Если n нечетное, то n можно заменить на n + 1 или n - 1.

Пример
Рассмотрим пример для n = 8:
- 8 (четное) -> 8 / 2 = 4
- 4 (четное) -> 4 / 2 = 2
- 2 (четное) -> 2 / 2 = 1

Таким образом, для n = 8 требуется 3 операции.

Для решения этой задачи можно использовать рекурсию с мемоизацией (кешированием результатов) 
или итеративный подход с использованием очереди (BFS).


## Example 1:
````
Input: n = 8
Output: 3
Explanation: 8 -> 4 -> 2 -> 1
````


## Example 2:
````
Input: n = 7
Output: 4
Explanation: 7 -> 8 -> 4 -> 2 -> 1
or 7 -> 6 -> 3 -> 2 -> 1
````

## Example 3:
````
Input: n = 4
Output: 2
````


## Constraints:
- 1 <= n <= 2^31 - 1