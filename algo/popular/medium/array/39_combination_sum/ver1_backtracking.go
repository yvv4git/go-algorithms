package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	/*
		METHOD: Backtracking
		TIME COMPLEXITY: O(N^(T/M + 1)), где N - количество кандидатов, T - целевая сумма, M - минимальный кандидат.
		Это связано с тем, что для каждого кандидата мы можем выбирать его несколько раз, пока сумма не превысит целевую.
		SPACE COMPLEXITY: O(T/M), так как в худшем случае глубина рекурсии может быть равна T/M, и для каждого уровня рекурсии мы храним текущий список чисел.
	*/
	// Сортируем кандидатов для удобства
	sort.Ints(candidates)

	// Инициализируем переменные для результатов
	var result [][]int

	// Вызываем вспомогательную функцию для поиска комбинаций
	dfs(candidates, target, 0, []int{}, &result)

	return result
}

func dfs(candidates []int, target int, index int, path []int, result *[][]int) {
	// Если текущая сумма равна целевой, добавляем текущий путь в результаты
	if target == 0 {
		newPath := make([]int, len(path))
		copy(newPath, path)
		*result = append(*result, newPath)
		return
	}

	// Если текущая сумма больше целевой, прекращаем поиск
	if target < 0 {
		return
	}

	// Проходим по каждому кандидату, начиная с текущего индекса
	for i := index; i < len(candidates); i++ {
		// Добавляем текущий кандидат в путь
		path = append(path, candidates[i])

		// Рекурсивно ищем комбинации с новой суммой и новым путём
		dfs(candidates, target-candidates[i], i, path, result)

		// Удаляем последний добавленный кандидат из пути для проверки следующих комбинаций
		path = path[:len(path)-1]
	}
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	result := combinationSum(candidates, target)
	fmt.Println(result) // Должен вывести [[2, 2, 3], [7]]
}
