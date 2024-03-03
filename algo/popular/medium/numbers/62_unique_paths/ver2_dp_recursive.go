package main

import "fmt"

var cache map[string]int

func uniquePathsV2(m int, n int) int {
	/*
		METHOD: Dynamic programming with recursion
		TIME COMPLEXITY: O(m*n), где m и n - размеры сетки, так как мы проходим по всей сетке ровно один раз.
		SPACE COMPLEXITY: O(m*n), так как мы создаем дополнительный массив dp размером m*n.
	*/
	// Инициализация кэша, если он еще не был создан
	if cache == nil {
		cache = make(map[string]int)
	}

	// Если размеры сетки некорректны (меньше 1), возвращаем 0
	if m < 1 || n < 1 {
		return 0
	}

	// Если сетка состоит из одной клетки, возвращаем 1
	if m == 1 && n == 1 {
		return 1
	}

	// Формируем ключ для кэша на основе размеров сетки
	key := fmt.Sprintf("%d-%d", m, n)

	// Проверяем, есть ли уже вычисленный результат для этого ключа в кэше
	if v, ok := cache[key]; ok {
		return v
	}

	// Рекурсивно вычисляем количество уникальных путей для сетки без клетки справа
	a := 0
	if m-1 > 0 {
		a = uniquePathsV2(m-1, n)
	}

	// Рекурсивно вычисляем количество уникальных путей для сетки без клетки снизу
	b := 0
	if n-1 > 0 {
		b = uniquePathsV2(m, n-1)
	}

	// Складываем количество путей и сохраняем результат в кэше
	cache[key] = a + b

	// Возвращаем результат из кэша
	return cache[key]
}
