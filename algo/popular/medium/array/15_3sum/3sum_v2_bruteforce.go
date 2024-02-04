package _5_3sum

import "sort"

func threeSumV2(nums []int) [][]int {
	/*
		METHOD: Bruteforce.
		TIME COMPLEXITY: O(n^2)
		Space complexity: O(1)
	*/
	sort.Ints(nums) // O(n log n)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ { // O(n^2)
		if i > 0 && nums[i] == nums[i-1] { // Проверяем, что нет повторяющихся элементов
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s < 0 {
				l++
			} else if s > 0 {
				r--
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}

	return result
}
