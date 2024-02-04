package set_mismatch

func findErrorNumsV2(nums []int) []int {
	/*
		METHOD: Hash
		Time complexity: O(n)
		Space complexity: O(n)
	*/
	numsMap := map[int]int{}
	for _, num := range nums {
		numsMap[num]++
	}
	missing := -1
	duplicated := -1
	for i := 1; i <= len(nums); i++ {
		if count, found := numsMap[i]; found {
			if count >= 2 {
				duplicated = i
			}
		} else {
			missing = i
		}
	}
	return []int{duplicated, missing}
}
