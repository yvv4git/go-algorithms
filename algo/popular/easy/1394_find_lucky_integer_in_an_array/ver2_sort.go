//go:build ignore

package main

import (
	"fmt"
	"sort"
)

func findLucky(arr []int) int {
	/*
		METHOD: Sort + Linear sort
		Сортировка массива + линейный проход для подсчета частот и поиска счастливого числа.

		TIME COMPLEXITY: O(n log n) — из-за сортировки.
		SPACE COMPLEXITY: O(1) — если не учитывать память на сортировку.
	*/
	// Шаг 1: Сортировка массива
	sort.Ints(arr)

	// Шаг 2: Подсчет частот и поиск счастливого числа
	maxLucky := -1
	currentNum := arr[0]
	count := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] == currentNum {
			count++
		} else {
			// Проверяем, является ли текущее число счастливым
			if currentNum == count && currentNum > maxLucky {
				maxLucky = currentNum
			}
			// Переходим к следующему числу
			currentNum = arr[i]
			count = 1
		}
	}

	// Проверяем последнее число
	if currentNum == count && currentNum > maxLucky {
		maxLucky = currentNum
	}

	return maxLucky
}

func main() {
	// Примеры входных данных
	arr1 := []int{2, 2, 3, 4}
	arr2 := []int{1, 2, 2, 3, 3, 3}
	arr3 := []int{2, 2, 2, 3, 3}

	// Вызов функции findLucky и вывод результатов
	fmt.Println("Пример 1:")
	fmt.Println("Входные данные:", arr1)
	fmt.Println("Счастливое число:", findLucky(arr1)) // Ожидаемый вывод: 2

	fmt.Println("\nПример 2:")
	fmt.Println("Входные данные:", arr2)
	fmt.Println("Счастливое число:", findLucky(arr2)) // Ожидаемый вывод: 3

	fmt.Println("\nПример 3:")
	fmt.Println("Входные данные:", arr3)
	fmt.Println("Счастливое число:", findLucky(arr3)) // Ожидаемый вывод: -1
}
