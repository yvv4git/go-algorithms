package add_digits

func addDigitsV2(num int) int {
	/*
		Более сложный вариант решения задачи, нежели предыдущий V1.

		TIME COMPLEXITY: O(1)
		SPACE COMPLEXITY: O(1)
	*/
	if num <= 9 {
		return num // base case - база рекурсии
	} else { // Если же число состоит из более 1 знака.
		newVal := 0
		for num != 0 {
			newVal += num % 10
			num = num / 10
		}

		return addDigitsV2(newVal) // Рекурсивный случай
	}
}
