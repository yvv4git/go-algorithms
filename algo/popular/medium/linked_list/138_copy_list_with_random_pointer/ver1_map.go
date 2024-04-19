package main

import "fmt"

// copyRandomList копирует связный список с указателем на случайный узел.
func copyRandomList(head *Node) *Node {
	/*
		METHOD: Map
		Для решения этой задачи мы можем использовать два прохода по списку.
		В первом проходе мы создаем новые узлы и сохраняем их в хэш-таблице, где ключом является исходный узел, а значением - новый узел.
		Во втором проходе мы устанавливаем случайные указатели для новых узлов, используя информацию из хэш-таблицы.

		TIME COMPLEXITY: O(n), где n - количество узлов в списке, так как мы проходим по каждому узлу два раза.

		SPACE COMPLEXITY: O(n), так как мы храним новые узлы в хэш-таблице.
	*/
	if head == nil {
		return nil
	}

	// Создаем хэш-таблицу для хранения соответствия между исходными и копируемыми узлами.
	nodeMap := make(map[*Node]*Node)

	// Первый проход: создаем новые узлы и сохраняем их в хэш-таблице.
	cur := head
	for cur != nil {
		nodeMap[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	// Второй проход: устанавливаем случайные указатели для новых узлов.
	cur = head
	for cur != nil {
		nodeMap[cur].Next = nodeMap[cur.Next]
		nodeMap[cur].Random = nodeMap[cur.Random]
		cur = cur.Next
	}

	// Возвращаем новый список.
	return nodeMap[head]
}

func main() {
	// Создаем тестовый список.
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node1.Next = node2
	node2.Next = node3
	node1.Random = node3
	node2.Random = node1
	node3.Random = node2

	// Копируем список.
	copiedList := copyRandomList(node1)

	// Выводим копию списка.
	for copiedList != nil {
		fmt.Printf("Val: %d, Random: %v\n", copiedList.Val, copiedList.Random)
		copiedList = copiedList.Next
	}
}
