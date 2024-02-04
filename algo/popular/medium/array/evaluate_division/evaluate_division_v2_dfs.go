package evaluate_division

func calcEquationV2(equations [][]string, values []float64, queries [][]string) []float64 {
	/*
		METHOD: DFS (Depth-First Search)
		TIME COMPLEXITY: O(n^2)
		SPACE COMPLEXITY: O(n)
	*/
	graph := make(map[string]map[string]float64)
	for i := 0; i < len(equations); i++ {
		equation := equations[i]
		p := equation[0]
		q := equation[1]
		val := values[i]
		if _, ok := graph[p]; !ok {
			graph[p] = make(map[string]float64)
		}
		if _, ok := graph[q]; !ok {
			graph[q] = make(map[string]float64)
		}
		graph[p][q] = val
		graph[q][p] = 1 / val
	}

	result := []float64{}
	for _, query := range queries {
		p, q := query[0], query[1]
		visited := make(map[string]bool)
		result = append(result, dfs(p, q, graph, visited, 1))
	}

	return result
}

func dfs(begin, end string, graph map[string]map[string]float64, visited map[string]bool, value float64) float64 {
	if _, ok := visited[begin]; ok {
		return -1
	}

	if _, ok := graph[begin]; !ok {
		return -1
	}

	if begin == end {
		return value
	}

	visited[begin] = true
	v := float64(-1)
	for nb, nbValue := range graph[begin] {
		v = max(v, dfs(nb, end, graph, visited, value*nbValue))
	}
	delete(visited, begin)

	return v
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}
