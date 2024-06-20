# 211. Design Add and Search Words Data Structure


## Level - medium


## Task
Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the WordDictionary class:
- WordDictionary() Initializes the object.
- void addWord(word) Adds word to the data structure, it can be matched later.
- bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. 
word may contain dots '.' where dots can be matched with any letter.


## Объяснение
Задача требует от вас создать структуру данных, которая поддерживает добавление слов и поиск слов с использованием шаблонов. 
В частности, структура данных должна поддерживать две операции:
1. addWord(word): Добавляет слово в структуру данных.
2. search(word): Проверяет, существует ли слово в структуре данных. Слово может содержать точки ., которые заменяют любой одиночный символ.

Например, если вы добавили слова "bad", "dad", и "mad" в структуру данных, то:
- search("pad") должен вернуть false, потому что "pad" не было добавлено.
- search("bad") должен вернуть true, потому что "bad" было добавлено.
- search(".ad") должен вернуть true, потому что "bad", "dad", и "mad" соответствуют шаблону ".ad".
- search("b..") должен вернуть true, потому что "bad" соответствует шаблону "b..".

Для реализации этой структуры данных можно использовать префиксное дерево (trie). 
Префиксное дерево — это эффективная структура данных для хранения набора строк, позволяющая быстро выполнять операции добавления и поиска. 
Для обработки точек в шаблонах поиска потребуется рекурсивный или итеративный подход для проверки всех возможных символов на каждой позиции, 
где встречается точка.


## Example:
````
Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output
[null,null,null,null,false,true,true,true]

Explanation
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True
````


## Constraints:
- 1 <= word.length <= 25
- word in addWord consists of lowercase English letters.
- word in search consist of '.' or lowercase English letters.
- There will be at most 2 dots in word for search queries.
- At most 104 calls will be made to addWord and search.