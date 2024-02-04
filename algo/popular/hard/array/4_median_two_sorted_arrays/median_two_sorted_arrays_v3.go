package __median_two_sorted_arrays

func findMedianSortedArraysV3(nums1 []int, nums2 []int) float64 {
	/*
		METHOD: Merge and divide
		TIME COMPLEXITY: 	O(n)
		SPACE COMPLEXITY:	O(1)
	*/
	lenNums1, lenNums2 := len(nums1), len(nums2)

	if lenNums1+lenNums2 == 1 {
		if lenNums1 == 1 {
			return float64(nums1[0])
		}

		if lenNums2 == 1 {
			return float64(nums2[0])
		}
	}

	nums := make([]int, 2)
	i, j, k := 0, 0, 0

	for k <= (lenNums1+lenNums2)/2 {
		if i < lenNums1 && j < lenNums2 {
			if nums1[i] <= nums2[j] {
				nums[k&0x01] = nums1[i]
				i++
			} else {
				nums[k&0x01] = nums2[j]
				j++
			}
		} else if i < lenNums1 {
			nums[k&0x01] = nums1[i]
			i++
		} else {
			nums[k&0x01] = nums2[j]
			j++
		}

		k++
	}

	if (lenNums1+lenNums2)&0x01 == 0 {
		return float64(nums[1]+nums[0]) / 2.0
	} else {
		return float64(nums[(k-1)&0x01])
	}
}
