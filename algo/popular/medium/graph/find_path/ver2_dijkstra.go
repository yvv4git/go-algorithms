package main

import (
	"container/heap"
	"math"
)

// Edge represents an edge in the graph.
type Edge struct {
	to     int
	weight int
}

// Graph represents a graph as an adjacency list.
type Graph struct {
	nodes map[int][]Edge
}

// NewGraph creates a new graph.
func NewGraph() *Graph {
	return &Graph{nodes: make(map[int][]Edge)}
}

// AddEdge adds an edge to the graph.
func (g *Graph) AddEdge(from, to, weight int) {
	g.nodes[from] = append(g.nodes[from], Edge{to: to, weight: weight})
	// For an undirected graph, add the reverse edge as well.
	g.nodes[to] = append(g.nodes[to], Edge{to: from, weight: weight})
}

// Item represents an item in the priority queue.
type Item struct {
	value    int
	priority int
	index    int
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the item with the lowest priority, so we use less than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// Dijkstra implements the Dijkstra algorithm to find the shortest path from source to destination.
func Dijkstra(graph *Graph, source, destination int) (int, bool) {
	dist := make(map[int]int)
	prev := make(map[int]int)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for node := range graph.nodes {
		dist[node] = math.MaxInt32 // Infinity
		prev[node] = -1
		item := &Item{
			value:    node,
			priority: dist[node],
		}
		heap.Push(&pq, item)
	}

	dist[source] = 0
	heap.Fix(&pq, source)

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Item).value

		if u == destination {
			break
		}

		for _, edge := range graph.nodes[u] {
			v := edge.to
			alt := dist[u] + edge.weight
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				pq.update(&Item{value: v, priority: alt}, v, alt)
			}
		}
	}

	// If the distance to the destination is infinity, there is no path.
	if dist[destination] == math.MaxInt32 {
		return 0, false
	}

	return dist[destination], true
}

// validPath uses Dijkstra's algorithm to determine if there is a path between source and destination.
func validPathV2(n int, edges [][]int, source int, destination int) bool {
	/*
		METHOD: Dijkstra

		TIME COMPLEXITY: O((V+E) log V) = O(E log V), так как каждая вершина вставляется в очередь ровно один раз, и для каждого ребра выполняется операция извлечения минимума и обновления.

		SPACE COMPLEXITY: O(V), так как мы должны хранить расстояния до каждой вершины и предшественников в массивах или списках.
		Кроме того, для приоритетной очереди требуется дополнительная память, которая в худшем случае может быть O(V).
		Таким образом, общее пространственное использование алгоритма Дейкстры будет O(V).
	*/
	graph := NewGraph()
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1], 1) // Assuming all edges have a weight of 1
	}

	_, found := Dijkstra(graph, source, destination)
	return found
}
