package main

import (
	"fmt"
	"testing"
)

func TestKnapsack(t *testing.T) {
	tests := []struct {
		name     string
		weights  []int
		values   []int
		capacity int
		want     int
	}{
		{
			name:     "example from main",
			weights:  []int{2, 3, 4, 5},
			values:   []int{3, 4, 5, 6},
			capacity: 5,
			want:     7,
		},
		{
			name:     "empty items",
			weights:  []int{},
			values:   []int{},
			capacity: 10,
			want:     0,
		},
		{
			name:     "zero capacity",
			weights:  []int{2, 3},
			values:   []int{3, 4},
			capacity: 0,
			want:     0,
		},
		{
			name:     "one item fits",
			weights:  []int{5},
			values:   []int{10},
			capacity: 5,
			want:     10,
		},
		{
			name:     "one item too heavy",
			weights:  []int{5},
			values:   []int{10},
			capacity: 3,
			want:     0,
		},
		{
			name:     "all items fit",
			weights:  []int{1, 2, 3},
			values:   []int{2, 3, 5},
			capacity: 10,
			want:     10,
		},
		{
			name:     "choose best combination",
			weights:  []int{2, 3, 4, 5},
			values:   []int{3, 7, 2, 9},
			capacity: 5,
			want:     10,
		},
		{
			name:     "single item exact fit",
			weights:  []int{10},
			values:   []int{100},
			capacity: 10,
			want:     100,
		},
		{
			name:     "multiple same weight pick highest value",
			weights:  []int{3, 3, 3},
			values:   []int{5, 10, 7},
			capacity: 3,
			want:     10,
		},
		{
			name:     "large capacity all items fit",
			weights:  []int{1, 2, 3, 4, 5},
			values:   []int{1, 2, 3, 4, 5},
			capacity: 20,
			want:     15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := knapsack(tt.weights, tt.values, tt.capacity)
			if got != tt.want {
				t.Errorf("knapsack(%v, %v, %d) = %d, want %d",
					tt.weights, tt.values, tt.capacity, got, tt.want)
			}
		})
	}
}

func TestKnapsack_Research(t *testing.T) {
	weights := []int{1, 2, 3} // веса предметов
	values := []int{2, 3, 5}  // ценность предметов
	capacity := 5             // емкость рюкзака

	n := len(weights)

	// Создаем DP-таблицу (инициализирована нулями)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ { // i — сколько первых предметов рассматриваем
		for w := 0; w <= capacity; w++ { // w — текущая вместимость рюкзака
			if weights[i-1] > w {
				// Если предмет не помещается, просто пропускаем его
				dp[i][w] = dp[i-1][w]
			} else {
				// Выбираем максимум между "взять предмет" и "не взять"
				dp[i][w] = max(dp[i-1][w], values[i-1]+dp[i-1][w-weights[i-1]])
			}

			fmt.Printf("i=%d w=%d dp=%v\n", i, w, dp)
		}
	}
}
