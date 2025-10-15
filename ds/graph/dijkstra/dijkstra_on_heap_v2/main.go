package main

import (
	"container/heap"
	"fmt"
)

// Edge представляет ребро графа
type Edge struct {
	To     string // Вершина, в которую ведет ребро
	Weight int    // Вес ребра
}

// Graph представляет граф в виде списка смежности
type Graph map[string][]Edge

// Item представляет элемент очереди с приоритетами
type Item struct {
	Vertex string // Вершина
	Dist   int    // Текущее расстояние до вершины
	index  int    // Индекс элемента в очереди (используется для обновления)
}

// PriorityQueue реализует очередь с приоритетами
type PriorityQueue []*Item

// Len возвращает длину очереди
func (pq PriorityQueue) Len() int { return len(pq) }

// Less сравнивает два элемента очереди по расстоянию
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Dist < pq[j].Dist
}

// Swap меняет местами два элемента очереди
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push добавляет элемент в очередь
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop извлекает элемент из очереди
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// update модифицирует элемент очереди с приоритетами
func (pq *PriorityQueue) update(item *Item, vertex string, dist int) {
	item.Vertex = vertex
	item.Dist = dist
	heap.Fix(pq, item.index)
}

// Dijkstra реализует алгоритм Дейкстры
func Dijkstra(graph Graph, start string) map[string]int {
	dist := make(map[string]int)            // Массив расстояний до вершин
	priorityQueue := make(PriorityQueue, 0) // Очередь с приоритетами
	heap.Init(&priorityQueue)

	// Инициализация: устанавливаем расстояние до начальной вершины в 0,
	// а до всех остальных вершин в максимальное значение
	for vertex := range graph {
		if vertex == start {
			dist[vertex] = 0
			heap.Push(&priorityQueue, &Item{Vertex: vertex, Dist: 0})
		} else {
			dist[vertex] = 1<<31 - 1 // Максимальное значение int32
			heap.Push(&priorityQueue, &Item{Vertex: vertex, Dist: 1<<31 - 1})
		}
	}

	// Основной цикл алгоритма Дейкстры
	for priorityQueue.Len() > 0 {
		// Извлекаем вершину с наименьшим расстоянием
		item := heap.Pop(&priorityQueue).(*Item)
		currentVertex := item.Vertex
		currentDist := item.Dist

		// Обновляем расстояния до соседей
		for _, edge := range graph[currentVertex] {
			neighbor := edge.To
			distance := currentDist + edge.Weight
			if distance < dist[neighbor] {
				dist[neighbor] = distance
				// Обновляем очередь с приоритетами
				for _, item := range priorityQueue {
					if item.Vertex == neighbor {
						priorityQueue.update(item, neighbor, distance)
						break
					}
				}
			}
		}
	}

	return dist
}

func main() {
	// Пример графа в виде словаря смежности
	graph := Graph{
		"A": {{"B", 1}, {"C", 4}},
		"B": {{"A", 1}, {"C", 2}, {"D", 5}},
		"C": {{"A", 4}, {"B", 2}, {"D", 1}},
		"D": {{"B", 5}, {"C", 1}},
	}

	startVertex := "A"                                // Начальная вершина
	shortestDistances := Dijkstra(graph, startVertex) // Запуск алгоритма Дейкстры
	fmt.Println(shortestDistances)                    // Вывод результатов map[A:0 B:1 C:3 D:4]
}
