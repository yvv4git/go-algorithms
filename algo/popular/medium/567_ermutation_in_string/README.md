# 567. Permutation in String


## Level - medium


## Task
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.


## Объяснение
Задача заключается в том, чтобы определить, существует ли в строке s2 любая перестановка строки s1.
Например, если s1 = "ab" и s2 = "eidbaooo", то ответ будет true, потому что s2 содержит перестановку s1 ("ba").

Требования к решению:
1. Входные данные: Две строки s1 и s2.
2. Выходные данные: Логическое значение true, если s2 содержит любую перестановку s1, и false в противном случае.
3. Ограничения: Длина строк s1 и s2 не превышает 100. Строки состоят только из строчных букв английского алфавита.


## Example 1:
````
Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
````


## Example 2:
````
Input: s1 = "ab", s2 = "eidboaoo"
Output: false
````


## Constraints:
- 1 <= s1.length, s2.length <= 10^4
- s1 and s2 consist of lowercase English letters.