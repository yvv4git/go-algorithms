package find_mode_in_binary_search_tree

import (
	"sort"
)

func findModeV2(root *TreeNode) []int {
	/*
		Method: DFS + Recursion.

		Time: O(n), where n is the size of the given tree.
		Space: O(n), where n is the size of the stack during the DFS.
	*/
	countMap := map[int]int{}
	dfs(root, countMap) // Здесь рекурсивно вычисляется сколько раз, встречается значение в графе.

	result := []int{} // Здесь будет храниться результат.
	count := [][]int{}
	for k, v := range countMap {
		val := []int{k, v}
		count = append(count, val)
	}

	sort.Slice(count, func(i, j int) bool {
		return count[i][1] > count[j][1] // Важно понимать, что count[i][1] - это второй элемент и он туда попадает из countMap(значение).
	})

	for _, v := range count {
		if count[0][1] == v[1] {
			result = append(result, v[0])
		}
	}

	// Сортируют итоговый slice, в котором храниться результат.
	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})

	return result
}

func dfs(node *TreeNode, countMap map[int]int) {
	/*
		В данном случае, алгоритм dfs используется для того, чтобы обойти весь граф и
		посчитать сколько раз какое значение встречается в графе.
	*/
	if node == nil {
		return
	}

	countMap[node.Val]++ // Увеличиваем счетчик.
	dfs(node.Left, countMap)
	dfs(node.Right, countMap)
}
