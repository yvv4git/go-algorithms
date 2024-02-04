package _875_minimum_size_subarray_in_infinite_array

/*
	METHOD: Window
	Time complexity: O(n)
	Space complexity: O(1)

	Time complexity
	Алгоритм проходит по бесконечному массиву, поэтому временная сложность - O(n),
	где n - количество элементов в бесконечном массиве.

	Space complexity
	Алгоритм использует дополнительное пространство для хранения суммы, минимальной длины и левой границы окна,
	поэтому пространственная сложность - O(1).
*/

// Функция minSubArrayLen находит минимальный под массив, сумма элементов которого не меньше target
func minSubArrayLenV1(nums []int, target int) int {
	//left, sum, minLen := 0, 0, math.MaxInt32
	//for right := range nums {
	//	sum += nums[right]
	//	// Прибавляем значение по индексу right до тех пор, пока сумма(sum) меньше желаемого(target).
	//	// Если же сумма(sum) превысила target, значит эта последовательность элементов массива нам не подходит.
	//	// Сдвигаем окно слева, увеличивая на 1 значение left окна.
	//	for sum > target {
	//		minLen = min(minLen, right-left+1)
	//		sum -= nums[left]
	//		left++
	//	}
	//}
	//
	//if sum < target {
	//	return -1
	//}
	//
	//return minLen
	result := 0
	left := 0
	sum := 0
	for right := range nums {
		sum += nums[right]
		result++

		for sum > target {
			sum -= nums[left]
			left++
			result--

			if sum == target {
				return result
			}
		}
	}

	if sum == target {
		return result
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
