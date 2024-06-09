package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	/*
		METHOD: Recursion

		TIME COMPLEXITY:  O(N * 2^N), где N - количество элементов в исходном наборе.
		Это связано с тем, что для каждого элемента мы можем выбирать его или не выбирать, что дает нам 2^N возможных подмножеств.

		SPACE COMPLEXITY: O(N * 2^N), так как мы должны хранить все возможные подмножества.
	*/
	sort.Ints(nums) // Сортируем исходный массив для обработки дубликатов
	return dfs(nums, []int{}, 0, make(map[string]bool))
}

func dfs(nums []int, path []int, index int, visited map[string]bool) [][]int {
	// Если мы достигли конца массива, то добавляем текущее подмножество в результат
	if index == len(nums) {
		key := fmt.Sprint(path) // Создаем ключ для хэш-таблицы
		if _, ok := visited[key]; !ok {
			visited[key] = true // Помечаем подмножество как посещенное
			return [][]int{path}
		}
		return [][]int{}
	}

	// Рекурсивно строим подмножества, выбирая текущий элемент или не выбирая
	withCurrent := dfs(nums, append([]int{}, path...), index+1, visited)
	withoutCurrent := dfs(nums, append(path, nums[index]), index+1, visited)

	// Объединяем результаты и возвращаем
	return append(withCurrent, withoutCurrent...)
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums)) // Выводит: [[],[1],[1,2],[1,2,2],[2],[2,2]]
}
