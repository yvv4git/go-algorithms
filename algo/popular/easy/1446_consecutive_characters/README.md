# 1446. Consecutive Characters


## Level - easy


## Task
The power of the string is the maximum length of a non-empty substring that contains only one unique character.

Given a string s, return the power of s.


## Объяснение
Суть задачи заключается в том:
1. Нам нужно найти "силу строки" (power of string)
2. Сила строки определяется как максимальная длина подстроки, которая состоит из одинаковых символов

Например:
- Для строки "leetcode" сила равна 2, потому что самая длинная подстрока из одинаковых символов это "ee"
- Для строки "abbcccddddeeeeedcba" сила равна 5, так как самая длинная подстрока это "eeeee"

Это классическая задача на поиск максимальной последовательности одинаковых символов в строке. 
Решение обычно включает в себя проход по строке с подсчетом последовательных одинаковых символов и обновлением максимального значения.


## Example 1:
```
Input: s = "leetcode"
Output: 2
Explanation: The substring "ee" is of length 2 with the character 'e' only.
```

## Example 2:
```
Input: s = "abbcccddddeeeeedcba"
Output: 5
Explanation: The substring "eeeee" is of length 5 with the character 'e' only.
```


## Constraints:
- 1 <= s.length <= 500
- s consists of only lowercase English letters.