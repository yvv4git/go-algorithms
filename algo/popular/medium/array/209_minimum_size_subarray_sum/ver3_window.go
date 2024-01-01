package _09_minimum_size_subarray_sum

func minSubArrayLenV3(target int, nums []int) int {
	/*
		Method: Window
		Time complexity: O(n)
		Space complexity: O(1)

		Time complexity
		Алгоритм проходит по массиву один раз, поэтому временная сложность - O(n), где n - количество элементов в массиве.

		Space complexity
		Алгоритм использует дополнительное пространство для хранения суммы, минимальной длины и левой границы окна,
		поэтому пространственная сложность - O(1).
	*/
	// Инициализируем переменные для отслеживания текущей суммы, минимальной длины и левой границы окна
	sum, minLen, left := 0, len(nums)+1, 0

	// Итерируемся по массиву
	for right := 0; right < len(nums); right++ {
		// Добавляем текущий элемент к сумме
		sum += nums[right]

		// Пока сумма больше или равна target, мы пытаемся уменьшить окно, удаляя элементы слева
		for sum >= target {
			// Обновляем минимальную длину, если необходимо
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			// Удаляем элемент слева и уменьшаем сумму
			sum -= nums[left]
			left++
		}
	}

	// Если минимальная длина не изменилась, то возвращаем 0, так как подмассив не найден
	if minLen == len(nums)+1 {
		return 0
	}

	// Возвращаем минимальную длину подмассива
	return minLen
}
