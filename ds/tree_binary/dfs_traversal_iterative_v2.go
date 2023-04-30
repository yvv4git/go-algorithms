package tree_binary

func dfsTraversalIterativeV2(root *Node) []int {
	/*
		Алгоритм обхода(Traversal) графа в грубину(DFS - Depth-First Search).
		Time complexity: O(V + E)
		Space complexity: O(V + E)
		Where V is the number of vertices and E is the number of edges in the graph.
	*/
	out := []int{}
	var parents []*Node

	for next := root; next != nil; {
		if next.Left != nil {
			// if we have a left child push this node onto the parent stack and set next to the left child
			parents = append(parents, next)
			next = next.Left
		} else {
			out = append(out, next.Val)
			// otherwise, add parents until a right child is found
			for next = next.Right; len(parents) > 0 && next == nil; {
				index := len(parents) - 1
				// pop the next parent off the stack
				parent := parents[index]
				parents = parents[:index]
				// add the parent to the output
				out = append(out, parent.Val)
				// set the next value to the parent's right subtree
				next = parent.Right
			}
		}
	}

	return out
}
