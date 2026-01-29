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
