# 647. Palindromic Substrings


## Level - medium


## Task
Given a string s, return the number of palindromic substrings in it.

A string is a palindrome when it reads the same backward as forward.

A substring is a contiguous sequence of characters within the string.


## Объяснение
Задача на сайте LeetCode заключается в том, чтобы найти все уникальные подстроки, которые являются палиндромами в заданной строке.

Палиндром - это слово, фраза, число или другая последовательность символов, которая читается одинаково в обоих направлениях, 
игнорируя пробелы, знаки препинания и регистр букв.

Например, в строке "abc", палиндромы будут: "a", "b", "c", "aa", "bb", "cc", "aba".

Ваша задача - написать функцию, которая будет принимать строку в качестве входных данных и возвращать количество уникальных палиндромных подстрок в этой строке.

## Example 1:
````
Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".
````

## Example 2:
````
Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
````


## Constraints:
- 1 <= s.length <= 1000
- s consists of lowercase English letters.
