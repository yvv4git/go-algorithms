//go:build ignore

package main

import (
	"container/list"
	"fmt"
)

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	/*
		Approach: BFS traversal with memoization
		findFrequentTreeSum finds the most frequent subtree sums in a binary tree
		Time complexity: O(n) where n is number of nodes
		Space complexity: O(n) for hash maps and queue

		Подход: BFS-обход с мемоизацией
		1. Используем очередь для обхода в ширину
		2. Сохраняем суммы поддеревьев в отдельной карте
		3. Для каждого узла вычисляем сумму как значение узла + суммы поддеревьев
		4. Подсчитываем частоту сумм и находим наиболее частые
	*/
	if root == nil {
		return nil
	}

	// Создаем карту для хранения сумм поддеревьев
	sumMap := make(map[*TreeNode]int)
	// Очередь для BFS
	queue := list.New()
	queue.PushBack(root)
	// Стек для обработки узлов в обратном порядке
	stack := list.New()

	// Первый проход: BFS и сохранение узлов в стек
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		stack.PushBack(node)

		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}

	// Второй проход: вычисление сумм в обратном порядке
	freq := make(map[int]int)
	maxFreq := 0

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		sum := node.Val

		if node.Left != nil {
			sum += sumMap[node.Left]
		}
		if node.Right != nil {
			sum += sumMap[node.Right]
		}

		sumMap[node] = sum
		freq[sum]++
		if freq[sum] > maxFreq {
			maxFreq = freq[sum]
		}
	}

	// Собираем результаты
	var result []int
	for sum, count := range freq {
		if count == maxFreq {
			result = append(result, sum)
		}
	}

	return result
}

func main() {
	// Example 1
	root1 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: -3,
		},
	}
	fmt.Println("Example 1:", findFrequentTreeSum(root1)) // Output: [2, -3, 4]

	// Example 2
	root2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: -5,
		},
	}
	fmt.Println("Example 2:", findFrequentTreeSum(root2)) // Output: [2]
}
