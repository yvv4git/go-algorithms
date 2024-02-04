package _436_destination_city

func destCityV2(paths [][]string) string {
	/*
		METHOD: Graph
		TIME COMPLEXITY: O(V+E)
		Space complexity: O(V+E)

		Time complexity
		Временная сложность также составляет O(V + E), так как алгоритм проходит по всем вершинам и ребрам графа ровно один раз.

		Space complexity
		Пространственная сложность этого алгоритма составляет O(V + E), где V - количество вершин, E - количество ребер.
		Это связано с тем, что алгоритм проходит по всем вершинам и ребрам графа ровно один раз.
	*/
	graph := NewGraph()
	for _, path := range paths {
		graph.AddVertex(path[0])
		graph.AddVertex(path[1])
		graph.AddEdge(path[0], path[1])
	}
	return graph.GetDestination()
}

// Graph - структура графа
type Graph struct {
	// Вершины графа
	vertices map[string]bool
	// Ребра графа
	edges map[string][]string
}

// NewGraph - создание нового графа
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]bool),
		edges:    make(map[string][]string),
	}
}

// AddVertex - добавление вершины в граф
func (g *Graph) AddVertex(vertex string) {
	g.vertices[vertex] = true
}

// AddEdge - добавление ребра в граф
func (g *Graph) AddEdge(vertex1, vertex2 string) {
	g.edges[vertex1] = append(g.edges[vertex1], vertex2)
}

// GetDestination - получение конечного города
func (g *Graph) GetDestination() string {
	// Проходим по всем вершинам графа
	for vertex, _ := range g.vertices {
		// Если у вершины нет исходящих ребер, то это конечный город
		if _, ok := g.edges[vertex]; !ok {
			return vertex
		}
	}
	return ""
}
