# B. Comparing Subtrees

## Info

Дано корневое дерево, каждой вершине приписано число, нужно посчитать количество пар изоморфных корневых поддеревьев.
Ключевая идея - построить для каждого поддерева каноническую форму снизу вверх и сгруппировать одинаковые формы.
Итоговая сложность - O(n*log n) по времени и O(n) по памяти.

## Level - medium

## Task

You are given a tree rooted at vertex 1. Each vertex has a number written on it.

Count the number of pairs of isomorphic rooted subtrees.

A rooted subtree consists of a chosen root vertex and all of its descendants.

Two rooted subtrees are isomorphic if there is a one-to-one correspondence between their vertices.
Every vertex of the first subtree must match exactly one vertex of the second subtree such that:

- (a) every vertex of the second subtree is matched by exactly one vertex of the first subtree with the same value;
- (b) the root of the first subtree is matched to the root of the second subtree;
- (c) two vertices of the first subtree are connected by an edge if and only if their matches are connected by an edge.

### Input

The first line contains an integer n (2 <= n <= 10^5) - the number of vertices in the tree.

Each of the next (n-1) lines contains two integers u_i and v_i - the vertices connected by the i-th edge (1 <= u_i, v_i <= n, u_i != v_i).

The last line contains n integers c_i (1 <= c_i <= 10^9) - the values written on the vertices.

The root of the tree is vertex 1.

### Output

Print a single number - the number of pairs of isomorphic rooted subtrees.

## Объяснение

Корневое поддерево вершины v состоит из самой вершины v и всех её потомков.
Изоморфизм поддеревьев - это взаимно однозначное соответствие вершин, которое сохраняет значения, корень и рёбра.
Дети в дереве не упорядочены, поэтому два поддерева изоморфны, если совпадают значения корней и наборы канонических форм их детей.
Каноническую форму можно строить снизу вверх: сначала для листьев, затем для их родителей.
Для каждой вершины нужно отсортировать канонические формы детей и объединить их с собственным значением.
Если одна и та же каноническая форма встретилась cnt раз, то она даёт cnt*(cnt-1)/2 пар изоморфных поддеревьев.

Возможные подходы:

- Bruteforce (попарное сравнение сигнатур). Построить строковые сигнатуры всех поддеревьев и сравнить все пары, время O(n^3) в худшем случае.
- Canonical IDs (канонические идентификаторы). Присваивать каждой новой форме целочисленный идентификатор через хеш-таблицу, время O(n*log n).
- Subtree hash (хеш поддерева). Считать две 64-битные хеш-суммы снизу вверх, время O(n*log n) в среднем.

## Example 1

```text
Input:
4
1 2
1 3
3 4
1 2 3 2

Output:
1
```

Explanation: the leaves 2 and 4 both have value 2 and no children, so they form one isomorphic pair.

## Example 2

```text
Input:
6
1 2
1 3
1 4
3 5
4 6
1 2 3 4 2 2

Output:
3
```

Explanation: the leaves 2, 5, and 6 all have value 2, so they form C(3, 2) = 3 isomorphic pairs.

## Constraints

- 2 <= n <= 10^5
- 1 <= u_i, v_i <= n and u_i != v_i for every edge.
- 1 <= c_i <= 10^9 for every vertex value.
- The root of the tree is vertex 1.

## См. также

- [Подходы и их сравнение](approachs.md)
- [Bruteforce](solve1_bruteforce)
- [Canonical IDs](solve2_canonical)
- [Subtree hash](solve3_hash)
