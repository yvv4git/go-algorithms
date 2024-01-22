# 12. Integer to Roman


## Level - medium


## Task
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example, 2 is written as II in Roman numeral, just two one's added together. 
12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. 
However, the numeral for four is not IIII. Instead, the number four is written as IV. 
Because the one is before the five we subtract it making four. 
The same principle applies to the number nine, which is written as IX. 
There are six instances where subtraction is used:
- I can be placed before V (5) and X (10) to make 4 and 9.
- X can be placed before L (50) and C (100) to make 40 and 90.
- C can be placed before D (500) and M (1000) to make 400 and 900.

Given an integer, convert it to a roman numeral.


## Объяснение
Задача "12. Integer to Roman" заключается в том, чтобы преобразовать целое число в римское число. 
Римские цифры представлены следующими символами:
I: 1
V: 5
X: 10
L: 50
C: 100
D: 500
M: 1000

Также существуют особые случаи, когда меняются правила:
- I можно поставить перед V (5) и X (10), чтобы получить 4 и 9.
- X можно поставить перед L (50) и C (100), чтобы получить 40 и 90.
- C можно поставить перед D (500) и M (1000), чтобы получить 400 и 900.

Для решения этой задачи мы можем использовать подход, основанный на массивах, где индексы соответствуют значениям римских цифр, 
а значения - соответствующим римским цифрам.


## Example 1:
````
Input: num = 3
Output: "III"
Explanation: 3 is represented as 3 ones.
````


## Example 2:
````
Input: num = 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.
````


## Example 3:
````
Input: num = 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
````


## Constraints:
- 1 <= num <= 3999