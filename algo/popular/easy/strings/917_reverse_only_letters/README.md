# 917. Reverse Only Letters


## Level - easy


## Task
Given a string s, reverse the string according to the following rules:
- All the characters that are not English letters remain in the same position.
- All the English letters (lowercase or uppercase) should be reversed.

Return s after reversing it.


## Объяснение
Задача состоит в том, чтобы перевернуть только буквы в строке, при этом не переворачивать другие символы, 
такие как цифры, знаки препинания и т.д.

Например, если входная строка "ab-cd", то результатом должна быть строка "dc-ba".

Основная идея решения заключается в том, чтобы использовать два указателя, один на начало строки, другой на конец, 
и двигать их к центру, меняя местами символы, которые являются буквами. 
Если символы, на которые указывают указатели, не являются буквами, то мы просто двигаем соответствующий указатель.

Важно отметить, что задача требует, чтобы буквы были перевернуты, а другие символы остались на своих местах. 
Это означает, что алгоритм должен быть эффективным и не использовать дополнительную память, кроме тех, 
которая используется для хранения входной и выходной строк.


## Example 1:
````
Input: s = "ab-cd"
Output: "dc-ba"
````


## Example 2:
````
Input: s = "a-bC-dEf-ghIj"
Output: "j-Ih-gfE-dCba"
````


## Example 3:
````
Input: s = "Test1ng-Leet=code-Q!"
Output: "Qedo1ct-eeLg=ntse-T!"
````

## Constraints:
- 1 <= s.length <= 100
- s consists of characters with ASCII values in the range [33, 122].
- s does not contain '\"' or '\\'.