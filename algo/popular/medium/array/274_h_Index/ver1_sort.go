package main

import (
	"fmt"
	"sort"
)

// Функция для вычисления H-индекса
func hIndex(citations []int) int {
	/*
		METHOD: Sort

		TIME COMPLEXITY: O(n log n) из-за сортировки массива

		SPACE COMPLEXITY: O(1)
	*/
	// Сортируем массив цитирований в порядке убывания
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	// Проходимся по отсортированному массиву
	for i, citation := range citations {
		// Если цитирование меньше или равно индексу плюс единица,
		// то возвращаем этот индекс как H-индекс
		if citation < i+1 {
			return i
		}
	}

	// Если мы прошли весь массив и не нашли подходящего индекса,
	// то возвращаем длину массива, так как H-индекс не может быть больше количества цитирований
	return len(citations)
}

func main() {
	citations := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(citations)) // Выведет: 3
}
