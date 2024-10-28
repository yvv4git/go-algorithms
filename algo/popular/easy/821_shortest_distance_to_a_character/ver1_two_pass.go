package main

import (
	"fmt"
	"math"
)

func shortestToChar(s string, c byte) []int {
	/*
		METHOD: Двухпроходный алгоритм (Two-pass algorithm)
		- Первый проход: слева направо
		- Второй проход: справа налево
		- На каждом проходе вычисляем расстояние до ближайшего символа `c`
		- Выбираем минимальное расстояние из двух проходов

		TIME COMPLEXITY: O(n), где n — длина строки s
		- Мы проходим по строке два раза, что дает нам линейную временную сложность O(n)

		SPACE COMPLEXITY: O(n), где n — длина строки s
		- Мы используем дополнительный массив `result` длины `n` для хранения результатов
	*/
	n := len(s)
	result := make([]int, n)

	// Инициализируем массив result бесконечностями
	for i := range result {
		result[i] = math.MaxInt32
	}

	// Прямой проход: слева направо
	pos := -math.MaxInt32 // Изначально позиция ближайшего символа c равна -бесконечности
	for i := 0; i < n; i++ {
		if s[i] == c {
			pos = i // Обновляем позицию ближайшего символа c
		}
		result[i] = i - pos // Вычисляем расстояние до ближайшего символа c слева
	}

	// Обратный проход: справа налево
	pos = math.MaxInt32 // Изначально позиция ближайшего символа c равна +бесконечности
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			pos = i // Обновляем позицию ближайшего символа c
		}
		result[i] = min(result[i], pos-i) // Выбираем минимальное расстояние из двух проходов
	}

	return result
}

// Функция для нахождения минимума двух чисел
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := "loveleetcode"
	c := byte('e')
	result := shortestToChar(s, c)
	fmt.Println(result) // Вывод: [3 2 1 0 1 0 0 1 2 2 1 0]
}
