# 316. Remove Duplicate Letters


## Level - medium


## Task
Given a string s, remove duplicate letters so that every letter appears once and only once. You must make sure your result is
the smallest in lexicographical order
among all possible results.


## Объяснение
Задача предлагает удалить все повторяющиеся буквы из строки, сохраняя порядок букв.

Например, если входная строка "bcabc", то результатом должна быть строка "abc", поскольку "b" и "c" повторяются, 
а "a" - нет.

Также, если входная строка "cbacdcbc", то результатом должна быть строка "acdb", 
поскольку "b" и "c" повторяются, а "a", "d" - нет.

Важно также отслеживать количество вхождений каждой буквы в строку, чтобы можно было определить, 
когда была добавлена последняя вхождение буквы, и ее можно удалить из стека.

В итоге, стек будет содержать уникальные буквы в правильном порядке.


## Example 1:
````
Input: s = "bcabc"
Output: "abc"
````


## Example 2:
````
Input: s = "cbacdcbc"
Output: "acdb"
````


## Constraints:
- 1 <= s.length <= 104
- s consists of lowercase English letters.