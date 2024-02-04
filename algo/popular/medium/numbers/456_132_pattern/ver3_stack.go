package _56_132_pattern

import "math"

func find132patternV3(nums []int) bool {
	/*
		METHOD: Two pointers + stack
		TIME COMPLEXITY: O(n), где n - размер входного массива, потому что мы проходим по массиву один раз.
		SPACE COMPLEXITY: O(n), потому что мы создаем стек для хранения возможных второй части 132-последовательности. Стек может содержать до n элементов в худшем случае.
	*/
	// Инициализируем стек и переменную second
	stack := make([]int, 0)
	second := math.MinInt64

	// Проходим по массиву справа налево
	for i := len(nums) - 1; i >= 0; i-- {
		// Если текущий элемент меньше second, то мы нашли 132-последовательность
		if nums[i] < second {
			return true
		}

		// Если стек не пуст и элемент на вершине стека меньше текущего элемента,
		// то обновляем second и удаляем элементы из стека, пока элемент на вершине стека меньше текущего элемента
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			second = max(second, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}

		// Добавляем текущий элемент в стек
		stack = append(stack, nums[i])
	}

	// Если мы не нашли 132-последовательность, то возвращаем false
	return false
}

// Функция max возвращает максимальное значение из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
