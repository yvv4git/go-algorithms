package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	/*
		METHOD: Dynamic programming
		Мы используем динамическое программирование для решения этой задачи. Основная идея заключается в том,
		чтобы создать массив dp, где dp[i] будет представлять минимальное количество монет, необходимое для набора суммы i.
		Мы заполняем этот массив, используя значения из предыдущих подзадач.

		TIME COMPLEXITY: O(n * k), где n — это сумма amount, а k — количество различных номиналов монет.
		Мы проходим по всем суммам от 1 до amount и для каждой суммы проверяем все номиналы монет.

		SPACE COMPLEXITY: O(n), где n — это сумма amount.
		Мы используем массив dp размером amount + 1 для хранения результатов подзадач.
	*/
	// Создаем массив dp размером amount + 1 и заполняем его значением math.MaxInt32,
	// чтобы обозначить, что эти значения еще не вычислены.
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	// Для суммы 0 не нужно ни одной монеты.
	dp[0] = 0

	// Для каждой суммы от 1 до amount вычисляем минимальное количество монет.
	for i := 1; i <= amount; i++ {
		// Для каждого номинала монеты проверяем, можно ли улучшить текущее значение dp[i].
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// Если dp[amount] равно math.MaxInt32, значит, невозможно набрать эту сумму.
	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

// Функция для нахождения минимума двух чисел.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount)) // Вывод: 3
}
