package main

import "fmt"

// connect устанавливает указатель next для каждого узла на том же уровне
func connect(root *Node) *Node {
	/*
		METHOD: BFS
		Для решения этой задачи мы будем использовать подход на основе поиска в ширину (BFS),
		так как он позволяет нам обходить дерево по уровням, что позволяет легко установить указатель next для каждого узла на том же уровне.

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, так как мы проходим по каждому узлу ровно один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем хранить в очереди половину узлов дерева.
	*/
	if root == nil {
		return nil
	}

	// Очередь для обхода в ширину
	queue := []*Node{root}

	for len(queue) > 0 {
		// Количество узлов на текущем уровне
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			// Получаем первый узел из очереди
			node := queue[0]
			queue = queue[1:]

			// Устанавливаем указатель next для текущего узла
			if i < levelSize-1 {
				node.Next = queue[0]
			}

			// Добавляем потомков текущего узла в очередь
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}

// printLevelOrder печатает узлы дерева в порядке уровней
func printLevelOrder(root *Node) {
	for root != nil {
		cur := root
		for cur != nil {
			fmt.Printf("%d -> ", cur.Val)
			cur = cur.Next
		}
		fmt.Println("NULL")
		root = root.Left
	}
}

func main() {
	// Создаем дерево
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}

	// Устанавливаем указатели next
	root = connect(root)

	// Печатаем узлы дерева в порядке уровней
	printLevelOrder(root)
}
