package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	/*
		METHOD: Brute Force
		Метод заключается в проверке, равны ли элементы на диагоналях матрицы.
		Если элементы на диагоналях равны, то возвращаем true, иначе false.

		TIME COMPLEXITY: O(n^2), где n - количество строк и столбцов в матрице
		SPACE COMPLEXITY: O(1)
	*/
	// Получаем количество строк и столбцов в матрице
	rows := len(matrix)
	cols := len(matrix[0])

	// Проходим по всем элементам матрицы, кроме последней строки и последнего столбца
	for i := 0; i < rows-1; i++ {
		for j := 0; j < cols-1; j++ {
			// Проверяем, равен ли текущий элемент элементу на следующей диагонали
			if matrix[i][j] != matrix[i+1][j+1] {
				// Если не равны, возвращаем false
				return false
			}
		}
	}

	// Если все элементы на диагоналях равны, возвращаем true
	return true
}

func main() {
	// Пример матрицы
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}

	// Вызываем функцию и выводим результат
	fmt.Println(isToeplitzMatrix(matrix)) // Ожидаемый вывод: true
}
