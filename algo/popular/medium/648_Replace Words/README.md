# 648. Replace Words


## Level - medium


## Task
In English, we have a concept called root, which can be followed by some other word to form another longer word - let's call this word derivative. 
For example, when the root "help" is followed by the word "ful", we can form a derivative "helpful".

Given a dictionary consisting of many roots and a sentence consisting of words separated by spaces, 
replace all the derivatives in the sentence with the root forming it. 
If a derivative can be replaced by more than one root, replace it with the root that has the shortest length.

Return the sentence after the replacement.


## Объяснение
английском языке у нас есть понятие корня (root), который можно дополнить другими словами или другим корнем, 
чтобы сформировать более длинное слово. Например, корень "help" может быть дополнен до слова "helping".

Теперь, дан словарь dictionary, состоящий из корней, и предложение sentence, состоящее из слов, разделенных пробелами. 
Нужно заменить все слова в предложении таким образом, что если слово в предложении начинается с какого-либо корня из словаря, 
то оно заменяется на этот корень. Если у слова несколько возможных корней, то выбирается самый короткий из них.

Вход:
- dictionary = ["cat","bat","rat"]
- sentence = "the cattle was rattled by the battery"

Выход:
- "the cat was rat by the bat"

Объяснение:
- "cattle" → начинается с "cat", поэтому заменяется на "cat".
- "rattled" → начинается с "rat", поэтому заменяется на "rat".
- "battery" → начинается с "bat", поэтому заменяется на "bat".


## Example 1:
````
Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
Output: "the cat was rat by the bat"
````


## Example 2:
````
Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
Output: "a a b c"
````


## Constraints:
- 1 <= dictionary.length <= 1000
- 1 <= dictionary[i].length <= 100
- dictionary[i] consists of only lower-case letters.
- 1 <= sentence.length <= 106
- sentence consists of only lower-case letters and spaces.
- The number of words in sentence is in the range [1, 1000]
- The length of each word in sentence is in the range [1, 1000]
- Every two consecutive words in sentence will be separated by exactly one space.
- sentence does not have leading or trailing spaces.