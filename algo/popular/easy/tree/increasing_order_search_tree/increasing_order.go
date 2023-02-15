package increasing_order_search_tree

/*
Задача в том, чтобы развернуть дерево в односвязный список.
*/
func increasingBST(root *TreeNode) *TreeNode {
	return dfs(root, nil)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root, tail *TreeNode) *TreeNode {
	if root == nil {
		return tail
	}

	res := dfs(root.Left, root)
	root.Left = nil
	root.Right = dfs(root.Right, tail)

	return res
}
