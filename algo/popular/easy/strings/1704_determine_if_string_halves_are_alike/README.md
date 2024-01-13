# 1704. Determine if String Halves Are Alike


## Level - easy


## Task
You are given a string s of even length. Split this string into two halves of equal lengths, 
and let a be the first half and b be the second half.

Two strings are alike if they have the same number of vowels ('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'). 
Notice that s contains uppercase and lowercase letters.

Return true if a and b are alike. Otherwise, return false.


## Объяснение

Задача заключается в том, чтобы определить, совпадают ли гласные буквы в двух половинах строки.
Гласными в английском языке считаются буквы "a", "e", "i", "o", "u", а также их заглавные аналоги "A", "E", "I", "O", "U".

Например, если у нас есть строка "book", то мы разделяем ее на две половины - "bo" и "ok". 
В обеих половинах есть одинаковое количество гласных букв ("o"), поэтому функция должна вернуть true.

Если у нас есть строка "textbook", то мы разделяем ее на две половины - "tex" и "book". 
В первой половине есть одна гласная буква ("e"), а во второй - две ("o" и "o"). 
Поэтому функция должна вернуть false.

Задача проверяет, можно ли разделить строку на две равные по длине части так, чтобы в обеих частях было одинаковое количество гласных букв.

## Example 1:
````
Input: s = "book"
Output: true
Explanation: a = "bo" and b = "ok". a has 1 vowel and b has 1 vowel. Therefore, they are alike.
````


## Example 2:
````
Input: s = "textbook"
Output: false
Explanation: a = "text" and b = "book". a has 1 vowel whereas b has 2. Therefore, they are not alike.
Notice that the vowel o is counted twice.
````


## Constraints:
- 2 <= s.length <= 1000
- s.length is even.
- s consists of uppercase and lowercase letters.