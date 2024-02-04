package _67_valid_perfect_square

func isPerfectSquareV2(num int) bool {
	/*
		METHOD: Iterative
		TIME COMPLEXITY: O(sqrt(n))
		SPACE COMPLEXITY: O(1)
	*/
	i := 1
	for i*i < num {
		i++
	}

	return i*i == num
}
