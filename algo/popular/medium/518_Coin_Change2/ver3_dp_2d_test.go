package _518_Coin_Change2

import "testing"

func TestChangeV3(t *testing.T) {
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
		got := changeV3(tt.amount, tt.coins)
		if got != tt.want {
			t.Errorf("changeV3(%d, %v) = %d, want %d", tt.amount, tt.coins, got, tt.want)
		}
	}
}
