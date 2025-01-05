package main

import "fmt"

// Функция kidsWithCandies определяет, у каких детей будет наибольшее количество конфет
// после выдачи дополнительных конфет.
func kidsWithCandies(candies []int, extraCandies int) []bool {
	/*
		METHOD: Linear Scan
		1. Находим максимальное количество конфет в массиве candies.
		2. Для каждого ребенка проверяем, станет ли его количество конфет (candy + extraCandies)
		   больше или равно максимальному количеству.
		3. Результат сохраняем в массив булевых значений.

		TIME COMPLEXITY: O(n)
		- Поиск максимума: O(n)
		- Проверка условия для каждого ребенка: O(n)
		Итого: O(n) + O(n) = O(n)

		SPACE COMPLEXITY: O(n)
		- Массив result занимает O(n) памяти.
	*/
	// Находим максимальное количество конфет в исходном массиве
	maxCandies := 0
	for _, candy := range candies {
		if candy > maxCandies {
			maxCandies = candy
		}
	}

	// Создаем массив для хранения результата
	result := make([]bool, len(candies))

	// Проверяем для каждого ребенка, станет ли его количество конфет максимальным
	// после добавления extraCandies
	for i, candy := range candies {
		result[i] = (candy + extraCandies) >= maxCandies
	}

	return result
}

func main() {
	// Пример входных данных
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	// Вызов функции и вывод результата
	result := kidsWithCandies(candies, extraCandies)
	fmt.Println(result) // Ожидаемый вывод: [false false true false false]
}
