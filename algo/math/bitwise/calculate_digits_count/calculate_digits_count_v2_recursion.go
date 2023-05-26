package calculate_digits_count

func CalculateDigitsNumberV2(num int) int {
	/*
		Method Recursion
		Time complexity: O(log(n))
		Space complexity: O(log(n))
	*/
	if num/10 == 0 {
		return 1
	}

	return 1 + CalculateDigitsNumberV2(num/10)
}
