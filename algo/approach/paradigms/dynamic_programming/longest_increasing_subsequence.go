//go:build ignore

package main

import (
	"fmt"
)

// Функция для нахождения длины наибольшей возрастающей подпоследовательности
func longestIncreasingSubsequence(arr []int) int {
	/*
		METHOD: Dynamic Programming
		1. Создаем массив dp[], где dp[i] — длина LIS, заканчивающейся в arr[i].
		2. Инициализируем dp[i] = 1, так как каждый элемент сам по себе — LIS длины 1.
		3. Для каждого i проверяем все j < i и, если arr[j] < arr[i], обновляем dp[i].
		4. Возвращаем максимальное значение из dp[].

		TIME COMPLEXITY: O(N^2) (вложенные циклы)
		SPACE COMPLEXITY: O(N) (массив dp[])
	*/

	n := len(arr)
	if n == 0 {
		return 0
	}

	// DP-массив, где dp[i] — длина LIS, заканчивающейся в i-м элементе
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	// Заполняем dp[]
	maxLIS := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLIS = max(maxLIS, dp[i])
	}

	return maxLIS
}

// Функция для нахождения максимума из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Дано N чисел, требуется найти длину наибольшей возрастающей подпоследовательности (LIS).
func main() {
	// Пример массива
	arr := []int{10, 22, 9, 33, 21, 50, 41, 60, 80}

	// Вызываем алгоритм
	result := longestIncreasingSubsequence(arr)

	// Выводим результат
	fmt.Printf("Длина LIS: %d\n", result)
}
