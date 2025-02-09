# 1417. Reformat The String


## Level - easy
You are given an alphanumeric string s. (Alphanumeric string is a string consisting of lowercase English letters and digits).

You have to find a permutation of the string where no letter is followed by another letter and no digit is followed by another digit. That is, no two adjacent characters have the same type.

Return the reformatted string or return an empty string if it is impossible to reformat the string.


## Объяснение
Задача заключается в том, чтобы преобразовать исходную строку, состоящую из букв и цифр, в новую строку, в которой буквы и цифры чередуются. 
Если это невозможно, нужно вернуть пустую строку.

Условие задачи:
- Дана строка s, которая содержит только строчные буквы английского алфавита (a-z) и/или цифры (0-9).
- Необходимо переформатировать строку так, чтобы буквы и цифры чередовались. Например, a1b2c3 или 1a2b3c.
- Если количество букв и цифр отличается более чем на 1, то переформатирование невозможно, и нужно вернуть пустую строку.


## Example 1:
```
Input: s = "a0b1c2"
Output: "0a1b2c"
Explanation: No two adjacent characters have the same type in "0a1b2c". "a0b1c2", "0a1b2c", "0c2a1b" are also valid permutations.
```


## Example 2:
```
Input: s = "leetcode"
Output: ""
Explanation: "leetcode" has only characters so we cannot separate them by digits.
```

## Example 3:
```
Input: s = "1229857369"
Output: ""
Explanation: "1229857369" has only digits so we cannot separate them by characters.
```


## Constraints:
- 1 <= s.length <= 500
- s consists of only lowercase English letters and/or digits.
