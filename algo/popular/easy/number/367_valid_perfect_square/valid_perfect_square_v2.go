package _67_valid_perfect_square

func isPerfectSquareV2(num int) bool {
	/*
		Method: Iterative
		Time complexity: O(sqrt(n))
		Space complexity: O(1)
	*/
	i := 1
	for i*i < num {
		i++
	}

	return i*i == num
}
