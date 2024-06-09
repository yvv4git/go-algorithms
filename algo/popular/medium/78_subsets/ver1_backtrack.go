package main

import "fmt"

// Функция для генерации всех подмножеств
func subsets(nums []int) [][]int {
	/*
		METHOD: Backtrack

		TIME COMPLEXITY: O(N * 2^N), где N - количество элементов в исходном наборе.
		Это связано с тем, что для каждого элемента мы можем взять его или не взять, что дает нам 2^N возможных подмножеств.

		SPACE COMPLEXITY: O(N * 2^N), так как мы храним все N * 2^N возможных подмножеств.
	*/
	var result [][]int                   // Срез для хранения результатов
	backtrack(nums, 0, []int{}, &result) // Вызов вспомогательной функции
	return result                        // Возврат результатов
}

// Вспомогательная функция для генерации подмножеств через backtracking
func backtrack(nums []int, start int, path []int, result *[][]int) {
	temp := make([]int, len(path))  // Создание временного среза для хранения текущего пути
	copy(temp, path)                // Копирование текущего пути в временный срез
	*result = append(*result, temp) // Добавление текущего пути в результаты

	// Перебор оставшихся элементов
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])       // Добавление текущего элемента в путь
		backtrack(nums, i+1, path, result) // Рекурсивный вызов для следующего индекса
		path = path[:len(path)-1]          // Удаление последнего элемента из пути для обхода других вариантов
	}
}

func main() {
	nums := []int{1, 2, 3}  // Исходный набор чисел
	result := subsets(nums) // Получение всех подмножеств
	fmt.Println(result)     // Вывод результатов
}
