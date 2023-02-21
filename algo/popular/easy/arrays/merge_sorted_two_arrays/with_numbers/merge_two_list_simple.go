package with_numbers

import "sort"

func mergeSimple(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	nums2 = nums2[:n]

	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
}
