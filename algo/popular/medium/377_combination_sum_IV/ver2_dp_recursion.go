package main

// CombinationSum4V2 вычисляет количество комбинаций чисел из nums, которые в сумме дают target.
// Этот вариант использует подход "сверху вниз" с мемоизацией для оптимизации.
func combinationSum4V2(nums []int, target int) int {
	/*
		METHOD: Top-Down Dynamic Programming with Memoization
		Этот метод использует рекурсию с мемоизацией для хранения уже вычисленных результатов и избежания повторных вычислений.

		TIME COMPLEXITY: O(target * n), где n — количество элементов в массиве nums.
		Мы проходим по всем значениям от 1 до target и для каждого значения проходим по всем числам в массиве nums.
		Мемоизация позволяет нам избежать повторных вычислений, улучшая общую производительность.

		SPACE COMPLEXITY: O(target)
		Мы используем дополнительное пространство для хранения результатов в кэше размером target + 1.
		Также учитываем стек вызовов рекурсии, который может достигать глубины target в худшем случае.
	*/
	return topDown(nums, target, make(map[int]int))
}

// topDown рекурсивно вычисляет количество комбинаций для заданного target, используя мемоизацию.
func topDown(nums []int, target int, cache map[int]int) int {
	// Если target равен 0, то существует только одна комбинация (пустая), которая в сумме дает 0.
	var res int
	if target == 0 {
		return 1
	}

	// Проверяем, есть ли уже вычисленный результат для данного target в кэше.
	if val, ok := cache[target]; ok {
		return val
	}

	// Перебираем все числа в массиве nums
	for _, item := range nums {
		// Если текущий target больше или равен item, то рекурсивно вычисляем количество комбинаций для target - item
		if target >= item {
			res += topDown(nums, target-item, cache)
		}
	}

	// Сохраняем результат в кэше
	cache[target] = res

	// Возвращаем результат
	return cache[target]
}
