package _518_Coin_Change2

import (
	"fmt"
	"testing"
)

func TestChangeV1(t *testing.T) {
	tests := []struct {
		amount int
		coins  []int
		want   int
	}{
		{amount: 5, coins: []int{1, 2, 5}, want: 4},
		{amount: 3, coins: []int{2}, want: 0},
		{amount: 10, coins: []int{10}, want: 1},
		{amount: 0, coins: []int{1}, want: 1},
		{amount: 7, coins: []int{1, 2, 3}, want: 8},
		{amount: 11, coins: []int{5, 7}, want: 0},
	}

	for _, tt := range tests {
		got := changeV1(tt.amount, tt.coins)
		if got != tt.want {
			t.Errorf("changeV1(%d, %v) = %d, want %d", tt.amount, tt.coins, got, tt.want)
		}
	}
}

func TestResearch_ChangeV1(t *testing.T) {
	// Задача: подсчитать количество комбинаций, которыми можно составить сумму amount,
	// используя монеты из coins неограниченное количество раз.
	// Порядок монет не имеет значения (комбинации, не перестановки).
	// Если сумму нельзя составить — вернуть 0.
	//
	// Пример: amount = 5, coins = [1, 2, 5]
	// Комбинации:
	//   5 = 5
	//   5 = 2 + 2 + 1
	//   5 = 2 + 1 + 1 + 1
	//   5 = 1 + 1 + 1 + 1 + 1
	// Итого: 4 комбинации
	amount := 5
	coins := []int{1, 2, 5}
	dp := make([]int, amount+1)
	dp[0] = 1 // Количество способов имеющимися монетами составить сумму 0

	i := 1

	for _, coin := range coins {
		for a := coin; a <= amount; a++ {
			dp[a] += dp[a-coin]
			fmt.Printf("[%d] coin=%d a=%d a-coin=%d dp=%v \n", i, coin, a, a-coin, dp)
			_ = 0
			i++ // only iterator for I know what is step number
		}
	}
}

func TestResearch_ChangeV1_2(t *testing.T) {
	amount := 5
	coins := []int{1, 2, 3}
	dp := make([]int, amount+1)
	dp[0] = 1 // Количество способов имеющимися монетами составить сумму 0

	i := 1

	for _, coin := range coins {
		for a := coin; a <= amount; a++ {
			dp[a] += dp[a-coin]
			fmt.Printf("[%d] coin=%d a=%d a-coin=%d dp=%v \n", i, coin, a, a-coin, dp)
			_ = 0
			i++ // only iterator for I know what is step number
		}
	}
}
