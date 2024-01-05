package _16_longest_palindromic_subsequence

// Функция longestPalindromeSubseqV2 находит длину наибольшей палиндромной подпоследовательности в строке s
func longestPalindromeSubseqV2(s string) int {
	/*
		Method: Dynamic programming + Recursion + Матрица смежности
		Time complexity: O(n^2)
		Space complexity: O(n^2)

		Time complexity
		Временная сложность этого алгоритма - O(n^2), где n - длина строки s.
		Это связано с двумя вложенными циклами, которые проходят по всем символам строки.

		Space complexity
		Пространственная сложность - O(n^2), так как мы используем двумерный массив dp для хранения результатов подсчета.
	*/

	// Создаем двумерный массив DP размером len(s) x len(s), заполненный -1
	DP := make([][]int, len(s)) // Двумерный массив или матрица смежности
	for i := 0; i < len(s); i++ {
		DP[i] = make([]int, len(s))
		for j := 0; j < len(s); j++ {
			DP[i][j] = -1
		}
	}
	// Вызываем функцию compute, которая вычисляет длину наибольшей палиндромной подпоследовательности
	return compute(s, 0, len(s)-1, DP)
}

// Функция max возвращает максимальное из двух чисел
func maxV2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Функция compute вычисляет длину наибольшей палиндромной подпоследовательности в строке s
func compute(s string, i, j int, DP [][]int) int {
	// Если индексы i и j равны, то возвращаем 1
	if j == i {
		return 1
	}
	// Если индекс j меньше i, то возвращаем 0
	if j < i {
		return 0
	}
	// Если значение DP[i][j] уже вычислено, то возвращаем его
	if DP[i][j] != -1 {
		return DP[i][j]
	}
	// Если символы i и j равны, то DP[i][j] = 2 + compute(s, i+1, j-1, DP)
	if s[i] == s[j] {
		DP[i][j] = 2 + compute(s, i+1, j-1, DP)
	} else {
		// Если символы i и j не равны, то DP[i][j] = max(compute(s, i+1, j, DP), compute(s, i, j-1, DP))
		DP[i][j] = max(compute(s, i+1, j, DP), compute(s, i, j-1, DP))
	}
	// Возвращаем значение DP[i][j]
	return DP[i][j]
}
