//go:build ignore

package main

import "fmt"

/**
 * Определение узла бинарного дерева.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Функция для поиска соответствующего узла в клонированном дереве.
 * original - оригинальное дерево.
 * cloned - клонированное дерево.
 * target - узел в оригинальном дереве, который нужно найти в клоне.
 * Возвращает соответствующий узел в клонированном дереве.
 */
func getTargetCopy(original *TreeNode, cloned *TreeNode, target *TreeNode) *TreeNode {
	/*
		METHOD: RECURSION (Depth-First Search, DFS)
		- Проверяем, является ли текущий узел целевым узлом.
		- Если да, возвращаем соответствующий узел в клонированном дереве.
		- Если нет, рекурсивно ищем в левом и правом поддереве.
		- Если узел не найден, возвращаем nil.

		TIME COMPLEXITY: O(N), где N - количество узлов в дереве.
		  В худшем случае мы обходим все узлы дерева, чтобы найти целевой узел.

		SPACE COMPLEXITY: O(H), где H - высота дерева.
		  Пространство используется для рекурсивного стека вызовов.
		  В худшем случае (вырожденное дерево) H = N, в сбалансированном дереве H = log(N).
	*/
	// Если оригинальное дерево пустое, возвращаем nil
	if original == nil {
		return nil
	}

	// Если текущий узел в оригинальном дереве равен целевому узлу,
	// возвращаем соответствующий узел в клонированном дереве
	if original == target {
		return cloned
	}

	// Рекурсивно ищем в левом поддереве
	left := getTargetCopy(original.Left, cloned.Left, target)
	if left != nil { // Если узел найден в левом поддереве, возвращаем его
		return left
	}

	// Рекурсивно ищем в правом поддереве
	right := getTargetCopy(original.Right, cloned.Right, target)
	if right != nil { // Если узел найден в правом поддереве, возвращаем его
		return right
	}

	// Если узел не найден, возвращаем nil
	return nil
}

func main() {
	// Создаем оригинальное дерево
	original := &TreeNode{Val: 1}
	original.Left = &TreeNode{Val: 2}
	original.Right = &TreeNode{Val: 3}
	original.Left.Left = &TreeNode{Val: 4}
	original.Left.Right = &TreeNode{Val: 5}

	// Клонируем дерево (в реальности клон будет создан отдельно, здесь для примера используем то же дерево)
	cloned := original

	// Целевой узел в оригинальном дереве (например, узел со значением 5)
	target := original.Left.Right

	// Находим соответствующий узел в клоне
	result := getTargetCopy(original, cloned, target)
	fmt.Println(result.Val) // Вывод: 5
}
