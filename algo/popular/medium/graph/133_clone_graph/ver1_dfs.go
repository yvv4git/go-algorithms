package main

import "fmt"

// CloneGraph копирует граф.
func CloneGraph(node *Node) *Node {
	/*
		METHOD: DFS
		В данном решении используется подход на основе глубокого копирования графа с использованием рекурсии и хэш-таблицы для отслеживания уже посещенных узлов.
		Этот подход называется "Depth-First Search" (DFS) и используется для обхода графа в глубину.

		TIME COMPLEXITY: O(N + M), где N - количество узлов в графе, а M - количество рёбер (связей между узлами).
		Это связано с тем, что функция проходит по каждому узлу и каждому ребру ровно один раз.

		SPACE COMPLEXITY: O(N), так как в хэш-таблице visited для отслеживания уже посещенных узлов хранится один узел для каждого узла в исходном графе.
		Также учитывается рекурсивная глубина стека, которая в худшем случае может достигать O(H),
		где H - максимальная глубина рекурсии (в графе может быть цепочка из узлов, связанных в цикле).
		Однако, в среднем и в лучшем случае глубина стека будет не больше O(H),
		где H - количество уровней в графе (в ширину).
		Поэтому, в среднем и в лучшем случае, пространственная сложность будет O(N).
	*/
	if node == nil {
		return nil
	}

	// Хэш-таблица для отслеживания уже посещенных узлов.
	visited := make(map[*Node]*Node)

	// Рекурсивная функция для копирования узла и его соседей.
	var clone func(node *Node) *Node
	clone = func(node *Node) *Node {
		if v, ok := visited[node]; ok {
			return v
		}

		cloneNode := &Node{Val: node.Val, Neighbors: make([]*Node, len(node.Neighbors))}
		visited[node] = cloneNode

		for i := 0; i < len(node.Neighbors); i++ {
			cloneNode.Neighbors[i] = clone(node.Neighbors[i])
		}

		return cloneNode
	}

	return clone(node)
}

func main() {
	// Создаем тестовый граф.
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	// Копируем граф.
	clonedGraph := CloneGraph(node1)

	// Выводим значения узлов в скопированном графе.
	fmt.Println(clonedGraph.Val)
	for _, neighbor := range clonedGraph.Neighbors {
		fmt.Println(neighbor.Val)
	}
}
