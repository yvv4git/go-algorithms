package _09_fibonacci_number

func fibonacciNumberIterative1(n int) int {
	/*
		Формула из задачи.
		F(0) = 0, F(1) = 1
		F(n) = F(n - 1) + F(n - 2), for n > 1.

		Вроде как последовательность чисел Фибоначчи, 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946,...

		Method: DP iterative
		Time complexity : O(n)
		Space complexity : O(1)
	*/
	if n < 2 {
		return n
	}

	a, b := 1, 1
	c := 1

	for i := 3; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}

	return c
}
