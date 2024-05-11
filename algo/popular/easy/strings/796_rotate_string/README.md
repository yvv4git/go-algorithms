# 796. Rotate String


## Level - easy


## Task
Given two strings s and goal, return true if and only if s can become goal after some number of shifts on s.

A shift on s consists of moving the leftmost character of s to the rightmost position.
- For example, if s = "abcde", then it will be "bcdea" after one shift.


## Объяснение
Это задача на проверку, является ли одна строка циклическим сдвигом другой строки.

Циклическим сдвигом строки называется операция, при которой последний символ строки перемещается в начало строки. 
Например, если у нас есть строка "abcde", то одним из ее циклических сдвигов будет "eabcd".

Задача требует написать функцию, которая принимает две строки и возвращает true, 
если одна строка является циклическим сдвигом другой, и false в противном случае.

Например:
````
Input: A = "abcde", B = "cdeab"
Output: true
````


## Example 1:
````
Input: s = "abcde", goal = "cdeab"
Output: true
````


## Example 2:
````
Input: s = "abcde", goal = "abced"
Output: false
````


## Constraints:
- 1 <= s.length, goal.length <= 100
- s and goal consist of lowercase English letters.