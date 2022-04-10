package numbersdisappearedinarray

func findDisappearedNumbers(nums []int) []int {
	abs := func(n int) int {
		if n > 0 {
			return n
		}

		return -n
	}

	ans := []int{}

	// Хитрая закономерность.
	// Перебираем числа и на месте, где они должны стоять, ставим минус.
	// Соответственно, все индексы, по которым остануться положительные числа
	// можно добавить в ответ, т.к. в массиве отсутствует соответствующее значение.
	for _, v := range nums {
		idx := abs(v) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	for i, v := range nums {
		if v > 0 {
			ans = append(ans, i+1)
		}
	}

	return ans
}
