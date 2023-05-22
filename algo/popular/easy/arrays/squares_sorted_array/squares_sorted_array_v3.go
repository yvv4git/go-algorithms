package squares_sorted_array

func sortedSquaresV3(nums []int) []int {
	/*
		Method: Clean code
		Time complexity: O(n)
		Space complexity: O(n)

		В задаче есть закономерность - слева будет самое большое отрицательное число, а справа самое большое положительное.
		Следовательно, после возведения в квадрат, их можно будет сравнивать - самое левое и самое правое.
		Остается только смещать pointer(left, right) и добавлять новое самое большое число в результирующий массив.
	*/
	result := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for i := 0; left <= right; i++ {
		if abs(nums[left]) > abs(nums[right]) {
			result[len(nums)-1-i] = nums[left] * nums[left]
			left++
		} else {
			result[len(nums)-1-i] = nums[right] * nums[right]
			right--
		}
	}

	return result
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
