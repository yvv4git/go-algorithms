//go:build ignore

package main

import (
	"fmt"
)

func heightChecker(heights []int) int {
	/*
		METHOD: Counting Sort
		TIME COMPLEXITY: O(n + k)
		SPACE COMPLEXITY: O(k)

		Подробное объяснение:
		- МЕТОД: Counting Sort
		  Мы используем Counting Sort для сортировки массива, так как диапазон значений роста студентов известен и ограничен.
		  Это позволяет нам сортировать массив за линейное время.

		- ВРЕМЕННАЯ СЛОЖНОСТЬ: O(n + k)
		  Counting Sort занимает O(n + k) времени, где n — количество элементов в массиве, а k — диапазон значений элементов.
		  После сортировки мы проходим по массивам один раз, что занимает O(n) времени.
		  Таким образом, общая временная сложность составляет O(n + k).

		- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ: O(k)
		  Мы используем дополнительный массив размером k для подсчета частоты встречаемости каждого элемента.
		  Таким образом, общая пространственная сложность составляет O(k).
	*/

	// Предположим, что диапазон значений роста студентов от 1 до 100
	k := 101
	count := make([]int, k)

	// Подсчитываем частоту встречаемости каждого элемента
	for _, height := range heights {
		count[height]++
	}

	// Сортируем массив с использованием подсчитанных частот
	sortedHeights := make([]int, 0, len(heights))
	for i := 1; i < k; i++ {
		for j := 0; j < count[i]; j++ {
			sortedHeights = append(sortedHeights, i)
		}
	}

	// Инициализируем счетчик несовпадающих элементов
	countMismatches := 0

	// Проходим по обоим массивам и сравниваем элементы
	for i := 0; i < len(heights); i++ {
		if heights[i] != sortedHeights[i] {
			countMismatches++
		}
	}

	// Возвращаем количество несовпадающих элементов
	return countMismatches
}

func main() {
	// Пример использования функции
	heights := []int{1, 1, 4, 2, 1, 3}
	result := heightChecker(heights)
	fmt.Println(result) // Вывод: 3
}
