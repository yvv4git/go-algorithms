# 318. Maximum Product of Word Lengths


## Level - medium


## Task
Given a string array words, return the maximum value of length(word[i]) * length(word[j]) 
where the two words do not share common letters. If no such two words exist, return 0.


## Объяснение
Задача требует от нас найти максимальное произведение длин двух слов в заданном списке, при условии, 
что эти два слова не имеют общих букв.

Формальное описание задачи:
- На вход подается список слов (строк).
- Нужно найти два слова, которые не имеют общих букв.
- Вычислить произведение длин этих двух слов.
- Вернуть максимальное значение такого произведения.

Пример:  
Input: ["abcw", "baz", "foo", "bar", "xtfn", "abcdef"]  
Output: 16  
Explanation:   
- "abcw" и "xtfn" не имеют общих букв.
- Длина "abcw" равна 4, длина "xtfn" равна 4.
- Произведение длин равно 4 * 4 = 16.

Для решения этой задачи можно использовать битовые маски, чтобы эффективно проверять наличие общих букв между словами. 
Каждое слово можно представить в виде битовой маски, где каждый бит соответствует наличию определенной буквы в слове. 
Если два слова не имеют общих букв, то побитовое И между их битовыми масками будет равно 0.


## Example 1:
````
Input: words = ["abcw","baz","foo","bar","xtfn","abcdef"]
Output: 16
Explanation: The two words can be "abcw", "xtfn".
````


## Example 2:
````
Input: words = ["a","ab","abc","d","cd","bcd","abcd"]
Output: 4
Explanation: The two words can be "ab", "cd".
````


## Example 3:
````
Input: words = ["a","aa","aaa","aaaa"]
Output: 0
Explanation: No such pair of words.
````


## Constraints:
- 2 <= words.length <= 1000
- 1 <= words[i].length <= 1000
- words[i] consists only of lowercase English letters.