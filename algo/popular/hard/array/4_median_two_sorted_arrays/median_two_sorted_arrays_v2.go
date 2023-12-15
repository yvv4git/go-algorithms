package __median_two_sorted_arrays

func findMedianSortedArraysV2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
	}

	idxMin, idxMax, halfLen := 0, m, (m+n+1)/2
	for idxMin <= idxMax {
		i := (idxMin + idxMax) / 2
		j := halfLen - i

		if i < m && nums2[j-1] > nums1[i] {
			idxMin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			idxMax = i - 1
		} else {
			maxLeftValue, minRightValue := 0, 0
			if i == 0 {
				maxLeftValue = nums2[j-1]
			} else if j == 0 {
				maxLeftValue = nums1[i-1]
			} else {
				maxLeftValue = maxValue(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(maxLeftValue)
			}

			if i == m {
				minRightValue = nums2[j]
			} else if j == n {
				minRightValue = nums1[i]
			} else {
				minRightValue = minValue(nums1[i], nums2[j])
			}

			return float64(maxLeftValue+minRightValue) / 2.0
		}
	}

	return 0.0
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
