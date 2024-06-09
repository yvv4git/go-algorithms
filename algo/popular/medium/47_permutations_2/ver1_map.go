package main

import "fmt"

func permuteUnique(nums []int) [][]int {
	// Инициализируем счетчик для каждого числа в массиве
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	// Инициализируем результат и вызываем вспомогательную функцию
	var result [][]int
	dfs(nums, counter, []int{}, &result)
	return result
}

// Функция dfs рекурсивно строит все уникальные перестановки чисел.
// nums - исходный массив чисел,
// counter - словарь, где ключ - число из массива, значение - количество таких чисел в массиве,
// path - текущий путь перестановки,
// result - указатель на результирующий срез с перестановками.
func dfs(nums []int, counter map[int]int, path []int, result *[][]int) {
	// Если путь содержит все элементы, добавляем его в результат
	if len(path) == len(nums) {
		*result = append(*result, append([]int{}, path...))
		return
	}

	// Проходим по каждому уникальному числу
	for num, count := range counter {
		if count > 0 {
			// Уменьшаем счетчик для текущего числа
			counter[num]--
			// Добавляем число в путь и вызываем dfs рекурсивно
			dfs(nums, counter, append(path, num), result)
			// Возвращаем счетчик в исходное состояние
			counter[num]++
		}
	}
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums)) // Вывод: [[1,1,2] [1,2,1] [2,1,1]]
}
