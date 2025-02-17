# 944. Delete Columns to Make Sorted


## Level - easy


## Task
You are given an array of n strings strs, all of the same length.

The strings can be arranged such that there is one on each line, making a grid.

For example, strs = ["abc", "bce", "cae"] can be arranged as follows:
- abc
- bce
- cae

You want to delete the columns that are not sorted lexicographically. In the above example (0-indexed), columns 0 ('a', 'b', 'c') and 2 ('c', 'e', 'e') are sorted, while column 1 ('b', 'c', 'a') is not, so you would delete column 1.

Return the number of columns that you will delete.


## Объяснение
Вам дан массив строк strs, где каждая строка имеет одинаковую длину. 
Эти строки можно представить как столбцы в таблице. 

Для strs = ["abc", "bce", "cae"]:
- Первый столбец: a, b, c — отсортирован.
- Второй столбец: b, c, a — не отсортирован.
- Третий столбец: c, e, e — отсортирован.

Нужно удалить только второй столбец, поэтому ответ: 1.

Решение:
- Проходим по каждому столбцу (по индексам символов в строках).
- Для каждого столбца проверяем, отсортированы ли символы сверху вниз.
- Если столбец не отсортирован, увеличиваем счетчик удаляемых столбцов.
- Возвращаем общее количество удаляемых столбцов.