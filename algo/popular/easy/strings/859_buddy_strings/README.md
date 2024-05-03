# 859. Buddy Strings


## Level - easy


## Task
Given two strings s and goal, return true if you can swap two letters in s so the result is equal to goal, otherwise, return false.

Swapping letters is defined as taking two indices i and j (0-indexed) such that i != j and swapping the characters at s[i] and s[j].
- For example, swapping at indices 0 and 2 in "abcd" results in "cbad".


## Объяснение
Данная задача находится в контексте алгоритмических задач на LeetCode, и её название "Buddy Strings" означает, 
что даны две строки A и B, и необходимо определить, можно ли изменить строку A таким образом, чтобы она стала строкой B, 
переставляя местами ровно два символа.

Например, если A = "ab", B = "ba", то ответ будет "true", так как можно изменить A так, чтобы она стала B, 
переставляя символы 'a' и 'b'.


## Example 1:
````
Input: s = "ab", goal = "ba"
Output: true
Explanation: You can swap s[0] = 'a' and s[1] = 'b' to get "ba", which is equal to goal.
````


## Example 2:
````
Input: s = "ab", goal = "ab"
Output: false
Explanation: The only letters you can swap are s[0] = 'a' and s[1] = 'b', which results in "ba" != goal.
````


## Example 3:
````
Input: s = "aa", goal = "aa"
Output: true
Explanation: You can swap s[0] = 'a' and s[1] = 'a' to get "aa", which is equal to goal.
````


## Constraints:
- 1 <= s.length, goal.length <= 2 * 10^4
- s and goal consist of lowercase letters.