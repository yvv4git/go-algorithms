package _18_pascal_triangle

func generateV3(numRows int) [][]int {
	/*
		Method: Iterative
		Time complexity : O(n)
		Space complexity : O(n)
	*/
	res := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
	}

	res[0][0] = 1

	for r := 1; r < numRows; r++ {
		for c := 0; c < len(res[r]); c++ {
			if c == 0 || c == len(res[r])-1 {
				res[r][c] = 1
			} else {
				res[r][c] = res[r-1][c-1] + res[r-1][c]
			}
		}
	}

	return res
}
