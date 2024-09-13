# 819. Most Common Word


## Level - easy


## Task
Given a string paragraph and a string array of the banned words banned, 
return the most frequent word that is not banned. It is guaranteed there is at least one word that is not banned, 
and that the answer is unique.

The words in paragraph are case-insensitive and the answer should be returned in lowercase.

## Объяснение
Дана строка, содержащая текст, и список запрещенных слов (бан-лист). 
Необходимо найти наиболее часто встречающееся слово в тексте, которое не входит в список запрещенных слов.

Пример:
````
Вход:
Текст: "Bob hit a ball, the hit BALL flew far after it was hit."
Запрещенные слова: ["hit"]
Выход: "ball"
````

Для решения задачи можно использовать различные подходы, 
включая использование хеш-таблиц (словарей) для подсчета частоты слов и фильтрации запрещенных слов.

Важные моменты:
- Регистронезависимость. Слова должны быть приведены к нижнему регистру, чтобы игнорировать различия в написании.
- Фильтрация запрещенных слов. Слова из списка запрещенных не должны учитываться при подсчете частоты.
- Обработка пунктуации. Пунктуация должна быть удалена или проигнорирована при разбиении текста на слова.


## Example 1:
````
Input: paragraph = "Bob hit a ball, the hit BALL flew far after it was hit.", banned = ["hit"]
Output: "ball"
Explanation:
"hit" occurs 3 times, but it is a banned word.
"ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph.
Note that words in the paragraph are not case sensitive,
that punctuation is ignored (even if adjacent to words, such as "ball,"),
and that "hit" isn't the answer even though it occurs more because it is banned.
````


## Example 2:
````
Input: paragraph = "a.", banned = []
Output: "a"
````


## Constraints:
- 1 <= paragraph.length <= 1000
- paragraph consists of English letters, space ' ', or one of the symbols: "!?',;.".
- 0 <= banned.length <= 100
- 1 <= banned[i].length <= 10
- banned[i] consists of only lowercase English letters.