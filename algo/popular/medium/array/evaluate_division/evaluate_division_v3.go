package evaluate_division

type resultHelper struct {
	result map[string]float64
}

func calcEquationV3(equations [][]string, values []float64, queries [][]string) []float64 {
	/*
		Method: Arithmetic

		FormulaA
		A/B = X
		B/A = 1/X

		FormulaB
		A/B = X
		B/C = Y
		A/C = X * Y
	*/
	result := make([]float64, 0)
	memo := make(map[string]resultHelper, 0)

	// В этой функции регистрируем значения и сохраняем в memo.
	registerQuery := func(q []string, value float64) {
		//register the value
		if memo[q[0]].result == nil {
			res := memo[q[0]]
			res.result = make(map[string]float64, 0)
			memo[q[0]] = res
		}
		memo[q[0]].result[q[1]] = value

		//register the formulaA result
		if memo[q[1]].result == nil {
			res := memo[q[1]]
			res.result = make(map[string]float64, 0)
			memo[q[1]] = res
		}
		memo[q[1]].result[q[0]] = 1 / value
	}

	for i, e := range equations {
		registerQuery(e, values[i])
	}

	// В этой функции ищем значения в memo.
	var searchQuery func(q []string, visited map[string]bool) float64
	searchQuery = func(q []string, visited map[string]bool) float64 {
		if visited[q[0]] {
			return -1
		}
		if memo[q[0]].result == nil {
			return -1
		}
		//try to query it if the result empty try to recursive search
		if memo[q[0]].result[q[1]] == 0 {
			// mark current value is already visited
			visited[q[0]] = true
			//loop all posible formulaB
			for k, v := range memo[q[0]].result {
				// if already visited skip it
				if visited[k] {
					continue
				}
				// try to search the value we wanted
				queryResult := searchQuery([]string{k, q[1]}, visited)
				if queryResult != -1 {
					//found the result, then we register the result
					registerQuery([]string{q[0], q[1]}, v*queryResult)
					break
				}
			}
		}
		//if it still not found after recursive search then return -1
		if memo[q[0]].result[q[1]] == 0 {
			return -1
		}
		// return the result
		return memo[q[0]].result[q[1]]
	}

	for _, q := range queries {
		result = append(result, searchQuery(q, make(map[string]bool)))
	}

	return result
}
