package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Codec содержит методы для сериализации и десериализации бинарного дерева.
type Codec struct {
}

// Constructor создает новый экземпляр Codec.
func Constructor() Codec {
	return Codec{}
}

// Serialize преобразует бинарное дерево в строку.
// Временная сложность: O(n), где n - количество узлов в дереве.
// Пространственная сложность: O(h), где h - высота дерева (для рекурсивных вызовов).
func (this *Codec) Serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return strconv.Itoa(root.Val) + "," + this.Serialize(root.Left) + "," + this.Serialize(root.Right)
}

// Deserialize преобразует строку обратно в бинарное дерево.
// Временная сложность: O(n), где n - количество узлов в дереве.
// Пространственная сложность: O(n) для хранения узлов в слайсе.
func (this *Codec) Deserialize(data string) *TreeNode {
	values := strings.Split(data, ",")
	var buildTree func() *TreeNode
	buildTree = func() *TreeNode {
		if len(values) == 0 {
			return nil
		}
		val := values[0]
		values = values[1:]
		if val == "null" {
			return nil
		}
		intVal, _ := strconv.Atoi(val)
		node := &TreeNode{Val: intVal}
		node.Left = buildTree()
		node.Right = buildTree()
		return node
	}
	return buildTree()
}

// Пример использования.
func main() {
	/*
		Функции Serialize и Deserialize в предоставленном коде используют обход в глубину (Depth-First Search, DFS).

		Объяснение:
		1. TreeNode.
		Структура, представляющая узел бинарного дерева.

		2. Codec.
		Структура, содержащая методы для сериализации и десериализации.

		3. Constructor.
		Функция для создания нового экземпляра Codec.

		4. Serialize.
		Рекурсивная функция, которая обходит дерево в предварительном порядке (pre-order) и преобразует его в строку.
		Если узел пустой, добавляется "null".

		5. Deserialize.
		Рекурсивная функция, которая строит дерево из строки. Используется вспомогательная функция buildTree,
		которая извлекает значения из слайса и строит дерево.

	*/
	// Создаем бинарное дерево для примера.
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	// Создаем экземпляр Codec.
	ser := Constructor()
	deser := Constructor()

	// Сериализуем дерево в строку.
	data := ser.Serialize(root)
	fmt.Println("Serialized tree:", data)

	// Десериализуем строку обратно в дерево.
	newRoot := deser.Deserialize(data)
	fmt.Println("Deserialized tree:", ser.Serialize(newRoot))
}
