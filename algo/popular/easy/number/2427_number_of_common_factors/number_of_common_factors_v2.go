package _427_number_of_common_factors

func commonFactorsV2(a int, b int) int {
	// Это делается для того, чтобы числа были отсортированы в правильном порядке a < b.
	if a > b {
		a, b = b, a
	}

	// Отсчет начинаем с минимального из чисел.
	result := 0
	for i := a; i > 0; i-- {
		if b%i == 0 && a%i == 0 {
			result++
		}
	}

	return result
}
