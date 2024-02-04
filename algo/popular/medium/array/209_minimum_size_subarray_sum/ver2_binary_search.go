package _09_minimum_size_subarray_sum

// Функция minSubArrayLen находит минимальный подмассив, сумма элементов которого не меньше target
func minSubArrayLenV2(target int, nums []int) int {
	/*
		METHOD: Binary search
		Time complexity: O(n)
		Space complexity: O(n)

		Time complexity
		Создание префиксных сумм занимает O(n) времени, где n - количество элементов в массиве.
		Бинарный поиск выполняется для каждого префиксной суммы, и в худшем случае он выполняет O(log n) операций,
		где n - количество элементов в массиве. Таким образом, общее время выполнения алгоритма - O(n log n).

		Space complexity
		Алгоритм использует дополнительное пространство для хранения префиксных сумм, поэтому пространственная сложность - O(n),
		где n - количество элементов в массиве.
	*/

	// Получаем количество элементов в массиве
	n := len(nums)

	// Создаем массив префиксных сумм
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	// Инициализируем минимальную длину подмассива
	minLen := n + 1

	// Проходим по всем префиксным суммам
	for i := 0; i <= n; i++ {
		// Используем бинарный поиск для нахождения конечного индекса подмассива
		end := binarySearch(i, n, sums[i]+target, sums)

		// Если конечный индекс меньше или равен n и длина подмассива меньше minLen,
		// то обновляем minLen
		if end <= n && end-i < minLen {
			minLen = end - i
		}
	}

	// Если minLen не изменился, то возвращаем 0, так как подмассив не найден
	if minLen == n+1 {
		return 0
	}

	// Возвращаем минимальную длину подмассива
	return minLen
}

// Функция binarySearch используется для нахождения конечного индекса подмассива,
// сумма элементов которого не меньше target
func binarySearch(start, end, target int, sums []int) int {
	left, right := start, end
	for left < right {
		mid := (left + right) / 2
		if sums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
