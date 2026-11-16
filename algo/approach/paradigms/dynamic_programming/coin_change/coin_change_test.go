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
				// Рекуррентное соотношение:
				//   dp[i] = min(dp[i], dp[i-coin] + 1)
				//
				// Обоснование:
				//   Если мы уже знаем, сколько монет нужно для суммы (i-coin),
				//   то для суммы i можно взять этот набор и добавить одну монету coin.
				//   Получаем dp[i-coin] + 1 монет.
				//   Из всех таких вариантов (по разным coin) выбираем минимум.
				//
				// Пример: coins=[1,2,5], i=11
				//   coin=1: dp[10] + 1 (известно, что dp[10]=2 → 5+5 → 3 монеты)
				//   coin=2: dp[9]  + 1
				//   coin=5: dp[6]  + 1 (известно, что dp[6]=2 → 5+1 → 3 монеты)
				//   min = 3, то есть 5+5+1.
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
