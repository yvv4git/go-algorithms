package main

func searchMatrix(matrix [][]int, target int) bool {
	/*
		METHOD: Binary search
		Двоичный поиск с использованием свойств отсортированности&
		Почему выбран этот подход:
		- Матрица отсортирована как по строкам, так и по столбцам, что позволяет нам использовать эти свойства для оптимизации поиска.
		- Начиная с правого верхнего угла, мы можем эффективно сужать область поиска,
		перемещаясь влево или вниз в зависимости от значения текущего элемента и целевого значения.
		- Этот подход обеспечивает линейное время выполнения, что является оптимальным для данной задачи.

		Этот метод позволяет нам избежать полного перебора всех элементов матрицы, что делает решение более эффективным.

		TIME COMPLEXITY: O(m + n), где m — количество строк, а n — количество столбцов в матрице.
		Это связано с тем, что в худшем случае мы пройдем по всем строкам и столбцам матрицы.

		SPACE COMPLEXITY: O(1). Мы используем только константное количество дополнительной памяти для хранения переменных row, col, m и n.
	*/
	// Проверяем, что матрица не пуста
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	// Получаем размеры матрицы
	m, n := len(matrix), len(matrix[0])
	// Начинаем с правого верхнего угла матрицы
	row, col := 0, n-1

	// Пока индекс строки в пределах матрицы и индекс столбца неотрицателен
	for row < m && col >= 0 {
		// Если текущий элемент равен целевому значению, возвращаем true
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			// Если текущий элемент больше целевого значения, перемещаемся на один столбец влево
			col--
		} else {
			// Если текущий элемент меньше целевого значения, перемещаемся на одну строку вниз
			row++
		}
	}

	// Если целевое значение не найдено, возвращаем false
	return false
}