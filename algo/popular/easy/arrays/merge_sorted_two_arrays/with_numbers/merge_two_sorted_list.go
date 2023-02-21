package with_numbers

func mergeV3(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	j, k := m-1, n-1

	for j >= 0 && k >= 0 {
		if nums1[j] > nums2[k] {
			nums1[i] = nums1[j]
			j--
		} else {
			nums1[i] = nums2[k]
			k--
		}
		i--
	}

	if k >= 0 {
		copy(nums1[:k+1], nums2[:k+1])
	}
}
