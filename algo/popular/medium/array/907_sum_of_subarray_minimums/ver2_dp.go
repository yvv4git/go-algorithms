package _07_sum_of_subarray_minimums

func sumSubarrayMinsV2(A []int) int {
	/*
		Method: Dynamic programming
		Time complexity: O(n^2), где n - размер входного массива. Это связано с двумя вложенными циклами, которые проходят по всем элементам массива.
		Space complexity: O(n^2), так как мы создаем матрицу размера n x n для хранения минимальных элементов подмассивов.
	*/
	MOD := 1000000007
	N := len(A)

	// Создаем матрицу для хранения минимальных элементов подмассивов
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, N)
	}

	// Заполняем матрицу
	for i := 0; i < N; i++ {
		dp[i][i] = A[i]
		for j := i + 1; j < N; j++ {
			dp[i][j] = min(dp[i][j-1], A[j])
		}
	}

	// Суммируем минимальные элементы подмассивов
	sum := 0
	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			sum = (sum + dp[i][j]) % MOD
		}
	}

	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
