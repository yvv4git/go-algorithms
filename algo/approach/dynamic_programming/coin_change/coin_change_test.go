package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestCointChange(t *testing.T) {
	coinChangeDebug([]int{1, 2, 5}, 11) // 3
}

func coinChangeDebug(coins []int, amount int) {
	res := coinChangeResearch(coins, amount)
	fmt.Println("Res: ", res)
}

func coinChangeResearch(coins []int, amount int) int {
	// Инициализация dp массива с большими значениями
	// Инициализируем dp[i] как math.MaxInt для всех i от 1 до amount,
	// чтобы обозначить, что изначально решение не найдено (бесконечность).
	// dp[0] = 0, поскольку для суммы 0 монет не требуется.
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0 // Для суммы 0 монет не нужно

	spew.Dump(dp)

	// Перебираем все суммы от 1 до amount
	for i := 1; i <= amount; i++ {
		// Для каждой суммы i (от 1 до amount) мы перебираем все доступные монеты coin.
		for _, coin := range coins {
			// Проверяем, что номинал монеты не превышает текущую сумму i,
			// чтобы избежать отрицательных индексов в dp[i-coin].
			if coin <= i {
				// Обновляем dp[i], беря минимум между текущим значением и dp[i-coin] + 1,
				// что означает добавление одной монеты coin к решению для суммы i-coin.
				dp[i] = min(dp[i], dp[i-coin]+1)
				fmt.Printf("dp[%d] = %v coin = %d \n", i, dp[i], coin)
				fmt.Println("dp: ", dp)
			}
		}
	}

	// Если dp[amount] равно бесконечности, значит решение невозможно
	if dp[amount] == math.MaxInt {
		return -1
	}

	// Возвращаем минимальное количество монет для суммы amount.
	return dp[amount]
}
