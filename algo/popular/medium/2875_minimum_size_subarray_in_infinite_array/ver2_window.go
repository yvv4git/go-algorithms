package _875_minimum_size_subarray_in_infinite_array

import "math"

func minSizeSubarrayV2(nums []int, target int) int {
	n := len(nums)
	var sum int
	for _, num := range nums {
		sum += num
	}

	repeats := target / sum
	target = target % sum
	if target == 0 {
		return repeats * n
	}
	res := math.MaxInt32
	var left, subsum int

	for right := 0; right < 2*len(nums); right++ {
		subsum += nums[right%n]
		for left < right && subsum > target {
			subsum -= nums[left%n]
			left++
		}
		if subsum == target {
			res = min(res, right-left+1)
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return repeats*n + res
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
