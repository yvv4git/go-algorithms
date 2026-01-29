package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, k int
	if _, err := fmt.Fscan(in, &n, &k); err != nil {
		return
	}

	fmt.Println(countWays(n, k))
}

func countWays(n, k int) int {
	/*
		METHOD: Dynamic Programming (одномерный массив)
		Суть задачи: Кузнечик прыгает по доске из N клеток, может прыгнуть на 1..k клеток вперёд.
		Нужно найти количество способов добраться из первой клетки в последнюю.

		dp[i] - количество способов добраться до клетки i.
		Переход: dp[i] = sum(dp[i-j]) для всех j от 1 до k, где i-j >= 1.

		TIME COMPLEXITY: O(n*k) - для каждой клетки проверяем до k предыдущих
		SPACE COMPLEXITY: O(n) - массив dp размера n+1
	*/

	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if i-j >= 1 {
				dp[i] += dp[i-j]
			}
		}
	}

	return dp[n]
}
