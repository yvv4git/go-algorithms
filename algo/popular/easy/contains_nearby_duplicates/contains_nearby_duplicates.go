package main

func containsNearbyDuplicate(nums []int, k int) bool {
	tmp := make(map[int]int, len(nums))
	for i, num := range nums {
		if index, existValue := tmp[num]; existValue {
			if i-index <= k {
				return true
			}
		}
		tmp[num] = i
	}

	return false
}
