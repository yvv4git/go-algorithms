# 467. Unique Substrings in Wraparound String


## Level - Medium


## Task
We define the string base to be the infinite wraparound string of "abcdefghijklmnopqrstuvwxyz", so base will look like this:

"...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd....".
Given a string s, return the number of unique non-empty substrings of s are present in base.


## Объяснение
### Суть задачи:
Дана строка s и бесконечная циклическая строка base, которая представляет собой бесконечно повторяющийся алфавит "abcdefghijklmnopqrstuvwxyz...". Необходимо подсчитать количество уникальных непустых подстрок строки s, которые присутствуют в строке base.

Подстрока считается валидной для base, если символы в ней идут подряд по алфавиту (с учётом цикла: после z идёт a).

### Подходы к решению:
1. Полный перебор (Brute Force) — сгенерировать все подстроки s и проверить каждую в base. Имеет сложность O(n^2) по времени и O(n^2) по памяти, что неэффективно для n=10^5.
2. Динамическое программирование (DP) — для каждого символа строки вычисляется длина максимальной непрерывной последовательности, заканчивающейся на этом символе. 
Затем суммируются все эти длины. Сложность O(n) по времени и O(1) по памяти.
3. Оптимизированный DP с массивом — хранится массив длиной 26 для каждой буквы, отслеживающий максимальную длину последовательности, заканчивающейся на эту букву. 
Это позволяет учесть повторяющиеся символы.


## Example 1:
```
Input: s = "a"
Output: 1
Explanation: Only the substring "a" of s is in base.
```


## Example 2:
```
Input: s = "cac"
Output: 2
Explanation: There are two substrings ("a", "c") of s in base.
```

## Example 3:
```
Input: s = "zab"
Output: 6
Explanation: There are six substrings ("z", "a", "b", "za", "ab", and "zab") of s in base.
```


## Constraints:
- 1 <= s.length <= 10^5
- s consists of lowercase English letters.