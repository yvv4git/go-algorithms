package v3

func intersection(nums1 []int, nums2 []int) []int {
	/*
		METHOD: Hash
		Time complexity: O(n)
		Space complexity: O(n)
	*/
	tmp := make(map[int]int)
	temp := []int{}

	for _, v := range nums1 {
		for _, V := range nums2 {
			if v == V {
				tmp[v]++
			}
		}
	}

	for i := range tmp {
		temp = append(temp, i)
	}

	return temp

}
