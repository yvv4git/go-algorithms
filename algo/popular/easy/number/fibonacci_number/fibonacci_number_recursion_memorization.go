package fibonacci_number

func fibonacciNumberRecursionMemorization(n int) int {
	/*
		Method: Recursion
		Time complexity : O(n)
		Space complexity : O(n)
	*/
	sMap := make(map[int]int)
	return fibWithMem(n, sMap)

}

func fibWithMem(n int, memo map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}

	val, ok := memo[n]
	if ok {
		return val
	}
	memo[n] = fibWithMem(n-1, memo) + fibWithMem(n-2, memo)
	return fibWithMem(n, memo)
}
