package search_matrix_2d

func searchMatrix(matrix [][]int, target int) bool {
	/*
		METHOD: Binary search.

		Time complexity : O(log(n*m))
		Space complexity : O(1)
	*/
	left, right := 0, len(matrix)*len(matrix[0])
	for left < right {
		middle := (left + right) / 2
		i, j := middle/len(matrix[0]), middle%len(matrix[0]) // Способ найти средний под массив(i), средний элемент в под массиве(j).

		if target < matrix[i][j] {
			right = middle
			continue
		}

		if target > matrix[i][j] {
			left = middle + 1
			continue
		}

		return true
	}

	return false
}
