# E. Nice Rectangles

## Info

Дана картина N×M из строчных латинских букв, нужно посчитать прямоугольные части, симметричные относительно своего центра.
Ключевая идея - сравнить 2D префиксный хеш прямоугольника с хешем его повёрнутой копии.
Итоговая сложность - O(N^2*M^2) по времени и O(N*M) по памяти.

## Level - medium

## Task

The famous artist Vasya has just finished his new masterpiece and wants to know how much he can get for his work.

The picture is a rectangle N by M centimeters, divided into small 1 by 1 centimeter squares with sides parallel to the sides of the picture.
To achieve harmony, Vasya painted each of these squares in one of 26 special colors denoted by lowercase Latin letters.

The cost of the picture is exactly the number of "nice" parts in it.
A part of the picture is any rectangle that can be cut out of it along the square borders.
A part is "nice" if applying symmetry about its center produces a rectangle colored the same as the original part.
For example, in the picture colored as:

```text
abc
acb
```

the nice parts are all parts consisting of a single square (there are 6), as well as the parts:

```text
bc
cb
```

and:

```text
a
a
```

Write a program that determines the cost of Vasya's masterpiece from the information about it.

### Input

The first line contains two numbers N and M (1 <= N, M <= 100).
The next N lines contain strings of M lowercase Latin letters.
The symbol in the i-th row and j-th column determines the color of the corresponding square of the picture.

### Output

Print the cost of the masterpiece - the number of parts symmetric about their center.

## Объяснение

Прямоугольная часть задаётся строками от r1 до r2 и столбцами от c1 до c2.
Симметрия относительно центра означает, что клетка (r, c) совпадает с клеткой (r1+r2-r, c1+c2-c) для всех клеток части.
Иначе говоря, если повернуть часть на 180 градусов, получится та же самая раскраска.
Всего существует O(N^2*M^2) прямоугольных частей, и каждую из них можно проверить отдельно.
Прямая проверка сравнивает все клетки части, что даёт O(N^3*M^3) в худшем случае и слишком медленно при N, M <= 100.
Быстрее сравнивать не клетки, а хеши: для каждой части считается хеш исходного прямоугольника и хеш того же прямоугольника в повёрнутом изображении.
Если два 2D префиксных хеша совпали, часть считается симпатичной.
Для уменьшения вероятности коллизий используются две независимые пары оснований.

Возможные подходы:

- Bruteforce (полный перебор). Для каждой части сравнивать все пары симметричных клеток, время O(N^3*M^3).
- 2D prefix hash (двумерный префиксный хеш). Сравнивать хеши частей за O(1), время O(N^2*M^2).

## Example 1

```text
Input:
2 3
abc
acb

Output:
8
```

Explanation: the six single squares are nice, the 2×1 part in the first column is nice, and the 2×2 part in the last two columns is nice.

## Constraints

- 1 <= N, M <= 100
- Every cell is a lowercase Latin letter from 'a' to 'z'.
- A part is any axis-aligned subrectangle of the picture.

## См. также

- [Подходы и их сравнение](approachs.md)
- [Bruteforce](solve1_bruteforce)
- [2D prefix hash](solve2_2d_prefix_hash)
