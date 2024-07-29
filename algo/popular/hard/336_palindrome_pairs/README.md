# 336. Palindrome Pairs


## Level - medium


## Task
You are given a 0-indexed array of unique strings words.

A palindrome pair is a pair of integers (i, j) such that:
- 0 <= i, j < words.length,
- i != j, and
- words[i] + words[j] (the concatenation of the two strings) is a palindrome.

Return an array of all the palindrome pairs of words.

You must write an algorithm with O(sum of words[i].length) runtime complexity.


## Объяснение
У вас есть массив строк words, и вам нужно найти все пары индексов (i, j) такие, что конкатенация строк words[i] + words[j] образует палиндром. 
Палиндром — это строка, которая читается одинаково как слева направо, так и справа налево.

Для решения этой задачи можно использовать следующий подход:
1. Создать словарь.  
Создать словарь, где ключом будет строка, а значением — её индекс в массиве words. 
Это позволит быстро проверять, существует ли строка, которая может образовать палиндром с текущей строкой.

2. Проверка каждой пары.  
Для каждой строки в массиве words проверять все возможные подстроки, 
которые могут образовать палиндром при конкатенации с другой строкой.

3. Проверка палиндрома.  
Для каждой пары строк проверять, является ли результат их конкатенации палиндромом.

4. Добавление результата.  
Если пара строк образует палиндром, добавить индексы этих строк в результирующий список.

## Example 1:
````
Input: words = ["abcd","dcba","lls","s","sssll"]
Output: [[0,1],[1,0],[3,2],[2,4]]
Explanation: The palindromes are ["abcddcba","dcbaabcd","slls","llssssll"]
````


## Example 2:
````
Input: words = ["bat","tab","cat"]
Output: [[0,1],[1,0]]
Explanation: The palindromes are ["battab","tabbat"]
````


## Example 3:
````
Input: words = ["a",""]
Output: [[0,1],[1,0]]
Explanation: The palindromes are ["a","a"]
````


## Constraints:
- 1 <= words.length <= 5000
- 0 <= words[i].length <= 300
- words[i] consists of lowercase English letters.