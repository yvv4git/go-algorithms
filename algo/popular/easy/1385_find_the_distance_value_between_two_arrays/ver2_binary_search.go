package main

import (
	"sort"
)

// Функция findTheDistanceValue принимает два массива целых чисел arr1 и arr2,
// а также целое число d. Она возвращает количество элементов в arr1,
// которые не имеют пары в arr2, такой что абсолютная разница между элементами пары больше d.
func findTheDistanceValueV2(arr1 []int, arr2 []int, d int) int {
	/*
		METHOD: Sorting and Binary Search
		1. Сортировка arr2: Сначала сортируем arr2, чтобы использовать бинарный поиск.
		2. Бинарный поиск: Для каждого элемента x в arr1 находим ближайший элемент y в отсортированном arr2 с помощью бинарного поиска.
		3. Проверка условия: Проверяем, выполняется ли условие |x - y| > d для найденного элемента y.
		4. Подсчет подходящих элементов: Считаем количество элементов x в arr1, которые удовлетворяют условию.

		TIME COMPLEXITY: O(n log m), где n - количество элементов в arr1, а m - количество элементов в arr2.
		Это обусловлено сортировкой arr2 (O(m log m)) и последующим бинарным поиском для каждого элемента arr1 (O(n log m)).

		SPACE COMPLEXITY: O(1)
		Временная сложность алгоритма составляет O(1), так как не используется дополнительная память,
		кроме нескольких переменных (count, x, index).
	*/
	// Сортируем arr2
	sort.Ints(arr2)

	// Инициализируем счетчик подходящих элементов
	count := 0

	// Проходим по каждому элементу x в arr1
	for _, x := range arr1 {
		// Используем бинарный поиск для поиска ближайшего элемента в arr2
		index := sort.SearchInts(arr2, x)

		// Проверяем, выполняется ли условие |x - y| > d для найденного элемента и его соседей
		valid := true
		if index < len(arr2) {
			if abs(x-arr2[index]) <= d {
				valid = false
			}
		}
		if index > 0 {
			if abs(x-arr2[index-1]) <= d {
				valid = false
			}
		}

		// Если элемент x подходит, увеличиваем счетчик
		if valid {
			count++
		}
	}

	// Возвращаем количество подходящих элементов
	return count
}

// Вспомогательная функция для вычисления абсолютного значения
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// func main() {
// 	// Пример использования функции
// 	arr1 := []int{4, 5, 8}
// 	arr2 := []int{10, 9, 1, 8}
// 	d := 2

// 	result := findTheDistanceValue(arr1, arr2, d)
// 	fmt.Println(result) // Вывод: 2
// }
