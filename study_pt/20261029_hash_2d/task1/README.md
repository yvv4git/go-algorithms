# A. Count Distinct Sets

## Info

Дано N множеств по K элементов, нужно посчитать количество различных множеств в наборе.
Ключевая идея - привести каждое множество к каноническому виду и сложить ключи в хеш-таблицу.
Итоговая сложность - O(N*K*log K) по времени и O(N*K) по памяти.

## Level - medium

## Task

You are given N sets, each containing K elements.
Determine how many distinct sets are contained in the collection.

### Input

The first line contains two integers N and K (1 <= N*K <= 10^6).
Each of the next N lines contains K distinct integers from 0 to 10^9 separated by spaces.

### Output

Print a single number - the number of distinct sets in the collection.

## Объяснение

Множество - это набор чисел, в котором порядок не важен и нет повторяющихся элементов.
Поэтому наборы `1 2 3` и `3 2 1` задают одно и то же множество, а `1 2 3` и `11 12 13` - разные.
Чтобы сравнивать множества быстро, каждое из них нужно привести к одному каноническому виду.
Самый простой канонический вид - отсортировать элементы по возрастанию.
После сортировки одинаковые множества становятся одинаковыми массивами чисел.
Остаётся сложить эти массивы в хеш-таблицу и посчитать количество различных ключей.

Возможные подходы:

- Bruteforce (полный перебор). Отсортировать каждое множество и сравнивать его со всеми предыдущими множествами, время O(N^2*K).
- Sort and hash (сортировка и хеш-таблица). Отсортировать множество, собрать ключ и добавить его в хеш-таблицу, время O(N*K*log K).
- Set hash (хеш множества, не зависящий от порядка). Считать две суммы псевдослучайных значений элементов, время O(N*K) в среднем.

## Example 1

```text
Input:
4 3
1 2 3
3 2 1
11 12 13
2 1 3

Output:
2
```

Explanation: the first, second, and fourth sets equal `{1, 2, 3}`, while the third set is `{11, 12, 13}`.
Therefore there are two distinct sets.

## Constraints

- 1 <= N*K <= 10^6
- Each set contains exactly K distinct integers.
- 0 <= a_i <= 10^9 for every element a_i.

## См. также

- [Подходы и их сравнение](approachs.md)
- [Bruteforce](solve1_bruteforce)
- [Sort and hash](solve2_sort_hash)
- [Set hash](solve3_set_hash)
