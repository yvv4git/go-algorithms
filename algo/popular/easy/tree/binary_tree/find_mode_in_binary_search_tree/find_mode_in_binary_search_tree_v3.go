package find_mode_in_binary_search_tree

func findModeV3(root *TreeNode) []int {
	/*
		Method: BFS + map of records.

		Time complexity : O(n) - поиск в ширину с помощью BFS.
		Space complexity : O(n) - нужно будет хранить n элементов в памяти.
	*/
	var nodesList []*TreeNode
	var result []int
	var records = map[int]int{}
	var maxCount = 0

	if root == nil {
		return result
	}

	nodesList = append(nodesList, root)

	for node := root; len(nodesList) != 0; {
		node, nodesList = nodesList[0], nodesList[1:]

		if node != nil {
			records[node.Val]++
			nodesList = append(nodesList, node.Right, node.Left)
			if records[node.Val] > maxCount {
				maxCount = records[node.Val]
			}
		}
	}

	for ele, times := range records {
		if times == maxCount {
			result = append(result, ele)
		}
	}

	return result
}
