package main

import (
	"fmt"
	"sort"
)

func deleteAndEarnGreedy(nums []int) int {
	/*
		METHOD: Greedy (sort by sum descending, pick if not adjacent to last picked)
		1. Группируем числа по значению и считаем сумму для каждого уникального числа.
		2. Сортируем уникальные числа по сумме убыванию.
		3. Проходим по отсортированному списку, берем число, если оно не сосед (отличается на 1) с последним взятым.

		TIME COMPLEXITY: O(n log n) (сортировка уникальных чисел)
		SPACE COMPLEXITY: O(n) (map для сумм)
	*/
	if len(nums) == 0 {
		return 0
	}

	// Группируем числа по значению и считаем сумму
	sumMap := make(map[int]int)
	for _, num := range nums {
		sumMap[num] += num
	}

	// Создаем слайс пар (сумма, число) для сортировки по сумме убыванию
	type pair struct {
		sum int
		num int
	}

	pairs := make([]pair, 0, len(sumMap))
	for num, sum := range sumMap {
		pairs = append(pairs, pair{sum, num})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].sum > pairs[j].sum
	})

	total := 0
	lastPicked := -100 // большое отрицательное, чтобы первое всегда взять

	for _, p := range pairs {
		if abs(p.num-lastPicked) != 1 {
			total += p.sum
			lastPicked = p.num
		}
	}

	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Примеры
	fmt.Println(deleteAndEarnGreedy([]int{3, 4, 2}))          // 6
	fmt.Println(deleteAndEarnGreedy([]int{2, 2, 3, 3, 3, 4})) // 9
}
