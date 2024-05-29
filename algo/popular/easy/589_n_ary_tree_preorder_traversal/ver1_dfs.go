package main

import "fmt"

// Node определяет структуру узла N-арного дерева
type Node struct {
	Val      int
	Children []*Node
}

// preorder обходит N-арное дерево в прямом порядке
func preorder(root *Node) []int {
	/*
		METHOD: DFS
		Используемый подход - это рекурсивный (depth-first search), так как мы сначала посещаем корень,
		а затем обходим каждое поддерево. Этот подход хорошо подходит для обхода деревьев,
		так как он позволяет легко реализовать и понять.

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, потому что мы посещаем каждый узел ровно один раз.


		SPACE COMPLEXITY: O(n), так как в худшем случае (когда дерево полностью несбалансированно) нам может понадобиться хранить n вызовов функции в стеке.
	*/
	// Базовый случай: если узел пустой, возвращаем пустой список
	if root == nil {
		return []int{}
	}

	// Инициализируем список для хранения результата
	result := []int{root.Val}

	// Обходим каждого из дочерних узлов
	for _, child := range root.Children {
		result = append(result, preorder(child)...)
	}

	return result
}

func main() {
	// Создаем тестовое N-арное дерево
	root := &Node{Val: 1}
	root.Children = []*Node{
		{Val: 3, Children: []*Node{
			{Val: 5},
			{Val: 6},
		}},
		{Val: 2},
		{Val: 4},
	}

	// Выполняем обход в прямом порядке и выводим результат
	fmt.Println(preorder(root)) // Вывод: [1 3 5 6 2 4]
}
