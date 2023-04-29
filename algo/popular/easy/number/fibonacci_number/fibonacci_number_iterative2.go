package fibonacci_number

func fibonacciNumberIterative2(n int) int {
	/*
		Method: DP iterative
		Time complexity : O(n)
		Space complexity : O(1)
	*/
	// В массиве будут храниться 2 числа n-1 и n-2.s
	d := [2]int{0, 1}

	// Проверка, что надо будет что-то вычислять.
	if n <= 1 {
		return d[n]
	}

	for i := 2; i <= n; i++ {
		d[1], d[0] = d[1]+d[0], d[1]
	}

	return d[1]
}
