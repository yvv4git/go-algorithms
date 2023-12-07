## 343. Integer Break


## Level - medium


## Task
Given an integer n, break it into the sum of k positive integers, where k >= 2, and maximize the product of those integers.

Return the maximum product you can get.


## Объяснение
Это задача на динамическое программирование.
Но ее оптимальнее решить с помощью математики.

Функция integerBreak принимает целое число n и возвращает максимальное произведение,
которое можно получить, разбив число n на сумму натуральных чисел.

## Example 1:
````
Input: n = 2
Output: 1
Explanation: 2 = 1 + 1, 1 × 1 = 1.
````


## Example 2:
````
Input: n = 10
Output: 36
Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
````


## Constraints:
- 2 <= n <= 58