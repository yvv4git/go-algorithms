# 482. License Key Formatting


## Level - easy


## Task
You are given a license key represented as a string s that consists of only alphanumeric characters and dashes. 
The string is separated into n + 1 groups by n dashes. You are also given an integer k.

We want to reformat the string s such that each group contains exactly k characters, except for the first group, 
which could be shorter than k but still must contain at least one character. 
Furthermore, there must be a dash inserted between two groups, and you should convert all lowercase letters to uppercase.

Return the reformatted license key.


## Объснение
Задача "License Key Formatting" предлагает нам переформатировать строку, которая представляет собой ключ лицензии, для удобства чтения. 
Ключ лицензии состоит из групп символов, разделенных дефисами. Например, если у нас есть ключ "5F3Z-2e-9-w", 
мы можем переформатировать его, чтобы он выглядел так: "5F3Z-2E9W".

Задача требует, чтобы мы разделили ключ на группы, каждая из которых состоит из K символов, 
за исключением первой группы, которая может содержать меньшее количество символов. 
Все буквы должны быть в верхнем регистре. Если в ключе есть дефисы, они должны быть удалены.

Например, если K = 4, а ключ "2-5g-3-J", то мы должны получить "2-5G3J" или "25G3J", если мы не учитываем дефисы.


## Example 1:
````
Input: s = "5F3Z-2e-9-w", k = 4
Output: "5F3Z-2E9W"
Explanation: The string s has been split into two parts, each part has 4 characters.
Note that the two extra dashes are not needed and can be removed.
````


## Example 2:
````
Input: s = "2-5g-3-J", k = 2
Output: "2-5G-3J"
Explanation: The string s has been split into three parts, each part has 2 characters except the first part as it could be shorter as mentioned above.
````


## Constraints:
- 1 <= s.length <= 10^5
- s consists of English letters, digits, and dashes '-'.
- 1 <= k <= 10^4