package difference_two_array

// FindDifferenceWithHashSet - used for find difference of two arrays.
func FindDifferenceWithHashSet(nums1 []int, nums2 []int) [][]int {
	res := make([][]int, 2)
	m := make(map[int]int)

	for _, v := range nums1 {
		m[v] = -1
	}
	for _, v := range nums2 {
		m[v] += 2
	}
	for k, v := range m {
		if v == -1 {
			res[0] = append(res[0], k)
		} else if v%2 == 0 {
			res[1] = append(res[1], k)
		}
	}

	return res
}

// FindDifferenceWithSlices - used for find difference of two arrays(O(n)).
func FindDifferenceWithSlices(nums1 []int, nums2 []int) [][]int {
	var dif2 []int
	var dif1 []int
	nums1Set, nums2Set := map[int]struct{}{}, map[int]struct{}{}
	for _, num := range nums1 {
		nums1Set[num] = struct{}{}
	}
	for _, num := range nums2 {
		nums2Set[num] = struct{}{}
	}

	for num, _ := range nums1Set {
		if _, ok := nums2Set[num]; !ok {
			dif1 = append(dif1, num)
		}
	}

	for num, _ := range nums2Set {
		if _, ok := nums1Set[num]; !ok {
			dif2 = append(dif2, num)
		}
	}

	return [][]int{
		dif1,
		dif2,
	}
}
