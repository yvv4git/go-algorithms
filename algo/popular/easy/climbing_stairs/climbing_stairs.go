package main

func climbStairs(n int) int {
	return fibonacci(n + 1)
}

func fibonacci(n int) int {
	var x, y int = 0, 1
	for n > 0 {
		x, y = y, x+y
		n--
	}
	return x
}
