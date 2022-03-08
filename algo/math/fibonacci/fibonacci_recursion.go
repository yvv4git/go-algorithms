package fibonacci

// fibonacciByRecursion - находит по порядковому номеру значение числа Фибоначи
func fibonacciByRecursion(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacciByRecursion(n-1) + fibonacciByRecursion(n-2)
}
