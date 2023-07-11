package network_delay_time

func networkDelayTime(times [][]int, n int, k int) int {
	type edge struct{ to, weight int }

	w := make(map[int][]edge)
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		w[from] = append(w[from], edge{to, weight})
	}

	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = -1
	}

	intree := make([]bool, n+1)

	intree[k] = true
	dist[k] = 0

	v := k
	for {
		intree[v] = true

		for _, e := range w[v] {
			if dist[e.to] == -1 || (dist[e.to] > dist[v]+e.weight) {
				dist[e.to] = dist[v] + e.weight
			}
		}

		minv, mind := -1, -1
		for i := 1; i <= n; i++ {
			if !intree[i] && (mind == -1 || dist[i] < mind) {
				minv, mind = i, dist[i]
			}
		}

		if minv == -1 {
			break
		}

		v = minv
	}

	max := dist[1]
	for i := 1; i <= n; i++ {
		if dist[i] == -1 {
			return -1
		}
		if dist[i] > max {
			max = dist[i]
		}
	}

	return max
}
