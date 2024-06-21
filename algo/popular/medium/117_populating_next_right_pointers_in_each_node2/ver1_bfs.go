package main

import "fmt"

// Функция для установки указателей next для каждого узла
func connect(root *Node) *Node {
	/*
		METHOD: BFS
		Почему выбран BFS:
		1. Обход в ширину позволяет нам посещать узлы по уровням, что идеально подходит для установки указателей next,
		так как нам нужно связать узлы одного уровня друг с другом.

		2. Использование очереди позволяет нам управлять узлами в порядке их уровней.
		Мы добавляем узлы в очередь и обрабатываем их, начиная с корня и двигаясь вниз по уровням.

		3. Для каждого уровня мы обрабатываем все узлы этого уровня, устанавливая указатель next для каждого узла,
		чтобы он указывал на следующий узел на том же уровне. Если узел является последним на своем уровне,
		его указатель next устанавливается в null.

		Этот подход гарантирует, что все узлы будут правильно связаны,
		и мы достигнем желаемого результата за линейное время относительно количества узлов в дереве.

		TIME COMPLEXITY: Временная сложность: O(n), где n — количество узлов в дереве.
		Мы посещаем каждый узел ровно один раз во время обхода в ширину.

		SPACE COMPLEXITY: O(m), где m — максимальное количество узлов на одном уровне.
		В худшем случае, это будет количество узлов на самом широком уровне дерева.
	*/
	if root == nil {
		return nil
	}

	// Используем очередь для обхода в ширину
	queue := []*Node{root}

	for len(queue) > 0 {
		// Размер текущего уровня
		size := len(queue)
		var prev *Node

		// Обрабатываем все узлы текущего уровня
		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]

			// Устанавливаем указатель next для предыдущего узла
			if prev != nil {
				prev.Next = current
			}
			prev = current

			// Добавляем дочерние узлы в очередь для следующего уровня
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}

	return root
}

// Вспомогательная функция для печати дерева (для проверки)
func printTree(root *Node) {
	if root == nil {
		return
	}

	current := root
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("null")

	if root.Left != nil {
		printTree(root.Left)
	} else if root.Right != nil {
		printTree(root.Right)
	}
}

func main() {
	// Пример дерева
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Right = &Node{Val: 7}

	root = connect(root)
	printTree(root)
}
