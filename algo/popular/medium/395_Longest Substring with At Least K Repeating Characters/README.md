# 395. Longest Substring with At Least K Repeating Characters


## Level - medium


## Task
Given a string s and an integer k, return the length of the longest substring of s such that the frequency of each character in this substring is greater than or equal to k.

if no such substring exists, return 0.


## Объяснение
Задача требует найти длину самой длинной подстроки в заданной строке, в которой каждый символ встречается не менее K раз.
Дана строка s и целое число k. Нужно найти длину самой длинной подстроки s, в которой каждый символ встречается не менее k раз. 
Если такой подстроки нет, вернуть 0.


## Example 1:
````
Input: s = "aaabb", k = 3
Output: 3
Explanation: The longest substring is "aaa", as 'a' is repeated 3 times.
````

## Example 2:
````
Input: s = "ababbc", k = 2
Output: 5
Explanation: The longest substring is "ababb", as 'a' is repeated 2 times and 'b' is repeated 3 times.
````


## Constraints:
- 1 <= s.length <= 10^4
- s consists of only lowercase English letters.
- 1 <= k <= 10^5