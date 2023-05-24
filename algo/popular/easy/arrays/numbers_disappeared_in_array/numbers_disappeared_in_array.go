package numbersdisappearedinarray

func findDisappearedNumbers(nums []int) []int {
	/*
		Method: Arithmetic
		Time complexity: O(n)
		Space complexity: O(n) - из-за того, что используем дополнительный массив result, где хранится ответ
	*/
	abs := func(n int) int {
		if n > 0 {
			return n
		}

		return -n
	}

	result := []int{}

	// Хитрая закономерность.
	// Перебираем числа и на месте, где они должны стоять, ставим минус.
	// Соответственно, все индексы, по которым останутся положительные числа
	// можно добавить в ответ, т.к. в массиве отсутствует соответствующее значение.
	for _, v := range nums {
		idx := abs(v) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	for i, v := range nums {
		if v > 0 {
			result = append(result, i+1)
		}
	}

	return result
}
