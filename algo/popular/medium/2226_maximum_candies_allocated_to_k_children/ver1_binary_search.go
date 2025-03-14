package main

import (
	"fmt"
)

// Функция для нахождения максимального количества конфет,
// которое можно выдать каждому ребенку
func maxCandies(candies []int, k int64) int {
	/*
		APPROACH: BINARY SEARCH
		1. Найти максимальное количество конфет в одном мешке
		2. Найти максимальный размер порции, который можно выдать k детям
		3. Найти минимальный размер порции, который дает k детям максимальное количество конфет
		4. Вернуть размер порции в качестве результата

		TIME COMPLEXITY: O(nlogn) - для нахождения максимального количества конфет в одном мешке
		SPACE COMPLEXITY: O(1) - не используется дополнительной памяти
	*/
	if k == 0 {
		return 0
	}

	// Поиск максимального количества конфет в одном мешке
	maxCandy := 0
	for _, c := range candies {
		if c > maxCandy {
			maxCandy = c
		}
	}

	// Границы бинарного поиска
	left, right := 1, maxCandy
	best := 0

	// Бинарный поиск по ответу
	for left <= right {
		mid := (left + right) / 2

		// Подсчёт количества детей, которых можно накормить порцией mid
		var count int64 = 0
		for _, c := range candies {
			count += int64(c) / int64(mid) // Приводим к int64
		}

		// Проверяем, хватает ли порций для k детей
		if count >= k {
			best = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return best
}

func main() {
	// Пример использования
	candies := []int{5, 8, 6}
	var k int64 = 3                     // k объявляем как int64
	fmt.Println(maxCandies(candies, k)) // Ожидаемый результат: 5
}
