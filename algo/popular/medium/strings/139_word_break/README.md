# 139. Word Break


## Level - medium


## Task
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.


## Объяснение
Это задача на динамическое программирование, которая заключается в том, чтобы определить, 
можно ли разбить данную строку на слова, присутствующие в заданном словаре.

Например, если у нас есть строка "leetcode" и словарь ["leet", "code"], мы можем разбить строку на "leet" и "code", поэтому ответ будет true.

Алгоритм должен быть эффективным, чтобы работать с большими строками и большими словарями. 
Одним из подходов к решению этой задачи является использование динамического программирования. 
Мы будем использовать массив dp, где dp[i] будет true, если строка s может быть разбита на слова из словаря в позиции i.


## Example 1:
````
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
````


## Example 2:
````
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.
````


## Example 3:
````
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false
````


## Constraints:
- 1 <= s.length <= 300
- 1 <= wordDict.length <= 1000
- 1 <= wordDict[i].length <= 20
- s and wordDict[i] consist of only lowercase English letters.
- All the strings of wordDict are unique.