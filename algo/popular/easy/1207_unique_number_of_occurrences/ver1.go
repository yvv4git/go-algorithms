package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	/*
		METHOD:
		- Мы используем два map: Frequency Counting
			1. Первый map (countMap) для подсчета количества вхождений каждого числа в массиве.
			2. Второй map (occurrenceSet) для проверки уникальности количеств вхождений.
		- Сначала мы проходим по массиву и заполняем countMap.
		- Затем мы проверяем, все ли количества вхождений уникальны, используя occurrenceSet.
		- Если какое-то количество вхождений уже встречалось, возвращаем false.
		- Если все количества уникальны, возвращаем true.

		TIME COMPLEXITY: O(n)
		- Мы проходим по массиву один раз для подсчета вхождений (O(n)).
		- Затем мы проходим по countMap, чтобы проверить уникальность (O(m), где m — количество уникальных чисел, но m <= n).
		- Таким образом, общая сложность составляет O(n).

		SPACE COMPLEXITY: O(n)
		- Мы используем два map:
			1. countMap может содержать до n элементов (если все числа уникальны).
			2. occurrenceSet может содержать до n элементов (если все количества вхождений уникальны).
		- Таким образом, общая пространственная сложность составляет O(n).
	*/
	// Создаем map для подсчета количества вхождений каждого числа
	countMap := make(map[int]int)

	// Подсчитываем количество вхождений
	for _, num := range arr {
		countMap[num]++
	}

	// Создаем set (множество) для хранения уникальных количеств вхождений
	occurrenceSet := make(map[int]bool)

	// Проверяем, все ли количества вхождений уникальны
	for _, count := range countMap {
		if occurrenceSet[count] {
			return false // Если количество уже встречалось, возвращаем false
		}
		occurrenceSet[count] = true // Добавляем количество в множество
	}

	return true // Все количества уникальны
}

func main() {
	arr1 := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(arr1)) // Вывод: true

	arr2 := []int{1, 2}
	fmt.Println(uniqueOccurrences(arr2)) // Вывод: false

	arr3 := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	fmt.Println(uniqueOccurrences(arr3)) // Вывод: true
}
