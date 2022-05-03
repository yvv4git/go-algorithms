package main

// maxSubArray - найти subarray с максимальной суммой элементов
func maxSubArray(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	max := nums[0]
	sum := max

	for _, n := range nums[1:] {
		if sum+n > n {
			sum += n
		} else {
			sum = n
		}

		if max < sum {
			max = sum
		}
	}

	return max
}
