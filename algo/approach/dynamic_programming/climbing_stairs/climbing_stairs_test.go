package main

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "n=0", n: 0, want: 1},
		{name: "n=1", n: 1, want: 1},
		{name: "n=2", n: 2, want: 2},
		{name: "n=3", n: 3, want: 3},
		{name: "n=4", n: 4, want: 5},
		{name: "n=5", n: 5, want: 8},
		{name: "n=6", n: 6, want: 13},
		{name: "n=7", n: 7, want: 21},
		{name: "n=10", n: 10, want: 89},
		{name: "n=20", n: 20, want: 10946},
		{name: "n=30", n: 30, want: 1346269},
		{name: "negative", n: -5, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)
			if got != tt.want {
				t.Errorf("climbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestClimbStairs_Research(t *testing.T) {
	// Задача: есть лестница из n ступенек.
	// За шаг можно подняться на 1 или 2 ступеньки.
	// Нужно посчитать, сколько всего способов добраться до вершины.
	//
	// Решение:
	// dp[i] — количество способов подняться на i-ю ступеньку.
	// На i-ю ступеньку можно прийти либо с (i-1), либо с (i-2).
	// Значит: dp[i] = dp[i-1] + dp[i-2] есть рекурентное соотношение (Recurrence relation).
	// В контексте DP его ещё называют переходом (transition) — формула, которая связывает состояние dp[i] с предыдущими состояниями.
	// Это то же самое, что числа Фибоначчи, но в задаче про способы.
	//
	// Начальные значения:
	// dp[0] = 1 (стоять внизу — 1 способ, ничего не делать)
	// dp[1] = 1 (только шаг на 1)

	n := 10

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	fmt.Println("Заполняем DP-таблицу:")
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]

		// i — текущая ступенька
		// dp[i-1] — способы добраться до предыдущей
		// dp[i-2] — способы добраться до пред-предыдущей
	}

	fmt.Println("Готовая DP-таблица:")
	for i, v := range dp {
		fmt.Printf("  dp[%d] = %d способов\n", i, v)
	}
	fmt.Printf("Ответ: %d способов подняться на %d ступенек\n", dp[n], n)
}

func TestClimbStairs_EdgeCases(t *testing.T) {
	if got := climbStairs(1); got != 1 {
		t.Errorf("climbStairs(1) = %d, want 1", got)
	}
	if got := climbStairs(0); got != 1 {
		t.Errorf("climbStairs(0) = %d, want 1", got)
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		climbStairs(30)
	}
}
