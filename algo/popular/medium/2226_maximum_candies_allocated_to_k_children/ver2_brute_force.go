//go:build ignore

package main

import (
	"fmt"
)

// Функция полного перебора (Brute Force)
func bruteForceMaxCandies(candies []int, k int64) int {
	/*
		APPROACH: BRUTE FORCE
		1. Найти максимальное количество конфет в одном мешке
		2. Перебрать все возможные размеры порции от 1 до max(candies)
		3. Для каждого размера порции подсчитать, сколько детей можно накормить
		4. Найти наибольший размер порции, при котором можно накормить не менее k детей
		5. Вернуть найденный размер порции

		TIME COMPLEXITY: O(n * max_candies) - проверяем все возможные размеры порций
		SPACE COMPLEXITY: O(1) - не используем дополнительной памяти
	*/
	if k == 0 {
		return 0
	}

	// Находим максимальное количество конфет в одном мешке
	maxCandy := 0
	for _, c := range candies {
		if c > maxCandy {
			maxCandy = c
		}
	}

	// Перебираем возможные размеры порции от 1 до maxCandy
	best := 0
	for x := 1; x <= maxCandy; x++ {
		var count int64 = 0
		for _, c := range candies {
			count += int64(c) / int64(x) // Подсчитываем количество возможных порций
		}
		if count >= k {
			best = x // Обновляем лучший найденный результат
		}
	}

	return best
}

func main() {
	// Тестовый пример
	candies := []int{5, 8, 6}
	var k int64 = 3
	fmt.Println(bruteForceMaxCandies(candies, k)) // Ожидаемый результат: 5
}
