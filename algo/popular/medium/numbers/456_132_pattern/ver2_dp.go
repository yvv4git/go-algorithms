package _56_132_pattern

import "math"

func find132patternV2(nums []int) bool {
	/*
		Method Dynamic programming
		Time complexity: O(n)
		Space complexity: O(n)

		Time complexity
		Временная сложность этого алгоритма O(n), где n - размер входного массива, потому что мы проходим по массиву дважды:
		один раз для создания массива min, а второй раз для поиска 132-последовательности.

		Space complexity
		Пространственная сложность также O(n), потому что мы создаем два дополнительных массива:
		min для хранения минимальных значений слева от каждого элемента, и stack для хранения возможных второй части 132-последовательности.
	*/
	// Если длина массива меньше 3, то 132-последовательность невозможна, поэтому мы возвращаем false.
	if len(nums) < 3 {
		return false
	}

	// Создаем массив min, в котором для каждого элемента хранится минимальное значение слева от него.
	min := make([]int, len(nums))
	min[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		min[i] = minVal(min[i-1], nums[i])
	}

	// Создаем стек для хранения возможных второй части 132-последовательности.
	stack := make([]int, 0)
	max := math.MinInt64
	// Проходим по массиву справа налево.
	for j := len(nums) - 1; j >= 0; j-- {
		// Если текущий элемент больше минимального значения слева, то проверяем, может ли он быть средним элементом в 132-последовательности.
		if nums[j] > min[j] {
			// Удаляем из стека все элементы, которые меньше или равны минимальному значению слева.
			for len(stack) > 0 && stack[len(stack)-1] <= min[j] {
				stack = stack[:len(stack)-1]
			}
			// Если стек не пуст и последний элемент стека меньше текущего элемента, то мы нашли 132-последовательность.
			if len(stack) > 0 && stack[len(stack)-1] < nums[j] {
				return true
			}
			// Добавляем текущий элемент в стек.
			stack = append(stack, nums[j])
			// Обновляем максимальное значение.
			max = maxVal(max, nums[j])
		}
	}

	// Если мы не нашли 132-последовательность, то возвращаем false.
	return false
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
