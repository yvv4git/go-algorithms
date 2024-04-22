# 520. Detect Capital


## Level - easy


## Task
We define the usage of capitals in a word to be right when one of the following cases holds:
- All letters in this word are capitals, like "USA".
- All letters in this word are not capitals, like "leetcode".
- Only the first letter in this word is capital, like "Google".

Given a string word, return true if the usage of capitals in it is right.


## Объяснение
Задача заключается в определении правильного использования регистра букв в слове. 
В зависимости от правильного использования регистра букв в слове, слово может быть написано как: все прописные буквы, 
все строчные буквы или первая буква прописная, остальные строчные.

Например, слова "USA", "leetcode" и "Google" должны быть написаны как: "USA", "leetcode" и "Google" соответственно.

Требуется написать функцию, которая принимает на вход строку слова и возвращает True, 
если все буквы в слове написаны правильно, и False, если нет.

Обратите внимание, что в слове, которое состоит из одной буквы, регистр буквы может быть любым.


## Example 1:
````
Input: word = "USA"
Output: true
````


## Example 2:
````
Input: word = "FlaG"
Output: false
````

## Constraints:
- 1 <= word.length <= 100
- word consists of lowercase and uppercase English letters.