package _53_find_minimum_in_rotated_sorted_array

// используя модифицированный бинарный поиск.
// Временная сложность: O(log n) - поскольку мы каждый раз делим массив на половины.
// Пространственная сложность: O(1) - поскольку мы не используем дополнительное пространство,
// которое зависит от размера входных данных.
func findMinV2(nums []int) int {
	/*
		METHOD: Binary search optimized
		TIME COMPLEXITY: O(log n)
		Space complexity: O(1)
	*/
	// Инициализируем две переменные, которые будут указывать на начало и конец массива.
	left, right := 0, len(nums)-1

	// Если первый элемент меньше последнего, то массив не повернут.
	if nums[left] < nums[right] {
		return nums[left]
	}

	// Пока левый указатель меньше правого, выполняем бинарный поиск.
	for left < right {
		// Находим середину.
		mid := left + (right-left)/2

		// Если средний элемент больше следующего, то следующий элемент - минимальный.
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		// Если средний элемент меньше предыдущего, то средний элемент - минимальный.
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		// Если средний элемент больше первого элемента, то минимальный элемент находится в правой половине.
		if nums[mid] > nums[0] {
			left = mid + 1
		} else {
			// Иначе минимальный элемент находится в левой половине.
			right = mid - 1
		}
	}

	// Если мы не нашли минимального элемента, возвращаем первый элемент.
	return nums[0]
}
