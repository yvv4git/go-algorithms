# B. Cyclic Shift?

## Info

Нужно проверить, является ли вторая строка циклическим сдвигом первой, и найти минимальный сдвиг.
Ключевая идея: искать вторую строку в удвоенной первой с помощью хешей или KMP.
Итоговая сложность O(n) по времени и O(n) по памяти.

## Level - easy

## Task

Kirill once wrote a line of uppercase and lowercase Latin letters, and then went to play football.
When he returned, he found that his friend Dima had written another line of the same length under his.
Dima claims that his line is Kirill's string cyclically shifted several positions to the right.
For example, a cyclic shift of abcde by 2 positions to the right gives deabc.
However, Dima is known to make mistakes in many calculations, so Kirill wonders whether to believe him.
Given these strings, output the minimal possible shift size, or -1 if Dima is mistaken.
Hint: checking that y is a cyclic shift of x means checking that y occurs in xx.

## Объяснение

Пусть n - длина обеих строк.
Сдвиг строки x вправо на k позиций даёт строку x[n-k:] + x[:n-k].
Этот результат равен y тогда и только тогда, когда y встречается в z = x + x начиная с позиции n - k.
Значит, нужно найти вхождения y в z с началом в диапазоне [0, n] и взять максимальную позицию pos.
Ей соответствует минимальный сдвиг k = n - pos.
Если вхождений нет, Дима ошибся и ответ равен -1.
Случай pos = n даёт k = 0, то есть строки совпадают.

Подходы:

- Bruteforce (полный перебор сдвигов). Для каждого k от 0 до n - 1 сравнить y с соответствующим сдвигом x посимвольно.
  Время O(n^2), дополнительная память O(1).
- Rolling hash (полиномиальный хеш). Построить хеши строки x + x и сравнивать окно длины n с y.
  Время O(n), память O(n), но возможны коллизии.
- KMP (алгоритм Кнута - Морриса - Пратта). Найти все вхождения y в x + x за один проход без коллизий.
  Время O(n), память O(n).
- Z-function (Z-функция). Вычислить Z-функцию строки y + "#" + x + x и найти позиции с Z[i] >= n.
  Время O(n), память O(n).

Для ограничения n <= 10000 достаточно даже полного перебора, но хеши и KMP дают линейное решение.

## Example 1

```text
Input:
zabcd
abcdz

Output:
4
```

Explanation: the string abcdz is a cyclic shift of zabcd to the right by 4 positions.

## Constraints

- 1 <= |x| = |y| <= 10000
- The strings consist of uppercase and lowercase Latin letters
- Time limit: 0.3 seconds
- Memory limit: 256 MB
- Input: standard input
- Output: standard output

## См. также

- [Подходы к задаче "Циклический сдвиг"](approachs.md)
- [Bruteforce](solve1_bruteforce/main.go)
- [Rolling hash](solve2_rolling_hash/main.go)
- [KMP](solve3_kmp/main.go)
- [Z-function](solve4_z_function/main.go)
