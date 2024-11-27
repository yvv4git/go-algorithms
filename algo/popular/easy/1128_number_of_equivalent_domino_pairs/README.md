# 1128. Number of Equivalent Domino Pairs


## Level - easy


## Task
Given a list of dominoes, dominoes[i] = [a, b] is equivalent to dominoes[j] = [c, d] if 
and only if either (a == c and b == d), or (a == d and b == c) - that is, one domino can be rotated to be equal to another domino.

Return the number of pairs (i, j) for which 0 <= i < j < dominoes.length, and dominoes[i] is equivalent to dominoes[j].


## Обяснение
У нас есть набор домино, где каждое домино представлено парой чисел (a, b), где a и b — целые числа от 1 до 9. 
Два домино считаются эквивалентными, если они могут быть представлены одинаковой парой чисел, независимо от порядка. 
Например, домино (1, 2) эквивалентно домино (2, 1).

Примеры:
- Домино (1, 2) и (2, 1) считаются эквивалентными.
- Домино (3, 4) и (4, 3) считаются эквивалентными.
- Домино (1, 1) и (1, 1) считаются эквивалентными.

Нам нужно найти количество пар домино, которые являются эквивалентными.


## Example 1:
```
Input: dominoes = [[1,2],[2,1],[3,4],[5,6]]
Output: 1
```


## Example 2:
```
Input: dominoes = [[1,2],[1,2],[1,1],[1,2],[2,2]]
Output: 3
```


## Constraints:
- 1 <= dominoes.length <= 4 * 10^4
- dominoes[i].length == 2
- 1 <= dominoes[i][j] <= 9