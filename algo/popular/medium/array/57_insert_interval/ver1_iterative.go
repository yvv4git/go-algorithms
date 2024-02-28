package main

func insert(intervals [][]int, newInterval []int) [][]int {
	/*
		METHOD: Iterative
		TIME COMPLEXITY: O(n), где n - количество интервалов в массиве intervals. Это связано с тем, что функция проходит по каждому интервалу только один раз.
		SPACE COMPLEXITY: O(n), так как в худшем случае может потребоваться хранить все исходные интервалы в результирующем массиве result.
		Однако, в целом, так как мы используем дополнительную память только для результирующего массива,
		а не для дополнительных переменных, пространственная сложность составляет O(n), а не O(n + m),
		где m - количество дополнительных переменных.
	*/
	var result [][]int
	i := 0

	// Добавляем все интервалы, которые заканчиваются до начала нового интервала
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// Объединяем все интервалы, которые пересекаются с новым интервалом
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	// Добавляем объединенный интервал
	result = append(result, newInterval)

	// Добавляем оставшиеся интервалы
	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
