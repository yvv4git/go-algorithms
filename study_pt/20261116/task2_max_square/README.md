# 2. Large Square

## Info

Город - прямоугольник N x M с вертикальными и горизонтальными улицами, под офис нужен квадратный участок без улиц.
Ключевая идея - разбить город на ячейки между соседними улицами и вычислить минимум из максимальных зазоров по x и y.
Итоговая сложность - O((Vert+Hor)*log(Vert+Hor)) по времени и O(Vert+Hor) по памяти.

## Level - medium

## Task

The city is a rectangle of size N x M with vertical and horizontal streets.
A street is a line: vertical streets are given by x-coordinates, horizontal streets by y-coordinates.
You can buy any square plot that does not intersect the streets and lies completely inside the city.
Find the maximum side length of such a square.

### Input

The first line contains two integers N and M (1 <= N, M <= 10^9) - the width and the height of the city.
The second line contains an integer Vert (0 <= Vert <= 100000) - the number of vertical streets.
The third line contains Vert integers x_i (0 <= x_i <= N) - the coordinates of the vertical streets.
This line may be empty if Vert = 0.
The fourth line contains an integer Hor (0 <= Hor <= 100000) - the number of horizontal streets.
The fifth line contains Hor integers y_i (0 <= y_i <= M) - the coordinates of the horizontal streets.
This line may be empty if Hor = 0.
Street coordinates are given in arbitrary order.

### Output

Print a single number - the maximum side length of a square that does not intersect any street.

## Объяснение

Улица - это линия нулевой ширины: вертикальная улица стоит на координате x, горизонтальная - на координате y.
Квадратный участок не должен пересекать эти линии и должен целиком лежать в городе.
Значит, квадрат обязан поместиться в одной ячейке - прямоугольнике между двумя соседними вертикальными улицами и двумя соседними горизонтальными.
Границы города x = 0, x = N, y = 0 и y = M тоже считаются границами ячеек, хотя улицами не являются.

Чтобы найти ячейки, отсортируем координаты вертикальных улиц и допишем к ним границы 0 и N.
Разности соседних отсортированных координат - это ширины ячеек по x, максимальная из них назовём maxX.
Точно так же по горизонтали получаем maxY.
Сторона квадрата не может превышать ни maxX, ни maxY, и равна min(maxX, maxY).

Проверим на примере N = 10, M = 10, вертикальная улица x = 3, горизонтальная y = 5.
По x координаты с границами: 0, 3, 10, зазоры 3 и 7, значит maxX = 7.
По y координаты с границами: 0, 5, 10, зазоры 5 и 5, значит maxY = 5.
Ответ min(7, 5) = 5, как в первом примере.

Порядок улиц на входе не важен, потому что перед подсчётом зазоров координаты всё равно сортируются.

Возможные подходы:

- Sort gaps (сортировка и зазоры). Отсортировать координаты с границами и взять максимальный зазор по каждой оси, время O((Vert+Hor)*log(Vert+Hor)).
- Binary search (бинарный поиск по ответу). Искать сторону k и проверять зазор не меньше k по обеим осям, время O((Vert+Hor)*log(Vert+Hor)+log(min(N,M))*(Vert+Hor)).

## Example 1

```text
Input:
10 10
1
3
1
5

Output:
5
```

Explanation: the vertical street is at x = 3, so the vertical gaps are 3 and 7.
The horizontal street is at y = 5, so the horizontal gaps are 5 and 5.
The maximum square side is min(7, 5) = 5.

## Example 2

```text
Input:
4 7
0

0

Output:
4
```

Explanation: there are no streets at all.
The whole city is one cell of size 4 x 7, so the answer is min(4, 7) = 4.

## Example 3

```text
Input:
5 7
2
5 0
3
3 1 2

Output:
4
```

Explanation: the vertical streets are at x = 0 and x = 5, so the vertical gaps are 0 and 5, and maxX = 5.
The horizontal streets are at y = 1, 2, 3, so the horizontal gaps are 1, 1, 1, 4, and maxY = 4.
The answer is min(5, 4) = 4.

## Constraints

- 1 <= N, M <= 10^9
- 0 <= Vert, Hor <= 100000
- 0 <= x_i <= N and 0 <= y_i <= M
- Street coordinates are given in arbitrary order.
- Street coordinate lines may be empty when the corresponding count is 0.

## См. также

- [Подходы и их сравнение](approachs.md)
- [Sort gaps](solve1_sort_gaps)
- [Binary search](solve2_binary_search)
