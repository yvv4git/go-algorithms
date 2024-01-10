package main

func numTreesV2(n int) int {
	m := map[int]int{}

	return numTreesHelper(n, m)
}

func numTreesHelper(n int, m map[int]int) int {
	if n == 0 { // empty is also one kind, for example, when left child tree is empty, right is not.
		return 1
	}
	if n == 1 {
		return 1
	}

	if c, ok := m[n]; ok {
		return c
	}

	var l, r, sum int
	for root := 1; root <= n; root++ {
		l = numTreesHelper(root-1, m)
		r = numTreesHelper(n-root, m)
		sum += l * r
	}

	m[n] = sum

	return sum
}
