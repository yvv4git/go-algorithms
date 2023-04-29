package fibonacci

func fibonacciClosure() func() int {
	x1, x2 := -1, 1
	return func() int {
		x1, x2 = x2, x1+x2
		return x2
	}
}

func fibonacciByClousure(n int) int {
	result := 0
	fibo := fibonacciClosure()

	for i := 0; i <= n; i++ {
		result = fibo()
	}

	return result
}
