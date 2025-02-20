//go:build ignore

package main

import (
	"fmt"
)

// Функция для нахождения максимальной суммы подмассива
func maxSubArraySum(arr []int) int {
	/*
		METHOD: Dynamic Programming
		1. Для каждого элемента массива dp[i] сохраняем максимальную сумму подмассива, заканчивающегося на i.
		2. Заполняем dp[] по формуле: dp[i] = max(arr[i], arr[i] + dp[i-1]).
		3. Ответом будет максимальное значение в dp[].

		TIME COMPLEXITY: O(N) (один цикл по массиву)
		SPACE COMPLEXITY: O(N) (массив dp[])
	*/

	n := len(arr)
	if n == 0 {
		return 0
	}

	// Инициализируем dp массив
	dp := make([]int, n)
	dp[0] = arr[0]

	maxSum := dp[0]

	// Заполняем dp[] для всех элементов массива
	for i := 1; i < n; i++ {
		dp[i] = max(arr[i], arr[i]+dp[i-1])
		maxSum = max(maxSum, dp[i])
	}

	return maxSum
}

// Функция для нахождения максимума из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Дано массив чисел, нужно найти максимальную сумму подмассива.
func main() {
	// Пример массива
	arr := []int{1, -2, 3, 4, -1, 2, 1, -5, 4}

	// Вызываем алгоритм
	result := maxSubArraySum(arr)

	// Выводим результат
	fmt.Printf("Максимальная сумма подмассива: %d\n", result)
}
