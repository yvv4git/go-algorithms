package main

// Функция, которая выводит элементы матрицы по спирали
func spiralOrder(matrix [][]int) []int {
	/*
		METHOD: Iterative
		TIME COMPLEXITY: O(n), где n - общее количество элементов в матрице, так как мы проходим по каждому элементу ровно один раз.
		SPACE COMPLEXITY: O(n), так как мы сохраняем результат в новый срез.
	*/
	// Проверяем, что матрица не пуста
	if len(matrix) == 0 {
		return []int{}
	}

	// Инициализируем переменные для границ и результата
	var result []int
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	// Пока не закончится спираль
	for {
		// Выводим верхнюю строку
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		// Выводим правый столбец
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if left > right {
			break
		}

		// Выводим нижнюю строку
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}

		// Выводим левый столбец
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return result
}

//func main() {
//	matrix := [][]int{
//		{1, 2, 3},
//		{4, 5, 6},
//		{7, 8, 9},
//	}
//	fmt.Println(spiralOrder(matrix)) // Выводит: [1 2 3 6 9 8 7 4 5]
//}
