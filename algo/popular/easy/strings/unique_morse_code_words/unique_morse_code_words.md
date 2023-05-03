# 804. Unique Morse Code Words


## TASK
International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes, as follows:
- 'a' maps to ".-",
- 'b' maps to "-...",
- 'c' maps to "-.-.", and so on.

Given an array of strings words where each word can be written as a concatenation of the Morse code of each letter.
- For example, "cab" can be written as "-.-..--...", which is the concatenation of "-.-.", ".-", and "-...". 
We will call such a concatenation the transformation of a word.
Return the number of different transformations among all words we have.


## Объяснение
В общем, задача простая, но имеет не очевидные трудности. Во-первых, я не знаком с азбукой Морзе и могу не понимать того,
что 2 разных слова могут выглядеть одинаково, если записать их в азбуке Морзе. Во-вторых, собеседующий может внимательно
наблюдать за тем, как ты будешь конкатенировать строки.

Алгоритм выглядит так:
1. Итерируемся по символам слова.
2. Конвертируем каждый символ в представление азбуки Морзе и записываем во временную переменную.
3. Получившееся слово(в азбуке Морзе) сохраняем в хэшь, hash[morse_word] = true.
4. Вычисляем количество элементов в hash. Оно и будет равно количеству уникальных значений.


## Constraints:

- 1 <= words.length <= 100
- 1 <= words[i].length <= 12
- words[i] consists of lowercase English letters.


## Example 1:
``
Input: words = ["gin","zen","gig","msg"]
Output: 2
Explanation: The transformation of each word is:
"gin" -> "--...-."
"zen" -> "--...-."
"gig" -> "--...--."
"msg" -> "--...--."
There are 2 different transformations: "--...-." and "--...--.".
``


## Example 2:
``
Input: words = ["a"]
Output: 1
``