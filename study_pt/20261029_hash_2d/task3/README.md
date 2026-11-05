# C. Security System

## Info

Дано чёрно-белое изображение N×M и K контрольных элементов размером L×L, нужно посчитать число вхождений каждого элемента.
Ключевая идея - посчитать 2D скользящий хеш всех окон L×L и сравнить их с хешами элементов.
Итоговая сложность - O(N*M + K*L^2) по времени и O(N*M + K*L^2) по памяти.

## Level - hard

## Task

"PermOlympBank" has developed a highly reliable security system for valuable papers.
The upper right corner of every paper must contain a rectangular black-and-white image.
A library of control elements was created to verify document authenticity.
If a document is genuine, then every control element from the library occurs in its image a specified number of times.
The library contains no duplicate elements.

After scanning, the control system receives the image as a digital N×M array, where digit 1 means black and digit 0 means white.
Then the system searches for the control elements in the received array.

A control element is represented by an L×L array of digits, each equal to 0 or 1.
The library contains K control elements.
A library element must match some part of the image exactly.
Rotations are not allowed when comparing the image and control elements.

Determine how many times each library element occurs in the image.

### Input

The first line contains the integers N, M, K, L separated by spaces (N <= 100000, M <= 1000, L <= 50, K <= 20).

Then K blocks follow, one per library element.
Each block consists of L lines with L digits (zero or one).
Each block is followed by an empty line.

The next N lines contain M digits each and describe the image.

### Output

The output must contain K lines.

Each line contains two numbers: the index of the control element in the library and the number of its occurrences (0 if the element is not found).

The index of the first element is 1, the second is 2, and so on.
All numbers in a line must be separated by a space.

## Объяснение

Изображение - это матрица из нулей и единиц, а контрольный элемент - квадратная матрица L×L из нулей и единиц.
Элемент встречается в изображении, если он совпадает с некоторым квадратным окном L×L без поворотов и отражений.
Нужно проверить все возможные положения окна: левый верхний угол может находиться в любой позиции от (0, 0) до (N-L, M-L).
Прямое сравнение каждого окна с каждым элементом требует O(K*N*M*L^2) операций, что при N*M = 10^8 и L = 50 слишком медленно.
Вместо этого можно один раз посчитать хеш каждого окна и сравнить его с хешами контрольных элементов.
Двумерный скользящий хеш позволяет обновлять хеш окна при сдвиге на одну строку или один столбец за O(1).
Чтобы избежать случайных совпадений, используются две независимые 64-битные хеш-функции.

Возможные подходы:

- Bruteforce (полный перебор). Для каждого элемента и каждого окна сравнивать все L×L клеток, время O(K*N*M*L^2).
- 2D rolling hash (двумерный скользящий хеш). Один раз пройти по изображению и обновлять хеши окон, время O(N*M + K*L^2).

## Example 1

```text
Input:
10 10 2 3
010
111
010

100
011
001

0010010010
0100001111
0000000100
0000100000
0001111000
0000101000
0100000010
0001000110
0011111000
0101000000

Output:
1 2
2 1
```

Explanation: the first control element occurs twice, and the second control element occurs once.

## Constraints

- 1 <= N <= 100000
- 1 <= M <= 1000
- 1 <= L <= 50
- 1 <= K <= 20
- Every cell of the image and every control element is 0 or 1.
- Rotations and reflections are not allowed.

## См. также

- [Подходы и их сравнение](approachs.md)
- [Bruteforce](solve1_bruteforce)
- [2D rolling hash](solve2_2d_rolling_hash)
