# 210. Course Schedule II


## Level - medium


## Task
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. 
You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

- For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. 
If it is impossible to finish all courses, return an empty array.


## Объяснение
Задача на LeetCode требует от нас найти порядок, в котором можно пройти все курсы в университете, учитывая их зависимости. 
Это классическая задача на топологическую сортировку графа.

Постановка задачи:  
У вас есть n курсов, пронумерованных от 0 до n-1. 
Некоторые курсы имеют предварительные условия, представленные в виде списка пар [a, b], где b должен быть пройден перед курсом a.

Необходимо определить порядок прохождения курсов, который позволит пройти все курсы, или определить, 
что это невозможно (в случае наличия цикла в зависимости курсов).

Пример 1:
````
Вход: numCourses = 2, prerequisites = [[1,0]]
Выход: [0,1]
Объяснение: Есть два курса. Чтобы пройти курс 1, нужно сначала пройти курс 0. Таким образом, правильный порядок — [0,1].
````

Пример 2:
````
Вход: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Выход: [0,1,2,3] или [0,2,1,3]
Объяснение: Есть четыре курса. Чтобы пройти курс 1 и 2, нужно сначала пройти курс 0. Чтобы пройти курс 3, нужно сначала пройти курсы 1 и 2. Возможные правильные порядки — [0,1,2,3] или [0,2,1,3].
````

Пример 3:
````
Вход: numCourses = 1, prerequisites = []
Выход: [0]
Объяснение: Есть только один курс, и он не имеет предварительных условий. Таким образом, правильный порядок — [0].
````


## Example 1:
````
Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. 
So the correct course order is [0,1].
````


## Example 2:
````
Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].
````


## Example 3:
````
Input: numCourses = 1, prerequisites = []
Output: [0]
````


## Constraints:
- 1 <= numCourses <= 2000
- 0 <= prerequisites.length <= numCourses * (numCourses - 1)
- prerequisites[i].length == 2
- 0 <= ai, bi < numCourses
- ai != bi
- All the pairs [ai, bi] are distinct.