package gcb

func gcb(m, n int) int {
	if m < n {
		return gcb(n, m)
	}

	if m%n == 0 {
		return n
	}

	return gcb(n, m%n)
}
