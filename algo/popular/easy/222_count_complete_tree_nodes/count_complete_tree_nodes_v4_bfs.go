package _22_count_complete_tree_nodes

func countNodesV4(root *TreeNode) int {
	/*
		Method: BFS
		Time complexity : O(V+E)
		Space complexity : O(1)

		Объяснение Time Complexity:
		Причина, по которой временная сложность равна O(V+E), заключается в том, что, когда BFS проходит по графику,
		он помечает узлы как посещенные, чтобы знать, что к ним нельзя возвращаться.
		Поскольку каждое ребро имеет только две конечные точки, и каждую конечную точку можно посетить только один раз, ребро можно наблюдать не более двух раз.
		Таким образом, BFS никогда не посещает какую-либо "вещь" более двух раз
		(узлы, которых V много, он может просмотреть только один раз, а ребра, которых E много, можно просмотреть только дважды).
	*/
	cnt := 0
	if root == nil {
		return cnt
	}

	var q []*TreeNode
	q = append(q, root)

	for len(q) > 0 {
		size := len(q)

		for idx := 0; idx < size; idx++ {
			node := q[0]
			q = q[1:]
			cnt++

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return cnt
}
