//go:build ignore

package main

import (
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
		Approach: Iterative post-order traversal with stack
		findFrequentTreeSum finds the most frequent subtree sums in a binary tree
		Time complexity: O(n) where n is number of nodes (visit each node once)
		Space complexity: O(n) for hash map and stack

		Подход: Итеративный пост-ордерный обход со стеком
		1. Используем стек для эмуляции рекурсивного обхода
		2. Для каждого узла сохраняем его в стеке вместе с флагом посещения
		3. Когда узел извлекается из стека с флагом посещения, вычисляем сумму поддерева
		4. Подсчитываем частоту сумм и находим наиболее частые
	*/
	if root == nil {
		return nil
	}

	freq := make(map[int]int)
	maxFreq := 0
	sumMap := make(map[*TreeNode]int)
	stack := []struct {
		node    *TreeNode
		visited bool
	}{{root, false}}

	for len(stack) > 0 {
		// Извлекаем последний элемент стека
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.visited {
			// Узел уже посещен, вычисляем сумму
			sum := top.node.Val
			if left, ok := sumMap[top.node.Left]; ok {
				sum += left
			}
			if right, ok := sumMap[top.node.Right]; ok {
				sum += right
			}

			sumMap[top.node] = sum
			freq[sum]++
			if freq[sum] > maxFreq {
				maxFreq = freq[sum]
			}
		} else {
			// Помещаем узел обратно в стек с флагом посещения
			stack = append(stack, struct {
				node    *TreeNode
				visited bool
			}{top.node, true})

			// Добавляем правый и левый дочерние узлы
			if top.node.Right != nil {
				stack = append(stack, struct {
					node    *TreeNode
					visited bool
				}{top.node.Right, false})
			}
			if top.node.Left != nil {
				stack = append(stack, struct {
					node    *TreeNode
					visited bool
				}{top.node.Left, false})
			}
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
