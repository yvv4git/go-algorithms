# 208. Implement Trie (Prefix Tree)


## Level - medium


## Task
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. 
There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:
- Trie() Initializes the trie object.
- void insert(String word) Inserts the string word into the trie.
- boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
- boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.


## Объяснение
Задача предлагает нам реализовать структуру данных "Префиксное дерево" или "Trie".

Префиксное дерево (Trie) - это древовидная структура данных, которая используется для хранения ассоциативных массивов, 
где ключи - это обычно строки. Каждая вершина дерева представляет префикс строки, а листья соответствуют словам, 
которые были вставлены в дерево.

В этой задаче нам нужно реализовать следующие методы:
- insert(word): Вставка слова в дерево.
- search(word): Проверка наличия слова в дереве.
- startsWith(prefix): Проверка наличия любого слова, начинающегося с префикса.

Также, важно отметить, что в Trie каждая вершина может содержать несколько дочерних вершин, и каждая вершина содержит одну букву. 
Это позволяет эффективно хранить и искать строки.



## Где используется?
Префиксное дерево (Trie) - это особая вид древовидной структуры данных, которая используется для эффективного хранения 
и извлечения ключей в наборах строк. Ключевым преимуществом Trie является эффективность поиска, вставки и удаления строк.

Префиксное дерево используется в различных областях компьютерных наук, включая информационные технологии, компьютерные науки, 
интеллектуальный анализ данных и другие. Ниже приведены несколько примеров, где префиксное дерево может быть полезно:
1. Автозаполнение и предложения поиска
Trie используется для реализации функции автозаполнения в поисковых системах, где он помогает предложить слова или фразы, 
которые могут быть интересны пользователю на основе их текущего ввода.
2. Распознавание слов
Trie используется в обработке естественного языка для распознавания слов в тексте. 
Это полезно в задачах, таких как поиск по ключевым словам, анализ тональности текста и другие.
3. Хранение IP-адресов 
Trie может быть использовано для эффективного хранения и поиска IP-адресов в сетевых приложениях.
4. Хранение маршрутов
В сетевых протоколах, таких как BGP (Border Gateway Protocol), Trie используется для эффективного хранения и поиска маршрутов.
5. Хранение структуры файловой системы
Trie может использоваться для эффективного хранения и поиска файлов и директорий в файловой системе

## Example 1:
````
Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]
Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True

````


## Constraints:
- 1 <= word.length, prefix.length <= 2000
- word and prefix consist only of lowercase English letters.
- At most 3 * 104 calls in total will be made to insert, search, and startsWith.