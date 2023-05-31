# 342. Power of Four


## Level - easy


## Task
Given an integer n, return true if it is a power of four. Otherwise, return false.

An integer n is a power of four, if there exists an integer x such that n == 4^x.

Follow up: Could you solve it without loops/recursion?


## Объяснение
Задача простая и интересная, способов решения очень много. Можно действовать в лоб, можно
использовать бинарный поиск, а можно моментально за O(1) получить результат с помощью bitwise
закономерностей.


## Constraints:
- -231 <= n <= 231 - 1


## Example 1:
````
Input: n = 16
Output: true
````


## Example 2:
````
Input: n = 5
Output: false
````


## Example 3:
````
Input: n = 1
Output: true
````