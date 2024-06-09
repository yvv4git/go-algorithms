# 866. Prime Palindrome


## Level - medium


## Task
Given an integer n, return the smallest prime palindrome greater than or equal to n.

An integer is prime if it has exactly two divisors: 1 and itself. Note that 1 is not a prime number.
- For example, 2, 3, 5, 7, 11, and 13 are all primes.

An integer is a palindrome if it reads the same from left to right as it does from right to left.
- For example, 101 and 12321 are palindromes.

The test cases are generated so that the answer always exists and is in the range [2, 2 * 10^8].


## Объяснение
Задача (Палиндромное простое число) - это задача на нахождение наименьшего простого палиндрома, 
которое больше или равно заданному числу N.

Число называется простым, если оно больше 1 и не имеет положительных делителей, отличных от 1 и самого себя.

Число называется палиндромом, если оно одинаково читается как слева направо, так и справа налево.
Например, числа 1, 2, 3, 11, 121, 12321 и т.д. являются палиндромами.

Требуется написать функцию, которая принимает на вход число N и возвращает наименьшее простое палиндромное число, 
которое больше или равно N.


## Example 1:
````
Input: n = 6
Output: 7
````


## Example 2:
````
Input: n = 8
Output: 11
````

## Example 3:
````
Input: n = 13
Output: 101
````

## Constraints:
- 1 <= n <= 10^8