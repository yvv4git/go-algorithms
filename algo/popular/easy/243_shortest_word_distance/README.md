# 243. Shortest Word Distance


## Level - easy


## Task
Дан массив слов wordsDict, который содержит список слов, и два слова word1 и word2. 
Вам нужно найти кратчайшее расстояние между этими двумя словами в массиве.

Вам нужно пройтись по массиву wordsDict и отслеживать индексы, на которых встречаются word1 и word2. 
Затем вычислить минимальную разницу между этими индексами, чтобы найти кратчайшее расстояние между двумя словами.


## Example-1
```
wordsDict = ["practice", "makes", "perfect", "coding", "makes"]
word1 = "coding"
word2 = "practice"
```