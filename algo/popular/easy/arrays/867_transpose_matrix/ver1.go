package main

import "fmt"

// Функция для транспонирования матрицы
func transpose(matrix [][]int) [][]int {
	/*
		METHOD: Math

		TIME COMPLEXITY: O(m*n), где m и n - количество строк и столбцов исходной матрицы соответственно.
		Это связано с тем, что мы проходим по каждому элементу матрицы.

		SPACE COMPLEXITY: O(m*n), так как мы создаем новую матрицу, которая имеет тот же размер, что и исходная матрица.
	*/
	// Получаем размеры исходной матрицы
	rows := len(matrix)
	cols := len(matrix[0])

	// Создаем новую матрицу с размером, равным количеству столбцов исходной матрицы на количество строк исходной матрицы
	transposed := make([][]int, cols)
	for i := range transposed {
		transposed[i] = make([]int, rows)
	}

	// Проходим по каждому элементу исходной матрицы и помещаем его в соответствующую позицию в транспонированной матрице
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	// Возвращаем транспонированную матрицу
	return transposed
}

func main() {
	// Исходная матрица
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Транспонируем матрицу
	transposed := transpose(matrix)

	// Выводим транспонированную матрицу
	for _, row := range transposed {
		fmt.Println(row)
	}
}
