package main

import "fmt"

// Проверяет, является ли дерево s поддеревом дерева t
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	/*
		METHOD: Recursion
		Подход к решению этой задачи заключается в рекурсивном сравнении деревьев.
		Мы начинаем сравнение с корня дерева root и поддерева subRoot.
		Если текущие узлы одинаковы (т.е., они имеют одинаковые значения и структуру дочерних узлов),
		мы считаем, что subRoot является поддеревом root.
		Если текущие узлы не равны, мы продолжаем сравнение с левым и правым дочерними узлами root.

		Функция isSameTree используется для проверки того, что два дерева идентичны.
		Она сравнивает значения узлов и рекурсивно вызывает себя для левого и правого дочерних узлов.
		Если все значения узлов и их дочерние узлы совпадают, функция возвращает true, указывая на то, что деревья идентичны.

		Таким образом, основная функция isSubtree использует isSameTree для проверки того, что subRoot является поддеревом root,
		и если нет, то продолжает сравнение с левым и правым дочерними узлами root.
		Если subRoot не найден в root, функция возвращает false.

		TIME COMPLEXITY: O(m*n), временная сложность алгоритма зависит от количества узлов в деревьях root и subRoot.
		В худшем случае, когда subRoot является поддеревом root, мы должны сравнить каждый узел root с корнем subRoot.
		Это означает, что мы можем посетить каждый узел в каждом дереве.
		Таким образом, временная сложность составляет O(n * m), где n и m - количество узлов в деревьях root и subRoot соответственно.

		SPACE COMPLEXITY: O(n) дополнительной памяти для стека вызовов, где n - количество узлов в большем из деревьев.
		Однако, важно отметить, что в большинстве случаев глубина рекурсии будет меньше, чем количество узлов в большем из деревьев,
		поэтому фактическая пространственная сложность может быть меньше O(n).
	*/
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// Проверяет, являются ли два дерева одинаковыми
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// Основная функция для поиска поддерева
func isSubtreeMain(s *TreeNode, t *TreeNode) bool {
	// Если s пустое, то поддерево t не может быть его поддеревом
	if s == nil {
		return false
	}
	// Проверяем, является ли дерево s поддеревом дерева t
	// или является ли дерево t поддеревом левого или правого поддерева дерева s
	return isSubtree(s, t) || isSubtreeMain(s.Left, t) || isSubtreeMain(s.Right, t)
}

func main() {
	// Создаем деревья s и t для примера
	s := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	t := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	// Проверяем, является ли t поддеревом s
	fmt.Println(isSubtreeMain(s, t)) // Должно вывести: true
}
