package dfs

func dfsTraversalIterativeV1(root *Node) []int {
	/*
		Алгоритм обхода(Traversal) графа в грубину(DFS - Depth-First Search) без использования рекурсии.
		TIME COMPLEXITY: O(V + E)
		Space complexity: O(V + E)
		Where V is the number of vertices and E is the number of edges in the graph.
	*/
	var out []int
	prev := root
	stack := []*Node{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node.Left != nil && node.Left != prev && node.Right != prev {
			stack = append(stack, node.Left)
		} else if node.Right != nil && node.Right != prev {
			stack = append(stack, node.Right)
		} else {
			prev = node
			stack = stack[:len(stack)-1]
			out = append(out, node.Val)
		}
	}

	return out
}
