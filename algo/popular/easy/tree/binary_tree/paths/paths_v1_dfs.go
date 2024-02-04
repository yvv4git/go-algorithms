package paths

import "strconv"

func binaryTreePathsV1(root *TreeNode) []string {
	/*
		METHOD: DFS / Recursion
		TIME COMPLEXITY: O(n^2)
		SPACE COMPLEXITY: O(n^2)
	*/
	if root == nil {
		return []string{}
	}

	result := []string{}
	cur := ""

	findPathsDFS(root, cur, &result)

	return result
}

func findPathsDFS(node *TreeNode, current string, paths *[]string) {
	if node == nil {
		return
	}

	current += strconv.Itoa(node.Val) + "->"

	if node.Left == nil && node.Right == nil {
		*paths = append(*paths, current[:len(current)-2])
		return
	}

	findPathsDFS(node.Left, current, paths)
	findPathsDFS(node.Right, current, paths)
}
