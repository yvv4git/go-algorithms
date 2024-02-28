package main

import "fmt"

func spiralOrderV2(matrix [][]int) []int {
	/*
		METHOD: Recursive
		TIME COMPLEXITY: O(n), где n - общее количество элементов в матрице, так как мы проходим по каждому элементу ровно один раз.
		SPACE COMPLEXITY: O(n), так как мы сохраняем результат в новый срез.
	*/
	if len(matrix) == 0 {
		return []int{}
	}

	var result []int
	spiralHelper(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1, &result)
	return result
}

func spiralHelper(matrix [][]int, top, bottom, left, right int, result *[]int) {
	if top > bottom || left > right {
		return
	}

	// Выводим верхнюю строку
	for i := left; i <= right; i++ {
		*result = append(*result, matrix[top][i])
	}
	top++

	// Выводим правый столбец
	for i := top; i <= bottom; i++ {
		*result = append(*result, matrix[i][right])
	}
	right--

	// Если спираль не закончилась, выводим нижнюю строку
	if top <= bottom {
		for i := right; i >= left; i-- {
			*result = append(*result, matrix[bottom][i])
		}
		bottom--
	}

	// Если спираль не закончилась, выводим левый столбец
	if left <= right {
		for i := bottom; i >= top; i-- {
			*result = append(*result, matrix[i][left])
		}
		left++
	}

	spiralHelper(matrix, top, bottom, left, right, result)
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(matrix)) // Выводит: [1 2 3 6 9 8 7 4 5]
}
