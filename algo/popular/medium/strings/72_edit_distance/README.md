## 72. Edit Distance


## Level - medium


## Task
Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

You have the following three operations permitted on a word:
- Insert a character
- Delete a character
- Replace a character


## Объяснение
Это задача динамического программирования, которая заключается в нахождении минимального количества операций (вставки, удаления или замены символа), 
необходимых для превращения одной строки в другую.

Например, если у нас есть две строки "kitten" и "sitting", минимальное количество операций, 
необходимых для превращения первой строки во вторую, равно 3: заменить 'k' на 's', заменить 'e' на 'i', и добавить 'g'.


Задача "Edit Distance" является важной в компьютерной лингвистике и текстовой инженерии, 
где она используется для определения схожести между двумя текстами, автоматического исправления ошибок в тексте и других задач, 
связанных с обработкой текста.


## Example 1:
````
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e') 
````


## Example 2:
````
Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
````


## Constraints:
- 0 <= word1.length, word2.length <= 500
- word1 and word2 consist of lowercase English letters.