package arrays_intersection

func intersection(nums1 []int, nums2 []int) []int {
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