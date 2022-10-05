package main

func twoSumSorted(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	if nums[0] > target {
		return false
	}

	length := len(nums)
	min := 0
	max := length - 1
	for {
		sum := nums[min] + nums[max]

		if sum == target {
			return true
		} else if min+1 == length {
			return false
		} else if max-1 == -1 {
			return false
		} else if sum < target {
			min += 1
		} else if sum > target {
			max -= 1
		} else if min == max {
			return false
		}
	}
}
