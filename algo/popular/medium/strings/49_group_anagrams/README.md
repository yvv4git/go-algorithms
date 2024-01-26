# 49. Group Anagrams


# Level - medium


## Task
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.


## Объяснение
Задача заключается в группировке слов, которые являются анаграммами друг другу. 
Анаграмма - это слово, образованное путём перестановки букв другого слова. 
Например, слова "eat" и "ate" являются анаграммами, так как они состоят из одних и тех же букв, но в разном порядке.

Цель задачи - написать функцию, которая принимает на вход массив строк и группирует анаграммы вместе. 
Результатом должен быть массив массивов строк, где каждый вложенный массив содержит группу анаграмм.

Например, для входного массива ["eat", "tea", "tan", "ate", "nat", "bat"], 
функция должна вернуть [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]], так как "eat", "tea" и "ate" - это анаграммы, 
"tan" и "nat" - анаграммы, а "bat" не имеет анаграмм в данном массиве.


## Example 1:
````
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
````


## Example 2:
````
Input: strs = [""]
Output: [[""]]
````


## Example 3:
````
Input: strs = ["a"]
Output: [["a"]]
````

## Constraints:
- 1 <= strs.length <= 10^4
- 0 <= strs[i].length <= 100
- strs[i] consists of lowercase English letters.