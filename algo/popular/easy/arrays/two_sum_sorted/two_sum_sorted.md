# Two sum sorted
Написать оптимальный алгоритм, который может проверить, имеется ли в отсортированном массиве два числа, сумма которых равна третьему.

## Вариант-1
Можно решить задачу с квадратичной сложность. Т.е. два цикла, в одном берем элемент и проходим по всем элементам.

## Вариант-2
Это вариант крутой. Используем закономерность. Если слева число + правое число меньше суммы, то надо правый индекс сдвинуть вправо.
А если сумма этих числе больше целевого, то сдвинуть правый индекс в лево. Получится линейная сложность.
