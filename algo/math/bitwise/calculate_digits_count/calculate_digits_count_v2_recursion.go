package calculate_digits_count

func CalculateDigitsNumberV2(num int) int {
	/*
		Method Recursion
		TIME COMPLEXITY: O(log(n))
		SPACE COMPLEXITY: O(log(n))
	*/
	if num/10 == 0 {
		return 1
	}

	return 1 + CalculateDigitsNumberV2(num/10)
}
