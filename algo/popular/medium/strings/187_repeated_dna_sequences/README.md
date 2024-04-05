# 187. Repeated DNA Sequences


## Level - medium


## Task
The DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T'.
- For example, "ACGAATTCCG" is a DNA sequence.

When studying DNA, it is useful to identify repeated sequences within the DNA.

Given a string s that represents a DNA sequence, return all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule. 
You may return the answer in any order.


## Объяснение
Задача состоит в том, чтобы найти все повторяющиеся последовательности ДНК длиной 10 символов в строке ДНК. 
Последовательность ДНК состоит из четырех типов символов: 'A', 'C', 'G' и 'T'.

Пример:
````
Вход: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
Выход: ["AAAAACCCCC", "CCCCCAAAAA"]
````

В этом примере "AAAAACCCCC" и "CCCCCAAAAA" являются повторяющимися последовательностями ДНК.

Основная идея решения заключается в использовании хеш-таблицы для отслеживания повторяющихся последовательностей ДНК. 
Мы проходим по строке ДНК и для каждой последовательности длины 10 символов вычисляем ее хеш-значение. 
Если такой хеш-значение уже есть в хеш-таблице, то мы добавляем последовательность в результат. 
В противном случае, мы добавляем хеш-значение в хеш-таблицу.


## Example 1:
````
Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
Output: ["AAAAACCCCC","CCCCCAAAAA"]
````


## Example 2:
````
Input: s = "AAAAAAAAAAAAA"
Output: ["AAAAAAAAAA"]
````


## Constraints:
- 1 <= s.length <= 10^5
- s[i] is either 'A', 'C', 'G', or 'T'.