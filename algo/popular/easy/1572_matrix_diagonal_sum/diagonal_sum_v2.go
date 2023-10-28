package main

func diagonalSumV2(mat [][]int) int {
	sum := 0

	i := len(mat) - 1
	j := 0

	for i < len(mat) && j < len(mat) {
		sum += mat[i][j] + mat[j][j]

		if i == j {
			sum -= mat[i][j]
		}

		i--
		j++
	}

	return sum
}
