package basis

func majorityElement(nums []int) int {
	cnt := 0
	result := 0

	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			result = nums[i]
			cnt++
		} else if nums[i] == result {
			cnt++
		} else {
			cnt--
		}
	}

	return result
}
