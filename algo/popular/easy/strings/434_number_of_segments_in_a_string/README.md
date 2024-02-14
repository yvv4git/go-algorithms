# 434. Number of Segments in a String


## Level - easy


## Task
Given a string s, return the number of segments in the string.

A segment is defined to be a contiguous sequence of non-space characters.


## Объяснение
Дана строка s, состоящая из нескольких слов, разделенных пробелами. 
Каждое слово представляет собой последовательность символов, не содержащих пробелов. 
Верните количество сегментов в строке s.

Сегмент - это непрерывная последовательность символов, разделенная одним или несколькими пробелами.

Например, если входные данные будут строкой "Hello, my name is John", то ответ должен быть 5, 
потому что в этой строке есть 5 сегментов: "Hello,", "my", "name", "is" и "John".

Задача предполагает написание функции, которая будет принимать строку s в качестве входных данных и возвращать количество сегментов в этой строке.`


## Example 1:
````
Input: s = "Hello, my name is John"
Output: 5
Explanation: The five segments are ["Hello,", "my", "name", "is", "John"]
````


## Example 2:
````
Input: s = "Hello"
Output: 1
````


## Constraints:
- 0 <= s.length <= 300
- s consists of lowercase and uppercase English letters, digits, or one of the following characters "!@#$%^&*()_+-=',.:".
- The only space character in s is ' '.