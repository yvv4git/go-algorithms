# 1374. Generate a String With Characters That Have Odd Counts


## Level - easy


## Task
Given an integer n, return a string with n characters such that each character in such string occurs an odd number of times.

The returned string must contain only lowercase English letters. If there are multiples valid strings, return any of them.  


## Объяснение
Необходимо сгенерировать строку из n символов, где каждый символ встречается нечётное количество раз. 
Строка должна состоять только из строчных букв английского алфавита. 
Если существует несколько подходящих строк, можно вернуть любую из них.

Пример:
- Если n = 4, допустимыми строками могут быть:
- - "aaab" (символ 'a' встречается 3 раза, 'b' — 1 раз).
- - "bbbc" (символ 'b' встречается 3 раза, 'c' — 1 раз).

- Если n = 2, допустимой строкой может быть:
- - "ab" (символы 'a' и 'b' встречаются по 1 разу).


## Example 1:
```
Input: n = 4
Output: "pppz"
Explanation: "pppz" is a valid string since the character 'p' occurs three times and the character 'z' occurs once. Note that there are many other valid strings such as "ohhh" and "love".
```


## Example 2:
```
Input: n = 2
Output: "xy"
Explanation: "xy" is a valid string since the characters 'x' and 'y' occur once. Note that there are many other valid strings such as "ag" and "ur".
```


## Example 3:
```
Input: n = 7
Output: "holasss"
```

## Constraints:
- 1 <= n <= 500