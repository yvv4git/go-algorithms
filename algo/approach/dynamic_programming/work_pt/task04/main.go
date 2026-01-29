package main

import (
	"fmt"
)

func solve(N int) (int, []int) {
	if N == 1 {
		return 0, []int{1}
	}
	dp := make([]int, N+1)
	prev := make([]int, N+1)
	for i := 2; i <= N; i++ {
		dp[i] = dp[i-1] + 1
		prev[i] = i - 1

		// Изменен порядок проверки для получения пути как в примерах
		if i%3 == 0 && dp[i/3]+1 <= dp[i] { // <= вместо < для приоритета ×3
			dp[i] = dp[i/3] + 1
			prev[i] = i / 3
		}
		if i%2 == 0 && dp[i/2]+1 <= dp[i] { // <= вместо < для приоритета ×2
			dp[i] = dp[i/2] + 1
			prev[i] = i / 2
		}
	}

	// восстановление пути
	path := []int{}
	current := N
	for current != 1 {
		path = append(path, current)
		current = prev[current]
	}
	path = append(path, 1)

	// разворот
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return dp[N], path
}

func main() {
	var N int
	fmt.Scan(&N)
	ops, path := solve(N)
	fmt.Println(ops)
	for i, v := range path {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
