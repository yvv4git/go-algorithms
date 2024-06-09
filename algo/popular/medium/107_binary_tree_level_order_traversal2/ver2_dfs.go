package main

func levelOrderBottomV2(root *TreeNode) [][]int {
	/*
		METHOD: DFS

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, потому что мы посещаем каждый узел ровно один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае (когда дерево сбалансировано) глубина рекурсии может достигать n, и мы храним эти уровни в стеке вызовов.
	*/
	result := [][]int{}
	dfs(root, &result, 0)
	// Развернем результат, так как мы добавляли уровни сверху вниз
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func dfs(node *TreeNode, result *[][]int, level int) {
	if node == nil {
		return
	}

	// Если уровень еще не добавлен в результат, добавляем его
	if level == len(*result) {
		*result = append(*result, []int{})
	}

	// Добавляем значение узла на соответствующий уровень
	(*result)[level] = append((*result)[level], node.Val)

	// Рекурсивно обходим левое и правое поддерево
	dfs(node.Left, result, level+1)
	dfs(node.Right, result, level+1)
}

//func main() {
//	root := &TreeNode{
//		Val: 3,
//		Left: &TreeNode{
//			Val: 9,
//		},
//		Right: &TreeNode{
//			Val: 20,
//			Left: &TreeNode{
//				Val: 15,
//			},
//			Right: &TreeNode{
//				Val: 7,
//			},
//		},
//	}
//
//	levels := levelOrderBottom(root)
//	for _, level := range levels {
//		fmt.Println(level)
//	}
//}
