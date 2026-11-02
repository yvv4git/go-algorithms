# E. Maximum Substring

## Info

Нужно найти лексикографически максимальную непрерывную подстроку строки.
Ключевая идея: максимальная подстрока всегда является суффиксом, поэтому достаточно найти максимальный суффикс.
Итоговая сложность O(n) по времени и O(1) по памяти для алгоритма Дюваля.

## Level - medium

## Task

Given a string S consisting of lowercase Latin letters.
Find its lexicographically maximum contiguous substring.
That is the substring that comes last when all substrings are sorted in alphabetical order.

## Объяснение

Подстрока - это непрерывный кусок строки.
Если подстрока S[i..j] не доходит до конца строки, то S[i..j+1] длиннее и имеет её своим префиксом, поэтому она лексикографически больше.
Значит, максимальная подстрока всегда является суффиксом S[k..n-1].
Остаётся найти лексикографически максимальный суффикс строки.
Его можно найти перебором, суффиксным массивом, хешами с бинпоиском или алгоритмом Дюваля.

Подходы:

- Bruteforce (перебор суффиксов). Сравнить все суффиксы напрямую и взять максимальный; время O(n^2), память O(1).
- Rolling hash and binary search (хеши и бинпоиск). Сравнивать суффиксы за O(log n) по хешам; время O(n log n), память O(n), возможны коллизии.
- Suffix array (суффиксный массив). Построить суффиксный массив и взять последний суффикс; время O(n log n), память O(n).
- Duval (алгоритм Дюваля). Найти максимальный суффикс за один проход; время O(n), память O(1).

Для n <= 10^6 лучше всего подходит алгоритм Дюваля: он линейный и не требует дополнительной памяти.

## Example 1

```text
Input:
abacaba

Output:
caba
```

Explanation: the maximum substring is the suffix "caba".

## Example 2

```text
Input:
babb

Output:
bb
```

Explanation: the maximum substring is the suffix "bb".

## Constraints

- 1 <= |S| <= 10^6
- The string consists of lowercase Latin letters
- Time limit: 3 seconds
- Memory limit: 256 MB
- Input: standard input
- Output: standard output

## См. также

- [Подходы к задаче "Максимальная подстрока"](approachs.md)
- [Bruteforce](solve1_bruteforce/main.go)
- [Rolling hash and binary search](solve2_rolling_hash_binary_search/main.go)
- [Suffix array](solve3_suffix_array/main.go)
- [Duval](solve4_duval/main.go)
