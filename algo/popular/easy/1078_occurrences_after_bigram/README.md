# 1078. Occurrences After Bigram


## Level - easy


## Task
Given two strings first and second, consider occurrences in some text of the form "first second third", 
where second comes immediately after first, and third comes immediately after second.

Return an array of all the words third for each occurrence of "first second third".


## Объяснение
Даны две строки first и second. 
Необходимо найти все вхождения в тексте вида "first second third", где second следует сразу за first, а third следует сразу за second.

Решение:
Для решения задачи можно разбить текст на слова и пройтись по ним, проверяя, соответствуют ли текущее и следующее слово строкам first и second. Если соответствие найдено, добавляем следующее слово в результат.

Шаги решения:
- Разбить текст на массив слов.
- Пройтись по массиву слов, начиная с первого элемента и до предпоследнего.
- Проверить, совпадают ли текущее слово и следующее слово с first и second.
- Если совпадают, добавить слово, следующее за second, в результат.
- Вернуть результат.


## Example 1:
```
Input: text = "alice is a good girl she is a good student", first = "a", second = "good"
Output: ["girl","student"]
```


## Example 2:
```
Input: text = "we will we will rock you", first = "we", second = "will"
Output: ["we","rock"]
```


Constraints:
- 1 <= text.length <= 1000
- text consists of lowercase English letters and spaces.
- All the words in text are separated by a single space.
- 1 <= first.length, second.length <= 10
- first and second consist of lowercase English letters.
- text will not have any leading or trailing spaces.