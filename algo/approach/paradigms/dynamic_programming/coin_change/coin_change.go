package main

import (
	"fmt"
	"math"
)

// coinChange - функция, которая вычисляет минимальное количество монет для набора заданной суммы.
// coins - массив номиналов монет (например, [1, 2, 5]).
// amount - сумма, которую нужно набрать (например, 11).
func coinChange(coins []int, amount int) int {
	/*
		METHOD: Dynamic Programming
		1. Инициализируем dp[], где dp[i] — минимальное количество монет для суммы i.
		2. Для каждой суммы i от 1 до amount пробуем все монеты и обновляем dp[i].
		3. Если dp[amount] равно бесконечности, значит решить задачу невозможно.

		TIME COMPLEXITY: O(N * S) (N — количество монет, S — сумма)
		SPACE COMPLEXITY: O(S) (массив dp[])
	*/

	// Инициализация dp массива с большими значениями
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0 // Для суммы 0 монет не нужно

	// Перебираем все суммы от 1 до amount
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// Если dp[amount] равно бесконечности, значит решение невозможно
	if dp[amount] == math.MaxInt {
		return -1
	}

	return dp[amount]
}

// Функция для нахождения минимального из двух чисел
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Даны монеты разных номиналов и сумма S. Нужно найти минимальное количество монет для суммы S.
func main() {
	// Пример монет и суммы
	coins := []int{1, 2, 5}
	amount := 11

	// Вызываем алгоритм
	result := coinChange(coins, amount)

	// Выводим результат
	fmt.Printf("Минимальное количество монет: %d\n", result)
}
