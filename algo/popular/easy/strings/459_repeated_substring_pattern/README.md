# 459. Repeated Substring Pattern


## Level - easy


## Task
Given a string s, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.


## Методы решения
1. Brute force O(n^2)
2. Конкатенация и подстрока O(n)
3. Алгоритм Кнута-Морриса-Пратта(KMP) O(n)
4. Использовать хэшь-функцию O(n)
5. Конечный автомат O(n)
6. Регулярные выражения O(n)

## Example 1:
````
Input: s = "abab"
Output: true
Explanation: It is the substring "ab" twice.
```` 

## Example 2:
````
Input: s = "aba"
Output: false
````

## Example 3:
````
Input: s = "abcabcabcabc"
Output: true
Explanation: It is the substring "abc" four times or the substring "abcabc" twice.
````

## Constraints:
- 1 <= s.length <= 104
- s consists of lowercase English letters.