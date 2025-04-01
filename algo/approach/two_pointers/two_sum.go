package main

import "fmt"

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target)) // [0, 1]
}
