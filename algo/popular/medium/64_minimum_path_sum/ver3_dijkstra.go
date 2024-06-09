package main

import (
	"container/heap"
	"math"
)

// Структура для хранения координат ячейки
type Cell struct {
	x, y int
}

// Структура для хранения приоритетной очереди
type PriorityQueue []*Item

// Структура для хранения элемента очереди
type Item struct {
	cell  Cell
	cost  int
	index int
}

// Функция для создания приоритетной очереди
func newPriorityQueue() PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return pq
}

// Реализация интерфейса heap.Interface для PriorityQueue
func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Функция для нахождения минимальной суммы пути
func minPathSumV3(grid [][]int) int {
	/*
		METHOD: Dynamic programming

		TIME COMPLEXITY: O(V^2), где V - количество вершин (в нашем случае, это количество ячеек в матрице).
		Это происходит, когда каждая вершина (ячейка) помещается в очередь ровно один раз, и для каждой вершины проверяются все ее соседи.
		Однако, если использовать приоритетную очередь (heap), временная сложность может быть уменьшена до O((V+E)logV),
		где E - количество ребер (в нашем случае, это количество соседних ячеек для каждой ячейки).
		Это происходит, когда мы извлекаем минимальный элемент из очереди и обновляем расстояния только для тех соседних вершин,
		которые еще не были посещены.
		В итоге, если использовать приоритетную очередь, временная сложность может быть уменьшена, но в худшем случае она все равно будет O(V^2).
		Пространственная сложность остается O(V^2), так как мы должны хранить информацию о расстояниях до каждой вершины.

		SPACE COMPLEXITY: O(V^2), так как мы храним расстояния до каждой вершины в матрице.
		Это требуется для того, чтобы мы могли обновлять расстояния и определять, какие вершины уже были посещены.
	*/
	rows := len(grid)
	cols := len(grid[0])
	directions := []Cell{{0, 1}, {1, 0}} // Право и вниз

	// Создаем массив для хранения минимальных сумм пути
	distances := make([][]int, rows)
	for i := range distances {
		distances[i] = make([]int, cols)
		for j := range distances[i] {
			distances[i][j] = math.MaxInt32 // Задаем начальное значение как бесконечность
		}
	}
	distances[0][0] = grid[0][0] // Начальная ячейка имеет сумму равную значению в ней

	// Создаем приоритетную очередь и добавляем начальную ячейку
	pq := newPriorityQueue()
	heap.Push(&pq, &Item{cell: Cell{0, 0}, cost: grid[0][0]})

	// Алгоритм Дейкстры
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		current := item.cell

		// Если достигли цели, возвращаем минимальную сумму пути
		if current.x == rows-1 && current.y == cols-1 {
			return distances[rows-1][cols-1]
		}

		// Проверяем соседние ячейки
		for _, dir := range directions {
			next := Cell{current.x + dir.x, current.y + dir.y}
			if next.x < rows && next.y < cols {
				newCost := distances[current.x][current.y] + grid[next.x][next.y]
				if newCost < distances[next.x][next.y] {
					distances[next.x][next.y] = newCost
					heap.Push(&pq, &Item{cell: next, cost: newCost})
				}
			}
		}
	}

	// Если целевая ячейка не достижима, возвращаем -1
	return -1
}

//func main() {
//	grid := [][]int{
//		{1, 3, 1},
//		{1, 5, 1},
//		{4, 2, 1},
//	}
//	fmt.Println(minPathSum(grid)) // Выводит: 7
//}
