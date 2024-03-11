package main

import "sort"

func subsetsWithDupV4(nums []int) [][]int {
	/*
		METHOD: Bitmask

		TIME COMPLEXITY: O(N * 2^N), где N - количество элементов в массиве nums.
		Это связано с тем, что для каждого элемента мы можем выбрать или не выбрать (две возможные ветви рекурсии),
		и таких вариантов в итоге будет 2^N.

		SPACE COMPLEXITY: O(N * 2^N), так как в худшем случае мы можем получить 2^N подмножеств, каждое из которых может иметь длину N.
	*/
	sort.Ints(nums)
	n := len(nums)
	result := [][]int{}

	// Итерируемся по всем возможным битовым маскам
	for mask := 0; mask < (1 << n); mask++ {
		subset := []int{}
		duplicate := false

		// Проверяем каждый бит в битовой маске
		for i := 0; i < n; i++ {
			// Если i-й бит установлен в 1
			if mask&(1<<i) != 0 {
				// Если это не первый элемент и предыдущий элемент равен текущему
				// и предыдущий элемент не был включен в подмножество (через эту маску)
				// то это дубликат и мы пропускаем его
				if i > 0 && nums[i] == nums[i-1] && mask&(1<<(i-1)) == 0 {
					duplicate = true
					break
				}
				subset = append(subset, nums[i])
			}
		}

		// Если не было найдено дубликатов, добавляем подмножество в результат
		if !duplicate {
			result = append(result, subset)
		}
	}

	return result
}
