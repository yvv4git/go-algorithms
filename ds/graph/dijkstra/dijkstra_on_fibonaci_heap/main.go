//go:build ignore

package main

import (
	"fmt"
	"math"
)

// Node представляет узел фибоначчиевой кучи
type Node struct {
	key    float64 // Ключ (расстояние)
	vertex string  // Вершина
	degree int     // Степень узла
	parent *Node   // Родитель
	child  *Node   // Один из детей
	left   *Node   // Левый брат
	right  *Node   // Правый брат
	marked bool    // Флаг для cascading cut
}

// FibonacciHeap представляет фибоначчиеву кучу
type FibonacciHeap struct {
	min *Node // Указатель на минимальный узел
	n   int   // Количество узлов
}

// NewFibonacciHeap создает новую фибоначчиеву кучу
func NewFibonacciHeap() *FibonacciHeap {
	return &FibonacciHeap{}
}

// isEmpty проверяет, пуста ли куча
func (fh *FibonacciHeap) isEmpty() bool {
	return fh.min == nil
}

// insert вставляет новый узел в кучу
func (fh *FibonacciHeap) insert(node *Node) {
	if fh.min == nil {
		fh.min = node
		node.left = node
		node.right = node
	} else {
		fh.min.left.right = node
		node.left = fh.min.left
		node.right = fh.min
		fh.min.left = node
		if node.key < fh.min.key {
			fh.min = node
		}
	}
	fh.n++
}

// extractMin извлекает узел с минимальным ключом
func (fh *FibonacciHeap) extractMin() *Node {
	z := fh.min
	if z != nil {
		if z.child != nil {
			children := []*Node{}
			x := z.child
			for {
				children = append(children, x)
				x = x.right
				if x == z.child {
					break
				}
			}
			for _, child := range children {
				fh.min.left.right = child
				child.left = fh.min.left
				child.right = fh.min
				fh.min.left = child
				child.parent = nil
			}
		}
		z.left.right = z.right
		z.right.left = z.left
		if z == z.right {
			fh.min = nil
		} else {
			fh.min = z.right
			fh.consolidate()
		}
		fh.n--
	}
	return z
}

// consolidate консолидирует кучу после извлечения минимума
func (fh *FibonacciHeap) consolidate() {
	maxDegree := 50 // Достаточно для большинства случаев
	A := make([]*Node, maxDegree)
	w := fh.min
	end := fh.min
	if w != nil {
		for {
			x := w
			w = w.right
			d := x.degree
			for A[d] != nil {
				y := A[d]
				if x.key > y.key {
					x, y = y, x
				}
				fh.link(y, x)
				A[d] = nil
				d++
			}
			A[d] = x
			if w == end {
				break
			}
		}
		fh.min = nil
		for _, node := range A {
			if node != nil {
				if fh.min == nil {
					fh.min = node
					node.left = node
					node.right = node
				} else {
					fh.min.left.right = node
					node.left = fh.min.left
					node.right = fh.min
					fh.min.left = node
					if node.key < fh.min.key {
						fh.min = node
					}
				}
			}
		}
	}
}

// link связывает два узла
func (fh *FibonacciHeap) link(y, x *Node) {
	y.left.right = y.right
	y.right.left = y.left
	y.parent = x
	if x.child == nil {
		x.child = y
		y.left = y
		y.right = y
	} else {
		x.child.left.right = y
		y.left = x.child.left
		y.right = x.child
		x.child.left = y
	}
	y.marked = false
	x.degree++
}

// decreaseKey уменьшает ключ узла
func (fh *FibonacciHeap) decreaseKey(node *Node, newKey float64) {
	if newKey > node.key {
		return
	}
	node.key = newKey
	y := node.parent
	if y != nil && node.key < y.key {
		fh.cut(node, y)
		fh.cascadingCut(y)
	}
	if node.key < fh.min.key {
		fh.min = node
	}
}

// cut отрезает узел от родителя
func (fh *FibonacciHeap) cut(x, y *Node) {
	x.left.right = x.right
	x.right.left = x.left
	y.degree--
	if y.child == x {
		y.child = x.right
		if x.right == x {
			y.child = nil
		}
	}
	x.parent = nil
	x.marked = false
	fh.min.left.right = x
	x.left = fh.min.left
	x.right = fh.min
	fh.min.left = x
}

// cascadingCut выполняет каскадное отрезание
func (fh *FibonacciHeap) cascadingCut(y *Node) {
	z := y.parent
	if z != nil {
		if !y.marked {
			y.marked = true
		} else {
			fh.cut(y, z)
			fh.cascadingCut(z)
		}
	}
}

// Edge представляет ребро графа
type Edge struct {
	To     string  // Вершина, в которую ведет ребро
	Weight float64 // Вес ребра
}

// Graph представляет граф в виде списка смежности
type Graph map[string][]Edge

// Dijkstra реализует алгоритм Дейкстры с использованием фибоначчиевой кучи
func Dijkstra(graph Graph, start string) map[string]float64 {
	/*
		METHOD: Dijkstra's Algorithm with Fibonacci Heap
		1. Инициализируем фибоначчиеву кучу и словарь расстояний
		2. Вставляем все вершины в кучу с расстоянием ∞, кроме start с 0
		3. В цикле извлекаем вершину с мин расстоянием и обновляем соседей
		4. Используем decrease-key для обновления расстояний

		TIME COMPLEXITY: O(V log V + E) (амортизированная)
		SPACE COMPLEXITY: O(V) (куча и словарь расстояний)
	*/

	dist := make(map[string]float64)
	nodes := make(map[string]*Node)
	fh := NewFibonacciHeap()

	// Инициализация
	for vertex := range graph {
		if vertex == start {
			dist[vertex] = 0
		} else {
			dist[vertex] = math.Inf(1)
		}
		node := &Node{key: dist[vertex], vertex: vertex}
		nodes[vertex] = node
		fh.insert(node)
	}

	// Основной цикл
	for !fh.isEmpty() {
		minNode := fh.extractMin()
		u := minNode.vertex

		for _, edge := range graph[u] {
			v := edge.To
			if dist[u]+edge.Weight < dist[v] {
				dist[v] = dist[u] + edge.Weight
				fh.decreaseKey(nodes[v], dist[v])
			}
		}
	}

	return dist
}

func main() {
	// Пример графа
	graph := Graph{
		"A": {{"B", 1}, {"C", 4}},
		"B": {{"A", 1}, {"C", 2}, {"D", 5}},
		"C": {{"A", 4}, {"B", 2}, {"D", 1}},
		"D": {{"B", 5}, {"C", 1}},
	}

	startVertex := "A"
	shortestDistances := Dijkstra(graph, startVertex)
	fmt.Println(shortestDistances) // map[A:0 B:1 C:3 D:4]
}
