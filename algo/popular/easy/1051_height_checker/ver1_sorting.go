//go:build ignore

package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	/*
		METHOD: Sorting
		TIME COMPLEXITY: O(n log n)
		SPACE COMPLEXITY: O(n)

		Подробное объяснение:
		- МЕТОД: Сортировка
		  Мы используем метод сортировки для решения задачи. Сначала создаем копию исходного массива и сортируем её.
		  Затем сравниваем отсортированный массив с исходным, чтобы найти количество элементов, стоящих не на своих местах.

		- ВРЕМЕННАЯ СЛОЖНОСТЬ: O(n log n)
		  Сортировка массива занимает O(n log n) времени, где n — количество элементов в массиве.
		  После сортировки мы проходим по массивам один раз, что занимает O(n) времени.
		  Таким образом, общая временная сложность составляет O(n log n).

		- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ: O(n)
		  Мы создаем копию исходного массива, которая занимает O(n) дополнительного пространства.
		  Таким образом, общая пространственная сложность составляет O(n).
	*/
	// Создаем копию массива heights
	sortedHeights := make([]int, len(heights))
	copy(sortedHeights, heights)

	// Сортируем копию массива
	sort.Ints(sortedHeights)

	// Инициализируем счетчик несовпадающих элементов
	count := 0

	// Проходим по обоим массивам и сравниваем элементы
	for i := 0; i < len(heights); i++ {
		if heights[i] != sortedHeights[i] {
			count++
		}
	}

	// Возвращаем количество несовпадающих элементов
	return count
}

func main() {
	// Пример использования функции
	heights := []int{1, 1, 4, 2, 1, 3}
	result := heightChecker(heights)
	fmt.Println(result) // Вывод: 3
}
