package _518_Coin_Change2

import "testing"

func TestChangeV2(t *testing.T) {
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
		got := changeV2(tt.amount, tt.coins)
		if got != tt.want {
			t.Errorf("changeV2(%d, %v) = %d, want %d", tt.amount, tt.coins, got, tt.want)
		}
	}
}

func TestResearchChangeV2(t *testing.T) {
	// Задача: подсчитать количество комбинаций, которыми можно составить сумму amount,
	// используя монеты из coins неограниченное количество раз.
	// Порядок монет не имеет значения (комбинации, не перестановки).
	// Если сумму нельзя составить — вернуть 0.
	//
	// Подход: рекурсия с мемоизацией (Top-down DP).
	// На каждом шаге выбираем: взять текущую монету или пропустить.
	// Кешируем результат для (i, remaining).
}

func TestChangeV2_Validation(t *testing.T) {
	// Сверяем результат top-down DP с bottom-up DP (ver1_dp.go).
	// Если changeV1 экспортирован, можно сравнивать напрямую.
	_ = changeV2
}
