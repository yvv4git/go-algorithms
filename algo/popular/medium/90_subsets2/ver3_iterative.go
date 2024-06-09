package main

import "sort"

func subsetsWithDupV3(nums []int) [][]int {
	/*
		METHOD: Iterative

		TIME COMPLEXITY: O(N * 2^N), где N - количество элементов в массиве nums.
		Это связано с тем, что для каждого элемента мы можем выбрать или не выбрать (две возможные ветви рекурсии),
		и таких вариантов в итоге будет 2^N.

		SPACE COMPLEXITY: O(N * 2^N), так как в худшем случае мы можем получить 2^N подмножеств, каждое из которых может иметь длину N.
	*/
	sort.Ints(nums)
	result := [][]int{{}} // Начинаем с пустого подмножества
	startIndex := 0
	endIndex := 0

	for i := 0; i < len(nums); i++ {
		startIndex = 0
		// Если текущий элемент не дубликат предыдущего, то расширяем все подмножества
		// Иначе только добавляем к новым подмножествам
		if i > 0 && nums[i] == nums[i-1] {
			startIndex = endIndex
		}
		endIndex = len(result)
		for j := startIndex; j < endIndex; j++ {
			newSubset := append([]int{}, result[j]...)
			newSubset = append(newSubset, nums[i])
			result = append(result, newSubset)
		}
	}
	return result
}
