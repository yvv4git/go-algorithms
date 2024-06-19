# 207. Course Schedule


## Level - medium


## Task
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. 
You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

- For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

Return true if you can finish all courses. Otherwise, return false.


## Объяснение
Задача заключается в определении возможности завершить все курсы в учебном плане. 
В этой задаче вам предоставляется количество курсов numCourses и список предварительных требований в виде пар [a, b], 
где b должен быть завершен перед курсом a.

Основная проблема заключается в том, чтобы определить, существует ли циклическая зависимость между курсами. 
Если такая зависимость существует, то невозможно завершить все курсы. 
Это эквивалентно задаче поиска цикла в ориентированном графе, где каждый курс представляет собой вершину, 
а каждая пара [a, b] представляет собой ориентированное ребро от b к a.

Для решения этой задачи можно использовать различные алгоритмы, такие как:
- Топологическая сортировка
Этот метод используется для нахождения цикла в графе. 
Если граф содержит цикл, то топологическая сортировка невозможна.
- Поиск в глубину (DFS)
Можно использовать DFS для обхода графа и проверки наличия цикла. 
Если во время обхода мы встречаем уже посещенную вершину в текущем пути, это означает наличие цикла.

Пример входных данных:
````
numCourses = 2
prerequisites = [[1, 0]]
````
Это означает, что есть два курса и курс 0 должен быть завершен перед курсом 1. 
В этом случае ответ true, так как нет циклических зависимостей.


Пример с циклом:
````
numCourses = 2
prerequisites = [[1, 0], [0, 1]]
````
Здесь курс 0 зависит от курса 1, и наоборот, что создает цикл. В этом случае ответ false.

Таким образом, основная цель задачи — определить, можно ли завершить все курсы без циклических зависимостей.


## Example 1:
````
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So it is possible.
````


## Example 2:
````
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.
````


## Constraints:
- 1 <= numCourses <= 2000
- 0 <= prerequisites.length <= 5000
- prerequisites[i].length == 2
- 0 <= ai, bi < numCourses
- All the pairs prerequisites[i] are unique.