package fibonacci

func fibSeriesMemoization(n int) []int {
	a := make([]int, n)
	m := make(map[int]int)

	for i := 1; i <= n; i++ {
		a[i-1] = fibMemo(i, m)
	}

	return a
}

func fibMemo(x int, m map[int]int) int {
	if x < 2 {
		m[x] = x
		return x
	}
	_, ok := m[x-1]
	if !ok {
		m[x-1] = fibMemo(x-1, m)
	}

	_, ok = m[x-2]
	if !ok {
		m[x-2] = fibMemo(x-2, m)
	}

	return m[x-1] + m[x-2]
}
