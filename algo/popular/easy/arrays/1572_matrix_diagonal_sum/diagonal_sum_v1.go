package main

func diagonalSumV1(mat [][]int) int {
	/*
		METHOD: Iterate
		TIME COMPLEXITY: O(1)
		SPACE COMPLEXITY: O(n)
	*/
	var (
		n      = len(mat)
		result = 0
	)

	for i := 0; i < n; i++ {
		result += mat[i][i] + mat[n-i-1][i]
	}

	if n%2 != 0 {
		result -= mat[n/2][n/2]
	}

	return result
}
