package main

import (
	"fmt"
	"sort"
)

// Функция для сортировки интервалов по начальной точке
func sortIntervals(intervals [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
}

// Функция для объединения интервалов
func merge(intervals [][]int) [][]int {
	/*
		METHOD: Sorting and Merging
		TIME COMPLEXITY: O(n log n), где n - количество интервалов, так как мы сначала сортируем массив, затем проходимся по нему.
		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем получить новый массив, который будет содержать все интервалы, если они не пересекаются.
	*/
	if len(intervals) == 0 {
		return [][]int{}
	}

	// Сортируем интервалы
	sortIntervals(intervals)

	// Инициализируем результирующий массив
	merged := [][]int{intervals[0]}

	// Проходимся по интервалам
	for _, interval := range intervals {
		// Получаем последний интервал в результирующем массиве
		lastMerged := merged[len(merged)-1]

		// Если текущий интервал пересекается с последним интервалом в результирующем массиве
		if interval[0] <= lastMerged[1] {
			// Обновляем конец последнего интервала в результирующем массиве
			lastMerged[1] = max(lastMerged[1], interval[1])
		} else {
			// Если интервалы не пересекаются, добавляем его в результирующий массив
			merged = append(merged, interval)
		}
	}

	return merged
}

// Функция для нахождения максимума из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals)) // Должно вывести [[1 6] [8 10] [15 18]]
}
