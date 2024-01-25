package dfs

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func dfsTraversalRecursion(root *Node) []*Node {
	/*
		Алгоритм обхода(Traversal) графа в грубину(DFS - Depth-First Search) с использованием рекурсии.
		Time complexity: O(V + E)
		Space complexity: O(V + E)
		Where V is the number of vertices and E is the number of edges in the graph.
	*/
	// visited := []*Node{}
	var visited []*Node // Здесь будут храниться посещенные ноды

	if root == nil {
		return visited // База рекурсии
	}

	return recurse(root, visited) // Рекурсивный случай
}

func recurse(root *Node, visited []*Node) []*Node {
	visited = append(visited, root)

	if root.Left != nil {
		visited = recurse(root.Left, visited)
	}

	if root.Right != nil {
		visited = recurse(root.Right, visited)
	}

	return visited
}
