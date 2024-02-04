package path_sum2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	/*
		TIME COMPLEXITY: O(n)
		Space complexity: O(n)
	*/
	res := make([][]int, 0)
	path := make([]int, 0)
	dfs(root, targetSum, path, &res)

	return res
}

func dfs(node *TreeNode, targetSum int, path []int, res *[][]int) {
	if node == nil {
		return
	}

	path = append(path, node.Val)
	if node.Left == nil && node.Right == nil && targetSum == node.Val {
		*res = append(*res, append([]int{}, path...))
	}

	dfs(node.Left, targetSum-node.Val, path, res)
	dfs(node.Right, targetSum-node.Val, path, res)
	path = path[:len(path)-1]
}
