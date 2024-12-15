package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	/*
		TASK: Задача требует преобразовать массив чисел в массив их рангов.
		Ранг числа — это его позиция в отсортированном массиве, где числа с одинаковыми значениями имеют одинаковый ранг.

		METHOD: Sort
		1. Создаем копию массива и сортируем его.
		2. Создаем словарь для хранения рангов чисел.
		3. Проходим по отсортированному массиву и заполняем словарь, где ключом является число, а значением — его ранг.
		4. Проходим по исходному массиву и заменяем каждое число на его ранг из словаря.

		TIME COMPLEXITY: O(n log n)
		- Сортировка массива: O(n log n)
		- Проход по отсортированному массиву для создания словаря: O(n)
		- Проход по исходному массиву для замены чисел на ранги: O(n)
		Итоговая временная сложность: O(n log n)

		SPACE COMPLEXITY: O(n)
		- Хранение отсортированного массива: O(n)
		- Хранение словаря с рангами: O(n)
		- Хранение результирующего массива: O(n)
		Итоговая пространственная сложность: O(n)
	*/
	// Создаем копию массива и сортируем его
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	// Создаем словарь для хранения рангов
	rankMap := make(map[int]int)
	rank := 1

	// Проходим по отсортированному массиву и заполняем словарь
	for _, num := range sortedArr {
		// Если число еще не встречалось, добавляем его в словарь с текущим рангом
		if _, exists := rankMap[num]; !exists {
			rankMap[num] = rank
			rank++
		}
	}

	// Создаем результирующий массив с рангами
	result := make([]int, len(arr))
	for i, num := range arr {
		result[i] = rankMap[num]
	}

	return result
}

func main() {
	// Пример использования
	arr1 := []int{40, 10, 20, 30}
	fmt.Println(arrayRankTransform(arr1)) // Вывод: [4 1 2 3]

	arr2 := []int{100, 100, 100}
	fmt.Println(arrayRankTransform(arr2)) // Вывод: [1 1 1]

	arr3 := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	fmt.Println(arrayRankTransform(arr3)) // Вывод: [5 3 4 2 8 6 7 1 3]
}