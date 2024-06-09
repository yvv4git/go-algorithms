# 22. Generate Parentheses


## Level - medium


## Task
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.


## Объяснение
Задача "Generate Parentheses" заключается в том, чтобы генерировать все возможные комбинации правильных скобочных последовательностей, используя n пар скобок.

Например, если n = 3, то возможны следующие комбинации:  
[
"((()))",
"(()())",
"(())()",
"()(())",
"()()()"
]

## Example 1:
````
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
````


## Example 2:
````
Input: n = 1
Output: ["()"]
````


## Constraints:
- 1 <= n <= 8