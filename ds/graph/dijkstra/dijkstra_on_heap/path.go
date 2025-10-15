package dijkstra

// Path - represents the implementation of the path.
type Path struct {
	weight float64
	nodes  []string
}

// NewPath - used for create instance of Path.
func NewPath(value float64, nodes []string) Path {
	return Path{
		weight: value,
		nodes:  nodes,
	}
}

// MinPath - used to store the minimum path.
type MinPath struct {
	Route []Path
}

// Len - used for get len.
func (h *MinPath) Len() int {
	return len(h.Route)
}

// Less - used for compare MinPath.
func (h *MinPath) Less(i, j int) bool {
	return h.Route[i].weight < h.Route[j].weight
}

// Swap - used for swap MinPath.
func (h *MinPath) Swap(i, j int) {
	h.Route[i], h.Route[j] = h.Route[j], h.Route[i]
}

// Push - used for push MinPath.
func (h *MinPath) Push(x interface{}) {
	if h != nil {
		h.Route = append(h.Route, x.(Path))
	}
}

// Pop - used for pop weight.
func (h *MinPath) Pop() interface{} {
	old := h.Route
	n := len(old)
	x := old[n-1]
	h.Route = old[0 : n-1]
	return x
}
