# 953. Verifying an Alien Dictionary


## Level - easy


## Task
In an alien language, surprisingly, they also use English lowercase letters, but possibly in a different order. 
The order of the alphabet is some permutation of lowercase letters.

Given a sequence of words written in the alien language, and the order of the alphabet, 
return true if and only if the given words are sorted lexicographically in this alien language.


## Объяснение
Вам дан список слов words и строка order, которая представляет собой алфавит инопланетной расы. 
Вам нужно определить, отсортированы ли слова в списке в лексикографическом порядке согласно этому алфавиту.

Пример 1:
```
Вход: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
Выход: true
Объяснение: В заданном алфавите 'h' идет перед 'l', поэтому слово "hello" идет перед "leetcode".
```

Пример 2:
```
Вход: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
Выход: false
Объяснение: В заданном алфавите 'd' идет после 'l', поэтому слово "world" не может идти перед "word".
```


## Example 1:
```
Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
Output: true
Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.
```


## Example 2:
```
Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
Output: false
Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.
```


## Example 3:
```
Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
Output: false
Explanation: The first three characters "app" match, and the second string is shorter (in size.) According to lexicographical rules "apple" > "app", because 'l' > '∅', where '∅' is defined as the blank character which is less than any other character (More info).
```


## Constraints:
- 1 <= words.length <= 100
- 1 <= words[i].length <= 20
- order.length == 26
- All characters in words[i] and order are English lowercase letters.