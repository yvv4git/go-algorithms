package main

import (
	"fmt"
	"sort"
)

// Функция для поиска всех уникальных комбинаций чисел из заданного массива, которые дают заданную сумму.
func combinationSum2(candidates []int, target int) [][]int {
	/*
		METHOD: Backtracking
		TIME COMPLEXITY: O(2^n), где n - количество чисел в массиве candidates.
		Это связано с тем, что в худшем случае нам может потребоваться проверить все возможные подмножества чисел в массиве.
		SPACE COMPLEXITY: O(2^n), так как в худшем случае мы можем хранить в результате всевозможные комбинации чисел, которые дают заданную сумму.
	*/
	// Сортируем массив для упорядочивания чисел и упрощения процесса отсеивания повторяющихся комбинаций.
	sort.Ints(candidates)
	result := make([][]int, 0)
	backtrack(candidates, target, 0, []int{}, &result)
	return result
}

// Рекурсивная функция для поиска комбинаций.
func backtrack(candidates []int, target int, start int, current []int, result *[][]int) {
	// Если сумма текущей комбинации равна заданной сумме, то добавляем комбинацию в результат.
	if target == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	// Перебираем каждое число из массива.
	for i := start; i < len(candidates); i++ {
		// Пропускаем повторяющиеся числа, чтобы избежать дубликатов комбинаций.
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		// Если текущее число больше целевой суммы, то прерываем цикл.
		if candidates[i] > target {
			break
		}
		// Добавляем текущее число в комбинацию.
		current = append(current, candidates[i])
		// Рекурсивно вызываем функцию для оставшейся части массива.
		backtrack(candidates, target-candidates[i], i+1, current, result)
		// Удаляем последнее число из комбинации для продолжения перебора.
		current = current[:len(current)-1]
	}
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	result := combinationSum2(candidates, target)
	fmt.Println(result)
}
