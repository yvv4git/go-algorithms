package search_matrix_2d

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	/*
		Method: Binary search.

		Time complexity : O(log(n*m))
		Space complexity : O(1)
	*/
	fmt.Printf("=>matrix=%#v \n\n", matrix)

	left, right := 0, len(matrix)*len(matrix[0])
	for left < right {
		middle := (left + right) / 2
		fmt.Printf("=>left=%d right=%d middle=%d \n", left, right, middle)
		i, j := middle/len(matrix[0]), middle%len(matrix[0]) // Способ найти средний под массив(i), средний элемент в под массиве(j).
		fmt.Printf("=>i=%d, j=%d \n", i, j)

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
