# 3954. Sum of Compatible Numbers in Range I

## Level - easy

## Task

You are given two integers n and k.

A positive integer x is called compatible if it satisfies both of the following conditions:

abs(n - x) \<= k
(n & x) == 0
Return the sum of all compatible integers x.

## Note

Here, & denotes the bitwise AND operator.
The absolute difference between integers i and j is defined as abs(i - j)

## Объяснение

Даны два целых числа n и k. Нужно найти сумму всех положительных целых x, удовлетворяющих двум условиям:

1. `abs(n - x) <= k` — x находится в диапазоне [max(1, n-k), n+k].
1. `(n & x) == 0` — побитовое AND между n и x равно нулю, то есть у n и x не должно быть единиц в одних и тех же битовых позициях.

Диапазон поиска ограничен значением k, поэтому достаточно перебрать x от n-k до n+k (но не меньше 1).

## Подходы

### Подход 1: Полный перебор (Brute Force)

- Перебираем все x от max(1, n-k) до n+k.
- Для каждого x проверяем `(n & x) == 0`.
- Если условие выполняется, добавляем x к сумме.
- Сложность: O(k) по времени, O(1) по памяти.

### Подход 2: Перебор по битам (Bitwise Enumeration)

- Находим все битовые позиции, где в n стоит 0. В этих позициях x может иметь 0 или 1.
- В позициях, где в n стоит 1, x обязан иметь 0.
- Рекурсивно или итеративно перебираем все комбинации битов, формируя кандидатов x.
- Для каждого кандидата проверяем условие `abs(n - x) <= k`.
- Сложность: O(2^m) по времени (где m — количество нулевых бит в n), O(1) по памяти.

## Example 1

```text
Input: n = 2, k = 3

Output: 10

Explanation:

The compatible integers are:

x = 1, since abs(2 - 1) = 1 and 2 & 1 = 0.
x = 4, since abs(2 - 4) = 2 and 2 & 4 = 0.
x = 5, since abs(2 - 5) = 3 and 2 & 5 = 0.
Thus, the answer is 1 + 4 + 5 = 10.
```

## Example 2

```text
Input: n = 5, k = 1

Output: 0

Explanation:

There are no compatible integers in the range [4, 6]. Thus, the answer is 0.
```

## Constraints

- 1 \<= n \<= 100
- 1 \<= k \<= 100
