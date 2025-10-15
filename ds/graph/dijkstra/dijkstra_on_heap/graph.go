package dijkstra

// Edge - an implementation of edge.
type Edge struct {
	node   string
	weight float64
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
func (g *Graph) AddEdge(origin, destiny string, weight float64) {
	g.nodes[origin] = append(
		g.nodes[origin],
		Edge{node: destiny, weight: weight},
	)
	g.nodes[destiny] = append(
		g.nodes[destiny],
		Edge{node: origin, weight: weight},
	)
}

// GetEdges - used for getting edges of node.
func (g *Graph) GetEdges(node string) []Edge {
	return g.nodes[node]
}

// CalcMinPath - used for calculate min path.
func (g *Graph) CalcMinPath(origin, destiny string) (float64, []string) {
	h := NewHeap()
	h.Push(
		Path{
			weight: 0,
			nodes:  []string{origin},
		},
	)
	visited := make(map[string]bool)

	for len(h.MinPath.Route) > 0 {
		p := h.Pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			return p.weight, p.nodes
		}

		for _, e := range g.GetEdges(node) {
			if !visited[e.node] {
				h.Push(
					Path{
						weight: p.weight + e.weight,
						nodes:  append([]string{}, append(p.nodes, e.node)...),
					},
				)
			}
		}

		visited[node] = true
	}

	return 0, nil
}
