package _18_pascal_triangle

func generateV4(numRows int) [][]int {
	/*
		METHOD: Dynamic programming
		Time complexity : O(n)
		Space complexity : O(n)
	*/
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}

	return result
}
