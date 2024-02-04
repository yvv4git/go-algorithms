package _630_arithmetic_subarrays

import "sort"

func checkArithmeticSubArraysV2(nums []int, l []int, r []int) []bool {
	/*
		METHOD: ???
		TIME COMPLEXITY: O(mn log(n))
		Space complexity: O(n)
	*/
	result := []bool{}

	for i, _ := range l {
		beginIdx := l[i]
		endIdx := r[i]

		subArr := []int{}
		subArr = append(subArr, nums[beginIdx:endIdx+1]...)

		sort.Slice(subArr, func(i, j int) bool {
			return subArr[i] < subArr[j]
		})

		result = append(result, isArithmetic(subArr))

	}

	return result
}

func isArithmetic(arr []int) bool {
	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1]-arr[i] != arr[i]-arr[i+1] {
			return false
		}
	}

	return true
}
