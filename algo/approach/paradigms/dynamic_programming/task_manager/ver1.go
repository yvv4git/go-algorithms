package main

import (
	"fmt"
	"math"
)

// optimalTaskDistribution вычисляет минимальное время, за которое можно распределить задачи между сотрудниками.
func optimalTaskDistribution(n int, k int, tasks [][]int) int {
	/*
		METHOD: Dynamic programming
		TIME COMPLEXITY: O(n^2 * k), где n - количество задач, k - количество сотрудников.
		SPACE COMPLEXITY: O(n * k), так как мы используем дополнительный массив dp для хранения результатов
	*/
	// Инициализация таблицы dp
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	// Базовый случай: если задач нет, то время равно 0
	dp[0][0] = 0

	// Заполнение таблицы dp
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			maxTime := 0
			for m := i - 1; m >= 0; m-- {
				maxTime += tasks[m][1]
				dp[i][j] = min(dp[i][j], max(dp[m][j-1], maxTime))
			}
		}
	}

	// Результат будет находиться в dp[n][k]
	return dp[n][k]
}

// min возвращает минимальное значение из двух
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// max возвращает максимальное значение из двух
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	// Пример входных данных
	tasks := [][]int{
		{10, 3},
		{20, 2},
		{30, 4},
		{40, 1},
		{50, 5},
	}
	n := 5
	k := 2

	// Вызов функции и вывод результата
	result := optimalTaskDistribution(n, k, tasks)
	fmt.Println("Минимальное время выполнения задач:", result)
}
