# C. String Base

## Info

Нужно найти минимально возможную длину строки S, из повторений которой получен данный префикс.
Ключевая идея: эта длина равна n - pi[n-1], где pi - префикс-функция строки.
Итоговая сложность O(n) по времени и O(n) по памяти.

## Level - easy

## Task

A string S was written several times in a row, and then a prefix was taken from the resulting string.
You are given that prefix, and your task is to determine the minimum possible length of the original string S.

## Объяснение

Пусть t - входная строка длины n.
Будем искать минимальный период p, то есть такую длину, что t является префиксом повторений строки t[:p].
Префикс-функция pi[i] - это длина наибольшего собственного префикса t[0..i], который одновременно является суффиксом t[0..i].
Если последняя граница строки имеет длину b = pi[n-1], то период равен p = n - b.
Совпадение префикса и суффикса длины b означает, что t[i] = t[i - p] для всех i >= p, то есть t продолжается с периодом p.
Большая граница даёт меньший период, а pi[n-1] - максимальная граница, поэтому ответ равен n - pi[n-1].
Например, для t = "bcabcab" имеем pi[n-1] = 4, значит минимальная длина S равна 7 - 4 = 3.

Подходы:

- Bruteforce (перебор периода). Для каждого p проверить t[i] = t[i mod p]; первый подходящий p - ответ, время O(n^2), память O(1).
- Prefix function (префикс-функция). Ответ равен n - pi[n-1], где pi - префикс-функция строки; время O(n), память O(n).
- Z-function (Z-функция). Найти минимальное p, для которого z[p] >= n - p; время O(n), память O(n).
- Rolling hash (полиномиальный хеш). Проверять равенство блоков по хешу для каждого p; время O(n), память O(n), возможны коллизии.

Для ограничения n <= 50000 подходит любой линейный подход, но проще всего префикс-функция.

## Example 1

```text
Input:
zzz

Output:
1
```

Explanation: S = "z", because "zzz" is a prefix of "zzzz...".

## Example 2

```text
Input:
bcabcab

Output:
3
```

Explanation: S = "bca", because "bcabcab" is a prefix of "bcabcabc...".

## Constraints

- 1 <= |S| <= 50000
- The string consists of Latin letters
- Time limit: 1 second
- Memory limit: 256 MB
- Input: standard input or input.txt
- Output: standard output or output.txt

## См. также

- [Подходы к задаче "Основание строки"](approachs.md)
- [Bruteforce](solve1_bruteforce/main.go)
- [Prefix function](solve2_prefix_function/main.go)
- [Z-function](solve3_z_function/main.go)
- [Rolling hash](solve4_rolling_hash/main.go)
