package set_mismatch

func findErrorNumsV1(nums []int) []int {
	/*
		METHOD: Cycle sort
		Time complexity: O(n)
		Space complexity: O(1)
	*/

	// Какая-то сортировка
	i := 0
	for i < len(nums) {
		pos := nums[i] - 1

		if nums[i]-1 != i && nums[i] != nums[pos] {
			nums[i], nums[pos] = nums[pos], nums[i]
		} else {
			i++
		}
	}

	var result []int

	for n := 0; n < len(nums); n++ {
		if nums[n] != n+1 {
			result = append(result, nums[n])
			result = append(result, n+1)
			break
		}
	}

	return result
}
