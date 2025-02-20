//go:build ignore

package main

import (
	"fmt"
)

// Функция для решения задачи 0/1 рюкзака с использованием динамического программирования
func knapsack(weights []int, values []int, capacity int) int {
	/*
		METHOD: Dynamic Programming
		1. Создаем DP-таблицу dp[i][w], где i – количество предметов, w – текущий вес рюкзака.
		2. Заполняем таблицу, выбирая максимум из двух вариантов: взять или не взять предмет.

		TIME COMPLEXITY: O(N * W) (проходим по всем предметам и всем возможным весам)
		SPACE COMPLEXITY: O(N * W) (матрица для хранения результатов)
	*/

	n := len(weights)

	// Создаем DP-таблицу (инициализирована нулями)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// Заполняем DP-таблицу
	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if weights[i-1] > w {
				// Если предмет не помещается, просто пропускаем его
				dp[i][w] = dp[i-1][w]
			} else {
				// Выбираем максимум между "взять предмет" и "не взять"
				dp[i][w] = max(dp[i-1][w], values[i-1]+dp[i-1][w-weights[i-1]])
			}
		}
	}

	// Ответ – максимальная ценность, которую можно уместить в рюкзак
	return dp[n][capacity]
}

// Функция для нахождения максимума из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Дано N предметов с весом w_i и ценностью v_i. Нужно максимизировать ценность в рюкзаке ёмкостью W.
func main() {
	// Пример предметов (вес, ценность)
	weights := []int{2, 3, 4, 5}
	values := []int{3, 4, 5, 6}
	capacity := 5 // Вместимость рюкзака

	// Вызываем алгоритм
	maxValue := knapsack(weights, values, capacity)

	// Выводим результат
	fmt.Printf("Максимальная ценность рюкзака: %d\n", maxValue)
}
