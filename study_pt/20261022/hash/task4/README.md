# D. Distinct Substrings

## Info

Нужно посчитать количество различных подстрок строки, то есть различных непрерывных кусков.
Ключевая идея: подстроки можно перебирать, хранить в суффиксном боре или считать по суффиксному автомату.
Для n <= 1024 подойдут решения за O(n^2), а суффиксный автомат даёт O(n).

## Level - medium

## Task

Given a string S of N characters.
A substring S[i..j] is the string from the i-th to the j-th character inclusive (i <= j).
Your task is to count the number of distinct substrings of the given string.

## Объяснение

Подстрока S[i..j] - это непрерывный кусок строки с i-го по j-й символ включительно.
Всего подстрок с учётом повторов n(n+1)/2.
Различные подстроки - это те, которые отличаются хотя бы одним символом.
В суффиксном боре каждая подстрока соответствует пути от корня, поэтому число различных подстрок равно числу вершин без корня.
В суффиксном автомате каждое состояние отвечает за группу подстрок, и вклад состояния v равен len[v] - len[link[v]].
В суффиксном массиве число различных подстрок равно n(n+1)/2 минус сумма LCP соседних суффиксов.
При n <= 1024 можно даже перебрать все подстроки и сложить их хеши в множество.

Подходы:

- Bruteforce (перебор подстрок и множество хешей). Вычислить хеш каждой подстроки и добавить в множество; время O(n^2), память O(n^2).
- Suffix trie (суффиксный бор). Вставить все суффиксы в бор; ответ равен числу вершин без корня, время O(n^2), память O(n^2).
- Suffix array and LCP (суффиксный массив и LCP). Вычесть сумму LCP соседних суффиксов из n(n+1)/2; время O(n^2 log n), память O(n).
- Suffix automaton (суффиксный автомат). Суммировать len[v] - len[link[v]] по состояниям; время O(n), память O(n).

Для n <= 1024 подходит любой подход, но суффиксный автомат самый быстрый.

## Example 1

```text
Input:
abc

Output:
6
```

Explanation: the distinct substrings are "a", "b", "c", "ab", "bc", "abc".

## Example 2

```text
Input:
aaa

Output:
3
```

Explanation: the distinct substrings are "a", "aa", "aaa".

## Constraints

- 1 <= |S| <= 1024
- The string consists of lowercase Latin letters
- Time limit: 1.5 seconds
- Memory limit: 256 MB
- Input: standard input or input.txt
- Output: standard output or output.txt

## См. также

- [Подходы к задаче "Количество различных подстрок"](approachs.md)
- [Bruteforce](solve1_bruteforce/main.go)
- [Suffix trie](solve2_suffix_trie/main.go)
- [Suffix array and LCP](solve3_suffix_array_lcp/main.go)
- [Suffix automaton](solve4_suffix_automaton/main.go)
