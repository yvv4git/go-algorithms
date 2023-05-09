package find_mode_in_binary_search_tree

func inorder(root *TreeNode, maxFreq, currVal, freq *int, result *[]int) {
	/*
		Time complexity : O(n) - look at every node.
		Space complexity : O(1) - or the pointers maxFreq, freq, and currVal. Solution sol array does not count toward space complexity.
	*/
	if root.Left != nil {
		inorder(root.Left, maxFreq, currVal, freq, result)
	}

	if root.Val == (*currVal) || (*freq) == 0 { //either this is the first time we're counting (start of the algorithm) or the last seen value (currVal) is the current value (root.Val).
		(*freq)++
	} else if root.Val != (*currVal) { //we're seeing a new number
		(*freq) = 1
	}
	(*currVal) = root.Val

	if (*freq) > (*maxFreq) {
		(*maxFreq) = (*freq)
		(*result) = []int{root.Val} //flush and replace
	} else if (*freq) == (*maxFreq) {
		(*result) = append((*result), root.Val)
	}

	if root.Right != nil {
		inorder(root.Right, maxFreq, currVal, freq, result)
	}
}

func findModeV1(root *TreeNode) []int {
	maxFreq := 0
	freq := 0
	currVal := 0
	var result []int

	inorder(root, &maxFreq, &currVal, &freq, &result)

	return result
}
