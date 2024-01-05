# 214. Shortest Palindrome


# Level - medium


## Task
You are given a string s. You can convert s to a
palindrome
by adding characters in front of it.

Return the shortest palindrome you can find by performing this transformation.


## Объяснение
Задача заключается в том, чтобы найти самый короткий палиндром, 
который можно создать путем добавления некоторых символов в начало строки.

Например, если у нас есть строка "abcd", самый короткий палиндром, который можно создать, это "dcbabcd".

Эта задача часто используется в различных алгоритмах, таких как поиск в тексте, 
где нужно найти самый короткий палиндром, который содержит все символы строки.


## Example 1:
````
Input: s = "aacecaaa"
Output: "aaacecaaa"
````


## Example 2:
````
Input: s = "abcd"
Output: "dcbabcd"
````


## Constraints:
- 0 <= s.length <= 5 * 10^4
- s consists of lowercase English letters only.