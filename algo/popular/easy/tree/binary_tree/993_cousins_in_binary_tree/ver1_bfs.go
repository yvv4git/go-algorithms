package main

import "fmt"

func isCousins(root *TreeNode, x int, y int) bool {
	/*
		METHOD: BFS
		Для решения этой задачи мы будем использовать обход в ширину (BFS),
		так как он позволяет нам легко отслеживать уровень каждого узла и его родителя.

		Двойниками называются два узла, если:
		1. Они находятся на одном уровне в дереве (т.е., расстояние от корня до них одинаковое)
		2. У них нет общего родителя (т.е., один из родителей одного узла не является родителем другого узла).

		     1
			/ \
		   2   3
		  /   / \
		 4   5   6

		Здесь узлы со значениями 4 и 6 являются двойниками, потому что:
		- Оба находятся на 2-м уровне (расстояние от корня).
		- Узлы 4 и 6 не имеют общего родителя, т.е., узел 2 не является родителем узла 3, и наоборот.

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, так как мы проходим по каждому узлу ровно один раз.

		SPACE COMPLEXITY: O(w), где w - ширина дерева на последнем уровне, так как в худшем случае мы храним все узлы последнего уровня в очереди.
	*/
	if root == nil {
		return false
	}

	// Создаем очередь для BFS и добавляем корень
	queue := []*TreeNode{root}

	// Пока очередь не пуста
	for len(queue) > 0 {
		// Инициализируем флаги для проверки наличия x и y на текущем уровне
		foundX := false
		foundY := false

		// Проходим по всем узлам на текущем уровне
		for _, node := range queue {
			// Проверяем, является ли узел x или y
			if node.Val == x {
				foundX = true
			}
			if node.Val == y {
				foundY = true
			}

			// Проверяем, есть ли у узла x и y общий родитель
			if node.Left != nil && node.Right != nil {
				if (node.Left.Val == x && node.Right.Val == y) || (node.Left.Val == y && node.Right.Val == x) {
					return false
				}
			}
		}

		// Если x и y найдены на одном уровне, возвращаем true
		if foundX && foundY {
			return true
		}
		if foundX || foundY {
			// Если только один из них найден, возвращаем false
			return false
		}

		// Переходим на следующий уровень
		nextLevel := []*TreeNode{}
		for _, node := range queue {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		queue = nextLevel
	}

	return false
}

func main() {
	// Пример использования функции
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	fmt.Println(isCousins(root, 4, 5)) // Вывод: true
}
