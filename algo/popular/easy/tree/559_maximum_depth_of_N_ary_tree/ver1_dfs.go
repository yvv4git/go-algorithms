package main

import "fmt"

// Функция для нахождения максимальной глубины дерева
func maxDepth(root *Node) int {
	/*
		METHOD: DFS
		Для решения этой задачи мы будем использовать подход на основе обхода в глубину (Depth-First Search, DFS).
		Мы будем проходить по дереву, начиная с корня, и для каждого узла будем вычислять его глубину.
		Затем мы будем выбирать максимальную из полученных глубин.

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, так как мы обходим каждый узел ровно один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем потребовать дополнительное пространство для стека вызовов,
		который может достигать глубины дерева.

	*/
	// Базовый случай: если дерево пустое, то его глубина равна 0
	if root == nil {
		return 0
	}

	// Инициализация переменной для отслеживания максимальной глубины
	res := 0

	// Обход узлов дерева с использованием стека
	stack := []*Node{root}
	depths := []int{1}

	for len(stack) > 0 {
		// Извлечение узла и его глубины из стека
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		depth := depths[len(depths)-1]
		depths = depths[:len(depths)-1]

		// Обновление максимальной глубины, если текущая глубина больше
		if depth > res {
			res = depth
		}

		// Добавление дочерних узлов в стек с увеличением глубины на 1
		for _, child := range node.Children {
			stack = append(stack, child)
			depths = append(depths, depth+1)
		}
	}

	// Возврат максимальной глубины дерева
	return res
}

func main() {
	// Пример использования функции maxDepth
	root := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 3, Children: []*Node{
				{Val: 5},
				{Val: 6},
			}},
			{Val: 2},
			{Val: 4},
		},
	}

	fmt.Println(maxDepth(root)) // Вывод: 3
}
