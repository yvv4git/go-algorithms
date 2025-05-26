package main

type FactorialResult struct {
	Num       int
	Factorial int64
}

func factorialConcurrency(n int, ch chan FactorialResult) {
	var result int64 = 1

	for i := 1; i <= n; i++ {
		result *= int64(i)
	}

	ch <- FactorialResult{Num: n, Factorial: result}
}
