//go:build ignore

package main

import (
	"fmt"
)

// Функция для нахождения длины наибольшей общей подпоследовательности
func longestCommonSubsequence(str1, str2 string) int {
	/*
		METHOD: Dynamic Programming
		1. Создаем двумерный массив dp[][], где dp[i][j] — длина LCS для первых i символов str1 и j символов str2.
		2. Если символы совпадают, увеличиваем длину LCS на 1, иначе выбираем максимальное значение из двух возможных.
		3. Ответом будет dp[len(str1)][len(str2)].

		TIME COMPLEXITY: O(n * m) (где n — длина первой строки, m — длина второй)
		SPACE COMPLEXITY: O(n * m) (двумерный массив dp)
	*/

	n := len(str1)
	m := len(str2)

	// Создаем двумерный массив для хранения длин LCS
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// Заполняем dp[][], рассчитывая длину LCS
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Возвращаем длину LCS
	return dp[n][m]
}

// Функция для нахождения максимума из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Даны две строки, необходимо найти наибольшую общую подпоследовательность.
func main() {
	// Пример строк
	str1 := "ABCBDAB"
	str2 := "BDCABB"

	// Вызываем алгоритм
	result := longestCommonSubsequence(str1, str2)

	// Выводим результат
	fmt.Printf("Длина наибольшей общей подпоследовательности: %d\n", result)
}
