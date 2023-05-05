package sort_array_by_parity

func sortArrayByParity(nums []int) []int {
	/*
		Time complexity : O(n)
		Space complexity : O(1)

		Надо все четные числа переместить в начало массива.
		Этого можно достичь операцией swap, если пройтись по всему массиву.
	*/
	evenIdx := 0 // Четный
	for i := range nums {
		if nums[i]%2 == 0 {
			nums[evenIdx], nums[i] = nums[i], nums[evenIdx]
			evenIdx++
		}
	}

	return nums
}
