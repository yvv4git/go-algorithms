# 557. Reverse Words in a String III


## Level - easy


## Task
Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.


## Объяснение
Данная задача требует написать функцию, которая будет принимать строку в качестве входных данных и возвращать новую строку, 
в которой каждое слово будет обращено. Слова в строке разделяются одним или несколькими пробелами.

Например, если входные данные будут "Let's take LeetCode contest", то выходные данные должны быть "s'teL ekat edoCteeL tsetnoc".

Обратите внимание, что обращение каждого слова должно быть выполнено независимо, 
то есть порядок символов внутри каждого слова должен быть обращен, а не порядок слов в строке.


## Example 1:
````
Input: s = "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
````


## Example 2:
````
Input: s = "Mr Ding"
Output: "rM gniD"
````


## Constraints:
- 1 <= s.length <= 5 * 10^4
- s contains printable ASCII characters.
- s does not contain any leading or trailing spaces.
- There is at least one word in s.
- All the words in s are separated by a single space.