package squares_sorted_array

func sortedSquaresV4(nums []int) []int {
	/*
		METHOD: Two pointer
		TIME COMPLEXITY: O(n)
		Space complexity: O(n) - используется массив, в который помещается результат
	*/
	res := make([]int, len(nums))
	// Три указателя: left, right, i.
	// Переменная left - указывает на левую границу исходного массива.
	// Переменная light - указывает на правую границу исходного массива.
	// Переменная i - это указатель в конец нового массива, результирующего. Он нужен для добавления нового максимального элемента.
	left, right, i := 0, len(nums)-1, len(nums)-1
	for left <= right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[i] = nums[left] * nums[left]
			left++
		} else {
			res[i] = nums[right] * nums[right]
			right--
		}
		i--
	}

	return res
}
