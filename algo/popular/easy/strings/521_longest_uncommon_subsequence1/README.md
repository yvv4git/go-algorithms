# 521. Longest Uncommon Subsequence I


## Level - easy


## Task
Given two strings a and b, return the length of the longest uncommon subsequence between a and b. 
If no such uncommon subsequence exists, return -1.

An uncommon subsequence between two strings is a string that is a subsequence  of exactly one of them.


## Объяснение
Задача заключается в нахождении самой длинной необычной (не повторяющейся) подпоследовательности в двух строках.

Необычной подпоследовательностью называется такая подпоследовательность, которая встречается только один раз в строке. 
Если такой подпоследовательности не существует, то длина необычной подпоследовательности считается равной -1.

Например, если мы имеем две строки "abc" и "abc", то самая длинная необычная подпоследовательность будет "abc", 
так как она встречается только один раз в обеих строках. 
Если мы имеем две строки "abc" и "def", то самая длинная необычная подпоследовательность будет -1, 
так как ни одна из строк не имеет необычных подпоследовательностей.


## Example 1:
````
Input: a = "aba", b = "cdc"
Output: 3
Explanation: One longest uncommon subsequence is "aba" because "aba" is a subsequence of "aba" but not "cdc".
Note that "cdc" is also a longest uncommon subsequence.
````


## Example 2:
````
Input: a = "aaa", b = "bbb"
Output: 3
Explanation: The longest uncommon subsequences are "aaa" and "bbb".
````


## Example 3:
````
Input: a = "aaa", b = "aaa"
Output: -1
Explanation: Every subsequence of string a is also a subsequence of string b. 
Similarly, every subsequence of string b is also a subsequence of string a. So the answer would be -1.
````