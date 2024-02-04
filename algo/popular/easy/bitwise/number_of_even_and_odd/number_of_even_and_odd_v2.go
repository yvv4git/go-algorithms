package number_of_even_and_odd

func evenOddBitV2(n int) []int {
	/*
		METHOD: Hash
		TIME COMPLEXITY: O(Log(n))
		SPACE COMPLEXITY: O(1) - хоть мы и используем хэш, все равно размер хэша не зависит от n, это постоянное значение, константа

		В бинарном представлении even(четный) - 1, odd(нечетный) - 0
	*/
	idx := 0
	result := make([]int, 2)
	for n > 0 {
		lastDigit := n % 2 // Находим остаток от деления - это последний бит
		n /= 2             // Эквивалент Right shift, т.е. сдвигу вправо
		if lastDigit == 1 {
			result[idx%2] += 1
		}
		idx += 1
	}

	return result
}
