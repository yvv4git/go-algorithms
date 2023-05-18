# 1492. The kth Factor of n


## Level - Medium


## Task
You are given two positive integers n and k. A factor of an integer n is defined as an integer i where n % i == 0.

Consider a list of all factors of n sorted in ascending order, return the k^th factor in this list or return -1 if n has less than k factors.


## Что надо сделать?
Тут все элементарно: даны числа n и k. При этом понятие фактора i, которое определяется как n % i = 0.
Т.е. остаток от деления должен быть равен 0.
Например, у n=12 такие факторы: [1, 2, 3, 4, 6, 12].
Так как 12 % 1 = 0, 12 % 2 = 0, 12 % 3 = 0, 12 % 4 = 0, 12 % 6 = 0, 12 % 12 = 0.
Следующий этап - это число k. Оно определяет какой из факторов надо вернуть.


## Constraints:
- 1 <= k <= n <= 1000

## Example 1:
``
Input: n = 12, k = 3
Output: 3
Explanation: Factors list is [1, 2, 3, 4, 6, 12], the 3rd factor is 3.
``


## Example 2:
``
Input: n = 7, k = 2
Output: 7
Explanation: Factors list is [1, 7], the 2nd factor is 7.
``

## Example 3:
``
Input: n = 4, k = 4
Output: -1
Explanation: Factors list is [1, 2, 4], there is only 3 factors. We should return -1.
``