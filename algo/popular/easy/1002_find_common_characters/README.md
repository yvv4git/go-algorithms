# 1002. Find Common Characters


## Level - easy


## Task
Given a string array words, return an array of all characters that show up in all strings within the words (including duplicates). You may return the answer in any order.


## Объяснение
Задача требует от нас найти все символьные символы, которые являются общими для всех строк в заданном массиве строк.
Дан массив строк words. Нам нужно вернуть массив символов, которые являются общими для всех строк в words, с учетом кратности их вхождений.
Нам нужно найти символы, которые присутствуют во всех строках.
Если символ встречается несколько раз в одной строке, но только один раз в другой, то мы должны учитывать наименьшее количество вхождений.


## Example 1:
```
Input: words = ["bella","label","roller"]
Output: ["e","l","l"]
```

Example 2:
```
Input: words = ["cool","lock","cook"]
Output: ["c","o"]
```


## Constraints:
- 1 <= words.length <= 100
- 1 <= words[i].length <= 100
- words[i] consists of lowercase English letters.