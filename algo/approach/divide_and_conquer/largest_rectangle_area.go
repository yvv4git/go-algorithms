//go:build ignore

package main

import "fmt"

// Функция для нахождения наибольшей прямоугольной подматрицы, состоящей только из 1
func maximalRectangle(matrix [][]byte) int {
	/*
		Задача: Найти наибольшую прямоугольную подматрицу, состоящую только из 1.

		Подход:
		1. Используем динамическое программирование для вычисления высоты столбцов.
		2. Для каждой строки находим максимальную прямоугольную подматрицу.

		ТАЙМ-КОМПЛЕКСНОСТЬ: O(m * n), где m — количество строк, n — количество столбцов.
		СПЕЙС-КОМПЛЕКСНОСТЬ: O(m * n)
	*/

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	// Создаем массив для хранения высот столбцов
	height := make([]int, n)
	maxArea := 0

	// Проходим по каждой строке
	for i := 0; i < m; i++ {
		// Обновляем высоты столбцов для текущей строки
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}

		// Находим максимальную площадь прямоугольника для текущих высот
		maxArea = max(maxArea, largestRectangleArea(height))
	}

	return maxArea
}

// Функция для нахождения максимальной площади прямоугольника в гистограмме
func largestRectangleArea(heights []int) int {
	/*
		Задача: Найти максимальную площадь прямоугольника в гистограмме.

		Подход:
		1. Используем стек для поиска наибольшего прямоугольника, который можно создать.
		2. Для каждого столбца вычисляем максимальную возможную ширину прямоугольника.

		ТАЙМ-КОМПЛЕКСНОСТЬ: O(n), где n — количество столбцов.
		СПЕЙС-КОМПЛЕКСНОСТЬ: O(n)
	*/
	stack := []int{}
	maxArea := 0

	// Добавляем в конец высоту 0 для обработки оставшихся элементов
	heights = append(heights, 0)

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width -= stack[len(stack)-1] + 1
			}
			maxArea = max(maxArea, h*width)
		}
		stack = append(stack, i)
	}

	return maxArea
}

// Функция для нахождения максимального числа
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Пример: матрица с 0 и 1
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	// Вызываем функцию для нахождения наибольшей прямоугольной подматрицы
	result := maximalRectangle(matrix)

	// Выводим результат
	fmt.Printf("Наибольшая площадь прямоугольной подматрицы: %d\n", result)
}
