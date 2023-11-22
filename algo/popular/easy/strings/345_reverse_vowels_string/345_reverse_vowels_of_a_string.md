# 345. Reverse Vowels of a String


## Level - easy


## Task
Given a string s, reverse only all the vowels in the string and return it.
The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.


## Объяснение
В условии задачи требуется реализовать функцию, 
которая будет принимать строку s в качестве входного параметра и возвращать новую строку, 
в которой все гласные буквы будут перевернуты. Гласными считаются буквы 'a', 'e', 'i', 'o', 'u' в любом регистре, 
и они могут встречаться в строке более одного раза.
Например, если входная строка s будет "hello", то функция должна вернуть "holle".

Чтобы решить эту задачу, нужно выполнить следующие шаги:
1. Инициализируйте два указателя, один в начале строки (left) и один в конце строки (right).
2. Пройдитесь по строке с помощью двух указателей, меняя местами гласные, находящиеся в позициях, 
указанных указателями.
3. Продолжайте этот процесс до тех пор, пока указатели не встретятся в середине строки.


## Example 1:
````
Input: s = "hello"
Output: "holle"
````


## Example 2:
````
Input: s = "leetcode"
Output: "leotcede"
````


## Constraints:

- 1 <= s.length <= 3 * 10^5
- s consist of printable ASCII characters.