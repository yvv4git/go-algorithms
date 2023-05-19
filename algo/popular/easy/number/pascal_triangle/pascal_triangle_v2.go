package pascal_triangle

func generateV2(numRows int) [][]int {
	sol := make([][]int, 0, numRows)
	if numRows == 0 {
		return sol
	}

	sol = append(sol, []int{1})
	if numRows == 1 {
		return sol
	}

	sol = append(sol, []int{1, 1})
	if numRows == 2 {
		return sol
	}

	for i := 2; i < numRows; i++ {
		top, row := sol[len(sol)-1], NewRow(i+1)
		for j := 1; j < i; j++ {
			row[j] = top[j-1] + top[j]
		}
		sol = append(sol, row)
	}

	return sol
}

func NewRow(size int) []int {
	row := make([]int, size)
	row[0], row[size-1] = 1, 1
	return row
}
