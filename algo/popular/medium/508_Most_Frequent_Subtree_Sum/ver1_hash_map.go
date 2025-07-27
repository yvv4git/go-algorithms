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
		Approach: Post-order traversal with hash map
		findFrequentTreeSum finds the most frequent subtree sums in a binary tree
		Time complexity: O(n), где n - количество узлов (посетите каждый узел один раз)
		Space complexity: O(n), где n - количество узлов (посетите каждый узел один раз)

		Подход: Пост-ордерный обход с хеш-картой
		1. Рекурсивно вычисляем сумму каждого поддерева (значение узла + суммы левого и правого поддеревьев)
		2. Подсчитываем частоту каждой суммы с помощью хеш-карты
		3. Находим суммы с максимальной частотой встречаемости
		4. Возвращаем все суммы с максимальной частотой
	*/
	if root == nil {
		return nil
	}

	freq := make(map[int]int)
	maxFreq := 0

	var postOrder func(node *TreeNode) int
	postOrder = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftSum := postOrder(node.Left)
		rightSum := postOrder(node.Right)
		sum := node.Val + leftSum + rightSum

		freq[sum]++
		if freq[sum] > maxFreq {
			maxFreq = freq[sum]
		}

		return sum
	}

	postOrder(root)

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
