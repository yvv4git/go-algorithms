package calculate_digits_count

func CalculateDigitsNumberV1(num int) int {
	/*
		METHOD: Iterative
		TIME COMPLEXITY: O(log_10(n)) or O(num digits)
		Space complexity: O(1) or constant
	*/
	if num == 0 {
		return 1
	}

	var count int
	for num > 0 {
		num /= 10 // Убираем старший разряд(слева). Например, было 1500, станет 500
		count++   // Считаем сколько у нас разрядов в исходном числе
	}

	return count
}
