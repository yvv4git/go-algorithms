# 132. Palindrome Partitioning II


## Level - hard


## Task
Given a string s, partition s such that every substring of the partition is a palindrome.

Return the minimum cuts needed for a palindrome partitioning of s.


## Объяснение
Задача заключается в разделении заданной строки на набор палиндромов таким образом, 
чтобы минимальное количество разбиений было найдено. Палиндром - это строка, 
которая читается одинаково в обоих направлениях (слева направо и справа налево). 
Например, "aba" является палиндромом, а "abc" - нет.

Задача требует найти минимальное количество разбиений строки таким образом, что каждый разбитый фрагмент будет палиндромом.

Например, если входные данные - "aab", то один из возможных наборов разбиений - ["a","a","b"], 
где каждый элемент является палиндромом. Таким образом, ответ - 1.

Цель задачи - реализовать алгоритм, который будет эффективно находить минимальное количество разбиений.


## Example 1:
````
Input: s = "aab"
Output: 1
Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.
````


## Example 2:
````
Input: s = "a"
Output: 0
````


## Example 3:
````
Input: s = "ab"
Output: 1
````


## Constraints:
- 1 <= s.length <= 2000
- s consists of lowercase English letters only.