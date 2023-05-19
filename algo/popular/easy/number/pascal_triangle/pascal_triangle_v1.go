package pascal_triangle

func generateV1(numRows int) [][]int {
	return fill(generateBase(numRows))
}

func generateBase(numRows int) [][]int {
	sol := make([][]int, numRows)

	for col := 0; col < numRows; col++ {
		sol[col] = make([]int, col+1)
		sol[col][0], sol[col][col] = 1, 1
	}

	sol[0][0] = 1

	return sol
}

func fill(base [][]int) [][]int {
	numRows := len(base)
	for row := 1; row < numRows; row++ {
		for col := 1; col < len(base[row])-1; col++ {
			base[row][col] = base[row-1][col-1] + base[row-1][col]
		}
	}
	return base
}
