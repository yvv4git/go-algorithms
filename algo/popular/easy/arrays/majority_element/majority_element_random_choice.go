package main

import "math/rand"

func majorityElementRandomChoice(nums []int) int {
	/*
		METHOD: Random choice
		Time complexity : O(∞) - тут может быть по разному.
		Space complexity : O(1)
	*/
	size := len(nums)
	majorElem := size / 2

	if size == 0 {
		return 0
	}

	for {
		choiceIdx := rand.Intn(size)
		element := nums[choiceIdx]

		count := 0
		for _, num := range nums {
			if num == element {
				count++
			}
		}

		if count > majorElem {
			return element
		}
	}
}
