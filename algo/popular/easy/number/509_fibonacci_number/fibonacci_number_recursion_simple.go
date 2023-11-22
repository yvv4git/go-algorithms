package _09_fibonacci_number

func fibonacciNumberRecursion(n int) (res int) {
	/*
		Method: Recursion
		Time complexity : O(n)
		Space complexity : O(1)
	*/
	if n < 2 {
		return n
	}

	return fibonacciNumberRecursion(n-1) + fibonacciNumberRecursion(n-2)
}
