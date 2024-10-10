package main

import (
	"sort"
)

type indexedValue struct {
	value int
	index int
}

func maxWidthRamp(nums []int) int {
	/*
		METHOD: Сортировка и поиск максимальной ширины рампы.
		1. Создаем массив пар (значение, индекс)
		2. Сортируем массив по значению, а при равенстве значений - по индексу
		3. Инициализируем минимальный индекс и максимальную ширину
		4. Проходим по отсортированному массиву
		5. Вычисляем текущую ширину рампы
		6. Обновляем максимальную ширину
		7. Обновляем минимальный индекс
		8. Возвращаем максимальную ширину рампы

		TIME COMPLEXITY: O(n log n) - где n — количество элементов в массиве.
		Сортировка занимает O(n log n), а проход по отсортированному массиву — O(n).

		SPACE COMPLEXITY: O(n) - для хранения массива пар (значение, индекс).
	*/
	// Создаем массив пар (значение, индекс)
	indexedNums := make([]indexedValue, len(nums))
	for i, v := range nums {
		indexedNums[i] = indexedValue{value: v, index: i}
	}

	// Сортируем массив по значению, а при равенстве значений - по индексу
	sort.Slice(indexedNums, func(i, j int) bool {
		if indexedNums[i].value == indexedNums[j].value {
			return indexedNums[i].index < indexedNums[j].index
		}
		return indexedNums[i].value < indexedNums[j].value
	})

	// Инициализируем минимальный индекс и максимальную ширину
	minIndex := len(nums)
	maxWidth := 0

	// Проходим по отсортированному массиву
	for _, iv := range indexedNums {
		// Вычисляем текущую ширину рампы
		currentWidth := iv.index - minIndex
		// Обновляем максимальную ширину
		if currentWidth > maxWidth {
			maxWidth = currentWidth
		}
		// Обновляем минимальный индекс
		if iv.index < minIndex {
			minIndex = iv.index
		}
	}

	return maxWidth
}

// func main() {
// 	// Пример использования:
// 	nums := []int{6, 0, 8, 2, 1, 5}
// 	fmt.Println(maxWidthRamp(nums)) // Вывод: 4
// }
