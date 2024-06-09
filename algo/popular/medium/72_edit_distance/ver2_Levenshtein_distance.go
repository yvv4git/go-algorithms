package main

import (
	"math"
)

func minDistanceV2(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cost := 0
			if word1[i-1] != word2[j-1] {
				cost = 1
			}
			dp[i][j] = int(math.Min(float64(dp[i-1][j]+1), math.Min(float64(dp[i][j-1]+1), float64(dp[i-1][j-1]+cost))))
			if i > 1 && j > 1 && word1[i-1] == word2[j-2] && word1[i-2] == word2[j-1] {
				dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i-2][j-2]+1)))
			}
		}
	}

	return dp[m][n]
}

//
//func main() {
//	fmt.Println(minDistance("kitten", "sitting"))
//}
