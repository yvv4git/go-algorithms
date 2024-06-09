# 91. Decode Ways


## Level - medium


## Task
A message containing letters from A-Z can be encoded into numbers using the following mapping:
````
'A' -> "1"
'B' -> "2"
...
'Z' -> "26"

````

To decode an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the mapping above (there may be multiple ways). For example, "11106" can be mapped into:
- "AAJF" with the grouping (1 1 10 6)
- "KJF" with the grouping (11 10 6)

Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".
Given a string s containing only digits, return the number of ways to decode it.
The test cases are generated so that the answer fits in a 32-bit integer.


## Объяснение
Задача заключается в нахождении количества возможных декодированных сообщений, 
которые могут быть получены из закодированной строки цифр. 
Каждая цифра может быть сопоставлена с буквой от 'A' до 'I' (то есть, '1' может быть декодировано как 'A', '2' как 'B', и так далее). 
Также существует сопоставление для цифр от 'J' до 'Z' (то есть, '10' может быть декодировано как 'J', '11' как 'K', и так далее).

Для решения этой задачи можно использовать динамическое программирование. 
Вы можете создать массив dp, где dp[i] будет содержать количество возможных декодированных сообщений для подстроки s[0:i].

Вот шаги, которые вам нужно выполнить:
1. Инициализируйте dp[0] как 1, так как пустая строка может быть декодирована в одно пустое сообщение.
2. Пройдитесь по строке s слева направо, начиная с индекса 1. Для каждого индекса i, проверьте, может ли текущая цифра (s[i]) быть декодирована как отдельная буква. Если это возможно, добавьте dp[i-1] к dp[i].
3. Также проверьте, может ли текущая цифра (s[i]) и предыдущая цифра (s[i-1]) быть декодированы как одна буква. Если это возможно, добавьте dp[i-2] к dp[i].
4. После того, как вы прошли по всей строке, dp[n-1] будет содержать количество возможных декодированных сообщений для всей строки s.

Важно обратить внимание на случаи, когда текущая цифра равна '0', так как '0' не может быть декодирована как отдельная буква 
и всегда должна быть использована в паре с предыдущей цифрой.




## Example 1:
````
Input: s = "12"
Output: 2
Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
````


## Example 2:
````
Input: s = "226"
Output: 3
Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
````


## Constraints:
- 1 <= s.length <= 100
- s contains only digits and may contain leading zero(s).