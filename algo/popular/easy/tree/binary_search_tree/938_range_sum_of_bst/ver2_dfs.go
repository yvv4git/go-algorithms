package main

func rangeSumBSTV2(root *TreeNode, low int, high int) int {
	/*
		Method: DFS
		Time complexity: O(n), где n - количество узлов в дереве, так как мы проходим по каждому узлу.
		Space complexity: O(h), где h - высота дерева, так как в худшем случае глубина рекурсии может достигать высоты дерева.
	*/
	if root == nil {
		return 0
	}

	if root.Val >= low && root.Val <= high {
		return root.Val + rangeSumBSTV2(root.Left, low, high) + rangeSumBSTV2(root.Right, low, high)
	} else {
		return rangeSumBSTV2(root.Left, low, high) + rangeSumBSTV2(root.Right, low, high)
	}
}
