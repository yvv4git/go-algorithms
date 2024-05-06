package main

import (
	"fmt"
	"sort"
)

func findRelativeRanksV2(score []int) []string {
	/*
		METHOD: Map
		Мы можем отсортировать массив результатов и создать хеш-таблицу, где ключом будет результат,
		а значением - его индекс в отсортированном массиве.
		Затем мы можем пройтись по исходному массиву и использовать хеш-таблицу для определения индекса в отсортированном массиве,
		что позволит нам определить его относительный ранг.

		TIME COMPLEXITY: O(n log n), так как мы сортируем массив.

		SPACE COMPLEXITY: O(n), так как мы храним отсортированный массив и хеш-таблицу.
	*/
	// Создаем копию массива и сортируем его
	sortedNums := make([]int, len(score))
	copy(sortedNums, score)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedNums)))

	// Создаем хеш-таблицу для быстрого поиска индекса результата в отсортированном массиве
	rankMap := make(map[int]int)
	for i, num := range sortedNums {
		rankMap[num] = i + 1
	}

	// Определяем относительные ранги
	ranks := make([]string, len(score))
	for i, num := range score {
		switch rankMap[num] {
		case 1:
			ranks[i] = "Gold Medal"
		case 2:
			ranks[i] = "Silver Medal"
		case 3:
			ranks[i] = "Bronze Medal"
		default:
			ranks[i] = fmt.Sprintf("%d", rankMap[num])
		}
	}

	return ranks
}

//func main() {
//	score := []int{5, 4, 3, 2, 1}
//	ranks := findRelativeRanks(score)
//	fmt.Println(ranks)
//}
