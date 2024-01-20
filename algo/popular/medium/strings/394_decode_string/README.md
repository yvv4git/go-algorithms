# 394. Decode String

## Task
Given an encoded string, return its decoded string.
The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. 
Note that k is guaranteed to be a positive integer.
You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. 
Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. 
For example, there will not be input like 3a or 2[4].
The test cases are generated so that the length of the output will never exceed 105.

## Объяснение
Это задача на расшифровку строки, которая содержит повторяющиеся подстроки в формате, 
который выглядит как "количество[подстрока]". Например, строка "3[a]2[bc]" расшифровывается как "aaabcbc".

Задача требует написать функцию, которая принимает на вход такую закодированную строку и возвращает расшифрованную строку.

## Решение с помощью стека
Одним из подходов к решению этой задачи является использование стека. 
Стек будет использоваться для отслеживания повторяющихся подстрок и их количеств. 
Алгоритм будет проходить по строке слева направо и выполнять следующие действия:
1. Если текущий символ является цифрой, то он добавляется в числовое значение, которое представляет количество повторений.
2. Если текущий символ является открывающей скобкой [, то сохраняется текущее число повторений и текущая строка в стеке, 
а затем обновляются значения для следующей подстроки.
3. Если текущий символ является закрывающей скобкой ], то из стека извлекаются последние сохраненные число повторений и строка, 
и текущая строка повторяется указанное количество раз, а затем добавляется к расшифрованной строке.

Если текущий символ является буквой, он добавляется к текущей строке.

## Constraints:
- 1 <= s.length <= 30
- s consists of lowercase English letters, digits, and square brackets '[]'.
- s is guaranteed to be a valid input.
- All the integers in s are in the range [1, 300].


## Example 1:
````
Input: s = "3[a]2[bc]"
Output: "aaabcbc"
````


## Example 2:
````
Input: s = "3[a2[c]]"
Output: "accaccacc"
````


## Example 3:
````
Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"
````
