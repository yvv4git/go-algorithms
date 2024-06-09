# 131. Palindrome Partitioning


## Level - medium


## Task
Given a string s, partition s such that every substring of the partition is a palindrome. 
Return all possible palindrome partitioning of s.


## Объяснение
Задача "131. Palindrome Partitioning" - это задача на разбиение строки на части так, чтобы каждая часть была палиндромом.  
Например, если мы имеем строку "aab", то один из возможных разбиений этой строки на палиндромы будет ["a", "a", "b"].

Требуется написать функцию, которая будет принимать строку в качестве входного параметра и возвращать все возможные разбиения этой строки на палиндромы.  
Функция должна возвращать список списков строк, где каждый вложенный список представляет собой одно из возможных разбиений.

Например, если входная строка "aab", функция должна вернуть [["a", "a", "b"], ["aa", "b"]].

## Example 1:
````
Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]
````


## Example 2:
````
Input: s = "a"
Output: [["a"]]
````


## Constraints:
- 1 <= s.length <= 16
- s contains only lowercase English letters.