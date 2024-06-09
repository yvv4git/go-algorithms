# 2396. Strictly Palindromic Number


## Level - medium


## Task
An integer n is strictly palindromic if, for every base b between 2 and n - 2 (inclusive), the string representation of the integer n in base b is palindromic.

Given an integer n, return true if n is strictly palindromic and false otherwise.

A string is palindromic if it reads the same forward and backward.


## Объяснение
Задача состоит в том, чтобы определить, является ли число n строго палиндромным. 
Число n называется строго палиндромным, если для каждой системы счисления b, где 2 <= b <= n - 2, 
представление числа n в системе счисления b является палиндромом.

Функция isStrictlyPalindromic(n int) bool должна возвращать true, если n является строго палиндромным, и false в противном случае.

Строка называется палиндромом, если она читается одинаково вперед и назад.


## Example 1:
````
Input: n = 9
Output: false
Explanation: In base 2: 9 = 1001 (base 2), which is palindromic.
In base 3: 9 = 100 (base 3), which is not palindromic.
Therefore, 9 is not strictly palindromic so we return false.
Note that in bases 4, 5, 6, and 7, n = 9 is also not palindromic.
````


## Example 2:
````
Input: n = 4
Output: false
Explanation: We only consider base 2: 4 = 100 (base 2), which is not palindromic.
Therefore, we return false.
````


## Constraints:
- 4 <= n <= 10^5