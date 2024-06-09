# 264. Ugly Number II


## Level - medium


## Task
An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

Given an integer n, return the nth ugly number.


## Объяснение
Задача "264. Ugly Number II" заключается в нахождении n-го числа, которое является "ugly number". "Ugly number" - это натуральное число, которое делится нацело только на 2, 3 или 5. Первые 10 "ugly numbers" это 1, 2, 3, 4, 5, 6, 8, 9, 10, 12.

Чтобы решить эту задачу, можно использовать динамическое программирование. Мы будем хранить три указателя для 2, 3 и 5, 
и мы будем увеличивать каждый из них, когда мы находим следующее "ugly number".


## Example 1:
````
Input: n = 10
Output: 12
Explanation: [1, 2, 3, 4, 5, 6, 8, 9, 10, 12] is the sequence of the first 10 ugly numbers.
````


## Example 2:
````
Input: n = 1
Output: 1
Explanation: 1 has no prime factors, therefore all of its prime factors are limited to 2, 3, and 5.
````


## Constraints:
- 1 <= n <= 1690