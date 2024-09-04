package main

import (
	"fmt"
)

// Функция для нахождения всех неубывающих подпоследовательностей
func findSubsequences(nums []int) [][]int {
	/*
		METHOD: Backtracking
		Используется метод бэктрекинга (обратного отслеживания) для генерации всех возможных подпоследовательностей.
		Рекурсивно перебираются элементы массива, и формируются подпоследовательности, которые затем проверяются на неубывание.

		TIME COMPLEXITY:
		Временная сложность составляет O(2^n * n), где n — длина входного массива. Это связано с тем, что для каждого элемента
		мы решаем, включать его в подпоследовательность или нет, что дает 2^n возможных комбинаций. Дополнительное умножение на n
		обусловлено необходимостью копирования подпоследовательностей при их добавлении в результат.

		SPACE COMPLEXITY:
		Пространственная сложность составляет O(n) для хранения текущей подпоследовательности и O(2^n) для хранения всех
		подпоследовательностей в результате. Таким образом, общая пространственная сложность составляет O(2^n * n).
	*/
	var result [][]int
	var path []int
	backtrack(nums, 0, path, &result)
	return result
}

// Рекурсивная функция для генерации подпоследовательностей
func backtrack(nums []int, start int, path []int, result *[][]int) {
	// Если длина текущей подпоследовательности >= 2, добавляем её в результат
	if len(path) >= 2 {
		*result = append(*result, append([]int(nil), path...))
	}

	// Множество для хранения уникальных элементов на текущем уровне рекурсии
	used := make(map[int]bool)

	// Перебираем элементы, начиная с индекса start
	for i := start; i < len(nums); i++ {
		// Пропускаем элементы, которые уже были использованы на текущем уровне рекурсии
		if used[nums[i]] {
			continue
		}

		// Если текущая подпоследовательность пуста или последний элемент меньше или равен текущему
		if len(path) == 0 || path[len(path)-1] <= nums[i] {
			// Добавляем текущий элемент в подпоследовательность
			path = append(path, nums[i])
			// Отмечаем элемент как использованный
			used[nums[i]] = true
			// Рекурсивно вызываем backtrack для следующего элемента
			backtrack(nums, i+1, path, result)
			// Удаляем последний элемент из подпоследовательности (бэктрекинг)
			path = path[:len(path)-1]
		}
	}
}

func main() {
	nums := []int{4, 6, 7, 7}
	result := findSubsequences(nums)
	for _, seq := range result {
		fmt.Println(seq)
	}
}
