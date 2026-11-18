# A. Independent Set

## Info

Дан неориентированный граф, и нужно проверить, существует ли в нём независимое множество заданного размера S.
Независимое множество - это набор вершин, никакие две из которых не соединены ребром.
Ключевая идея - перебирать вершины с отсечениями, а итоговая сложность составляет O(2^(V/2) + V + E) по времени и памяти.

## Level - medium

## Task

An undirected graph without self-loops but possibly with multiple edges is given.
Check whether this graph contains an independent set of the given size.
An independent set is a set of vertices such that no two of them are connected by an edge.

### Input

The first line contains two integers V and S - the number of vertices and the size of the required independent set.
Both values satisfy 1 <= V <= 50 and 1 <= S <= 50.
The second line contains an integer E - the number of edges in the graph (0 <= E <= 1000).
Each of the next E lines contains a pair of vertices b_i, e_i describing an edge (1 <= b_i <= V, 1 <= e_i <= V, b_i != e_i).

### Output

Print YES if the graph contains an independent set of size S.
Otherwise print NO.

## Объяснение

Независимое множество - это такой набор вершин, внутри которого нет ни одного ребра.
Иными словами, нельзя выбрать две вершины из набора, соединённые ребром.
Кратные рёбра не усложняют задачу: они не меняют факт смежности двух вершин, а петли в условии запрещены.

Если в графе есть независимое множество размера больше S, то из него можно взять любое подмножество ровно из S вершин.
Поэтому проверка "есть ли множество размера хотя бы S" даёт тот же ответ, что и проверка ровно размера S.

Например, в первом примере вершины 1, 3 и 5 не соединены рёбрами друг с другом, поэтому они образуют независимое множество размера 3.
Во втором примере максимальное независимое множество имеет размер 3, поэтому множества размера 4 не существует.

Возможные подходы:

- Bruteforce (перебор подмножеств). Перебрать все наборы из S вершин. Время O(C(V, S) * S^2), память O(V + E).
- Backtracking with pruning (перебор с отсечениями). Отбрасывать ветки, где не хватает вершин до S. Время O(2^V), память O(V + E).
- Meet-in-the-middle (встречный перебор). Перебирать две половины вершин отдельно. Время O(2^(V/2) + V + E), память O(2^(V/2) + V + E).

## Example 1

```text
Input:
5 3
3
1 2
3 4
4 5

Output:
YES
```

Explanation: vertices 1, 3 and 5 are pairwise non-adjacent, so they form an independent set of size 3.

## Example 2

```text
Input:
5 4
3
1 2
3 4
4 5

Output:
NO
```

Explanation: the largest independent set in this graph has size 3, so an independent set of size 4 does not exist.

## Constraints

- 1 <= V <= 50
- 1 <= S <= 50
- 0 <= E <= 1000
- 1 <= b_i <= V
- 1 <= e_i <= V
- b_i != e_i
- Multiple edges between the same pair of vertices are allowed.

## См. также

- [Подходы и их сравнение](approachs.md)
- [Bruteforce](solve1_bruteforce)
- [Backtracking with pruning](solve2_backtracking)
- [Meet-in-the-middle](solve3_meet_in_the_middle)
