package _22_count_complete_tree_nodes

type NodeGetter func(root *TreeNode) *TreeNode

func count(root *TreeNode, nodeGetter NodeGetter) int {
	var ret int
	for root != nil {
		root = nodeGetter(root)
		ret++
	}

	return ret
}

func countNodesV1(root *TreeNode) int {
	/*
		METHOD: Recursion
		TIME COMPLEXITY: O(log N * logN)
		SPACE COMPLEXITY: O(log N)
	*/
	if root == nil {
		return 0
	}

	leftCount := count(root, func(r *TreeNode) *TreeNode { return r.Left })
	rightCount := count(root, func(r *TreeNode) *TreeNode { return r.Right })

	if leftCount == rightCount {
		return (1 << leftCount) - 1
	}

	return countNodesV1(root.Left) + countNodesV1(root.Right) + 1
}
