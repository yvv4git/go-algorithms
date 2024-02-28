package main

// Функция insertV3 принимает на вход массив интервалов и новый интервал, который нужно вставить.
// Она возвращает новый массив интервалов после вставки нового интервала.
func insertV3(intervals [][]int, newInterval []int) [][]int {
	/*
		METHOD: Binary search

		TIME COMPLEXITY: O(log n) для бинарного поиска, где n - количество интервалов в массиве intervals.
		Это связано с использованием бинарного поиска для поиска позиций вставки нового интервала,
		который выполняется за логарифмическое время.

		SPACE COMPLEXITY: O(n), так как в худшем случае может потребоваться хранить все исходные интервалы в результирующем массиве intervals.
		Однако, в целом, так как мы используем дополнительную память только для результирующего массива,
		пространственная сложность составляет O(n), а не O(n + m), где m - количество дополнительных переменных.
	*/
	// Используем бинарный поиск для поиска позиций, куда можно вставить новый интервал,
	// чтобы сохранить отсортированность по начальному значению интервала
	startIndex := binarySearch(intervals, newInterval[0])

	// Используем бинарный поиск для поиска позиций, куда можно вставить новый интервал,
	// чтобы сохранить отсортированность по конечному значению интервала
	endIndex := binarySearch(intervals, newInterval[1])

	// Создаем новый срез, который будет содержать результат
	result := append(make([][]int, 0, startIndex+len(intervals)-endIndex+1), intervals[:startIndex]...)

	// Если новый интервал пересекается с существующим интервалом, обновляем начало нового интервала
	if startIndex != len(intervals) && intervalInclude(intervals[startIndex], newInterval[0]) {
		newInterval[0] = intervals[startIndex][0]
	}

	// Если новый интервал пересекается с существующим интервалом, обновляем конец нового интервала
	if endIndex != len(intervals) && intervalInclude(intervals[endIndex], newInterval[1]) {
		newInterval[1] = intervals[endIndex][1]
	}

	// Вставляем новый интервал в результирующий срез
	result = append(result, newInterval)

	// Если новый интервал пересекается с существующим интервалом, увеличиваем индекс конца
	if endIndex != len(intervals) && intervalInclude(intervals[endIndex], newInterval[1]) {
		endIndex++
	}

	// Добавляем оставшиеся интервалы в результирующий срез и возвращаем его
	return append(result, intervals[endIndex:]...)
}

// Функция binarySearch выполняет бинарный поиск позиции для вставки нового интервала
func binarySearch(intervals [][]int, val int) int {
	var (
		left  = 0
		right = len(intervals) - 1
		mid   = right / 2
	)

	for left <= right {
		if val < intervals[mid][0] {
			right = mid - 1
		} else if val > intervals[mid][1] {
			left = mid + 1
		} else {
			return mid
		}

		mid = (right-left)/2 + left
	}

	return left
}

// Функция intervalInclude проверяет, входит ли значение в интервал
func intervalInclude(interval []int, val int) bool {
	return interval[0] <= val && interval[1] >= val
}
