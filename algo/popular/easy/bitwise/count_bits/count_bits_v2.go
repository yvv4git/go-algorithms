package count_bits

func countBitsV2(n int) []int {
	/*
		Функция позволяет посчитать количество 1 в бинарном представлении числа.
	*/
	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		result[i] = cntOnes(i)
	}

	return result
}

// Как посчитать количество 1 в бинарном представлении числа?
func cntOnes(number int) int {
	count := 0
	for number > 0 {
		count += number & 1 // number & 1 возвращает 1, если число оканчивается на 1 и возвращает 0, если число оканчивается на 0.
		number >>= 1
	}

	return count
}
