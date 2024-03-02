package main

import "fmt"

// Функция generateMatrix генерирует двумерную матрицу размера n x n, заполненную по спирали.
func generateMatrix(n int) [][]int {
	/*
		METHOD: Iterative
		TIME COMPLEXITY: O(n^2), где n - размер матрицы. Это связано с тем, что алгоритм проходит по каждому элементу матрицы ровно один раз.
		SPACE COMPLEXITY: O(n^2), так как алгоритм создает новую матрицу размера n x n для хранения результата. Таким образом, пространственная сложность линейно зависит от размера входных данных.
	*/
	// Инициализация матрицы заполненной нулями.
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// Начальное значение для заполнения матрицы.
	num := 1
	// Определение границ для текущего слоя спирали.
	top, bottom, left, right := 0, n-1, 0, n-1

	// Заполнение матрицы по спирали.
	for num <= n*n {
		// Заполнение верхней строки слева направо.
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		// Заполнение правого столбца сверху вниз.
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		// Заполнение нижней строки справа налево.
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--

		// Заполнение левого столбца снизу вверх.
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix
}

func main() {
	// Размер матрицы.
	n := 3
	// Получение заполненной по спирали матрицы.
	result := generateMatrix(n)
	// Вывод матрицы на экран.
	for _, row := range result {
		fmt.Println(row)
	}
}
