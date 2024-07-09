# 32. Longest Valid Parentheses


## Level - hard


## Task
Given a string containing just the characters '(' and ')', return the length of the longest valid (well-formed) parentheses
substring.


## Объяснение
Задача требует от нас найти длину самой длинной подстроки из правильно расставленных скобок в заданной строке, 
состоящей только из символов '(' и ')'. Нам нужно найти самую длинную подстроку, в которой все скобки правильно сбалансированы. 
Правильно сбалансированная подстрока означает, что каждая открывающая скобка '(' имеет соответствующую закрывающую скобку ')', 
и они расположены в правильном порядке.

Варианты решения:
1. Использование стека.
2. Использование двух проходов.


## Example 1:
````
Input: s = "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()".
````


## Example 2:
````
Input: s = ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()".
````


## Example 3:
````
Input: s = ""
Output: 0
````


## Constraints:
- 0 <= s.length <= 3 * 10^4
- s[i] is '(', or ')'.