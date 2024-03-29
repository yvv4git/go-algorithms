# 172. Factorial Trailing Zeroes


## Level - medium


## Task
Given an integer n, return the number of trailing zeroes in n!.

Note that n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.


## Объяснение
В английском языке - "Нули в конце факториала". 
Задача заключается в нахождении количества нулей в конце факториала заданного числа. 
Факториал числа n (обозначается как n!) - это произведение всех натуральных чисел от 1 до n.

Например, факториал числа 5 (5!) равен 120, который заканчивается на ноль. 
Таким образом, в этом случае ответ будет 1.

Задача может быть решена различными способами, но одним из самых эффективных является подсчет количества пар чисел (2, 5), 
которые входят в состав чисел от 1 до n. Так как 2 встречается чаще, чем 5, нас интересует только количество 5 в составе чисел.


## Example 1:
````
Input: n = 3
Output: 0
Explanation: 3! = 6, no trailing zero.
````


## Example 2:
````
Input: n = 5
Output: 1
Explanation: 5! = 120, one trailing zero.
````


## Example 3:
````
Input: n = 0
Output: 0
````


## Constraints:
- 0 <= n <= 10^4


Follow up: Could you write a solution that works in logarithmic time complexity?s