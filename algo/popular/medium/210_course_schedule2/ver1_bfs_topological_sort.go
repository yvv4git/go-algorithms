package main

import "fmt"

// Функция findOrder принимает количество курсов и список предварительных условий,
// и возвращает порядок прохождения курсов или пустой массив, если это невозможно.
// Параметр prerequisites - это список курсов, которые надо пройти до этого курса.
// В задаче "Course Schedule II" список prerequisites представляет собой массив пар, где каждая пара [a, b] указывает,
// что курс a зависит от курса b. Цель задачи — найти такой порядок прохождения курсов,
// чтобы все зависимости были соблюдены, или определить, что это невозможно из-за наличия циклов в зависимостях.
func findOrder(numCourses int, prerequisites [][]int) []int {
	/*
		METHOD: Topological sort on BFS
		Мы используем алгоритм топологической сортировки, основанный на BFS (Breadth-First Search). Основные шаги:
		1. Построение графа
		Создаем граф зависимостей и массив входящих степеней для каждой вершины.

		2. Инициализация очереди
		Добавляем в очередь все вершины с нулевой входящей степенью.

		3. Топологическая сортировка
		Пока очередь не пуста, извлекаем вершину, добавляем её в результат и уменьшаем входящую степень для всех её соседей.
		Если у соседа входящая степень становится нулевой, добавляем его в очередь.

		4. Проверка на цикл
		Если размер результата не равен количеству курсов, значит, в графе есть цикл, и возвращаем пустой массив.

		TIME COMPLEXITY: O(V + E), где V — количество вершин (курсов), E — количество ребер (зависимостей).
		Мы посещаем каждую вершину и каждое ребро ровно один раз.

		SPACE COMPLEXITY: O(V + E). Мы храним граф в виде списка смежности и массив входящих степеней для каждой вершины.
	*/
	// Граф зависимостей
	graph := make([][]int, numCourses)
	// Массив входящих степеней для каждой вершины
	inDegree := make([]int, numCourses)

	// Построение графа и подсчет входящих степеней
	for _, pre := range prerequisites {
		course, prereq := pre[0], pre[1]
		graph[prereq] = append(graph[prereq], course)
		inDegree[course]++
	}

	// Очередь для вершин с нулевой входящей степенью
	queue := []int{}
	for course, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, course)
		}
	}

	// Результат топологической сортировки
	result := []int{}

	// Пока очередь не пуста
	for len(queue) > 0 {
		// Извлекаем вершину из очереди
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		// Уменьшаем входящую степень для всех соседей
		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			// Если у соседа входящая степень стала нулевой, добавляем его в очередь
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Если размер результата равен количеству курсов, возвращаем результат.
	// В этом коде проверяется, была ли успешно выполнена топологическая сортировка графа зависимостей курсов.
	// Если размер результата (result) равен количеству курсов (numCourses), это означает,
	// что все курсы были успешно добавлены в результат в правильном порядке, и зависимости были соблюдены.
	// В этом случае возвращается массив result.
	if len(result) == numCourses {
		return result
	}

	// В противном случае возвращаем пустой массив.
	// Однако, если размер результата не равен количеству курсов, это означает, что в графе зависимостей присутствует цикл,
	// и невозможно найти порядок прохождения курсов, который бы удовлетворял всем зависимостям.
	// В таком случае возвращается пустой массив []int{}, что является индикатором наличия цикла и невозможности пройти все курсы в заданных условиях.
	return []int{}
}

func main() {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	result := findOrder(numCourses, prerequisites)
	fmt.Println(result) // Вывод: [0, 1, 2, 3] или [0, 2, 1, 3]
}