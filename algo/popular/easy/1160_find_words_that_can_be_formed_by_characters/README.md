# 1160. Find Words That Can Be Formed by Characters


## Level - easy


## Task
You are given an array of strings words and a string chars.

A string is good if it can be formed by characters from chars (each character can only be used once).

Return the sum of lengths of all good strings in words.


## Объяснение
Вам дан массив строк words и строка chars. 
Задача состоит в том, чтобы найти сумму длин всех строк из массива words, 
которые можно составить, используя только символы из строки chars. 
Каждый символ в строке chars может быть использован только один раз для каждой строки.

Пример:  
```
Input: words = ["cat","bt","hat","tree"], chars = "atach"
Output: 6
Explanation: 
- "cat" можно составить из символов "a", "t", "c" из строки "atach".
- "bt" нельзя составить, так как в "atach" нет символа "b".
- "hat" можно составить из символов "h", "a", "t" из строки "atach".
- "tree" нельзя составить, так как в "atach" нет символов "r" и "e".
Сумма длин строк, которые можно составить: 3 (cat) + 3 (hat) = 6.
```

Решение:
1. Подсчитать количество каждого символа в строке chars.
2. Для каждой строки в массиве words проверить, можно ли составить эту строку из символов, имеющихся в chars, учитывая их количество.
3. Если строку можно составить, добавить её длину к общей сумме.


## Example 1:
```
Input: words = ["cat","bt","hat","tree"], chars = "atach"
Output: 6
Explanation: The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.
```


## Example 2:
```
Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
Output: 10
Explanation: The strings that can be formed are "hello" and "world" so the answer is 5 + 5 = 10.
```


## Constraints:
- 1 <= words.length <= 1000
- 1 <= words[i].length, chars.length <= 100
- words[i] and chars consist of lowercase English letters.