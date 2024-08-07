# DFS


## INFO
DFS позволяет понять, что граф связный. Это одно из применений dfs.
Можно просто проверить, что все вершины попали в массив результирующий.

Как запустить dfs из всех вершин графа?
Пройдемся по всем вершинам, от 0 до n-1.
Если какая-то очередная вершинка не использована, то запускаем на ней dfs.
Тем самым мы можем запустить dfs на всех не связанных графах и тем самым понять
количество компонент в графе.

Краеугольный камень алгоритмов на графах - а как хранить граф.
Можно хранить в матрице смежности. 
Можно в словаре ребер. Это самый лучший способ хранения.

## Как можно хранить граф смежности
1. Матрица смежности. Это просто храним табличку NxN, где 1 значит есть ребро, а 0 - нет ребра.
Временная сложность прохода по вершинам - O(n^2). 
Так как для каждой вершины сделаем проход, а какие же ребра ведут из вершины.
2. Список ребер.
Временная сложность прохода по вершинам - O(n*m). Но из каждой вершины, где запустим dfs, придется пройтись
по всем ребрам.
3. Списки смежности.
Временная сложность будет O(n+m). Смотрим на соседей ровно столько раз, сколько у вершины есть соседей.
Это наиболее оптимальный способ хранения. Но если граф полный, тогда проще хранить в матрице смежности. Так как она
хранит только бит,а список смежности хранит числа.


Если используем рекурсию, то есть неявные затраты. На стек фреймах будут храниться все вершины.
Это займет памяти O(n). Но dfs жестко эксплуатирует стек рекурсии и про это не стоит запоминать.
Можно написать не рекурсивный DFS. Можно в стеке явным образом хранить вершины, в которых находимся.



## Задачи для DFS
1. Топологическая сортировка.


## Связанность в ориентированном графе.
1. Слабая.
Это связность в неориентированном графе. Он разбивается на компоненты связанности.
2. Сильная связанность.
Граф сильно связан, если из любой вершины можно добраться в любую другую.

Все вершинки графа разбиваются на компоненты сильной связанности.
А между компонентами сильной связанности можно в некотором порядке ходить.
Это называется конденсация графа. Это когда вершины разбивают на компоненты сильной связанности.
Эти компоненты объединяются в единую вершину.

Одно из важнейших свойств конденсата - ацикличность.



## Как найти компоненты сильной связанности?
Берем граф и делаем топологическую сортировку.
В графе могут быть циклы, но мы все равно делаем топологическую сортировку.
Получаем какой-то порядок, это не топологическая сортировка. Но получим какой-то порядок.
Потом идем по вершинам в порядке обратной топологической сортировки. 