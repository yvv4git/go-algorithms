package main

import "fmt"

func maxSumSubarray(nums []int, k int) int {
	maxSum := 0
	windowSum := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		windowSum += nums[right]

		if right >= k-1 {
			if windowSum > maxSum {
				maxSum = windowSum
			}
			windowSum -= nums[left]
			left++
		}
	}

	return maxSum
}

func main() {
	nums := []int{1, -1, 5, -2, 3}
	k := 3
	fmt.Println(maxSumSubarray(nums, k)) // 6
}
