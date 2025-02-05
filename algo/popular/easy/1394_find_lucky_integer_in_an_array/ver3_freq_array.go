//go:build ignore

package main

import (
	"fmt"
)

func findLucky(arr []int) int {
	/*
		METHOD: Frequencies array
		Использование массива частот для подсчета и поиска счастливого числа.

		TIME COMPLEXITY: O(n) — проход по массиву для подсчета частот и поиска счастливого числа.
		SPACE COMPLEXITY: O(1) — используется массив фиксированного размера (501 элемент).
	*/
	// Шаг 1: Создаем массив для подсчета частот
	freq := make([]int, 501) // По условию задачи числа в массиве <= 500

	// Шаг 2: Подсчет частот
	for _, num := range arr {
		freq[num]++
	}

	// Шаг 3: Поиск счастливого числа
	maxLucky := -1
	for num := 1; num <= 500; num++ {
		if freq[num] == num && num > maxLucky {
			maxLucky = num
		}
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
