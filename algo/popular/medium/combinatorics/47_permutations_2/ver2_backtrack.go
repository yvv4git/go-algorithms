package main

// Функция permuteUniqueV2 возвращает все уникальные перестановки чисел в nums.
func permuteUniqueV2(nums []int) [][]int {
	/*
		METHOD: Backtrack
		TIME COMPLEXITY: O(n * n!) где n - количество чисел в массиве, так как для каждого элемента мы генерируем все возможные перестановки, и таких перестановок n!, где n - количество элементов в массиве.
		SPACE COMPLEXITY: O(n * n!) для хранения результатов перестановок. Каждая перестановка занимает O(n) места, и их всего n!, где n - количество элементов в массиве.
	*/
	var res [][]int // Срез для хранения результатов перестановок.

	// Внутренняя функция backtrack выполняет рекурсивный обход для построения перестановок.
	var backtrack func(int)
	backtrack = func(idx int) {
		// Базовый случай: если индекс равен длине nums, то перестановка завершена.
		if idx == len(nums) {
			// Добавляем копию текущей перестановки в результаты.
			res = append(res, append([]int(nil), nums...))
			return
		}

		// Карта для отслеживания использованных чисел в текущей перестановке.
		used := make(map[int]bool)

		// Проходим по оставшимся числам для перестановки.
		for i := idx; i < len(nums); i++ {
			// Если число уже использовалось в текущей перестановке, пропускаем его.
			if used[nums[i]] {
				continue
			}

			// Меняем местами текущее число с числом на позиции idx.
			nums[i], nums[idx] = nums[idx], nums[i]
			// Рекурсивно строим перестановки для следующего индекса.
			backtrack(idx + 1)
			// Возвращаем исходное состояние перестановки.
			nums[i], nums[idx] = nums[idx], nums[i]

			// Помечаем число как использованное.
			used[nums[i]] = true
		}
	}

	// Начинаем построение перестановок с начального индекса 0.
	backtrack(0)

	// Возвращаем результаты.
	return res
}
