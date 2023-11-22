package _02_happy_number

func isHappyV2(n int) bool {
	/*
		Time complexity: O(log n)
		Space complexity: O(1)
	*/
	slow, fast := n, n
	for slow, fast = squareSum(slow), squareSum(squareSum(fast)); slow != fast; slow, fast = squareSum(slow), squareSum(squareSum(fast)) {
	}
	return slow == 1
}

func squareSum(n int) int {
	if n == 1 {
		return 1
	}
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
