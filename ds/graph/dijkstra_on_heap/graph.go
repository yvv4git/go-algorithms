/**

 */
package main

// Edge - an implementation of edge.
type Edge struct {
	node   string
	weight int
}

// Graph - an implementation of a graph.
type Graph struct {
	nodes map[string][]Edge
}

// NewGraph - used for create instance of the graph.
func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]Edge)}
}

// AddEdge - used for add edge.
func (g *Graph) AddEdge(origin, destiny string, weight int) {
	g.nodes[origin] = append(g.nodes[origin], Edge{node: destiny, weight: weight})
	g.nodes[destiny] = append(g.nodes[destiny], Edge{node: origin, weight: weight})
}

// GetEdges - used for getting edges of node.
func (g *Graph) GetEdges(node string) []Edge {
	return g.nodes[node]
}

// GetPath - used for calculate path.
func (g *Graph) GetPath(origin, destiny string) (int, []string) {
	h := NewHeap()
	h.Push(Path{value: 0, nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.Pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			return p.value, p.nodes
		}

		for _, e := range g.GetEdges(node) {
			if !visited[e.node] {
				// We calculate the total spent so far plus the cost and the Path of getting here
				h.Push(Path{value: p.value + e.weight, nodes: append([]string{}, append(p.nodes, e.node)...)})
			}
		}

		visited[node] = true
	}

	return 0, nil
}
