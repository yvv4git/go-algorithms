# 279. Perfect Squares


## Level - medium


## Task
Given an integer n, return the least number of perfect square numbers that sum to n.

A perfect square is an integer that is the square of an integer; in other words, it is the product of some integer with itself. 
For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.


## Объяснение
Задача требует от нас найти минимальное количество полных квадратов, сумма которых равна заданному числу n.
Нам нужно найти минимальное количество полных квадратов (чисел, которые могут быть представлены как 𝑘^2 k^2, где k — целое число), 
которые в сумме дают заданное число n. Это задача на динамическое программирование, 
где мы можем использовать массив для хранения минимального количества квадратов, необходимых для получения каждого числа от 1 до n.


Примеры:
````
Для n = 12:
12 можно представить как 4 + 4 + 4 (три квадрата числа 4).
Ответ: 3
````

````
Для n = 13:
13 можно представить как 4 + 9 (один квадрат числа 4 и один квадрат числа 9).
Ответ: 2
````


## Example 1:
````
Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.
````


## Example 2:
````
Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.
````


## Constraints:
- 1 <= n <= 10^4