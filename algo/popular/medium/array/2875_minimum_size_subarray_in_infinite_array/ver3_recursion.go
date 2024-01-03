package _875_minimum_size_subarray_in_infinite_array

import "math"

func minSubArrayLenV3(nums []int, target int) int {
	/*
		Method: Recursion
		Time complexity: O(n^2), это связано с тем, что для каждого элемента мы можем вызвать функцию рекурсивно до тех пор, пока сумма элементов не станет больше или равна целевому значению.
		Space complexity: O(n^2)
	*/
	return minSubArrayLenRecursive(nums, target, 0, 0, 0, math.MaxInt32)
}

func minSubArrayLenRecursive(nums []int, target int, left int, right int, sum int, minLen int) int {
	if right == len(nums) {
		if sum < target {
			if minLen == math.MaxInt32 {
				return -1
			}
			return minLen
		}
		return minSubArrayLenRecursive(nums, target, left, 0, 0, minLen)
	}

	if sum >= target {
		minLen = min(minLen, right-left)
		return minSubArrayLenRecursive(nums, target, left+1, right, sum-nums[left], minLen)
	}

	if right < len(nums) {
		return minSubArrayLenRecursive(nums, target, left, right+1, sum+nums[right], minLen)
	}

	return minLen
}
