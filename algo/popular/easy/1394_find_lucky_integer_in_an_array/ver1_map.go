package main

import (
	"fmt"
)

// Функция для поиска "счастливого числа" в массиве
func findLucky(arr []int) int {
	/*
		METHOD: Hash table / Map
		Используется хэш-таблица (map) для подсчета частот чисел в массиве, затем выполняется поиск числа, которое равно своей частоте.

		TIME COMPLEXITY: O(n), где n — длина массива. Проход по массиву для подсчета частот и проход по хэш-таблице для поиска счастливого числа.
		SPACE COMPLEXITY: O(n), так как используется хэш-таблица для хранения частот чисел.
	*/
	// Шаг 1: Подсчет частот
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}

	// Шаг 2: Поиск счастливого числа
	maxLucky := -1
	for num, count := range freq {
		if num == count {
			if num > maxLucky {
				maxLucky = num
			}
		}
	}

	// Шаг 3: Возврат результата
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
