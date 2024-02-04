package evaluate_division

type Pair struct {
	node string
	val  float64
}

func calcEquationV1(equations [][]string, values []float64, queries [][]string) []float64 {
	/*
		METHOD: BFS (Breadth-First Search)
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

	answer := []float64{}
	for _, query := range queries {
		p, q := query[0], query[1]
		visited := make(map[string]bool)
		que := []Pair{Pair{p, 1}}
		ans := float64(-1)
		for len(que) > 0 {
			pair := que[0]
			que = que[1:]
			node, val := pair.node, pair.val
			if _, ok := visited[node]; ok {
				continue
			}
			if _, ok := graph[node]; !ok {
				break
			}
			if node == q {
				ans = val
				break
			}
			visited[node] = true
			for nb, nb_val := range graph[node] {
				que = append(que, Pair{nb, val * nb_val})
			}
		}
		answer = append(answer, ans)
	}

	return answer
}
