package __longest_substring_without_repeating_characters

func lengthOfLongestSubstringV4(s string) int {
	/*
		METHOD: Dynamic programming
		TIME COMPLEXITY: O(n)
		Space complexity: O(1) or O(m), где m - длина алфавита
	*/
	// result - результат, который мы будем возвращать
	result := 0

	// last - массив, в котором мы храним последнее вхождение каждого символа
	last := [256]int{}
	for i := range last {
		last[i] = -1
	}

	// size - размер текущей подстроки без повторяющихся символов
	size := 0

	// min - функция, которая возвращает минимальное из двух чисел
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// Проходим по всем символам строки
	for i, b := range []byte(s) {
		// Обновляем размер текущей подстроки
		size = min(size+1, i-last[b])
		// Обновляем последнее вхождение текущего символа
		last[b] = i
		// Обновляем результат
		result = result + size - min(result, size) // max(result, size)
	}

	// Возвращаем результат
	return result
}
